#!/bin/bash

set -o pipefail

# config
default_semvar_bump=${DEFAULT_BUMP:-patch}
with_v=${WITH_V:-true}
release_branches=${RELEASE_BRANCHES:-}
custom_tag=${CUSTOM_TAG:-}
source=${SOURCE:-.}
dryrun=${DRY_RUN:-false}
initial_version=${INITIAL_VERSION:-0.0.0}
tag_context=${TAG_CONTEXT:-repo}
prerelease=${PRERELEASE:-true}
suffix=${PRERELEASE_SUFFIX:-SNAPSHOT}
verbose=${VERBOSE:-false}
major_string_token=${MAJOR_STRING_TOKEN:-#major}
minor_string_token=${MINOR_STRING_TOKEN:-#minor}
patch_string_token=${PATCH_STRING_TOKEN:-#patch}
none_string_token=${NONE_STRING_TOKEN:-#none}
# since https://github.blog/2022-04-12-git-security-vulnerability-announced/ runner uses?
git config --global --add safe.directory /github/workspace

cd "${GITHUB_WORKSPACE}/${source}" || exit 1

echo "*** CONFIGURATION ***"
echo -e "\tDEFAULT_BUMP: ${default_semvar_bump}"
echo -e "\tWITH_V: ${with_v}"
echo -e "\tRELEASE_BRANCHES: ${release_branches}"
echo -e "\tCUSTOM_TAG: ${custom_tag}"
echo -e "\tSOURCE: ${source}"
echo -e "\tDRY_RUN: ${dryrun}"
echo -e "\tINITIAL_VERSION: ${initial_version}"
echo -e "\tTAG_CONTEXT: ${tag_context}"
echo -e "\tPRERELEASE: ${prerelease}"
echo -e "\tPRERELEASE_SUFFIX: ${suffix}"
echo -e "\tVERBOSE: ${verbose}"
echo -e "\tMAJOR_STRING_TOKEN: ${major_string_token}"
echo -e "\tMINOR_STRING_TOKEN: ${minor_string_token}"
echo -e "\tPATCH_STRING_TOKEN: ${patch_string_token}"
echo -e "\tNONE_STRING_TOKEN: ${none_string_token}"

# verbose, show everything
if $verbose
then
    set -x
fi

setOutput() {
    echo "${1}=${2}" >> "${GITHUB_OUTPUT}"
}

current_branch=$(git rev-parse --abbrev-ref HEAD)

pre_release="$prerelease"
IFS=',' read -ra branch <<< "$release_branches"
for b in "${branch[@]}"; do
    # check if ${current_branch} is in ${release_branches} | exact branch match
    if [[ "$current_branch" == "$b" ]]
    then
        pre_release="false"
    fi
    # verify non specific branch names like  .* release/* if wildcard filter then =~
    if [ "$b" != "${b//[\[\]|.? +*]/}" ] && [[ "$current_branch" =~ $b ]]
    then
        pre_release="false"
    fi
done
echo "pre_release = $pre_release"

# fetch tags
git fetch --tags

tagFmt="^v?[0-9]+\.[0-9]+\.[0-9]+$"
preTagFmt="^v?[0-9]+\.[0-9]+\.[0-9]+(-$suffix\.[0-9]+)$"

# get latest tag that looks like a semver (with or without v)
case "$tag_context" in
    *repo*) 
        tag="$(git for-each-ref --sort=-v:refname --format '%(refname:lstrip=2)' | grep -E "$tagFmt" | head -n 1)"
        pre_tag="$(git for-each-ref --sort=-v:refname --format '%(refname:lstrip=2)' | grep -E "$preTagFmt" | head -n 1)"
        ;;
    *branch*) 
        tag="$(git tag --list --merged HEAD --sort=-v:refname | grep -E "$tagFmt" | head -n 1)"
        pre_tag="$(git tag --list --merged HEAD --sort=-v:refname | grep -E "$preTagFmt" | head -n 1)"
        ;;
    * ) echo "Unrecognised context"
        exit 1;;
esac

# if there are none, start tags at INITIAL_VERSION
if [ -z "$tag" ]
then
    if $with_v
    then
        tag="v$initial_version"
    else
        tag="$initial_version"
    fi
    if [ -z "$pre_tag" ] && $pre_release
    then
        if $with_v
        then
            pre_tag="v$initial_version"
        else
            pre_tag="$initial_version"
        fi
    fi
fi

# get current commit hash for tag
tag_commit=$(git rev-list -n 1 "$tag")

# get current commit hash
commit=$(git rev-parse HEAD)

if [ "$tag_commit" == "$commit" ]
then
    echo "No new commits since previous tag. Skipping..."
    setOutput "new_tag" "$tag"
    setOutput "tag" "$tag"
    exit 0
fi

# get the merge commit message looking for #bumps
log=$(git show -s --format=%s)
echo "Last commit message: $log"

case "$log" in
    *$major_string_token* ) new=$(semver -i major "$tag"); part="major";;
    *$minor_string_token* ) new=$(semver -i minor "$tag"); part="minor";;
    *$patch_string_token* ) new=$(semver -i patch "$tag"); part="patch";;
    *$none_string_token* ) 
        echo "Default bump was set to none. Skipping..."
        setOutput "new_tag" "$tag"
        setOutput "tag" "$tag"
        exit 0;;
    * ) 
        if [ "$default_semvar_bump" == "none" ]
        then
            echo "Default bump was set to none. Skipping..."
            setOutput "new_tag" "$tag"
            setOutput "tag" "$tag"
            exit 0 
        else 
            new=$(semver -i "${default_semvar_bump}" "$tag")
            part=$default_semvar_bump 
        fi 
        ;;
esac

if $pre_release
then
    # already a pre-release available, bump it
    newPreTagFmt="$new+(-$suffix\.[0-9]+)$"
    exists="$(git tag --list --merged HEAD --sort=-v:refname | grep -E "$newPreTagFmt" | head -n 1)"
    if [[ $exists != "" ]]
    then
      echo -e "Found parent to ${new} pre-tag ${exists}..."
      new=$(semver -i prerelease "${exists}" --preid "${suffix}")
      echo -e "Bumping ${suffix} pre-tag ${exists}. New pre-tag ${new}"
    elif [[ "$pre_tag" =~ $new ]] && [[ "$pre_tag" =~ $suffix ]]
    then
      new=$(semver -i prerelease "${pre_tag}" --preid "${suffix}")
      echo -e "Bumping ${suffix} pre-tag ${pre_tag}. New pre-tag ${new}"
    else
      new="$new-$suffix.0"
      echo -e "Setting ${suffix} pre-tag ${pre_tag} - With pre-tag ${new}"
    fi
    part="pre-$part"
else
  echo -e "Bumping tag ${tag} - New tag ${new}"
fi

if $with_v
then
    new="v$new"
fi

# as defined in readme if CUSTOM_TAG is used any semver calculations are irrelevant.
if [ -n "$custom_tag" ]
then
    new="$custom_tag"
fi

# set outputs
setOutput "new_tag" "$new"
setOutput "part" "$part"
setOutput "tag" "$new" # this needs to go in v2 is breaking change
setOutput "old_tag" "$tag"

# dry run exit without real changes
if $dryrun
then
    exit 0
fi

# create local git tag
git tag "$new"

# push new tag ref to github
dt=$(date '+%Y-%m-%dT%H:%M:%SZ')
full_name=$GITHUB_REPOSITORY
git_refs_url=$(jq .repository.git_refs_url "$GITHUB_EVENT_PATH" | tr -d '"' | sed 's/{\/sha}//g')

echo "$dt: **pushing tag $new to repo $full_name"

git_refs_response=$(
curl -s -X POST "$git_refs_url" \
-H "Authorization: token $GITHUB_TOKEN" \
-d @- << EOF

{
  "ref": "refs/tags/$new",
  "sha": "$commit"
}
EOF
)

git_ref_posted=$( echo "${git_refs_response}" | jq .ref | tr -d '"' )

echo "::debug::${git_refs_response}"
if [ "${git_ref_posted}" = "refs/tags/${new}" ]
then
    exit 0
else
    echo "::error::Tag was not created properly."
    exit 1
fi

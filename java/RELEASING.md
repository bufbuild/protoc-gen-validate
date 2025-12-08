# Releasing PGV Java components

These steps are for releasing the Java components of PGV:

- pgv-java-stub
- pgv-java-grpc
- pgv-artifacts

## Releasing using CI

Releasing from main is fully automated by CI:

Releasing from versioned tags is similar. To release version `vX.Y.Z`, first
create a Git tag called `vX.Y.Z` (preferably through the GitHub release flow),
then run the following to kick off a release build:

The [`Maven Deploy`](../.github/workflows/maven-publish.yaml) CI flow will use the version number from the tag to deploy
to the Maven repository.

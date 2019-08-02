These steps are for releasing the Java components of PGV:
- pgv-java-stub
- pgv-java-grpc
- pgv-artifacts

```
curl -X POST -H "Content-Type: application/json" -d '{
"build_parameters": {
    "CIRCLE_JOB": "javabuild", 
    "RELEASE": "<release-version>",
    "NEXT": "<next-version>-SNAPSHOT",
    "GIT_USER_EMAIL": "envoy-bot@users.noreply.github.com",
    "GIT_USER_NAME": "Via CircleCI"
}}' "https://circleci.com/api/v1.1/project/github/envoyproxy/protoc-gen-validate/tree/master?circle-token=<my-token>"
```

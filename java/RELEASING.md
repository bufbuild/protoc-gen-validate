These steps are for releasing the Java components of PGV:
- pgv-java-stub
- pgv-java-grpc
- pgv-artifacts

```
curl -X POST -H "Content-Type: application/json" -d '{
"build_parameters": {
    "RELEASE": "0.1.0",
    "NEXT": "0.2.0-SNAPSHOT",
    "GIT_USER_EMAIL": "foo@bar.org",
    "GIT_USER_NAME": "Via CircleCI"
}}' "https://circleci.com/api/v1.1/project/github/envoyproxy/java-control-plane/tree/master?circle-token=<my-token>"
```

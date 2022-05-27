# TEST-DAPR

```shell
docker login acrcontainerapps.azurecr.io

go install github.com/google/ko@latest
KO_DOCKER_REPO=acrcontainerapps.azurecr.io/test-dapr

az containerapp update \
  --name hello-world \
  --resource-group rg-container-apps \
  --image $(ko build ./)
```

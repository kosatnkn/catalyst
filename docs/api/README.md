## Mock

### Perishable Container

**Linux**
```bash
docker run --name sample_mock -it --rm -p 3000:8080 -v $PWD/openapi.yaml:/openapi/openapi.yaml -e "OPENAPI_MOCK_SPECIFICATION_URL=/openapi/openapi.yaml" muonsoft/openapi-mock
```

**Windows**
```bash
docker run --name sample_mock -it --rm -p 3000:8080 -v %cd%\\openapi.yaml:/openapi/openapi.yaml -e "OPENAPI_MOCK_SPECIFICATION_URL=/openapi/openapi.yaml" muonsoft/openapi-mock
```

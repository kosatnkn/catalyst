## Mock

### Using muonsoft/openapi-mock

**Linux**
```bash
docker run --name sample_mock -it --rm -p 3000:8080 -v $PWD/docs/api/openapi.yaml:/openapi/openapi.yaml -e "OPENAPI_MOCK_SPECIFICATION_URL=/openapi/openapi.yaml" muonsoft/openapi-mock
```

**Windows**
```bash
docker run --name sample_mock -it --rm -p 3000:8080 -v %cd%\\docs\\api\\openapi.yaml:/openapi/openapi.yaml -e "OPENAPI_MOCK_SPECIFICATION_URL=/openapi/openapi.yaml" muonsoft/openapi-mock
```

### Using swaggermock/swagger-mock

**Linux**
```bash
docker --name sample_mock run -it --rm -p 3000:8080 -v $PWD/docs/api/openapi.yaml:/openapi/openapi.yaml -e "SWAGGER_MOCK_SPECIFICATION_URL=/openapi/openapi.yaml" swaggermock/swagger-mock
```

**Windows**
```bash
docker run --name sample_mock -it --rm -p 3000:8080 -v %cd%\\docs\\api\\openapi.yaml:/openapi/openapi.yaml -e "SWAGGER_MOCK_SPECIFICATION_URL=/openapi/openapi.yaml" swaggermock/swagger-mock
```
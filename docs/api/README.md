## Mock

### Perishable Container

**Linux**
```bash
docker run -it --rm -p 3000:8080 -v $PWD/doc/api/openapi.yaml:/openapi/openapi.yaml -e "SWAGGER_MOCK_SPECIFICATION_URL=/openapi/openapi.yaml" swaggermock/swagger-mock
```

**Windows**
```bash
docker run -it --rm -p 3000:8080 -v %cd%\\doc\\api\\openapi.yaml:/openapi/openapi.yaml -e "SWAGGER_MOCK_SPECIFICATION_URL=/openapi/openapi.yaml" swaggermock/swagger-mock
```

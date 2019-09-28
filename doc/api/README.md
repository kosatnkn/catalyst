## Mock

### Perishable Container

```bash
# Need to be in `doc/api/`
cd doc/api

docker run -p 3000:8080 -v $PWD/openapi.yaml:/openapi/openapi.yaml -e "SWAGGER_MOCK_SPECIFICATION_URL=/openapi/openapi.yaml" --rm swaggermock/swagger-mock
```

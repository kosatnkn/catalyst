## Mock

### Using APISprout
[https://github.com/danielgtaylor/apisprout](https://github.com/danielgtaylor/apisprout)

**Linux**
```bash
docker run --name catalyst_mock -it --rm -p 3000:8000 -e "SPROUT_VALIDATE_REQUEST=1" -v $PWD/docs/api/openapi.yaml:/api.yaml danielgtaylor/apisprout /api.yaml
```

**Windows**
```bash
docker run --name catalyst_mock -it --rm -p 3000:8000 -e "SPROUT_VALIDATE_REQUEST=1" -v %cd%\\docs\\api\\openapi.yaml:/api.yaml danielgtaylor/apisprout /api.yaml
```

### Using OpenAPI-Mock
[https://github.com/muonsoft/openapi-mock](https://github.com/muonsoft/openapi-mock)

>**NOTE:** This is not working as expected at the moment.

**Linux**
```bash
docker run --name catalyst_mock -it --rm -p 3000:8080 -v $PWD/docs/api/openapi.yaml:/openapi/openapi.yaml -e "OPENAPI_MOCK_SPECIFICATION_URL=/openapi/openapi.yaml" muonsoft/openapi-mock
```

**Windows**
```bash
docker run --name catalyst_mock -it --rm -p 3000:8080 -v %cd%\\docs\\api\\openapi.yaml:/openapi/openapi.yaml -e "OPENAPI_MOCK_SPECIFICATION_URL=/openapi/openapi.yaml" muonsoft/openapi-mock
```

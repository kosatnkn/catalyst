# Mock

## Using Stoplight Prism
- [Documentation](https://meta.stoplight.io/docs/prism/README.md)
- [Installation](https://meta.stoplight.io/docs/prism/docs/getting-started/01-installation.md)
- [GitHub Repo](https://github.com/stoplightio/prism)

**Linux**
```bash
docker run --init --name catalyst_mock -it --rm -v $PWD/docs/api:/tmp -p 3000:4010 stoplight/prism mock -h 0.0.0.0 "/tmp/openapi.yaml"
```

**Windows**
```bash
docker run --init --name catalyst_mock -it --rm -v %cd%\\docs\\api:/tmp -p 3000:4010 stoplight/prism mock -h 0.0.0.0 "/tmp/openapi.yaml"
```
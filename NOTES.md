# Notes

## Creating a Git Tag without a GitHub Release
Since GitHub releases are automated you only need to create an annotated tag and push it upstream.
The automation will create a release draft for that tag.

```shell
# create an annotated tag
git tag -a v3.0.0 -m "Version 3.0.0"
# or
git tag --annotate v3.0.0 --message="Version 3.0.0"

# push tag upstream
git push origin v3.0.0
```

## Remove a Git Tag
If you need to remove a tag you will also have to remove the GitHub Release associated with that tag.

**Delete the remote tag**
```shell
git push origin --delete v3.0.0
```

**Delete the local tag**
```shell
git tag -d v3.0.0
```

**Tidy up**
```shell
git fetch
```

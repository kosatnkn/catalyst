# Notes

## Creating a Git Tag without a GitHub Release
Since GitHub releases are automated for Catalyst you only need to create an annotated tag from the `main` branch
and push it upstream.
The automation will create a release draft for that tag. Later you will have to review the draft and publish it.

```shell
# create an annotated tag
git tag -a v3.1.0 -m "Release 3.1.0"
# or, the long fom
git tag --annotate v3.1.0 --message="Release 3.1.0"

# push tag upstream
git push origin v3.1.0
```

## Remove a Git Tag
If you need to remove a tag you will also have to remove the GitHub Release associated with that tag.

**Delete the remote tag**
```shell
git push origin --delete v3.1.0
```

**Delete the local tag**
```shell
git tag -d v3.1.0
```

**Tidy up**
```shell
git fetch
```

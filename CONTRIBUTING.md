# Contributing

PR's welcome:

https://github.com/trycourier/courier-go

## Releasing New Versions

Make sure you have incremented the version string in `courier.go` to your new version string, hereafter referred to as `<VERSION>` and merged that commit. Then:

```bash
git tag -a v<VERSION> -m v<VERSION>
git push origin v<VERSION>
```
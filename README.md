# gonum - golang module example

Based on https://medium.com/rungo/anatomy-of-modules-in-go-c8274d215c16

## How it was done

This must be done **NOT** in your `GOPATH`.

```console
git clone git@github.com:elkaszcz/gonum.git
cd gonum
go mod init https://github.com/elkaszcz/gonum

```

Creating a new version
```console
git tag v1.0.0
git push --tags
```

## Run
```console
go run app.go
```
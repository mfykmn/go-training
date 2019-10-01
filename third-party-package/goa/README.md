# goa

https://github.com/goadesign/goa

```bash
mkdir -p calc/design
cd calc
go mod init calc
```

```bash
$ export GO111MODULE=on
$ go get -u goa.design/goa/v3
$ go install $GOPATH/pkg/mod/goa.design/goa/v3@v3.0.6/cmd/goa
$ goa gen calc/design
```
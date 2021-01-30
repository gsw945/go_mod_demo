# go_mod_demo

```bash
go mod init github.com/gsw945/go_mod_demo

go list -m all
go list -m -versions rsc.io/sampler
go list -m rsc.io/q...

go get rsc.io/sampler@v1.3

go doc rsc.io/quote/v3

go mod tidy
```

```bash
go test

go run .
# go run main.go

go build -o go_mod_demo.exe .
# go build -o go_mod_demo.exe main.go
```

## refer
- [The Go Blog - Using Go Modules](https://blog.golang.org/using-go-modules)
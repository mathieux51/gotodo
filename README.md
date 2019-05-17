# mem-crud

## Coverage report

```sh
go test -coverprofile temp/cover.out ./... && go tool cover -html=temp/cover.out
```

## Memory profile

```sh
go test -memprofilerate 1 -memprofile temp/mem.out ./model
# Get model from ...
go tool pprof -web temp/model.a temp/mem.out
```

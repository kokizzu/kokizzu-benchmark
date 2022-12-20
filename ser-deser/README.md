
# map/struct serialization deserialization benchmark

- encoding/json
- github.com/kokizzu/json5b/encoding/json5b
- github.com/goccy/go-json

## TL;DR

- `goccy/go-json` the fastest

## How to run?

```bash
go test -bench .
```

## How to add new test?

- create a function
- add on `TestVerify`
- create `Benchmark_X2X_RepoName_MethodName(b *testing.B)`
- add to `README.md`

```bash
go mod tidy
go test .
```
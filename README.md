Mathpractice is a web application that generates exercises for practicing basic
mathematics. The generated X and Y numbers are up to (not including) the integer
supplied as URL path.

Test

```
go test ./...
go test -fuzz=. -fuzztime=10s
```

Run locally

```
go run .
```

Build for a web server

```
GOOS=linux GOARCH=arm64 go build
```
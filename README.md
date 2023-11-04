Mathpractice is a web application that generates exercises for practicing basic
mathemathics. The result of the produced exercises is up to the integer supplied
as URL path.

Test

```
go test -fuzz=Fuzz -fuzztime 10s
```

Run locally

```
go run .
```

Build for a web server

```
GOOS=linux GOARCH=arm64 go build -o /tmp
```
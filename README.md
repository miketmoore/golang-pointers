# Playing with Pointers in Go

I was struggling to understand pointers in a Go program that was passing primitive values to and from various struct instances. I realized, while I thought I understood how pointers worked from very basic tutorials and examples, that I probably didn't!


## Run Tests

```
go test
```

## Run Code Coverage

```
go test -coverprofile=coverage.out && go tool cover -func=coverage.out
```

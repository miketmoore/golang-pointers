# Playing with Pointers in Go

I was struggling to understand pointers in a Go program that was passing primitive values to and from various struct instances. I realized, while I thought I understood how pointers worked from very basic tutorials and examples, that I probably didn't!

## Note

After some [comments](https://github.com/miketmoore/golang-pointers/commit/3776c65ded8c9c06b288e208d1753a6451acd985) from [@davecheney](https://twitter.com/davecheney), I now know I need to understand slices better. I updated the repo per his comments.

## Run Tests

```
go test
```

## Run Code Coverage

```
go test -coverprofile=coverage.out && go tool cover -func=coverage.out
```

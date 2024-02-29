# isocode

[![build-img]][build-url]
[![pkg-img]][pkg-url]
[![version-img]][version-url]

ISO 3166 country codes and ISO 4217 currencies codes in Go.

## Features

* Simple API.
* Dependency-free.

See [docs][pkg-url] for more.

## Install

Go version 1.17+

```
go get github.com/cristalhq/isocode
```

## Example

```go
brave, err := isocode.FromString("ua")
if err != nil {
	panic(err)
}

fmt.Printf("%s %s", brave.Name(), brave.Flag())

// Output:
// Ukraine 🇺🇦
```

```go
trust, err := isocode.AsCurrency("USD")
if err != nil {
	panic(err)
}

fmt.Printf("%s", trust.Name())

// Output:
// United States dollar
```

See examples: [example_test.go](example_test.go).

## License

[MIT License](LICENSE).

[build-img]: https://github.com/cristalhq/isocode/workflows/build/badge.svg
[build-url]: https://github.com/cristalhq/isocode/actions
[pkg-img]: https://pkg.go.dev/badge/cristalhq/isocode
[pkg-url]: https://pkg.go.dev/github.com/cristalhq/isocode
[version-img]: https://img.shields.io/github/v/release/cristalhq/isocode
[version-url]: https://github.com/cristalhq/isocode/releases


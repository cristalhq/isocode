# isocode

[![build-img]][build-url]
[![pkg-img]][pkg-url]
[![version-img]][version-url]

ISO 3166 country codes in Go.

## Features

* Simple API.

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

See examples: [example_test.go](example_test.go).

## License

[MIT License](LICENSE).

[build-img]: https://github.com/cristalhq/acmd/workflows/build/badge.svg
[build-url]: https://github.com/cristalhq/acmd/actions
[pkg-img]: https://pkg.go.dev/badge/cristalhq/acmd
[pkg-url]: https://pkg.go.dev/github.com/cristalhq/acmd
[version-img]: https://img.shields.io/github/v/release/cristalhq/acmd
[version-url]: https://github.com/cristalhq/acmd/releases


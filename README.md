![Droll](assets/logo.png)

## Supported Comics

| xkcd                                   | Indexed     | PHD Comics   |
|----------------------------------------|-------------|--------------|
| ![xkcd](https://xkcd.com/s/0b7742.png) | Coming Soon | Coming Soon  |

## Examples

### XKCD

**General**
```
query {
  xkcd {
    alt
    day
    image
    link
    month
    news
    num
    safeTitle
    title
    transcript
    year
  }
}
```

**With Limit and Offset**
```
query {
  xkcd(limit: 10, offset: 1) {
    alt
    day
    image
    link
    month
    news
    num
    safeTitle
    title
    transcript
    year
  }
}
```

## Getting Started

### Installation and Setup

- Download and Install [Golang](https://golang.org/dl/)

- Install dependencies

```bash
go get ./...
```

### Running

```bash
make run
```

### Testing

```bash
make test
```

### Code Coverage

```bash
make cover
```

## Contribute

Contributing Guide coming soon!

## Issues

Issues are managed via [GitHub Issues](https://github.com/prabhuomkar/droll-api/issues).

## License

This project is licensed under the Apache License. See the [LICENSE](LICENSE) file for details.

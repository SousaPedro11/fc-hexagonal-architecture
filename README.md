<h1 align="center">Welcome to fc-hexagonal-architecture  👋</h1>
<p>
</p>

> Repositório referente ao curso de Arquitetura hexagonal da Full Cycle

## Install

```sh
go mod download
```

## Usage

```sh
go run ./cmd/server/main.go
```

## Run tests

```sh
go test -coverprofile cover.out $(go list ./... | grep -v /application/mock | grep -v ./cmd/server) && go tool cover -html cover.out -o cover.html
```

## Author

👤 **Sousapedro11**

- Github: [@sousapedro11](https://github.com/sousapedro11)
- LinkedIn: [@sousapedro11](https://linkedin.com/in/sousapedro11)

## Show your support

Give a ⭐️ if this project helped you!

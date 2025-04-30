# xwag
[![Go Reference](https://pkg.go.dev/badge/github.com/zaidfadhil/xwag.svg)](https://pkg.go.dev/github.com/zaidfadhil/xwag)
[![Go Report Card](https://goreportcard.com/badge/github.com/zaidfadhil/xwag)](https://goreportcard.com/report/github.com/zaidfadhil/xwag)

**xwag** is a lightweight tool written in Go that serves Swagger API documentation using Swagger UI.

## Installation

### From Source

```bash
git clone https://github.com/zaidfadhil/xwag.git
go build -o xwag
```

## Usage
```bash
xwag -file <path_to_swagger_file>
```

- `-file`: Specifies the path to the Swagger file in YAML format. *(Required)*
- `-addr`: Specifies the server address. Defaults to `localhost` if not provided.
- `-port`: Specifies the server port. Defaults to `50166` if not provided.

Once the server is started, you can access the Swagger UI by opening a web browser at:

```
http://<server_address>:<server_port>.
```

## License
xwag is licensed under the [MIT License](https://github.com/zaidfadhil/xwag/blob/master/LICENSE).

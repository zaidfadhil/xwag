# Xwag
[![Go Reference](https://pkg.go.dev/badge/github.com/zaidfadhil/xwag.svg)](https://pkg.go.dev/github.com/zaidfadhil/xwag)
[![Go Report Card](https://goreportcard.com/badge/github.com/zaidfadhil/xwag)](https://goreportcard.com/report/github.com/zaidfadhil/xwag)

Xwag is a small tool built with Go that serves Swagger API documentation using the Swagger UI.

## Installation

```shell
git clone https://github.com/zaidfadhil/xwag.git
```

## Build
```shell
go build -o xwag
```

## Usage
```shell
xwag -file <path_to_swagger_file> [-addr <server_address>] [-port <server_port>]
```
- `-file`: Specifies the path to the Swagger file in YAML format. This is a required parameter.
- `-addr`: Specifies the server address. Defaults to `localhost` if not provided.
- `-port`: Specifies the server port. Defaults to `50166` if not provided.

Once the server is started, you can access the Swagger UI by opening a web browser and navigating to http://<server_address>:<server_port>.

## License
Xwag is licensed under the [MIT License](https://github.com/zaidfadhil/xwag/blob/master/LICENSE).

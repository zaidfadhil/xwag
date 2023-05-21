package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
)

var file string
var address string
var port int

func main() {

	flag.StringVar(&file, "file", "", "swagger file path.")
	flag.StringVar(&address, "addr", "localhost", "server address.")
	flag.IntVar(&port, "port", 50166, "server port.")
	flag.Parse()

	addr := fmt.Sprintf("%v:%v", address, port)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var scheme string
		if r.TLS != nil {
			scheme = "https://"
		} else {
			scheme = "http://"
		}
		hostURL := scheme + r.Host + r.URL.String() + "/swagger.yaml"
		html := fmt.Sprintf(htmlindex, hostURL)
		fmt.Fprintf(w, html)
	})

	http.HandleFunc("/swagger.yaml", func(w http.ResponseWriter, r *http.Request) {
		fileContent, err := ioutil.ReadFile(file)
		if err != nil {
			http.Error(w, "Failed to read Swagger file", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/yaml")
		w.Write(fileContent)
	})

	fmt.Println("Server started", addr)
	http.ListenAndServe(addr, nil)
}

var htmlindex = `
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8">
    <title>Swagger UI</title>
    <link rel="stylesheet" type="text/css" href="https://unpkg.com/swagger-ui-dist@4/swagger-ui.css" >
    <script src="https://unpkg.com/swagger-ui-dist@4/swagger-ui-bundle.js"> </script>
    <script src="https://unpkg.com/swagger-ui-dist@4/swagger-ui-standalone-preset.js"> </script>
		<style>
      html
      {
        box-sizing: border-box;
        overflow: -moz-scrollbars-vertical;
        overflow-y: scroll;
      }    
      *,
      *:before,
      *:after
      {
        box-sizing: inherit;
      }

      body
      {
        margin:0;
        background: #fafafa;
      }
      .errors-wrapper {
         display: none !IMPORTANT;
      }
    </style>
  </head>    
  <body>
    <div id="swagger-ui"></div>    
	<script>
    window.onload = function() {          
      const ui = SwaggerUIBundle({
        dom_id: "#swagger-ui",
        deepLinking: true,
        presets: [
          SwaggerUIBundle.presets.apis,
          SwaggerUIStandalonePreset
        ],
        plugins: [
          SwaggerUIBundle.plugins.DownloadUrl
        ],
        layout: "StandaloneLayout",
        validatorUrl: "https://validator.swagger.io/validator",
        urls: [
          {url: "%s", name: "swagger"},
        ],
        "urls.primaryName": "Patient"
      })
      window.ui = ui
    }
  </script>
  </body>
</html>
`

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	var file string
	var port int
	flag.StringVar(&file, "file", "", "swagger file path.")
	flag.IntVar(&port, "port", 50166, "swagger port.")
	flag.Parse()

	addr := fmt.Sprintf("%v:%v", "localhost", port)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		html := fmt.Sprintf(htmlindex, "http://"+addr+"/swagger.yaml")
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
<!-- HTML for static distribution bundle build -->
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8">
    <title>Swagger UI</title>
    <link rel="stylesheet" type="text/css" href="https://unpkg.com/swagger-ui-dist@3/swagger-ui.css" >
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
    <script src="https://unpkg.com/swagger-ui-dist@3/swagger-ui-bundle.js"> </script>
    <script src="https://unpkg.com/swagger-ui-dist@3/swagger-ui-standalone-preset.js"> </script>    <script>
    window.onload = function() {          
      // Begin Swagger UI call region
      const ui = SwaggerUIBundle({
        "dom_id": "#swagger-ui",
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

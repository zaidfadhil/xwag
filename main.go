package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

var (
	file    string
	address string
	port    int
)

func main() {
	parseFlags()

	http.HandleFunc("/", swaggerHandler)
	http.HandleFunc("/swagger.yaml", swaggerYAMLHandler)

	addr := fmt.Sprintf("%s:%d", address, port)
	log.Printf("Starting server at http://%s", addr)

	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

func parseFlags() {
	flag.StringVar(&file, "file", "", "Swagger YAML file path.")
	flag.StringVar(&address, "addr", "localhost", "Server address.")
	flag.IntVar(&port, "port", 50166, "Server port.")
	flag.Parse()

	if file == "" {
		log.Fatal("The -file flag is required and cannot be empty.")
	}
}

func swaggerHandler(w http.ResponseWriter, r *http.Request) {
	html := fmt.Sprintf(htmlTemplate, "swagger.yaml")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, err := w.Write([]byte(html))
	if err != nil {
		log.Printf("Failed to write response for /: %v", err)
	}
}

func swaggerYAMLHandler(w http.ResponseWriter, r *http.Request) {
	fileContent, err := os.ReadFile(file)
	if err != nil {
		log.Printf("Failed to read Swagger file (%s): %v", file, err)
		http.Error(w, "Failed to read Swagger file", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/yaml")
	_, err = w.Write(fileContent)
	if err != nil {
		log.Printf("Failed to write Swagger YAML response: %v", err)
	}
}

var htmlTemplate = `
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8">
    <title>Swagger UI</title>
    <link rel="stylesheet" type="text/css" href="https://unpkg.com/swagger-ui-dist@5/swagger-ui.css">
    <script src="https://unpkg.com/swagger-ui-dist@5/swagger-ui-bundle.js"></script>
    <script src="https://unpkg.com/swagger-ui-dist@5/swagger-ui-standalone-preset.js"></script>
    <style>
      html {
        box-sizing: border-box;
        overflow-y: scroll;
      }
      *, *:before, *:after {
        box-sizing: inherit;
      }
      body {
        margin: 0;
        background: #fafafa;
      }
      .errors-wrapper {
        display: none !important;
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
          urls: [
            {url: window.location.origin + "/%s", name: "swagger"}
          ],
          "urls.primaryName": "swagger",
          filter: true
        });
        window.ui = ui;
      }
    </script>
  </body>
</html>
`

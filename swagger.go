package swaggerui

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
)

// Handler returns a handler using the latest SwaggerUI that loads /swagger.json
func Handler() http.Handler {
	return CustomHandler("API Documentation", "/swagger.json", "latest")
}

// CustomHandler returns a handler to render the Swagger UI but allows for custom parameters to be passed
func CustomHandler(title string, specPath string, swaggerUIVersion string) http.Handler {
	opts := map[string]string{
		"Title":           title,
		"SwaggerUIScript": fmt.Sprintf("//unpkg.com/swagger-ui-dist@%s/swagger-ui-bundle.js", swaggerUIVersion),
		"SwaggerUIStyles": fmt.Sprintf("//unpkg.com/swagger-ui-dist@%s/swagger-ui.css", swaggerUIVersion),
		"SpecURL":         specPath,
	}

	tmpl := template.Must(template.New("swaggerUI").Parse(swaggerUITemplate))

	buf := bytes.NewBuffer(nil)
	_ = tmpl.Execute(buf, opts)

	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write(buf.Bytes())
	}

	return http.HandlerFunc(fn)
}

const (
	swaggerUITemplate = `<!DOCTYPE html>
<html>
<head>
<meta charset="UTF-8">
<title>{{ .Title }}</title>
<script src="{{ .SwaggerUIScript }}"></script>
<link rel="stylesheet" href="{{ .SwaggerUIStyles }}" />
</head>

<body>
    <div id="swagger-ui" />
    <script>
        var ui = SwaggerUIBundle({
            url: "{{ .SpecURL }}",
            dom_id: '#swagger-ui',
            presets: [
                SwaggerUIBundle.presets.apis,
                SwaggerUIBundle.SwaggerUIStandalonePreset
            ],
        })
    </script>
</body>
</html>
`
)

package docs

import (
	"net/http"

	"github.com/mvrilo/go-redoc"
)

func GetHandler() http.HandlerFunc {
	filePath := "./static/openapi.yaml"
	specPath := "static/openapi.yaml"

	doc := redoc.Redoc{
		Title:       "open api spec",
		Description: "open api spec",
		SpecFile:    filePath,
		SpecPath:    specPath,
	}

	return doc.Handler()
}

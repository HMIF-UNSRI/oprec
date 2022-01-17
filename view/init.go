package view

import (
	"embed"
	"html/template"
)

//go:embed *.gohtml
var embedTemplates embed.FS

var Templates = template.Must(template.ParseFS(embedTemplates, "./*.gohtml"))

package main

import "github.com/pwilliams-ck/snip-doc/internal/models"

type templateData struct {
	Snippet  models.Snippet
	Snippets []models.Snippet
}

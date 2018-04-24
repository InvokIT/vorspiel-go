package main

import (
  "html/template"
)

var Templates map[string]*template.Template

func init() {
  Templates = map[string]*template.Template{
    "authSuccess": loadTemplate("templates/auth-success.html"),
  }
}

func loadTemplate(path string) *template.Template {
  tmpl, err := template.ParseFiles(path)
  if err != nil {
    logger.Panic(err)
  }

  return tmpl
}

package ui

import "embed"

const (
  StripPath = "shoppinglist-ui/dist"
)

//go:embed shoppinglist-ui/dist
var Content embed.FS

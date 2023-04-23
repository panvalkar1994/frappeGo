package main

import (
	doc "github.com/panvalkar1994/frappeGo/document"
)

func main() {
	filePath := "./testdata/doctype_example.json"
	document, err := doc.ParseDocument(filePath)
	if err != nil {
		return
	}
	document.GetModified()
	document.CreateDocumentModel("")
}

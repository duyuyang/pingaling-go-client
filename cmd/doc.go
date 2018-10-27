package cmd

import (
	"log"

	"github.com/spf13/cobra/doc"
)

func MakeDoc() {

	err := doc.GenMarkdownTree(rootCmd, "./doc")
	if err != nil {
		log.Fatalf("failed to make doc %v", err)
	}

}

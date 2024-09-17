package main

import (
	"log"

	"github.com/spark-cli/spark/cmd"
	"github.com/spf13/cobra/doc"
)

func main() {
	err := doc.GenMarkdownTree(cmd.Cmd, "./docs/")
	if err != nil {
		log.Fatal(err)
	}
}

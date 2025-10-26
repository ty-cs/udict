package cmd

import (
	"fmt"
	"log"

	"github.com/ty-cs/udict/internal/api"
	"github.com/ty-cs/udict/internal/formatter"
)

func getAndDisplayDefinition(endpoint string, limit int, notFoundMsg string) {
	definition, err := api.FetchAndDecode(endpoint)
	if err != nil {
		log.Fatal(err)
	}

	if len(definition.List) == 0 {
		fmt.Println(notFoundMsg)
		return
	}

	fmt.Println(formatter.StyleDefinition(definition, limit))
}

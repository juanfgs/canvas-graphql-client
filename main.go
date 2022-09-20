package main

import (
	"context"
	"fmt"
	"github.com/juanfgs/canvas-client/canvas"
	"github.com/shurcooL/graphql"
	"log"
	"os"
)

func main() {
	url := os.Getenv("CANVAS_GRAPHQL_URL")
	client := graphql.NewClient(url, nil)
	var query struct {
		Canvases []canvas.Canvas
	}
	err := client.Query(context.Background(), &query, nil)
	if err != nil {
		log.Fatal(err)
	} else {
		for _, canvas := range query.Canvases {
			fmt.Println(fmt.Sprintf("Displaying canvas : %s", canvas.Name))
			canvas.Render()

		}

	}
}

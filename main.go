package main

import (
	"encoding/json"
	"fmt"
	"graphql-with-go/models"
	"log"

	"github.com/graphql-go/graphql"
)

func main() {

	var commentType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Comment",
			Fields: graphql.Fields{
				"body": &graphql.Field{
					Type: graphql.String,
				},
			},
		},
	)

	var authorType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Author",
			Fields: graphql.Fields{
				"Name": &graphql.Field{
					Type: graphql.String,
				},
				"Tutorials": &graphql.Field{
					Type: graphql.NewList(graphql.Int),
				},
			},
		},
	)

	var tutorialType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Tutorial",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"title": &graphql.Field{
					Type: graphql.String,
				},
				"author": &graphql.Field{
					Type: authorType,
				},
				"comments": &graphql.Field{
					Type: graphql.NewList(commentType),
				},
			},
		},
	)

	tutorials := populate()
	fields := graphql.Fields{
		"tutorial": &graphql.Field{
			Type:        tutorialType,
			Description: "Get tutorial by id",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {

				id, ok := p.Args["id"].(int)
				if ok {
					for _, tutorial := range tutorials {
						if tutorial.Id == id {
							return tutorial, nil
						}
					}
				}
				return nil, nil
			},
		},
		"list": &graphql.Field{
			Type:        graphql.NewList(tutorialType),
			Description: "Get all tutorials",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return tutorials, nil
			},
		},
	}

	rootQuery := graphql.ObjectConfig{
		Name:   "RootQuery",
		Fields: fields,
	}

	schemaConfig := graphql.SchemaConfig{
		Query: graphql.NewObject(rootQuery),
	}

	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("Failed to create a new graphql schema %v", err)
	}

	query := `
	{
		tutorial(id:1) {
			title
			author {
				Name
				Tutorials
			}
		}
	}
	`

	params := graphql.Params{
		Schema:        schema,
		RequestString: query,
	}
	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		log.Fatalf("Failed to execute graphql operation, errors: %+v", r.Errors)
	}
	rJson, _ := json.Marshal(r)
	fmt.Printf("%s \n ", rJson)

}

func populate() []models.Tutorial {
	author := &models.Author{
		Name:      "Author 1",
		Tutorials: []int{1},
	}
	tutorial := models.Tutorial{
		Id:     1,
		Title:  "Go Graphql",
		Author: *author,
		Comments: []models.Comment{
			{
				Body: "First comment",
			},
		},
	}

	var tutorials []models.Tutorial
	tutorials = append(tutorials, tutorial)

	return tutorials
}

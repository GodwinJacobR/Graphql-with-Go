package main

import "graphql-with-go/models"

func Populate() []models.Tutorial {
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

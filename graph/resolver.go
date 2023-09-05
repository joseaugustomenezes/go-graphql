package graph

import "github.com/joseaugustomenezes/go-graphql/internal/database"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	CategoryDb *database.CategoryDb
	CourseDb *database.CourseDb
}

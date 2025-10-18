package main

import (
	"context"
	"database/sql"
	"fmt"

	// "github.com/google/uuid"
	"github.com/sergioc0sta/sqlc/internal/db"

	_ "github.com/go-sql-driver/mysql"

)

func main() {
	ctx := context.Background()
	dbConn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")

	if err != nil {
		panic(err)
	}

	defer dbConn.Close()

	queries := db.New(dbConn)

	// err = queries.CreateCategorie(ctx, db.CreateCategorieParams{
	// 	ID:          uuid.New().String(),
	// 	Name:        "sandes mistas",
	// 	Description: sql.NullString{String: "nem sei", Valid: true},
	// })
	//
	// if err != nil {
	// 	panic(err)
	// }
	//

	categories, err := queries.ListCategories(ctx)

	if err != nil {
		panic(err)
	}

	for _, category := range categories {

		fmt.Printf("o meu id 'e: %s com o nome: %s e a Description: %s", category.ID, category.Name, category.Description.String)
	}

}

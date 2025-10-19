package main

import (
	"context"
	"database/sql"
	"fmt"

	// "github.com/google/uuid"
	"github.com/google/uuid"
	"github.com/sergioc0sta/sqlc/internal/db"

	_ "github.com/go-sql-driver/mysql"
)

type CoursesDB struct {
	dbConn *sql.DB
	*db.Queries
}

type CourseParams struct {
	ID          string
	Name        string
	Description sql.NullString
	Price       float64
}

type CategoryParams struct {
	ID          string
	Name        string
	Description sql.NullString
}

func NewCourseDB(dbConn *sql.DB) *CoursesDB {
	return &CoursesDB{
		dbConn:  dbConn,
		Queries: db.New(dbConn),
	}
}

func (c *CoursesDB) callTx(ctx context.Context, fn func(*db.Queries) error) error {
	tx, err := c.dbConn.BeginTx(ctx, nil)

	if err != nil {
		return err
	}

	querie := db.New(tx)

	err = fn(querie)

	if err != nil {
		if rb := tx.Rollback(); rb != nil {
			return err
		}
		return err
	}

	return tx.Commit()
}

func (c *CoursesDB) CreateCoursesAndCategory(ctx context.Context, argsCategory CategoryParams, argsCourse CourseParams) error {

	err := c.callTx(ctx, func(q *db.Queries) error {
		var err error
		err = q.CreateCategorie(ctx, db.CreateCategorieParams{
			Name:        argsCategory.Name,
			Description: argsCategory.Description,
			ID:          argsCategory.ID,
		})

		if err != nil {
			return err
		}

		err = q.CreateCourses(ctx, db.CreateCoursesParams{
			Name:        argsCourse.Name,
			Description: argsCourse.Description,
			ID:          argsCourse.ID,
			CategoryID:  argsCategory.ID,
			Price:       argsCourse.Price,
		})

		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}
	return nil
}

func main() {
	ctx := context.Background()
	dbConn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")

	if err != nil {
		panic(err)
	}

	defer dbConn.Close()

	queries := db.New(dbConn)

	categoryArgs := CategoryParams{
		ID:          uuid.New().String(),
		Name:        "cenas",
		Description: sql.NullString{String: "panas", Valid: true},
	}

	courseArgs := CourseParams{
		ID:          uuid.New().String(),
		Name:        "cenas tipo cenas",
		Description: sql.NullString{String: "panas type panas", Valid: true},
		Price:       12.01,
	}

	courseDB := NewCourseDB(dbConn)

	err = courseDB.CreateCoursesAndCategory(ctx, categoryArgs, courseArgs)

	if err != nil {
		panic(err)
	}

	courses, err := queries.ListCourses(ctx)
	if err != nil {
		panic(err)
	}
	for _, course := range courses {
		fmt.Printf("Category: %s, Course ID: %s, Course Name: %s, Course Description: %s, Course Price: %f",
			course.CategoryName, course.ID, course.Name, course.Description.String, course.Price)
	}
}

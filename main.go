package main

import (
	"database/sql"

	"gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"
)

type Student struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Class string `json:"class"`
}

func GetStudentHandler (ctx *gofr.Context) (interface{}, error) {
	name := ctx.PathParam("name")

	if name == "" {
		return nil, errors.MissingParam{Param: []string{"name"}}
	}

	resp, err := GetByName(ctx, name)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func GetByName(ctx *gofr.Context, name string) (Student, error) {
	var resp Student

	err := ctx.DB().QueryRowContext(ctx, "SELECT * FROM students WHERE name=?", name).Scan(&resp.ID, &resp.Name, &resp.Class)

	switch err {
	case sql.ErrNoRows:
		return resp, errors.EntityNotFound{Entity: "student", ID: name}
	case nil:
		return resp, nil
	default:
		return resp, err
	}
}

func PostStudentHandler (ctx *gofr.Context) (interface{}, error) {
	var student Student

	// ctx.Bind() binds the incoming data from the HTTP request to a provided interface (i).
	if err := ctx.Bind(&student); err != nil {
		ctx.Logger.Errorf("error in binding: %v", err)
		return nil, errors.InvalidParam{Param: []string{"body"}}
	}

	resp, err := PostStudent(ctx, student)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func PostStudent(ctx *gofr.Context, student Student) (Student, error) {
	var resp Student
	
	ctx.DB().QueryRowContext(ctx, "INSERT INTO students (name, class) VALUES(?,?)", student.Name, student.Class)

	err := ctx.DB().QueryRowContext(ctx, "SELECT * FROM students WHERE name=?", student.Name).Scan(&resp.ID, &resp.Name, &resp.Class)

	if err != nil {
		return Student{}, errors.DB{Err: err}
	}

	return resp, nil
}

func main() {
    app := gofr.New()

	app.GET("/students/{name}", GetStudentHandler)
	app.POST("/students", PostStudentHandler)

	// starting the server on a custom port
	app.Server.HTTP.Port = 9092
    app.Start()
}
package main

import (
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
	// case sql.ErrNoRows:
	// 	return "", errors.EntityNotFound{Entity: "student", ID: name}
	case nil:
		return resp, nil
	default:
		return resp, err
	}
}

func main() {
    app := gofr.New()

	app.GET("/students/{name}", GetStudentHandler)

	// starting the server on a custom port
	app.Server.HTTP.Port = 9092
    app.Start()
}
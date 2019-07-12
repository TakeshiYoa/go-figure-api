package controllers

import (
	"net/http"
	"github.com/labstack/echo/v4"

	"api/app/submission/store"
)

type Submission struct {
	Id int `json:"id"`
}

func Submit(c echo.Context) error {
	//grab input from context

	//put values into storage

	//respond with id of new submission

	submission := &Submission{1}

	return c.JSON(http.StatusOK, submission)
}

func GetSubmit(c echo.Context) error {
	//grab input from context

	//put values into storage

	//respond with id of new submission

	store.New().Submit()

	submission := &Submission{1}

	return c.JSON(http.StatusOK, submission)
}
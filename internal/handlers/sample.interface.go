package handlers

import "github.com/labstack/echo"

type ISampleHandler interface {
	GetSample(c echo.Context) error
	CreateSample(c echo.Context) error
	UpdateSample(c echo.Context) error
	DeleteSample(c echo.Context) error
}

package main

import (
	
	"net/http"

	"github.com/labstack/echo/v4"
)

type Request struct {
	Num1      float64 `json:"num1"`
	Num2      float64 `json:"num2"`
	Operation string  `json:"operation"`
}

type Response struct {
	Result float64 `json:"result"`
	Error  string  `json:"error"`
}

func calculate(ctx echo.Context) error {

	var req Request
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, Response{Error: "invalid value"})
	}

	// Check for valid operation input
	if req.Operation == "" {
		return ctx.JSON(http.StatusBadRequest, Response{Error: "enter a valid operation"})
	}

	var result float64
	switch req.Operation {
	case "+":
		result = req.Num1 + req.Num2

	case "-":
		result = req.Num1 - req.Num2

	case "*":
		result = req.Num1 * req.Num2

	case "/":
		if req.Num2 == 0 {
			return ctx.JSON(http.StatusBadRequest, Response{Error: "cannot divide by zero"})
		}
		result = req.Num1 / req.Num2

	default:
		return ctx.JSON(http.StatusBadRequest, Response{Error: "invalid operation"})
	}

	return ctx.JSON(http.StatusOK, Response{Result: result})
}

func main() {
	e := echo.New()
	e.POST("/calculate", calculate) // Changed endpoint name to match functionality
	e.Start(":8080")
}

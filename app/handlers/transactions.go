package handlers

import (
	"github.com/chadiso/golang_tryout/app/repositories"
	"github.com/labstack/echo"
	"net/http"
)

func (hc Context) ListTransactions(c echo.Context) error {
	transactions := repositories.GetTransactions(hc.DB)

	res := echo.Map{"transactions": transactions}

	return c.JSONPretty(http.StatusOK, res, "  ")
}

func (hc Context) GetTransaction(c echo.Context) error {
	transaction := repositories.GetTransaction(hc.DB, c.Param("id"))
	res := echo.Map{"transaction": transaction}

	return c.JSONPretty(http.StatusOK, res, "  ")
}

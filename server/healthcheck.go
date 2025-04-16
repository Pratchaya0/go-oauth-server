package server

import (
	"net/http"

	"github.com/Pratchaya0/go-oauth-server/pkg/response"
	"github.com/labstack/echo/v4"
)

type healthCheck struct {
	App    string `json:"app"`
	Code   int    `json:"code"`
	Status string `json:"status"`
}

func (s *server) healthCheckService(c echo.Context) error {
	return response.SuccessReponse(c, &healthCheck{
		App:    s.cfg.App.Name,
		Code:   http.StatusOK,
		Status: "OK",
	})
}

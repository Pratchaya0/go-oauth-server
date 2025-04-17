package userhandler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/Pratchaya0/go-oauth-server/config"
	"github.com/Pratchaya0/go-oauth-server/modules/user"
	userusecase "github.com/Pratchaya0/go-oauth-server/modules/user/userUsecase"
	"github.com/Pratchaya0/go-oauth-server/pkg/request"
	"github.com/Pratchaya0/go-oauth-server/pkg/response"
	"github.com/labstack/echo/v4"
)

type (
	IUserHttpHandler interface {
		CreateOneUser(c echo.Context) error
		UpdateOneUserDetails(c echo.Context) error
	}

	userHttpHandler struct {
		cfg         *config.Config
		userUsecase userusecase.IUserUsecase
	}
)

func NewUserHandler(cfg *config.Config, userUsscase userusecase.IUserUsecase) IUserHttpHandler {
	return &userHttpHandler{cfg, userUsscase}
}

func (h *userHttpHandler) CreateOneUser(c echo.Context) error {

	ctx := context.Background()

	wrapper := request.ContextWrapper(c)

	var request user.CreateUserRequestDTO

	if err := wrapper.Bind(request); err != nil {
		return response.ErrorResponse[map[string]uint](c, http.StatusBadRequest, err.Error())
	}

	userID, err := h.userUsecase.CreateOneUser(ctx, &request)

	if err != nil {
		return response.ErrorResponse[map[string]uint](c, http.StatusBadRequest, err.Error())
	}

	return response.SuccessReponse(c, map[string]uint{
		"user_id": *userID,
	})

}

func (h *userHttpHandler) UpdateOneUserDetails(c echo.Context) error {
	ctx := context.Background()

	wrapper := request.ContextWrapper(c)

	var request user.UpdateUserRequestDTO

	idParam := c.Param("userId")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return response.ErrorResponse[map[string]uint](c, http.StatusBadRequest, err.Error())
	}

	if err := wrapper.Bind(request); err != nil {
		return response.ErrorResponse[map[string]uint](c, http.StatusBadRequest, err.Error())
	}

	result, err := h.userUsecase.UpdateOneUserDetails(ctx, uint(id), &request)

	if err != nil {
		return response.ErrorResponse[map[string]uint](c, http.StatusBadRequest, err.Error())
	}

	return response.SuccessReponse(c, map[string]uint{
		"user_id": *result,
	})
}

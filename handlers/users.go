package users

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/zaru/go-api-documents-demo/models"
)

type resultLists struct {
	Users []user.User `json:"users"`
}

type handler struct {
	UserModel user.UserModelImpl
}

func NewHandler(u user.UserModelImpl) *handler {
	return &handler{u}
}

func (h *handler) GetIndex(c echo.Context) error {
	lists := h.UserModel.FindAll()
	u := &resultLists{
		Users: lists,
	}
	return c.JSON(http.StatusOK, u)
}

package users

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/zaru/go-api-documents-demo/models"
)

type (
	UsersModelStub struct{}
)

func (u *UsersModelStub) FindByID(id string) user.User {
	return user.User{
		ID:   1,
		Name: "foo",
	}
}
func (u *UsersModelStub) FindAll() []user.User {
	users := []user.User{}
	users = append(users, user.User{
		ID:   100,
		Name: "foo",
		Rank: 1,
	})
	return users
}

func TestGetIndex(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/users")

	u := &UsersModelStub{}
	h := NewHandler(u)

	var userJSON = `{"users":[{"id":100,"name":"foo","rank":1}]}`

	if assert.NoError(t, h.GetIndex(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, userJSON, rec.Body.String())
	}
}

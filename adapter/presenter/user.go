package presenter

import (
	"ca/v2/entity"
	"ca/v2/usecase/port"
	"fmt"
	"net/http"
)

type User struct {
	w http.ResponseWriter
}

// Render implements port.UserOutputPort
func (u *User) Render(user *entity.User) {
	u.w.WriteHeader(http.StatusOK)
	fmt.Println(u.w, user.Name)
}

// RenderError implements port.UserOutputPort
func (u *User) RenderError(err error) {
	u.w.WriteHeader(http.StatusInternalServerError)
	fmt.Println(u.w, err)
}

func NewUserOutputPort(w http.ResponseWriter) port.UserOutputPort {
	return &User{
		w: w,
	}
}

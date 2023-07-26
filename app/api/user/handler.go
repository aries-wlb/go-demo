package user

import (
	"net/http"

	"github.com/google/wire"
	"github.com/uptrace/bunrouter"
	"patrick.com/abroad/app/domain"
)

var userSet = wire.NewSet(wire.Struct(new(UserHandler), "*"))

type UserHandler struct {
	UserDomain *domain.UserDomain
}

func (u *UserHandler) createUser(w http.ResponseWriter, req bunrouter.Request) error {
	return nil
}

func (u *UserHandler) getUser(w http.ResponseWriter, req bunrouter.Request) error {
	id, err := req.Params().Int32("id")
	if err != nil {
		return err
	}
	user := u.UserDomain.Get(id)
	return bunrouter.JSON(w, bunrouter.H{
		"user":  user,
		"route": req.Route(),
	})
}

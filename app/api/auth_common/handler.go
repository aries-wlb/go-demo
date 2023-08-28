package auth_common

import (
	"net/http"

	"github.com/google/wire"
	"github.com/uptrace/bunrouter"
	userDomain "patrick.com/abroad/app/domain/user"
	"patrick.com/abroad/app/middleware/jwt"
	"patrick.com/abroad/app/utils"
)

var AuthCommonSet = wire.NewSet(wire.Struct(new(AuthCommonHandler), "*"))

func (c *AuthCommonHandler) match(w http.ResponseWriter, req bunrouter.Request) error {
	ctx := req.Context()
	userInfo := ctx.Value(jwt.UserCtxKey{}).(*userDomain.Claims)
	var param MatchReq
	err := utils.DecodeJSONBody(w, req, &param)
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}
	return utils.GenSuccessResp(w, map[string]interface{}{"param": param, "user_info": userInfo})
}

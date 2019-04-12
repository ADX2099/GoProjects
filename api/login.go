package api

import (
	"lab01/api/classes"
	"lab01/api/security"
	"lab01/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Access(c *gin.Context, user classes.UserLogin) {

	var request_login classes.RequestLogin

	if user.User == "admin@admin.com" && user.Password == "admin" {

		tk := security.GetToken(user.User)

		if tk.Code == config.SUSSES_CREATE_TOKEN {
			c.Header("token", tk.Token)
			//c.Status(http.StatusOK)

			r := classes.RequestCode{Code: config.OK_PROCESS, Message: "OK process"}
			request_login.Response = r
			c.JSON(http.StatusOK, request_login)
		} else {
			r := classes.RequestCode{Code: config.ERROR_IN_SECURITY_PROCESS, Message: "Security error"}
			request_login.Response = r
			c.JSON(http.StatusUnauthorized, request_login)
		}

	} else {
		r := classes.RequestCode{Code: config.USER_NOT_AUTHORIZE, Message: "User  not authorized"}
		request_login.Response = r
		c.JSON(http.StatusUnauthorized, request_login)
	}

}

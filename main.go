package main

import (
	"fmt"
	"lab01/api"
	"lab01/api/classes"
	"lab01/api/security"
	"lab01/api/val"
	"lab01/config"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	consulapi "github.com/hashicorp/consul/api"
)

/*
*Go-micro
 */
func main() {
	registerServiceWithConsul()

	//LEVANTAMOS EL SERVIDOR
	//CREO UNA VARIABLE DEL TIPO GIN
	engine := gin.Default()
	val.IniValidation()
	//LE ASIGNO EL PUERTO

	//SERVICIO GET
	engine.GET("/info", func(c *gin.Context) {
		c.String(http.StatusOK, "GXSERVER Response")
	})
	//SERVICIO POST LOGIN
	engine.POST("/login", func(c *gin.Context) {
		var user classes.UserLogin
		var request_login classes.RequestLogin
		var err error

		if err = c.BindJSON(&user); err == nil {
			res_val := val.ValidationLogin(user)

			if res_val {
				api.Access(c, user)
			} else {
				r := classes.RequestCode{Code: config.ERROR_IN_DATA, Message: config.MSG_VERIFY_DATA}
				request_login.Response = r
				c.JSON(http.StatusBadRequest, request_login)
			}

			/*if user.User == "admin@admin" && user.Password == "admin" {
				r := classes.RequestCode{Code: config.OK_PROCESS, Message: "OK process"}
				request_login.Response = r
				c.JSON(http.StatusOK, request_login)

			} else {
				r := classes.RequestCode{Code: config.ERROR_IN_SECURITY_PROCESS, Message: "Security error"}
				request_login.Response = r
				c.JSON(http.StatusUnauthorized, request_login)

			}*/
		} else {
			r := classes.RequestCode{Code: config.ERROR_IN_JSON_STRING, Message: config.MSG_ERROR_JSON_STRING}
			request_login.Response = r
			c.JSON(http.StatusBadRequest, request_login)
		}

		c.String(http.StatusUnauthorized, "SIN ACCESO")
	})

	engine.GET("/books", func(c *gin.Context) {
		tk := c.GetHeader("Authorization")
		if tk == "" {

		} else {
			if security.ValidateToken(tk).Code == config.OK_PROCESS {
				api.GetBooks(c)
			} else {
				var request_sec classes.RequestLogin
				r := classes.RequestCode{Code: config.USER_NOT_AUTHORIZE, Message: config.MSG_USER_NOT_AUTHORIZE}
				request_sec.Response = r
				c.JSON(http.StatusUnauthorized, request_sec)
			}

		}

	})
	// wget  -qSO - http://localhost:9090/books/1234567890
	engine.GET("/books/:isbn", func(c *gin.Context) {
		tk := c.GetHeader("Authorization")
		if security.ValidateToken(tk).Code == config.OK_PROCESS {
			isbn := c.Params.ByName("isbn")
			api.GetBook(c, isbn)
		} else {
			var request_sec classes.RequestLogin
			r := classes.RequestCode{Code: config.USER_NOT_AUTHORIZE, Message: config.MSG_USER_NOT_AUTHORIZE}
			request_sec.Response = r
			c.JSON(http.StatusUnauthorized, request_sec)
		}
	})

	println("Start service host")
	engine.Run(":9010")

}

func registerServiceWithConsul() {
	config_default := consulapi.DefaultConfig()
	consul, err := consulapi.NewClient(config_default)
	if err != nil {
		fmt.Println(err)
	}

	var registration = new(consulapi.AgentServiceRegistration)

	registration.ID = config.CONSUL_REGISTRATION_ID
	registration.Name = config.CONSUL_REGISTRATION_NAME

	address := hostname()
	registration.Address = address
	port, _ := strconv.Atoi(port()[1:len(port())])
	registration.Port = port

	registration.Check = new(consulapi.AgentServiceCheck)
	registration.Check.HTTP = fmt.Sprintf(config.CONSUL_URL_CHECK, address, port)
	registration.Check.Interval = config.CONSUL_INTERVAL
	registration.Check.Timeout = config.CONSUL_TIMEOUT

	consul.Agent().ServiceRegister(registration)
}

func hostname() string {
	hostname, _ := os.Hostname()
	return hostname
}

func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = config.SYSTEM_PORT_DEFAULT
	}
	return ":" + port
}

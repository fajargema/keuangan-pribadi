package route

import (
	"keuangan-pribadi/controllers"
	m "keuangan-pribadi/middleware"
	"keuangan-pribadi/utils"

	"github.com/labstack/echo/v4"
	mid "github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	// create a new echo instance
	e := echo.New()

	loggerConfig := m.LoggerConfig{
		Format: "[${time_rfc3339}] ${status} ${method} ${host} ${path} ${latency_human}" + "\n",
	}
	loggerMiddleware := loggerConfig.Init()
	e.Use(loggerMiddleware)

	
	v1 := e.Group("/api/v1")
	eJwt := v1.Group("")
	eJwt.Use(mid.JWT([]byte(utils.GetConfig("JWT_SECRET_KEY"))))

	// Route / to handler function
	user := controllers.InitUserController()
	v1.POST("/users/login", user.Login)
	v1.POST("/users/register", user.Register)
	eJwt.GET("/users/:email", user.GetByEmail)
	eJwt.PUT("/users", user.Update)

	return e
}
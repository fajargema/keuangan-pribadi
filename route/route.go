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

	category := controllers.InitCategoryController()
	eJwt.GET("/categories", category.GetAll)
	eJwt.GET("/categories/:id", category.GetByID)
	eJwt.POST("/categories", category.Create)
	eJwt.PUT("/categories/:id", category.Update)
	eJwt.DELETE("/categories/:id", category.Delete)

	finance := controllers.InitFinanceController()
	eJwt.GET("/finances", finance.GetAll)
	eJwt.GET("/finances/:id", finance.GetByID)
	eJwt.GET("/finances/search", finance.Search)
	eJwt.POST("/finances", finance.Create)
	eJwt.PUT("/finances/:id", finance.Update)
	eJwt.DELETE("/finances/:id", finance.Delete)

	return e
}
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

	coffe := controllers.GetCoffeePrice
	v1.GET("/coffee", coffe)

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

	saving := controllers.InitSavingController()
	eJwt.GET("/savings", saving.GetAll)
	eJwt.GET("/savings/:id", saving.GetByID)
	eJwt.POST("/savings", saving.Create)
	eJwt.PUT("/savings/:id", saving.Update)
	eJwt.DELETE("/savings/:id", saving.Delete)

	detailSaving := controllers.InitDetailSavingController()
	eJwt.GET("/detail-savings", detailSaving.GetAll)
	eJwt.GET("/detail-savings/:id", detailSaving.GetByID)
	eJwt.POST("/detail-savings", detailSaving.Create)
	eJwt.PUT("/detail-savings/:id", detailSaving.Update)
	eJwt.DELETE("/detail-savings/:id", detailSaving.Delete)

	return e
}
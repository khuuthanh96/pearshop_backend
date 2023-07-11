package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	appConfig "pearshop_backend/app/config"
	_ "pearshop_backend/app/delivery/http/docs" // Import the docs package for swag-go
	"pearshop_backend/app/delivery/http/handler"
	"pearshop_backend/app/registry"
	"pearshop_backend/pkg/hashid"
)

// Handler define mapping routes
// @title pearshop backend
// @version 1.0
// @description Pearshop backend api docs
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath /api/v1
func Handler(cfg *appConfig.Config) *gin.Engine {
	router := gin.New()

	if cfg.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router.GET("/", root)
	router.GET("/api/healthz", health)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	productHandler := handler.NewProductHandler(
		registry.InjectedProductFindUsecase(),
		registry.InjectedProductUpdateUsecase(),
		registry.InjectedProductCreateUsecase(),
		hashid.GetIDHasher(),
	)

	// handlers
	v1Api := router.Group("/api/v1")

	v1Api.GET("/products", productHandler.Find)

	v1Api.POST("/products", productHandler.Create)

	v1Api.PUT("/products/:id", productHandler.Update)

	return router
}

func root(ctx *gin.Context) {
	svcInfo := struct {
		Version string `json:"version,omitempty"`
		Name    string `json:"name,omitempty"`
	}{
		Version: "v1",
		Name:    "PEARSHOP API",
	}

	ctx.JSON(http.StatusOK, svcInfo)
}

func health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
	})
}

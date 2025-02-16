// Package v1 implements routing paths. Each services in own file.
package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	// Swagger docs.
	_ "github.com/ozlemugur/go-clean-arch-tt/docs"
	"github.com/ozlemugur/go-clean-arch-tt/internal/usecase"
	"github.com/ozlemugur/go-clean-arch-tt/pkg/logger"
)

// NewRouter -.
// Swagger spec:
// @title       Automatic Message Sender API
// @description The system sends 2 messages every 2 minutes.
// @version     1.0
// @host        localhost:8080
// @BasePath    /v1
func NewRouter(handler *gin.Engine, l logger.Interface, a usecase.AutoMessager, t usecase.Messager) {
	// Options
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	// Swagger
	swaggerHandler := ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "DISABLE_SWAGGER_HTTP_HANDLER")
	handler.GET("/swagger/*any", swaggerHandler)

	// K8s probe --Checks if the container is still alive or stuck. If a container fails the liveness probe, Kubernetes will restart it.
	handler.GET("/healthz", func(c *gin.Context) { c.Status(http.StatusOK) })

	// Prometheus metrics
	handler.GET("/metrics", gin.WrapH(promhttp.Handler()))

	l.Info("http://localhost:8080/swagger/index.html")

	// Routers
	h := handler.Group("/v1")
	{
		newMessageRoutes(h, t, l)
		newAutoMessageSenderRoutes(h, a, l)
	}

}

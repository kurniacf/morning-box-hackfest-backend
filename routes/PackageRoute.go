package routes

import (
	"morning-box-hackfest-be/handler"
	"morning-box-hackfest-be/repository"
	"morning-box-hackfest-be/service"

	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
)

func SetupPackageRoutes(router *gin.Engine, client *firestore.Client) {
	packageRepo := repository.NewPackageRepository(client)
	packageService := service.NewPackageService(packageRepo)
	packageHandler := handler.NewPackageHandler(*packageService)
	packages := router.Group("/packages")
	{
		packages.GET("", packageHandler.GetAllPackages)
		packages.GET("/:id", packageHandler.GetPackage)
		packages.POST("", packageHandler.CreatePackage)
		packages.PUT("/:id", packageHandler.UpdatePackage)
		packages.DELETE("/:id", packageHandler.DeletePackage)
	}
}

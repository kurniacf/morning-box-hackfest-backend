package routes

import (
	"morning-box-hackfest-be/config"
	"morning-box-hackfest-be/handler"
	"morning-box-hackfest-be/repository"
	"morning-box-hackfest-be/service"

	"github.com/gin-gonic/gin"
)

func AddPackageRoutes(r *gin.Engine) {
	firestoreClient := config.GetFirestoreClient()
	packageRepo := repository.NewPackageRepository(firestoreClient)
	packageService := service.NewPackageService(packageRepo)
	packageHandler := handler.NewPackageHandler(*packageService)

	packageGroup := r.Group("/packages")
	{
		packageGroup.GET("", packageHandler.GetAllPackages)
		packageGroup.GET("/:id", packageHandler.GetPackage)
		packageGroup.POST("", packageHandler.CreatePackage)
		packageGroup.PUT("/:id", packageHandler.UpdatePackage)
		packageGroup.DELETE("/:id", packageHandler.DeletePackage)
	}
}

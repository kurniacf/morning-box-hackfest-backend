package handler

import (
	"morning-box-hackfest-be/config"
	"morning-box-hackfest-be/model"
	"morning-box-hackfest-be/repository"
	"morning-box-hackfest-be/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PackageHandler struct {
	service service.PackageService
}

func NewPackageHandler(service service.PackageService) *PackageHandler {
	return &PackageHandler{service: service}
}

func (h *PackageHandler) GetAllPackages(c *gin.Context) {

	packages, err := h.service.GetAllPackages()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get packages"})
		return
	}

	firestoreClient := config.GetFirestoreClient()
	foodRepo := repository.NewFoodRepository(firestoreClient)
	foodService := service.NewFoodService(foodRepo)

	for _, pack := range packages {
		pack.Foods.Monday, _ = foodService.GetFood(pack.Foods.Monday.Id)
		pack.Foods.Tuesday, _ = foodService.GetFood(pack.Foods.Tuesday.Id)
		pack.Foods.Wednesday, _ = foodService.GetFood(pack.Foods.Wednesday.Id)
		pack.Foods.Thursday, _ = foodService.GetFood(pack.Foods.Thursday.Id)
		pack.Foods.Friday, _ = foodService.GetFood(pack.Foods.Friday.Id)
		pack.Foods.Saturday, _ = foodService.GetFood(pack.Foods.Saturday.Id)
		pack.Foods.Sunday, _ = foodService.GetFood(pack.Foods.Tuesday.Id)
	}

	c.JSON(http.StatusOK, packages)
}

func (h *PackageHandler) GetPackage(c *gin.Context) {
	id := c.Param("id")
	pkg, err := h.service.GetPackage(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Package not found"})
		return
	}

	firestoreClient := config.GetFirestoreClient()
	foodRepo := repository.NewFoodRepository(firestoreClient)
	foodService := service.NewFoodService(foodRepo)

	pkg.Foods.Monday, _ = foodService.GetFood(pkg.Foods.Monday.Id)
	pkg.Foods.Tuesday, _ = foodService.GetFood(pkg.Foods.Tuesday.Id)
	pkg.Foods.Wednesday, _ = foodService.GetFood(pkg.Foods.Wednesday.Id)
	pkg.Foods.Thursday, _ = foodService.GetFood(pkg.Foods.Thursday.Id)
	pkg.Foods.Friday, _ = foodService.GetFood(pkg.Foods.Friday.Id)
	pkg.Foods.Saturday, _ = foodService.GetFood(pkg.Foods.Saturday.Id)
	pkg.Foods.Sunday, _ = foodService.GetFood(pkg.Foods.Sunday.Id)

	c.JSON(http.StatusOK, pkg)
}

func (h *PackageHandler) CreatePackage(c *gin.Context) {
	var req model.PackageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := h.service.CreatePackage(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func (h *PackageHandler) UpdatePackage(c *gin.Context) {
	id := c.Param("id")

	var req model.PackageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.UpdatePackage(id, req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

func (h *PackageHandler) DeletePackage(c *gin.Context) {
	id := c.Param("id")
	err := h.service.DeletePackage(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete package"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Package deleted successfully"})
}

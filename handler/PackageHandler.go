package handler

import (
	"morning-box-hackfest-be/model"
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
	c.JSON(http.StatusOK, packages)
}

func (h *PackageHandler) GetPackage(c *gin.Context) {
	id := c.Param("id")
	pkg, err := h.service.GetPackage(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Package not found"})
		return
	}
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

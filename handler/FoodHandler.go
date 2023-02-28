package handler

import (
	"context"
	"fmt"
	"net/http"

	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/iterator"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Food struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	ImageURL    string `json:"imageUrl"`
}

// CreateFood creates a new food document in Firestore.
func CreateFood(c *gin.Context) {
	// Get Firestore client from context
	client := c.MustGet("firestoreClient").(*firestore.Client)

	// Bind JSON payload to Food struct
	var food Food
	if err := c.ShouldBindJSON(&food); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Add food document to Firestore
	docRef, _, err := client.Collection("foods").Add(context.Background(), food)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Food created with ID: %s", docRef.ID)})
}

// GetFood retrieves a food document by ID from Firestore.
func GetFood(c *gin.Context) {
	// Get Firestore client from context
	client := c.MustGet("firestoreClient").(*firestore.Client)

	// Get food document from Firestore
	docRef := client.Collection("foods").Doc(c.Param("id"))
	docSnap, err := docRef.Get(context.Background())
	if err != nil {
		if status.Code(err) == codes.NotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Food not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	// Map Firestore document snapshot to Food struct
	var food Food
	if err := docSnap.DataTo(&food); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, food)
}

// GetAllFoods retrieves all food documents from Firestore.
func GetAllFoods(c *gin.Context) {
	// Get Firestore client from context
	client := c.MustGet("firestoreClient").(*firestore.Client)

	// Get all food documents from Firestore
	iter := client.Collection("foods").Documents(context.Background())
	defer iter.Stop()

	var foods []Food
	for {
		docSnap, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Map Firestore document snapshot to Food struct
		var food Food
		if err := docSnap.DataTo(&food); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		foods = append(foods, food)
	}

	c.JSON(http.StatusOK, foods)
}

// UpdateFood updates an existing food document in Firestore.
func UpdateFood(c *gin.Context) {
	// Get firestore client from context
	client := c.MustGet("firestoreClient").(*firestore.Client)

	// Get food ID from path
	foodID := c.Param("id")

	// Get food data from request body
	var updatedFood Food
	if err := c.ShouldBindJSON(&updatedFood); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update the food in the database
	_, err := client.Collection("foods").Doc(foodID).Set(context.Background(), updatedFood)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Food with ID %s has been updated successfully", foodID),
	})
}

// DeleteFood deletes an existing food document from Firestore.
func DeleteFood(c *gin.Context) {
	// Get Firestore client from context
	client := c.MustGet("firestoreClient").(*firestore.Client)

	// Get food ID from path
	foodID := c.Param("id")

	// Delete food document from Firestore
	if _, err := client.Collection("foods").Doc(foodID).Delete(context.Background()); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Food with ID %s has been deleted successfully", foodID)})
}

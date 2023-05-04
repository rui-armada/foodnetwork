package controllers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"foodnetwork/config"
	"foodnetwork/models"
)

// CreatePost creates a new post
func CreatePost(c *gin.Context) {
	db := config.InitDB()

	// Get token from Authorization header
	authHeader := c.Request.Header.Get("Authorization")
	tokenString := strings.Replace(authHeader, "Bearer ", "", -1)

	// Verify and decode token to extract user ID
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Verify that the signing method is correct
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// Return the secret key used to sign the token
		return []byte("j@I,1I'`Vno&NW(NiFV?]LXG#n3l*3A?"), nil
	})
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
		return
	}

	// Extract user ID from token claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token claims"})
		return
	}
	userID := uint(claims["user_id"].(float64))

	// Bind request body to Post struct
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set user ID and create Post
	post.UserID = userID
	db.Create(&post)

	c.JSON(http.StatusCreated, gin.H{"data": post})
}

// controllers/post.go
func GetAllPosts(c *gin.Context) {
	db := config.InitDB()

	var posts []models.Post
	if err := db.Find(&posts).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Records not found"})
		return
	}

	postData := make([]map[string]interface{}, len(posts))

	for i, post := range posts {
		userName, err := GetUserNameByID(post.UserID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}
		postData[i] = map[string]interface{}{
			"id":          post.ID,
			"userID":      post.UserID,
			"businessID":  post.BusinessID,
			"description": post.Description,
			"userName":    userName,
		}
	}

	c.JSON(http.StatusOK, gin.H{"data": postData})
}

// GetPost retrieves a single post by its ID
func GetPost(c *gin.Context) {
	db := config.InitDB()

	var post models.Post
	if err := db.Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": post})
}

// UpdatePost updates a post by its ID
func UpdatePost(c *gin.Context) {
	db := config.InitDB()

	var post models.Post
	if err := db.Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Save(&post)

	c.JSON(http.StatusOK, gin.H{"data": post})
}

// DeletePost deletes a post by its ID
func DeletePost(c *gin.Context) {
	db := config.InitDB()

	var post models.Post
	if err := db.Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	db.Delete(&post)

	c.JSON(http.StatusOK, gin.H{"data": "Record deleted successfully"})
}

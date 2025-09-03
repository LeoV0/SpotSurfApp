package user

import (
	"net/http"
	"surf_spots_app/model"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func (h *UserHandler) RegisterUsers(ctx *gin.Context) {

	var data map[string]string

	if ctx.BindJSON(&data) != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	if data["email"] == "" || data["password"] == "" || data["name"] == "" {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing required fields"})
        return
    }


	hash, err := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})
		return
	}
	

	user := model.Users{
		Name: data["name"], 
		Email: data["email"], 
		Password: hash, 
		Role: data["role"],
	}

	

	result := h.DB.Create(&user)
	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id":    user.ID,
		"name":  user.Name,
		"email": user.Email,
		"role": user.Role,
	})
}

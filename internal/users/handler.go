package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type UserHandler struct {
	*UserService
}

func (uh *UserHandler) CreateUser(ctx *gin.Context) {

	var user User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format: " + err.Error()})
		return
	}

	err := uh.UserService.CreateUser(ctx.Request.Context(), user)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user" + err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "User created successfully",
		"data":    user,
	})
}

func (uh *UserHandler) GetAllUsers(ctx *gin.Context) {

	users, err := uh.UserService.GetAll(ctx.Request.Context())

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get users" + err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Got Users successfully",
		"data":    users,
	})
}

func (uh *UserHandler) GetUserByID(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := bson.ObjectIDFromHex(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed: " + err.Error()})
		return
	}

	user, err := uh.UserService.GetUserByID(ctx.Request.Context(), id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user" + err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Got User successfully",
		"data":    user,
	})
}

func (uh *UserHandler) UpdateUser(ctx *gin.Context) {
	idstr := ctx.Param("id")
	id, err := bson.ObjectIDFromHex(idstr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID: " + err.Error()})
		return
	}

	var updateData bson.M
	if err := ctx.ShouldBindJSON(&updateData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON: " + err.Error()})
		return
	}

	err = uh.UserService.UpdateUser(ctx.Request.Context(), id, updateData)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user: " + err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":        "User updated successfully",
		"updated_fields": updateData,
	})
}

func (uh *UserHandler) DeleteUser(ctx *gin.Context) {
	idstr := ctx.Param("id")
	id, err := bson.ObjectIDFromHex(idstr)
	err = uh.UserService.DeleteUser(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete" + err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Delted User successfully",
		"data":    id,
	})
}

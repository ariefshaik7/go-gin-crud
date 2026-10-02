package users

import "github.com/gin-gonic/gin"



func RegisterRoutes(r *gin.Engine, han *UserHandler) {
	
	userRoutes := r.Group("/users")
	{
		userRoutes.POST("", han.CreateUser)
		userRoutes.GET("", han.GetAllUsers)
		userRoutes.GET("/:id", han.GetUserByID)
		userRoutes.PATCH("/:id", han.UpdateUser)
		userRoutes.DELETE("/:id", han.DeleteUser)
	}
}
package routes

import (
	"github.com/gin-gonic/gin"

	"foodnetwork/controllers"
)

func SetupRouter(router *gin.Engine) *gin.Engine {

	v1 := router.Group("/api/v1")

	// Routes for User entity
	v1.POST("/users", controllers.CreateUser)
	v1.GET("/users/:id", controllers.GetUser)
	v1.PUT("/users/:id", controllers.UpdateUser)
	v1.DELETE("/users/:id", controllers.DeleteUser)

	// Routes for Business entity
	v1.POST("/businesses", controllers.CreateBusiness)
	v1.GET("/businesses/:id", controllers.GetBusiness)
	v1.PUT("/businesses/:id", controllers.UpdateBusiness)
	v1.DELETE("/businesses/:id", controllers.DeleteBusiness)

	// Routes for BusinessUser entity
	v1.POST("/business_users", controllers.CreateBusinessUser)
	v1.GET("/business_users/:id", controllers.GetBusinessUser)
	v1.PUT("/business_users/:id", controllers.UpdateBusinessUser)
	v1.DELETE("/business_users/:id", controllers.DeleteBusinessUser)

	// Routes for BusinessType entity
	v1.POST("/business_types", controllers.CreateBusinessType)
	v1.GET("/business-types/:id", controllers.GetBusinessType)
	v1.PUT("/business-types/:id", controllers.UpdateBusinessType)
	v1.DELETE("/business-types/:id", controllers.DeleteBusinessType)

	// JobTitle routes
	v1.POST("/job_titles", controllers.CreateJobTitle)
	v1.GET("/job_titles/:id", controllers.GetJobTitle)
	v1.PUT("/job_titles/:id", controllers.UpdateJobTitle)
	v1.DELETE("/job_titles/:id", controllers.DeleteJobTitle)
	v1.GET("/job_titles", controllers.GetAllJobTitles)

	// JobPost routes
	v1.POST("/job_posts", controllers.CreateJobPost)
	v1.GET("/job_posts/:id", controllers.GetJobPost)
	v1.PUT("/job_posts/:id", controllers.UpdateJobPost)
	v1.DELETE("/job_posts/:id", controllers.DeleteJobPost)
	//v1.GET("/job_posts", controllers.ListJobPosts)

	// Product routes
	v1.POST("/products", controllers.CreateProduct)
	v1.GET("/products/:id", controllers.GetProduct)
	v1.PUT("/products/:id", controllers.UpdateProduct)
	v1.DELETE("/products/:id", controllers.DeleteProduct)
	//v1.GET("/products", controllers.ListProducts)

	// Service routes
	v1.POST("/services", controllers.CreateService)
	v1.GET("/services/:id", controllers.GetService)
	v1.PUT("/services/:id", controllers.UpdateService)
	v1.DELETE("/services/:id", controllers.DeleteService)
	//v1.GET("/services", controllers.ListServices)

	// Rating routes
	v1.POST("/ratings", controllers.CreateRating)
	v1.GET("/ratings/:id", controllers.GetRating)
	v1.PUT("/ratings/:id", controllers.UpdateRating)
	v1.DELETE("/ratings/:id", controllers.DeleteRating)
	//v1.GET("/ratings", controllers.ListRatings)

	// Review routes
	v1.POST("/reviews", controllers.CreateReview)
	v1.GET("/reviews/:id", controllers.GetReview)
	v1.PUT("/reviews/:id", controllers.UpdateReview)
	v1.DELETE("/reviews/:id", controllers.DeleteReview)
	//v1.GET("/reviews", controllers.ListReviews)

	// ProfessionalExperience routes
	v1.POST("/professional-experiences", controllers.CreateProfessionalExperience)
	v1.GET("/professional-experiences/:id", controllers.GetProfessionalExperience)
	v1.PUT("/professional-experiences/:id", controllers.UpdateProfessionalExperience)
	v1.DELETE("/professional-experiences/:id", controllers.DeleteProfessionalExperience)
	//v1.GET("/professional-experiences", controllers.ListProfessionalExperiences)

	// Routes for Post entity
	v1.POST("/posts", controllers.CreatePost)
	v1.GET("/posts", controllers.GetAllPosts)
	v1.GET("/posts/:id", controllers.GetPost)
	v1.PUT("/posts/:id", controllers.UpdatePost)
	v1.DELETE("/posts/:id", controllers.DeletePost)

	v1.POST("/login", controllers.Login)
	v1.POST("/register", controllers.Register)

	return router
}

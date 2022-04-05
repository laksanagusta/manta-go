package routes

import (
	"net/http"
	"novocaine-dev/auth"
	"novocaine-dev/handler"
	"novocaine-dev/helper"
	"novocaine-dev/product"
	"novocaine-dev/task"
	"novocaine-dev/taskHistory"
	"novocaine-dev/user"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB) gin.Engine {
	//Repository
	userRepository := user.NewRepository(db)
	productRepository := product.NewRepository(db)
	taskRepository := task.NewRepository(db)
	taskHistoryRepository := taskHistory.NewRepository(db)

	//Service
	productService := product.NewService(productRepository)
	userService := user.NewService(userRepository)
	taskService := task.NewService(taskRepository)
	taskHistoryService := taskHistory.NewService(taskHistoryRepository)
	authService := auth.NewService()

	//Handler
	userHandler := handler.NewUserHandler(userService, authService)
	productHandler := handler.NewProductHandler(productService)
	taskHandler := handler.NewTaskHandler(taskService, userService)
	taskHistoryHandler := handler.NewTaskHistoryHandler(taskHistoryService, taskService, userService)

	router := gin.Default()
	router.Static("/images", "./images")
	api := router.Group("/api/v1")

	//api users
	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.GET("/users/:id", userHandler.UserDetails)
	api.GET("/users", userHandler.UserFindAll)
	api.PUT("/users/:id", authMiddleware(authService, userService), userHandler.UpdateUser)
	api.DELETE("/users/:id", authMiddleware(authService, userService), userHandler.DeleteUser)

	//api task
	api.POST("/tasks", authMiddleware(authService, userService), taskHandler.CreateTask)
	api.PUT("/tasks/:id", authMiddleware(authService, userService), taskHandler.UpdateTask)
	api.GET("/tasks/:id", authMiddleware(authService, userService), taskHandler.FindTaskById)
	api.GET("/tasks", authMiddleware(authService, userService), taskHandler.CustomFilter)
	api.DELETE("/tasks/:id", authMiddleware(authService, userService), taskHandler.Delete)

	//api products
	api.GET("/products", productHandler.GetProducts)
	api.GET("/products/:id", productHandler.GetProductById)
	api.POST("/products", authMiddleware(authService, userService), productHandler.CreateProduct)
	api.PUT("/products/:id", authMiddleware(authService, userService), productHandler.UpdateProduct)
	api.PUT("/products/upload-image/:id", authMiddleware(authService, userService), productHandler.UploadImage)
	api.DELETE("/products/:id", authMiddleware(authService, userService), productHandler.DeleteProduct)

	//api taskHistory
	api.POST("/task-histories", authMiddleware(authService, userService), taskHistoryHandler.CreateTaskHistory)

	return *router
}

func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userID := int(claim["user_id"].(float64))

		user, err := userService.UserDetails(userID)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("currentUser", user)

	}
}

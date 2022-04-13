package routes

import (
	"net/http"
	"novocaine-dev/auth"
	"novocaine-dev/handler"
	"novocaine-dev/helper"
	"novocaine-dev/organization"
	"novocaine-dev/product"
	"novocaine-dev/task"
	"novocaine-dev/taskHistory"
	"novocaine-dev/taskUser"
	"novocaine-dev/transaction"
	"novocaine-dev/transactionProduct"
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
	taskUserRepository := taskUser.NewRepository(db)
	transactionRepository := transaction.NewRepository(db)
	transactionProductRepository := transactionProduct.NewRepository(db)
	organizationRepository := organization.NewRepository(db)

	//Service
	productService := product.NewService(productRepository)
	userService := user.NewService(userRepository)
	taskService := task.NewService(taskRepository, taskHistoryRepository, taskUserRepository)
	taskHistoryService := taskHistory.NewService(taskHistoryRepository)
	transactionService := transaction.NewService(transactionRepository, transactionProductRepository)
	transactionProductService := transactionProduct.NewService(transactionProductRepository)
	organizationService := organization.NewService(organizationRepository)
	authService := auth.NewService()

	//Handler
	userHandler := handler.NewUserHandler(userService, authService)
	productHandler := handler.NewProductHandler(productService)
	taskHandler := handler.NewTaskHandler(taskService, userService)
	taskHistoryHandler := handler.NewTaskHistoryHandler(taskHistoryService, taskService, userService)
	transactionHandler := handler.NewTransactionHandler(transactionService)
	transactionProductHandler := handler.NewTransactionProductHandler(transactionProductService)
	organizationHandler := handler.NewOrganizationHandler(organizationService)

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
	api.POST("/tasks/process", authMiddleware(authService, userService), taskHandler.ProcessTask)
	api.POST("/tasks/assign", authMiddleware(authService, userService), taskHandler.AssignTask)
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

	//transaction
	api.GET("/transactions/:id", transactionHandler.FindById)
	//transaction
	api.GET("/transactions/userId/:userId", transactionHandler.FindTransactionByUser)
	//transaction product
	api.POST("/transaction-product/add", authMiddleware(authService, userService), transactionHandler.AddToCart)       //atc
	api.DELETE("/transaction-product/:id", authMiddleware(authService, userService), transactionProductHandler.Delete) //delete items from cart

	//api taskHistory
	api.POST("/task-histories", authMiddleware(authService, userService), taskHistoryHandler.CreateTaskHistory)

	//api organizations
	api.POST("/organizations", authMiddleware(authService, userService), organizationHandler.Save)
	api.GET("/organizations/:id", authMiddleware(authService, userService), organizationHandler.FindById)

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

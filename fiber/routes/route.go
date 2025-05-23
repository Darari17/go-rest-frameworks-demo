package routes

import (
	"log"
	"os"

	"github.com/go-rest-frameworks-demo/fiber/config"
	"github.com/go-rest-frameworks-demo/fiber/internal/controllers"
	"github.com/go-rest-frameworks-demo/fiber/internal/jwt"
	"github.com/go-rest-frameworks-demo/fiber/internal/middleware"
	"github.com/go-rest-frameworks-demo/fiber/internal/repositories"
	"github.com/go-rest-frameworks-demo/fiber/internal/services"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Routes struct {
	app *fiber.App
	db  *gorm.DB

	userController controllers.UserController

	authMiddleware middleware.AuthMiddleware
	jwtHandler     jwt.JWTHandler
}

func (r *Routes) setupRoutes() {
	api := r.app.Group("/api/v1")

	api.Post("/login", r.userController.Login)
	api.Post("/register", r.userController.Register)
}

func (r *Routes) Run(port string) {
	r.setupRoutes()
	log.Println("Server is running on port", port)

	if err := r.app.Listen(port); err != nil {
		log.Fatalf("Failed to start server on port %s: %v", port, err)
	}
}

func (r *Routes) Close() {
	sqlDB, err := r.db.DB()
	if err != nil {
		log.Println("Failed to get database instance:", err)
		return
	}

	if err := sqlDB.Close(); err != nil {
		log.Println("Failed to close database connection:", err)
	} else {
		log.Println("Database connection closed successfully")
	}
}

func NewServer() *Routes {
	db, err := config.ConnDB()
	if err != nil {
		log.Fatal("Could not connect to database:", err)
	}

	// migrate

	secretKey := os.Getenv("SECRET_KEY")
	if secretKey == "" {
		log.Fatal("SECRET_KEY is not set in environment")
	}

	app := fiber.New()

	jwtHandler := jwt.NewJWTHandler(secretKey)
	authMiddleware := middleware.NewAuthMiddleware(jwtHandler)

	userRepo := repositories.NewUserRepo(db)

	userService := services.NewUserService(userRepo, jwtHandler)

	userController := controllers.NewUserController(userService)

	return &Routes{
		app:            app,
		db:             db,
		userController: userController,
		authMiddleware: authMiddleware,
		jwtHandler:     jwtHandler,
	}
}

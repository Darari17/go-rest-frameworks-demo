package routes

import (
	"log"
	"os"

	"github.com/Darari17/go-rest-frameworks-demo/gin/config"
	"github.com/Darari17/go-rest-frameworks-demo/gin/internal/controllers"
	"github.com/Darari17/go-rest-frameworks-demo/gin/internal/jwt"
	"github.com/Darari17/go-rest-frameworks-demo/gin/internal/middleware"
	"github.com/Darari17/go-rest-frameworks-demo/gin/internal/repositories"
	"github.com/Darari17/go-rest-frameworks-demo/gin/internal/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Routes struct {
	app            *gin.Engine
	db             *gorm.DB
	authMiddleware middleware.AuthMiddleware
	jwtHandler     jwt.JWTHandler

	userController controllers.UserController
}

func (r *Routes) userRoute() {
	auth := r.app.Group("/api/v1")
	auth.POST("/login", r.userController.Login)
	auth.POST("/register", r.userController.Register)
}

func (r *Routes) postRoute() {
	post := r.app.Group("/api/v1/posts")
	post.POST("/").Use(r.authMiddleware.RequiredToken())
	post.GET("/")
	post.GET("/:id")
	post.GET("/:id").Use(r.authMiddleware.RequiredToken())
	post.DELETE("/:id").Use(r.authMiddleware.RequiredToken())
}

func (r *Routes) setupRoutes() {
	r.userRoute()
	r.postRoute()
}

func (r *Routes) Run(port string) {
	r.setupRoutes()
	log.Println("Server is running on port", port)

	if err := r.app.Run(port); err != nil {
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

	// if err := migrations.Migrate(db); err != nil {
	// 	log.Fatal("Database migration failed: ", err)
	// }

	secretKey := os.Getenv("SECRET_KEY")
	if secretKey == "" {
		log.Fatal("SECRET_KEY is not set in environment")
	}

	app := gin.Default()

	jwtHandler := jwt.NewJWTHandler(secretKey)
	middleware := middleware.NewAuthMiddleware(*jwtHandler)

	userRepo := repositories.NewUserRepo(db)

	userService := services.NewUserService(*userRepo, *jwtHandler)

	userController := controllers.NewUserController(*userService)

	return &Routes{
		app:            app,
		db:             &gorm.DB{},
		authMiddleware: *middleware,
		userController: *userController,
		jwtHandler:     *jwtHandler,
	}
}

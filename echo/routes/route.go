package routes

import (
	"log"

	"github.com/Darari17/go-rest-frameworks-demo/echo/config"
	"github.com/Darari17/go-rest-frameworks-demo/echo/internal/controllers"
	"github.com/Darari17/go-rest-frameworks-demo/echo/internal/jwt"
	"github.com/Darari17/go-rest-frameworks-demo/echo/internal/middleware"
	"github.com/Darari17/go-rest-frameworks-demo/echo/internal/repositories"
	"github.com/Darari17/go-rest-frameworks-demo/echo/internal/services"
	"github.com/Darari17/go-rest-frameworks-demo/echo/migrations"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Routes struct {
	app            *echo.Echo
	db             *gorm.DB
	authMiddleware middleware.AuthMiddleware
	jwtHandler     jwt.JWTHandler

	userController controllers.UserController
	postController controllers.PostController
}

func (r *Routes) userRoute() {
	auth := r.app.Group("/api/v1")
	auth.POST("/login", r.userController.Login)
	auth.POST("/register", r.userController.Register)
}

func (r *Routes) postRoute() {
	post := r.app.Group("/api/v1/posts")

	post.GET("/:id", r.postController.GetPostByPostID)

	post.POST("/", r.postController.CreatePost, r.authMiddleware.RequiredToken())
	post.GET("/", r.postController.GetPostsByUserID, r.authMiddleware.RequiredToken())
	post.PUT("/:id", r.postController.UpdatePost, r.authMiddleware.RequiredToken())
	post.DELETE("/:id", r.postController.DeletePost, r.authMiddleware.RequiredToken())
}

func (r *Routes) setupRoutes() {
	r.userRoute()
	r.postRoute()
}

func (r *Routes) Run(port string) {
	r.setupRoutes()
	log.Println("Server is running on port", port)

	if err := r.app.Start(port); err != nil {
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

	if err := migrations.Migrate(db); err != nil {
		log.Fatal("Database migration failed: ", err)
	}

	app := echo.New()

	jwtHandler := jwt.NewJWTHandler(config.AppConfig.JWTConfig.SecretKey)

	middleware := middleware.NewAuthMiddleware(*jwtHandler)

	userRepo := repositories.NewUserRepository(db)
	postRepo := repositories.NewPostRepository(db)

	userService := services.NewUserService(*userRepo, *jwtHandler)
	postService := services.NewPostService(*postRepo)

	userController := controllers.NewUserController(*userService)
	postController := controllers.NewPostController(*postService)

	return &Routes{
		app:            app,
		db:             db,
		authMiddleware: *middleware,
		jwtHandler:     *jwtHandler,

		userController: *userController,
		postController: *postController,
	}
}

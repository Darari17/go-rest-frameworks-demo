package routes

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/Darari17/go-rest-frameworks-demo/no-frameworks/config"
	"github.com/Darari17/go-rest-frameworks-demo/no-frameworks/internal/dtos"
	"github.com/Darari17/go-rest-frameworks-demo/no-frameworks/internal/handlers"
	"github.com/Darari17/go-rest-frameworks-demo/no-frameworks/internal/helper"
	"github.com/Darari17/go-rest-frameworks-demo/no-frameworks/internal/middleware"
	"github.com/Darari17/go-rest-frameworks-demo/no-frameworks/internal/repositories"
	"github.com/Darari17/go-rest-frameworks-demo/no-frameworks/internal/usecases"
	"github.com/Darari17/go-rest-frameworks-demo/no-frameworks/internal/util"
	"github.com/Darari17/go-rest-frameworks-demo/no-frameworks/migrations"
)

type Routes struct {
	db             *sql.DB
	authMiddleware middleware.AuthMiddleware
	jwtHandler     util.JwtHandler

	userHandler handlers.UserHandler
	postHandler handlers.PostHandler
}

func (r *Routes) UserRoutes() {
	http.HandleFunc("/api/v1/login", r.userHandler.Login)
	http.HandleFunc("/api/v1/register", r.userHandler.Register)
}

func (r *Routes) PostRoutes() {
	// /api/v1/posts        → GET (all), POST (create)
	http.HandleFunc("/api/v1/posts", func(w http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case http.MethodGet:
			r.authMiddleware.RequiredToken(http.HandlerFunc(r.postHandler.FindPostsByUserId)).ServeHTTP(w, req)
		case http.MethodPost:
			r.authMiddleware.RequiredToken(http.HandlerFunc(r.postHandler.CreatePost)).ServeHTTP(w, req)
		default:
			helper.JSON(w, http.StatusMethodNotAllowed, dtos.Response[any]{
				Code:   http.StatusMethodNotAllowed,
				Status: http.StatusText(http.StatusMethodNotAllowed),
				Error:  "Method not allowed",
			})
		}
	})

	// /api/v1/posts/{id}   → GET (by id), PUT, DELETE
	http.HandleFunc("/api/v1/posts/", func(w http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case http.MethodGet:
			r.postHandler.FindPostByPostId(w, req)
		case http.MethodPut:
			r.authMiddleware.RequiredToken(http.HandlerFunc(r.postHandler.UpdatePost)).ServeHTTP(w, req)
		case http.MethodDelete:
			r.authMiddleware.RequiredToken(http.HandlerFunc(r.postHandler.DeletePost)).ServeHTTP(w, req)
		default:
			helper.JSON(w, http.StatusMethodNotAllowed, dtos.Response[any]{
				Code:   http.StatusMethodNotAllowed,
				Status: http.StatusText(http.StatusMethodNotAllowed),
				Error:  "Method not allowed",
			})
		}
	})
}

func (r *Routes) InitRoutes() {
	r.UserRoutes()
	r.PostRoutes()
}

func (r *Routes) Run(port string) {
	r.InitRoutes()
	log.Println("Server is running on port", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Failed to start server on port %s: %v", port, err)
	}
}

func (r *Routes) Close() {
	if err := r.db.Close(); err != nil {
		log.Println("Failed to close database connection:", err)
	} else {
		log.Println("Database connection closed successfully")
	}
}

func NewServer() *Routes {

	db, err := config.ConnectDB(config.Cfg.DBConfig)
	if err != nil {
		log.Fatal("Could not connect to database:", err)
	}

	if err := migrations.Migrate(db); err != nil {
		log.Fatal("Database migration failed:", err)
	}

	jwtHandler := util.NewJwtHandler(config.Cfg.JWTConfig.JWTSecretKey)
	authMiddleware := middleware.NewAuthMiddleware(*jwtHandler)

	userRepo := repositories.NewUserRepo(db)
	postRepo := repositories.NewPostRepo(db)

	userUsecase := usecases.NewUserUsecase(*userRepo, *jwtHandler)
	postUsecase := usecases.NewPostUsecase(*postRepo)

	userHandler := handlers.NewUserHandler(*userUsecase)
	postHandler := handlers.NewPostHandler(*postUsecase)

	return &Routes{
		db:             db,
		authMiddleware: *authMiddleware,
		jwtHandler:     *jwtHandler,
		userHandler:    *userHandler,
		postHandler:    *postHandler,
	}
}

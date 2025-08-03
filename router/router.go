package router

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/hanzala211/CRUD/internal/api/handler"
	"github.com/hanzala211/CRUD/middlewares"
)

func SetupRouter(userHandler *handler.UserHandler, postHandler *handler.PostHandler, commentHandler *handler.CommentHandler) http.Handler {
	r := chi.NewRouter()
	r.Route("/api/v1", func(u chi.Router) {
		u.Post("/", userHandler.JustRecieveFile)
		u.Route("/users", func(s chi.Router) {
			s.Post("/", userHandler.CreateUser)
			s.Put("/{id}", userHandler.UpdateUser)
			s.Delete("/{id}", userHandler.DeleteUser)
			s.Get("/", userHandler.GetAllUsers)
		})
		u.Route("/auth", func(s chi.Router) {
			s.Post("/login", userHandler.Login)
			s.Group(func(r chi.Router) {
				r.Use(middlewares.JWTAuthorization)
				r.Get("/me", userHandler.Me)
			})
		})
		u.Route("/posts", func(s chi.Router) {
			s.Group(func(u chi.Router) {
				u.Use(middlewares.JWTAuthorization)
				u.Post("/", postHandler.CreatePost)
				u.Get("/{id}", postHandler.GetPostByID)
			})
			s.Route("/{id}/comments", func(s chi.Router) {
				s.Group(func(u chi.Router) {
					u.Use(middlewares.JWTAuthorization)
					u.Post("/", commentHandler.AddComment)
					u.Get("/", commentHandler.GetPostComments)
				})
			})
		})
	})
	return r
}

package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/hanzala211/CRUD/internal/api/models"
	"github.com/hanzala211/CRUD/internal/services"
	"github.com/hanzala211/CRUD/utils"
)

type PostHandler struct {
	postService *services.PostService
}

func NewPostHandler(postService *services.PostService) *PostHandler {
	return &PostHandler{
		postService: postService,
	}
}

func (h *PostHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(string)
	post := &models.Post{}
	err := json.NewDecoder(r.Body).Decode(post)
	if err != nil {
		utils.WriteError(w, 400, "Invalid Data")
		return
	}
	post.UserID = userID
	err = h.postService.CreatePost(r.Context(), post)
	if err != nil {
		utils.WriteError(w, 500, "Failed to create post")
		return
	}
	utils.WriteJSON(w, 200, post)
}

func (h *PostHandler) GetPostByID(w http.ResponseWriter, r *http.Request) {
	postID := chi.URLParam(r, "id")
	if postID == "" {
		utils.WriteError(w, 400, "Post ID is required")
		return
	}
	post := &models.Post{}
	err := h.postService.GetPostByID(r.Context(), post, postID)
	if err != nil {
		utils.WriteError(w, 500, err.Error())
		return
	}
	utils.WriteJSON(w, 200, post)
}

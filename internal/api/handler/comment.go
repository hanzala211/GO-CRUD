package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/hanzala211/CRUD/internal/api/models"
	"github.com/hanzala211/CRUD/internal/services"
	"github.com/hanzala211/CRUD/utils"
)

type CommentHandler struct {
	commentService *services.CommentService
}

func NewCommentHandler(commentService *services.CommentService) *CommentHandler {
	return &CommentHandler{
		commentService: commentService,
	}
}

func (h *CommentHandler) AddComment(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(string)
	postID := chi.URLParam(r, "id")
	if postID == "" {
		utils.WriteError(w, 400, "Post ID is required")
		return
	}
	comment := &models.Comment{
		PostId: postID,
		UserId: userID,
	}
	err := json.NewDecoder(r.Body).Decode(comment)
	if err != nil {
		utils.WriteError(w, 400, "Invalid Data")
		return
	}
	err = h.commentService.AddComment(r.Context(), comment)
	if err != nil {
		utils.WriteError(w, 500, "Failed to add comment")
		return
	}
	utils.WriteJSON(w, 200, comment)
}

func (h *CommentHandler) GetPostComments(w http.ResponseWriter, r *http.Request) {
	postID := chi.URLParam(r, "id")
	if postID == "" {
		utils.WriteError(w, 400, "Post ID is required")
		return
	}
	comments, err := h.commentService.GetPostComments(r.Context(), postID)
	if err != nil {
		utils.WriteError(w, 500, "Failed to get comments")
		return
	}
	utils.WriteJSON(w, 200, comments)
}

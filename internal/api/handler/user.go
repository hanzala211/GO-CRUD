package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/hanzala211/CRUD/internal/api/models"
	"github.com/hanzala211/CRUD/internal/services"
	"github.com/hanzala211/CRUD/utils"
)

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		utils.WriteError(w, 400, "Invalid Data")
	}

	err = h.userService.CreateUser(r.Context(), user)
	if err != nil {
		utils.WriteError(w, 500, "Failed to create user")
		return
	}

	utils.WriteJSON(w, 200, user)
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "id")
	if userId == "" {
		utils.WriteError(w, 400, "User ID is required")
		return
	}

	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		utils.WriteError(w, 400, "Invalid Data")
		return
	}

	err = h.userService.UpdateUser(r.Context(), user, userId)
	if err != nil {
		utils.WriteError(w, 500, "Failed to update user")
		return
	}
	utils.WriteJSON(w, 200, user)
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "id")
	if userId == "" {
		utils.WriteError(w, 400, "User ID is required")
		return
	}

	err := h.userService.DeleteUser(r.Context(), userId)
	if err != nil {
		utils.WriteError(w, 500, "Failed to delete user")
		return
	}
	utils.WriteJSON(w, 200, map[string]string{"message": "User deleted successfully"})
}

func (h *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.userService.GetAllUsers(r.Context())
	if err != nil {
		utils.WriteError(w, 500, "Failed to get users")
		return
	}
	utils.WriteJSON(w, 200, map[string]any{"data": users})
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		utils.WriteError(w, 400, "Invalid Data")
		return
	}
	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		utils.WriteError(w, 500, "Failed to generate token")
		return
	}
	utils.WriteJSON(w, 200, map[string]any{"token": token})
}

func (h *UserHandler) Me(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("user_id").(string)
	user := &models.User{}
	err := h.userService.GetUserByID(r.Context(), user, userId)
	if err != nil {
		utils.WriteError(w, 500, "Failed to get user")
		return
	}
	utils.WriteJSON(w, 200, map[string]any{"data": user, "message": "User fetched successfully"})
}

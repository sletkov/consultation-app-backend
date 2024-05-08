package handlers

import (
	"github.com/sletkov/consultation-app-backend/internal/api/http/v1/requests"
	"github.com/sletkov/consultation-app-backend/internal/models"
	"log/slog"
	"net/http"
)

func (c *Controller) SaveUser(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewSaveUserRequest(r)

	if err != nil {
		return
	}

	var user = &models.User{
		FullName: req.FullName,
		Email:    req.Email,
		Password: req.Password,
		Role:     models.UserRole(req.Role),
	}

	id, err := c.userService.SaveUser(r.Context(), user)
	if err != nil {
		slog.Log(r.Context(), slog.LevelError, "failed to save user", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))

		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(id))
}

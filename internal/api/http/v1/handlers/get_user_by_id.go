package handlers

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sletkov/consultation-app-backend/internal/api/http/v1/utils"
	"log/slog"
	"net/http"
)

func (c *Controller) GetUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]

	fmt.Println("userID:", userID)
	user, err := c.userService.GetUserByID(r.Context(), userID)
	user.RemoveSensitiveFields()
	if err != nil {
		slog.Log(r.Context(), slog.LevelError, "failed to get user", err)
		utils.ErrRespond(w, r, http.StatusInternalServerError, err)
		return
	}

	utils.Respond(w, r, http.StatusOK, user)
}

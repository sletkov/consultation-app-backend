package handlers

import (
	"fmt"
	"github.com/sletkov/consultation-app-backend/internal/api/http/v1/requests"
	"github.com/sletkov/consultation-app-backend/internal/api/http/v1/utils"
	"log/slog"
	"net/http"
)

func (c *Controller) Login(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewLoginRequest(r)

	if err != nil {
		return
	}

	fmt.Println("req:", req)
	u, err := c.userService.Login(r.Context(), req.Email, req.Password)
	fmt.Println("user:", u)
	if err != nil {
		slog.Log(r.Context(), slog.LevelError, "failed to authenticate user", err)
		utils.ErrRespond(w, r, http.StatusInternalServerError, err)
		return
	}

	session, err := c.sessionStore.Get(r, sessionName)
	fmt.Println("session:", session)
	if err != nil {
		utils.ErrRespond(w, r, http.StatusInternalServerError, err)
		return
	}

	session.Values["user_id"] = u.ID

	if err := c.sessionStore.Save(r, w, session); err != nil {
		utils.ErrRespond(w, r, http.StatusInternalServerError, err)
		return
	}

	utils.Respond(w, r, http.StatusOK, u)
}

package handlers

import (
	"fmt"
	"github.com/sletkov/consultation-app-backend/internal/api/http/v1/requests"
	"github.com/sletkov/consultation-app-backend/internal/api/http/v1/utils"
	"log/slog"
	"net/http"
)

func (c *Controller) SignupConsultation(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewSignupConsultationRequest(r)

	if err != nil {
		return
	}

	fmt.Println("req:", req)

	cons, err := c.consultationService.GetConsultationByID(r.Context(), req.ConsultationID)
	if err != nil {
		slog.Log(r.Context(), slog.LevelError, "consultation wasn't found", err)
		utils.ErrRespond(w, r, http.StatusInternalServerError, err)

		return
	}

	if cons.StudentsCount >= cons.Limit {
		slog.Log(r.Context(), slog.LevelError, "signup limit exceeded", err)
		utils.ErrRespond(w, r, http.StatusBadRequest, err)

		return
	}
	student, err := c.userService.GetUserByID(r.Context(), req.StudentID)
	if err != nil {
		slog.Log(r.Context(), slog.LevelError, "cannot find student", err)
		utils.ErrRespond(w, r, http.StatusBadRequest, err)

		return
	}

	err = c.consultationService.SignupConsultation(r.Context(), student, req.ConsultationID)

	if err != nil {
		slog.Log(r.Context(), slog.LevelError, "failed signup on consultation", err)
		utils.ErrRespond(w, r, http.StatusInternalServerError, err)

		return
	}

	utils.Respond(w, r, http.StatusOK, "success to signup on consultation")
}

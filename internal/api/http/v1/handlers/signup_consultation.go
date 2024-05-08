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

	err = c.consultationService.SignupConsultation(r.Context(), req.StudentID, req.ConsultationID)

	if err != nil {
		slog.Log(r.Context(), slog.LevelError, "failed signup on consultation", err)
		utils.ErrRespond(w, r, http.StatusInternalServerError, err)

		return
	}

	utils.Respond(w, r, http.StatusOK, "success to signup on consultation")
}

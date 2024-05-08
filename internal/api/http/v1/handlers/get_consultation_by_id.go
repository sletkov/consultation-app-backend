package handlers

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sletkov/consultation-app-backend/internal/api/http/v1/utils"
	"log/slog"
	"net/http"
)

func (c *Controller) GetConsultationByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	consultationID := vars["id"]

	fmt.Println("consultationID:", consultationID)
	consultation, err := c.consultationService.GetConsultationByID(r.Context(), consultationID)
	if err != nil {
		slog.Log(r.Context(), slog.LevelError, "failed to get consultations", err)
		utils.ErrRespond(w, r, http.StatusInternalServerError, err)
		return
	}

	utils.Respond(w, r, http.StatusOK, consultation)
}

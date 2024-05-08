package handlers

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sletkov/consultation-app-backend/internal/api/http/v1/utils"
	"log/slog"
	"net/http"
)

func (c *Controller) GetStudentConsultations(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	studentID := vars["id"]

	fmt.Println("studentID:", studentID)
	consultations, err := c.consultationService.GetStudentConsultations(r.Context(), studentID)
	if err != nil {
		slog.Log(r.Context(), slog.LevelError, "failed to get consultations", err)
		utils.ErrRespond(w, r, http.StatusInternalServerError, err)
		return
	}

	utils.Respond(w, r, http.StatusOK, consultations)
}

package handlers

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sletkov/consultation-app-backend/internal/api/http/v1/utils"
	"log/slog"
	"net/http"
)

func (c *Controller) GetTeacherConsultations(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	teacherID := vars["id"]

	fmt.Println("teacherID:", teacherID)
	consultations, err := c.consultationService.GetTeacherConsultations(r.Context(), teacherID)
	if err != nil {
		slog.Log(r.Context(), slog.LevelError, "failed to get consultations", err)
		utils.ErrRespond(w, r, http.StatusInternalServerError, err)
		return
	}

	utils.Respond(w, r, http.StatusOK, consultations)
}

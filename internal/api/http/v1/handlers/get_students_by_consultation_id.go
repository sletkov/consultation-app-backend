package handlers

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sletkov/consultation-app-backend/internal/api/http/v1/utils"
	"log/slog"
	"net/http"
)

func (c *Controller) GetStudentsByConsultationID(w http.ResponseWriter, r *http.Request) {
	slog.Info("[GetStudentsByConsultationID]")
	vars := mux.Vars(r)
	consultationID := vars["id"]

	fmt.Println("consultationID:", consultationID)
	students, err := c.consultationService.GetStudentsByConsultationID(r.Context(), consultationID)
	if err != nil {
		slog.Log(r.Context(), slog.LevelError, "failed to get students", err)
		utils.ErrRespond(w, r, http.StatusInternalServerError, err)
		return
	}

	utils.Respond(w, r, http.StatusOK, students)
}

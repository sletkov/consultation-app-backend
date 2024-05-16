package handlers

import (
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sletkov/consultation-app-backend/internal/api/http/v1/requests"
	"github.com/sletkov/consultation-app-backend/internal/api/http/v1/utils"
	"github.com/sletkov/consultation-app-backend/internal/models"
	"log/slog"
	"net/http"
	"slices"
)

func (c *Controller) UpdateConsultation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	consultationID := vars["id"]
	fmt.Println("consID", consultationID)

	req, err := requests.NewUpdateConsultationRequest(r)
	if err != nil {
		return
	}

	consultations, err := c.consultationService.GetTeacherConsultations(r.Context(), req.TeacherID)
	consIDs := make([]string, len(consultations))
	for _, consultation := range consultations {
		if consultation != nil {
			consIDs = append(consIDs, consultation.ID)

		}
	}
	fmt.Println("consIDS", consIDs)

	if !slices.Contains(consIDs, consultationID) {
		fmt.Println("not contain")
		utils.ErrRespond(w, r, http.StatusInternalServerError, errors.New("teacher not allowed to update this consultation"))
		return
	}

	consultation := &models.Consultation{
		Title:       req.Title,
		Description: req.Description,
		Format:      req.Format,
		Type:        req.Type,
		Date:        req.Date,
		Time:        req.Time,
		Campus:      req.Campus,
		Classroom:   req.Classroom,
		Link:        req.Link,
		Limit:       req.Limit,
		Draft:       req.Draft,
	}

	fmt.Println("req:", req)
	err = c.consultationService.UpdateConsultation(r.Context(), consultation, consultationID)
	if err != nil {
		slog.Log(r.Context(), slog.LevelError, "failed delete consultation", err)
		utils.ErrRespond(w, r, http.StatusInternalServerError, err)

		return
	}

	utils.Respond(w, r, http.StatusOK, consultationID)
}

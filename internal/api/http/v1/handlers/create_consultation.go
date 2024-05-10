package handlers

import (
	"fmt"
	"github.com/sletkov/consultation-app-backend/internal/api/http/v1/requests"
	"github.com/sletkov/consultation-app-backend/internal/api/http/v1/utils"
	"github.com/sletkov/consultation-app-backend/internal/models"
	"log/slog"
	"net/http"
)

func (c *Controller) CreateConsultation(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewCreateConsultationRequest(r)

	if err != nil {
		return
	}

	fmt.Println("req:", req)

	consultation := &models.Consultation{
		Title:         req.Title,
		Description:   req.Description,
		Type:          req.Type,
		Format:        req.Format,
		Date:          req.Date,
		Time:          req.Time,
		TeacherName:   req.TeacherName,
		TeacherID:     req.TeacherID,
		Campus:        req.Campus,
		Classroom:     req.Classroom,
		Link:          req.Link,
		Limit:         req.Limit,
		StudentsCount: req.StudentsCount,
		Draft:         req.Draft,
	}

	fmt.Println("cons after parsing from req:", consultation)

	id, err := c.consultationService.CreateConsultation(r.Context(), consultation, req.TeacherID)

	if err != nil {
		slog.Log(r.Context(), slog.LevelError, "failed create consultation", err)
		utils.ErrRespond(w, r, http.StatusInternalServerError, err)

		return
	}

	utils.Respond(w, r, http.StatusOK, id)
}

package handlers

import (
	"fmt"
	"github.com/sletkov/consultation-app-backend/internal/api/http/v1/requests"
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
	teacher, err := c.userService.GetUserByID(r.Context(), req.TeacherID)
	if err != nil {
		slog.Log(r.Context(), slog.LevelError, "failed get teacher", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))

		return
	}

	//limit, _ := strconv.Atoi(req.Limit)
	consultation := &models.Consultation{
		Title:         req.Title,
		Description:   req.Description,
		Type:          req.Type,
		Format:        req.Format,
		Date:          req.Date,
		Time:          req.Time,
		TeacherName:   teacher.FullName,
		Campus:        req.Campus,
		Classroom:     req.Classroom,
		Link:          req.Link,
		Limit:         req.Limit,
		StudentsCount: req.StudentsCount,
	}

	fmt.Println("cons after parsing from req:", consultation)

	id, err := c.consultationService.CreateConsultation(r.Context(), consultation, req.TeacherID)

	if err != nil {
		slog.Log(r.Context(), slog.LevelError, "failed create consultation", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))

		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(id))
}

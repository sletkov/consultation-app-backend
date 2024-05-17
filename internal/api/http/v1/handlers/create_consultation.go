package handlers

import (
	"fmt"
	"github.com/sletkov/consultation-app-backend/internal/api/http/v1/requests"
	"github.com/sletkov/consultation-app-backend/internal/api/http/v1/utils"
	"github.com/sletkov/consultation-app-backend/internal/models"
	"go.uber.org/zap"
	"net/http"
)

func (c *Controller) CreateConsultation(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewCreateConsultationRequest(r)

	if err != nil {
		return
	}
	zap.L().Info("[CreateConsultation] got request", zap.Any("req", req))
	//fmt.Println("req:", req)

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

	// Если при создании консультации произошла ошибка, логируем ошибку и возвращаем клиенту ответ с ошибкой
	if err != nil {
		zap.L().Error("failed create consultation", zap.Error(err))
		utils.ErrRespond(w, r, http.StatusInternalServerError, err)

		return
	}

	// Если консультация создана успешно, логируем успешное создание и возвращаем клиенту ответ
	zap.L().Info("[CreateConsultation] Success to create consultation")
	utils.Respond(w, r, http.StatusOK, id)
}

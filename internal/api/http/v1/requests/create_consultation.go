package requests

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
)

type CreateConsultationRequest struct {
	Title         string `json:"title"`
	Description   string `json:"description,omitempty""`
	Format        string `json:"format"`
	Type          string `json:"type"`
	TeacherID     string `json:"teacher_id"`
	TeacherName   string `json:"teacher_name"`
	Date          string `json:"date"`
	Time          string `json:"time"`
	Campus        string `json:"campus"`
	Classroom     string `json:"classroom"`
	Link          string `json:"link,omitempty"`
	Limit         int    `json:"limit,omitempty"`
	StudentsCount int    `json:"students_count ,omitempty"`
	Draft         bool   `json:"draft""`
}

func NewCreateConsultationRequest(r *http.Request) (*CreateConsultationRequest, error) {
	defer func() {
		err := r.Body.Close()
		if err != nil {
			slog.Log(r.Context(), slog.LevelError, "failed close body", err)
		}
	}()

	body, err := io.ReadAll(r.Body)
	fmt.Println(body)
	if err != nil {
		return nil, err
	}

	request := &CreateConsultationRequest{}
	if err := json.Unmarshal(body, request); err != nil {
		slog.Log(r.Context(), slog.LevelError, "unmarshal error", err)
	}

	//validate request
	return request, nil
}

package requests

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
)

type UpdateConsultationRequest struct {
	TeacherID   string `json:"teacher_id" validate:"required"`
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
	Format      string `json:"format"`
	Type        string `json:"type"`
	Date        string `json:"date"`
	Time        string `json:"time"`
	Campus      string `json:"campus"`
	Classroom   string `json:"classroom"`
	Link        string `json:"link,omitempty"`
	Limit       int    `json:"limit,omitempty"`
	Draft       bool   `json:"draft"`
}

func NewUpdateConsultationRequest(r *http.Request) (*UpdateConsultationRequest, error) {
	defer func() {
		err := r.Body.Close()
		if err != nil {
			slog.Log(r.Context(), slog.LevelError, "failed close body", err)
		}
	}()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	request := &UpdateConsultationRequest{}
	if err := json.Unmarshal(body, request); err != nil {
		slog.Log(r.Context(), slog.LevelError, "unmarshal error", err)
	}

	//validate request
	return request, nil
}

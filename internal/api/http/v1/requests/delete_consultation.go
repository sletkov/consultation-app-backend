package requests

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
)

type DeleteConsultationRequest struct {
	TeacherID string `json:"teacher_id" validate:"required"`
}

func NewDeleteConsultationRequest(r *http.Request) (*DeleteConsultationRequest, error) {
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

	request := &DeleteConsultationRequest{}
	if err := json.Unmarshal(body, request); err != nil {
		slog.Log(r.Context(), slog.LevelError, "unmarshal error", err)
	}

	//validate request
	return request, nil
}

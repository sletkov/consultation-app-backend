package requests

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
)

type SignupConsultationRequest struct {
	StudentID      string `json:"student_id"`
	ConsultationID string `json:"consultation_id"`
}

func NewSignupConsultationRequest(r *http.Request) (*SignupConsultationRequest, error) {
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

	request := &SignupConsultationRequest{}
	if err := json.Unmarshal(body, request); err != nil {
		slog.Log(r.Context(), slog.LevelError, "unmarshal error", err)
	}

	//validate request
	return request, nil
}

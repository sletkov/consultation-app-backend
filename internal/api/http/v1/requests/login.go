package requests

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
)

type LoginRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func NewLoginRequest(r *http.Request) (*LoginRequest, error) {
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

	request := &LoginRequest{}
	if err := json.Unmarshal(body, request); err != nil {
		slog.Log(r.Context(), slog.LevelError, "unmarshal error", err)
	}

	//validate request
	return request, nil
}

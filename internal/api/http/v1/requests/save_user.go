package requests

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
)

type SaveUserRequest struct {
	FullName string `json:"full_name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
	Role     string `json:"role" validate:"required"`
}

func NewSaveUserRequest(r *http.Request) (*SaveUserRequest, error) {
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

	request := &SaveUserRequest{}
	if err := json.Unmarshal(body, request); err != nil {
		slog.Log(r.Context(), slog.LevelError, "unmarshal error", err)
	}

	//validate request
	return request, nil
}
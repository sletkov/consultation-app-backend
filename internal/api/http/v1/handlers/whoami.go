package handlers

import (
	"net/http"
)

func (c *Controller) Whoami(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Login succeed"))
}

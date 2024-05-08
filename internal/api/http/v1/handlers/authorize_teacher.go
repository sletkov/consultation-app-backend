package handlers

import (
	"context"
	"net/http"
)

const teacherRole = "teacher"

func (c *Controller) AuthorizeTeacher(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := c.sessionStore.Get(r, sessionName)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		id, ok := session.Values["user_id"]
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(err.Error()))
			return
		}

		u, err := c.userService.GetUserByID(r.Context(), id.(string))

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(err.Error()))
			return
		}

		if u.Role != teacherRole {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(err.Error()))
			return
		}

		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), id, u)))
	})
}

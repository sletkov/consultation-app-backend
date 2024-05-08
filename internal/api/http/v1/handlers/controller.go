package handlers

import (
	"context"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/sletkov/consultation-app-backend/internal/models"
	"net/http"
)

type UserService interface {
	SaveUser(ctx context.Context, user *models.User) (string, error)
	Login(ctx context.Context, email, password string) (*models.User, error)
	GetUserByID(ctx context.Context, id string) (*models.User, error)
}

type ConsultationService interface {
	CreateConsultation(ctx context.Context, consultations *models.Consultation, teacherID string) (string, error)
	GetConsultations(ctx context.Context) ([]*models.Consultation, error)
	GetConsultationByID(ctx context.Context, consultationID string) (*models.Consultation, error)
	GetStudentConsultations(ctx context.Context, studentID string) ([]*models.Consultation, error)
	GetTeacherConsultations(ctx context.Context, teacherID string) ([]*models.Consultation, error)
	SignupConsultation(ctx context.Context, studentID, consultationID string) error
}

type Controller struct {
	userService         UserService
	consultationService ConsultationService
	sessionStore        sessions.Store
}

func NewV1Controller(userService UserService, consultationService ConsultationService, sessionStore sessions.Store) *Controller {
	return &Controller{
		userService,
		consultationService,
		sessionStore,
	}
}

func (c *Controller) InitRoutes() *mux.Router {
	r := mux.NewRouter()
	//r.Use(handlers.CORS(handlers.AllowedOrigins([]string{"http://localhost:3000"}), handlers.AllowedMethods([]string{http.MethodPost, http.MethodGet}), handlers.AllowCredentials()))
	r.Use(handlers.CORS(handlers.AllowedOrigins([]string{"http://localhost:3000"}), handlers.AllowCredentials()))
	r.HandleFunc("/sign-up", c.SaveUser).Methods(http.MethodPost)
	r.HandleFunc("/login", c.Login).Methods(http.MethodPost)

	private := r.PathPrefix("/private").Subrouter()
	private.Use(c.AuthenticateUser)
	private.HandleFunc("/whoami", c.Whoami).Methods(http.MethodGet)
	private.HandleFunc("/consultations", c.CreateConsultation).Methods(http.MethodPost)
	private.HandleFunc("/consultations", c.GetConsultations).Methods(http.MethodGet)
	private.HandleFunc("/consultations/{id}", c.GetConsultationByID).Methods(http.MethodGet)
	private.HandleFunc("/consultations/student/{id}", c.GetStudentConsultations).Methods(http.MethodGet)
	private.HandleFunc("/consultations/teacher/{id}", c.GetTeacherConsultations).Methods(http.MethodGet)
	private.HandleFunc("/consultations/signup", c.SignupConsultation).Methods(http.MethodPost)

	private.HandleFunc("/users/{id}", c.GetUserByID).Methods(http.MethodGet)
	return r
}

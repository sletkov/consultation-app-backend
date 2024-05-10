package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/sletkov/consultation-app-backend/internal/models"
)

type ConsultationRepository interface {
	CreateConsultation(ctx context.Context, consultation *models.Consultation) error
	GetAll(ctx context.Context) ([]*models.Consultation, error)
	GetConsultationByID(ctx context.Context, consultationID string) (*models.Consultation, error)
	GetStudentConsultationsIDs(ctx context.Context, studentID string) ([]string, error)
	GetTeacherConsultationsIDs(ctx context.Context, teacherID string) ([]string, error)
	GetUserConsultations(ctx context.Context, consultationIDs []string) ([]*models.Consultation, error)
	SignupConsultation(ctx context.Context, student *models.User, consultation *models.Consultation) error
	AddTeacherConsultation(ctx context.Context, teacherID, consultationID string) error
	DeleteConsultation(ctx context.Context, consultationID string) error
	UpdateConsultation(ctx context.Context, consultation *models.Consultation, consultationID string) error
	GetStudentsByConsultationID(ctx context.Context, consultationID string) ([]*models.User, error)
}

type ConsultationService struct {
	consultationRepository ConsultationRepository
}

func NewConsultationService(consultationRepository ConsultationRepository) *ConsultationService {
	return &ConsultationService{
		consultationRepository,
	}
}

func (service *ConsultationService) CreateConsultation(ctx context.Context, consultation *models.Consultation, teacherID string) (string, error) {
	//no nil pointer check better use values not pointers
	id := uuid.New().String()
	consultation.ID = id

	err := service.consultationRepository.CreateConsultation(ctx, consultation)
	if err != nil {
		return "", err
	}

	err = service.consultationRepository.AddTeacherConsultation(ctx, teacherID, consultation.ID)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (service *ConsultationService) GetConsultations(ctx context.Context) ([]*models.Consultation, error) {
	return service.consultationRepository.GetAll(ctx)
}

func (service *ConsultationService) GetConsultationByID(ctx context.Context, consultationID string) (*models.Consultation, error) {
	return service.consultationRepository.GetConsultationByID(ctx, consultationID)
}

func (service *ConsultationService) GetStudentConsultations(ctx context.Context, studentID string) ([]*models.Consultation, error) {
	consultationIDs, err := service.consultationRepository.GetStudentConsultationsIDs(ctx, studentID)
	if err != nil {
		return nil, err
	}
	return service.consultationRepository.GetUserConsultations(ctx, consultationIDs)
}

func (service *ConsultationService) GetTeacherConsultations(ctx context.Context, teacherID string) ([]*models.Consultation, error) {
	consultationIDs, err := service.consultationRepository.GetTeacherConsultationsIDs(ctx, teacherID)
	if err != nil {
		return nil, err
	}
	return service.consultationRepository.GetUserConsultations(ctx, consultationIDs)
}

func (service *ConsultationService) SignupConsultation(ctx context.Context, student *models.User, consultationID string) error {
	consultation, err := service.consultationRepository.GetConsultationByID(ctx, consultationID)
	if err != nil {
		return err
	}

	return service.consultationRepository.SignupConsultation(ctx, student, consultation)
}

func (service *ConsultationService) DeleteConsultation(ctx context.Context, consultationID string) error {
	return service.consultationRepository.DeleteConsultation(ctx, consultationID)
}

func (service *ConsultationService) UpdateConsultation(ctx context.Context, consultation *models.Consultation, consultationID string) error {
	return service.consultationRepository.UpdateConsultation(ctx, consultation, consultationID)
}

func (service *ConsultationService) GetStudentsByConsultationID(ctx context.Context, consultationID string) ([]*models.User, error) {
	return service.consultationRepository.GetStudentsByConsultationID(ctx, consultationID)
}

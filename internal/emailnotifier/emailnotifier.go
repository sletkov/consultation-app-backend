package emailnotifier

import (
	"crypto/tls"
	"fmt"
	"github.com/sletkov/consultation-app-backend/internal/config"
	"github.com/sletkov/consultation-app-backend/internal/models"
	"gopkg.in/gomail.v2"
)

type EmailNotifier struct {
	cfg    *config.Config
	dialer *gomail.Dialer
}

func New(cfg *config.Config) *EmailNotifier {
	d := gomail.NewDialer(cfg.SMTPHost, cfg.SMTPPort, cfg.SMTPEmail, cfg.SMTPPassword)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	return &EmailNotifier{
		cfg:    cfg,
		dialer: d,
	}
}

func (notifier *EmailNotifier) NotifyStudentSignup(email string, consultation *models.Consultation, student *models.User) error {
	m := gomail.NewMessage()
	m.SetHeader("From", notifier.dialer.Username)
	m.SetHeader("To", email)
	m.SetHeader("Subject", "На вашу консультацию записались!")
	consultationLink := fmt.Sprintf("https://%s/consultations/%s", notifier.cfg.ServerHost, consultation.ID)
	m.SetBody("text/plain", fmt.Sprintf("Студент %s из группы 201-321 записался на вашу %s по дисциплине %s: %s", student.FullName, consultation.Type, consultation.Title, consultationLink))

	if err := notifier.dialer.DialAndSend(m); err != nil {
		return err
	}

	return nil
}

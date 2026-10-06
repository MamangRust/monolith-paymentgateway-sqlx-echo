package mailer

import (
	"bytes"
	"context"
	"fmt"
	"net/smtp"

	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

// Mailer is a struct that implements the Mailer interface
type Mailer struct {
	ctx      context.Context
	server   string
	port     int
	user     string
	password string
	logger   logger.LoggerInterface
	tracer   trace.Tracer
}

type MailerInterface interface {
	Send(to, subject, body string) error
}

// NewMailer returns a new instance of Mailer with the given server, port, user,
// password, and logger.
//
// The context is used to create a new tracer that is stored in the Mailer
// instance. The Mailer instance is used to send emails.
func NewMailer(ctx context.Context, server string, port int, user string, password string, logger logger.LoggerInterface) MailerInterface {
	return &Mailer{
		ctx:      ctx,
		server:   server,
		port:     port,
		user:     user,
		password: password,
		tracer:   otel.Tracer("mailer"),
		logger:   logger,
	}
}

// Send sends an email to the given recipient with the given subject and body.
//
// This function starts a new span named "SendEmail" with the following attributes:
// - email.recipient: The recipient of the email.
// - email.subject: The subject of the email.
// - smtp.server: The SMTP server to use.
// - smtp.port: The SMTP port to use.
//
// If the email fails to send, this function logs the error and sets the span
// status to Error with the message "Failed to send email".
func (m *Mailer) Send(to, subject, body string) error {
	m.logger.Info("Sending email", zap.String("to", to), zap.String("subject", subject))

	_, span := m.tracer.Start(m.ctx, "SendEmail",
		trace.WithAttributes(
			attribute.String("email.recipient", to),
			attribute.String("email.subject", subject),
			attribute.String("smtp.server", m.server),
			attribute.Int("smtp.port", m.port),
		),
	)
	defer span.End()

	auth := smtp.PlainAuth("", m.user, m.password, m.server)

	headers := map[string]string{
		"From":         m.user,
		"To":           to,
		"Subject":      subject,
		"MIME-Version": "1.0",
		"Content-Type": `text/html; charset="UTF-8"`,
	}

	var msg bytes.Buffer
	for k, v := range headers {
		msg.WriteString(fmt.Sprintf("%s: %s\r\n", k, v))
	}
	msg.WriteString("\r\n")
	msg.WriteString(body)

	addr := fmt.Sprintf("%s:%d", m.server, m.port)

	err := smtp.SendMail(addr, auth, m.user, []string{to}, msg.Bytes())
	if err != nil {
		m.logger.Error("Failed to send email", zap.Error(err))

		span.RecordError(err)
		span.SetStatus(codes.Error, "Failed to send email")
	}

	m.logger.Info("Email sent", zap.String("to", to), zap.String("subject", subject))

	return err
}

package src

import (
	"log"

	"github.com/vitwit/matic-jagar/config"
	"github.com/vitwit/matic-jagar/src/alerting/alerts"
)

// SendTelegramAlert sends the alert to telegram chat
func SendTelegramAlert(msg string, cfg *config.Config) error {
	if err := alerts.NewTelegramAlerter().Send(msg, cfg.Telegram.BotToken, cfg.Telegram.ChatID); err != nil {
		log.Printf("failed to send telegram alert: %v", err)
		return err
	}
	return nil
}

// SendEmailAlert to send mail
func SendEmailAlert(msg string, cfg *config.Config) error {
	if err := alerts.NewEmailAlerter().Send(msg, cfg.SendGrid.SendGridAPIToken, cfg.SendGrid.ReceiverMailAddress); err != nil {
		log.Printf("failed to send email alert: %v", err)
		return err
	}
	return nil
}

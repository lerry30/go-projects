package webhook

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"time"
	"encoding/json"
	"net/http"
	"math/rand"
)

type WebhookAppscript struct {
	firstName string
	lastName string
	endpoint string
	status string
	spreadSheetID string
}

func NewWebhookAppscript(firstName, lastName, endpoint, status, spreadSheetID string) *WebhookAppscript {
	return &WebhookAppscript{
		firstName: firstName,
		lastName: lastName,
		endpoint: endpoint,
		status: status,
		spreadSheetID: spreadSheetID,
	}
}

func (w *WebhookAppscript) NotifyAppscript(ctx context.Context, action string) error {
	now := time.Now()
	if w.status == "stopped" {
		return fmt.Errorf("status for the current schedule was paused: %s", now)
	}

	log.Println("sending...")

	formattedSheetName := fmt.Sprintf("%d.%d", int(now.Month()), now.Day())

	payload := struct{
		TrackerID string `json:"tracker-id"`
		SheetName string `json:"sheet-name"`
		FirstName string `json:"first-name"`
		LastName string `json:"last-name"`
		Log string `json:"log"`
	}{
		TrackerID: w.spreadSheetID,
		SheetName: formattedSheetName,
		FirstName: w.firstName,
		LastName: w.lastName,
		Log: action,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	// Delay
	secondsDelay := rand.Intn(20) // 20 seconds
	time.Sleep(time.Duration(secondsDelay) * time.Second)

	resp, err := http.Post(w.endpoint, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	log.Printf("Webhook sent, status: %d", resp.StatusCode)

	b, _ := json.MarshalIndent(payload, "", "  ")
	fmt.Println(string(b))

	return nil
}

/*
{
    "tracker-id": "1ttTYPbzvW6JGi98lnHEUCI5FKrCMx538Ecm2kaQTfSs",
    "sheet-name": "10.11",
    "first-name": "John Lerry",
    "last-name": "Samson",
    "log": "1st-break-in"
}
*/
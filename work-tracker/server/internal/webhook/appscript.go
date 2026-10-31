package webhook

import (
	"context"
	//"net/http"
)
/*
{
    "tracker-id": "1ttTYPbzvW6JGi98lnHEUCI5FKrCMx538Ecm2kaQTfSs",
    "sheet-name": "10.11",
    "first-name": "John Lerry",
    "last-name": "Samson",
    "log": "1st-break-in"
}
*/

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
	
	return nil
}
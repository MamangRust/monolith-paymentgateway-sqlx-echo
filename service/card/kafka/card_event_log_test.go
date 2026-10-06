package kafka

import "testing"

func TestParseCardEventExtractsStableReferences(t *testing.T) {
	tests := []struct {
		name      string
		topic     string
		key       string
		payload   string
		card      string
		reference string
	}{
		{
			name:      "payment",
			topic:     "card.payment.posted",
			key:       "payment-key",
			payload:   `{"payment_id":42,"card_number":"4111"}`,
			card:      "4111",
			reference: "42",
		},
		{
			name:      "fraud alert",
			topic:     "card.fraud.alert",
			key:       "txn-key",
			payload:   `{"txn_id":"txn-1","card_number":"4111"}`,
			card:      "4111",
			reference: "txn-1",
		},
		{
			name:      "statement",
			topic:     "card.statement.generated",
			key:       "cycle-key",
			payload:   `{"billing_id":42,"billing_cycle_day":25}`,
			reference: "42",
		},
		{
			name:      "fallback key",
			topic:     "card.payment.posted",
			key:       "fallback",
			payload:   `{"card_number":"4111"}`,
			card:      "4111",
			reference: "fallback",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			event, err := parseCardEvent(tt.topic, []byte(tt.key), []byte(tt.payload))
			if err != nil {
				t.Fatalf("parseCardEvent returned error: %v", err)
			}
			if event.CardNumber != tt.card || event.ReferenceID != tt.reference {
				t.Fatalf("event metadata = card %q/reference %q, want card %q/reference %q", event.CardNumber, event.ReferenceID, tt.card, tt.reference)
			}
		})
	}
}

func TestParseCardEventRejectsInvalidPayload(t *testing.T) {
	if _, err := parseCardEvent("card.fraud.alert", nil, []byte(`[]`)); err == nil {
		t.Fatal("expected non-object payload to be rejected")
	}
	if _, err := parseCardEvent("card.fraud.alert", nil, []byte(`{`)); err == nil {
		t.Fatal("expected invalid JSON payload to be rejected")
	}
}

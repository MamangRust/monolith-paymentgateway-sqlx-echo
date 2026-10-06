// Code migrated from sqlc-generated sources; domain types kept for repository layer.
// source: card_event_logs.sql

package db

type InsertCardEventLogParams struct {
	Topic       string `db:"topic" json:"topic"`
	EventType   string `db:"event_type" json:"event_type"`
	CardNumber  string `db:"card_number" json:"card_number"`
	ReferenceID string `db:"reference_id" json:"reference_id"`
	Payload     []byte `db:"payload" json:"payload"`
}

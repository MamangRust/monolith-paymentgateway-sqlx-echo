// Code migrated from sqlc-generated sources; domain types kept for repository layer.
// source: outbox.sql

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type ClaimOutboxEventsRow struct {
	ID           int64  `db:"id" json:"id"`
	Topic        string `db:"topic" json:"topic"`
	MessageKey   string `db:"message_key" json:"message_key"`
	Payload      []byte `db:"payload" json:"payload"`
	Attempts     int32  `db:"attempts" json:"attempts"`
	ClaimVersion int64  `db:"claim_version" json:"claim_version"`
}

type ClaimOutboxEventsByTopicsParams struct {
	Limit   int32    `db:"limit" json:"limit"`
	Column2 []string `db:"column2" json:"column_2"`
}

type ClaimOutboxEventsByTopicsRow struct {
	ID           int64  `db:"id" json:"id"`
	Topic        string `db:"topic" json:"topic"`
	MessageKey   string `db:"message_key" json:"message_key"`
	Payload      []byte `db:"payload" json:"payload"`
	Attempts     int32  `db:"attempts" json:"attempts"`
	ClaimVersion int64  `db:"claim_version" json:"claim_version"`
}

type EnqueueOutboxEventParams struct {
	EventKey   string `db:"event_key" json:"event_key"`
	Topic      string `db:"topic" json:"topic"`
	MessageKey string `db:"message_key" json:"message_key"`
	Column4    []byte `db:"column4" json:"column_4"`
}

type MarkConsumerInboxProcessedParams struct {
	ConsumerName       string `db:"consumer_name" json:"consumer_name"`
	EventKey           string `db:"event_key" json:"event_key"`
	ReservationVersion int64  `db:"reservation_version" json:"reservation_version"`
}

type MarkOutboxFailedParams struct {
	ID           int64           `db:"id" json:"id"`
	Status       string          `db:"status" json:"status"`
	Column3      pgtype.Interval `db:"column3" json:"column_3"`
	LastError    string          `db:"last_error" json:"last_error"`
	ClaimVersion int64           `db:"claim_version" json:"claim_version"`
}

type MarkOutboxPublishedParams struct {
	ID           int64 `db:"id" json:"id"`
	ClaimVersion int64 `db:"claim_version" json:"claim_version"`
}

type ReleaseConsumerInboxParams struct {
	ConsumerName       string `db:"consumer_name" json:"consumer_name"`
	EventKey           string `db:"event_key" json:"event_key"`
	LastError          string `db:"last_error" json:"last_error"`
	ReservationVersion int64  `db:"reservation_version" json:"reservation_version"`
}

type ReserveConsumerInboxParams struct {
	ConsumerName  string `db:"consumer_name" json:"consumer_name"`
	EventKey      string `db:"event_key" json:"event_key"`
	Topic         string `db:"topic" json:"topic"`
	PartitionID   int32  `db:"partition_id" json:"partition_id"`
	MessageOffset int64  `db:"message_offset" json:"message_offset"`
}

type ReserveConsumerInboxRow struct {
	Reserved           bool  `db:"reserved" json:"reserved"`
	Processed          bool  `db:"processed" json:"processed"`
	ReservationVersion int64 `db:"reservation_version" json:"reservation_version"`
}

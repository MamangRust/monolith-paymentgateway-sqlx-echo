package errors

import (
	"encoding/json"
	"errors"
	"strings"
	"testing"
)

type sqlStateError struct {
	state string
}

func (e sqlStateError) Error() string    { return "database detail: " + e.state }
func (e sqlStateError) SQLState() string { return e.state }

func TestErrConstraintOrFailedMapsUniqueViolationToConflict(t *testing.T) {
	err := ErrConstraintOrFailed(sqlStateError{state: "23505"}, "Merchant", "create merchant")

	if err.Code != 409 {
		t.Fatalf("expected HTTP 409, got %d", err.Code)
	}
	if err.Type != ErrorTypeConflict {
		t.Fatalf("expected conflict type, got %s", err.Type)
	}
	if err.Message != "Merchant already exists" {
		t.Fatalf("unexpected public message: %q", err.Message)
	}
	if err.Internal == nil {
		t.Fatal("expected original database error to be retained internally")
	}
	if err.Error() == (sqlStateError{state: "23505"}).Error() {
		t.Fatal("database detail leaked into public error")
	}
}

func TestErrConstraintOrFailedKeepsNonConstraintErrorsInternal(t *testing.T) {
	err := ErrConstraintOrFailed(errors.New("database unavailable"), "Merchant", "create merchant")

	if err.Code != 500 {
		t.Fatalf("expected HTTP 500, got %d", err.Code)
	}
	if err.Message == "database unavailable" {
		t.Fatal("database detail leaked into public message")
	}
}

func TestAppErrorJSONDoesNotLeakInternalOrCode(t *testing.T) {
	// The client-visible payload must never carry the internal database error
	// (SQL detail, SQLSTATE) nor the raw HTTP code field.
	appErr := ErrConstraintOrFailed(sqlStateError{state: "23505"}, "Merchant", "create merchant")
	appErr.Internal = errors.New("SQLSTATE 23505: duplicate key value violates unique constraint \"merchants_api_key_key\"")

	b, err := json.Marshal(appErr)
	if err != nil {
		t.Fatalf("marshal app error: %v", err)
	}
	s := string(b)

	if strings.Contains(s, "SQLSTATE") || strings.Contains(s, "23505") || strings.Contains(s, "duplicate key") {
		t.Fatalf("internal database detail leaked into JSON response: %s", s)
	}
	if strings.Contains(s, `"code"`) || strings.Contains(s, `"internal"`) {
		t.Fatalf("internal fields serialized into JSON response: %s", s)
	}
	if !strings.Contains(s, `"message":"Merchant already exists"`) {
		t.Fatalf("public message missing from JSON response: %s", s)
	}
	if !strings.Contains(s, `"type":`) {
		t.Fatalf("error type missing from JSON response: %s", s)
	}
}

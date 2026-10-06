package outbox

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"errors"
	"strings"
	"testing"
)

type fakeExecutor struct {
	calls int
	query string
	args  []any
}

func (f *fakeExecutor) ExecContext(_ context.Context, query string, args ...any) (sql.Result, error) {
	f.calls++
	f.query = query
	f.args = args
	return driver.RowsAffected(1), nil
}

func (f *fakeExecutor) QueryRowContext(_ context.Context, query string, args ...any) *sql.Row {
	panic("fakeExecutor: real *sql.Row cannot be faked; use a driver-level fake for Reserve tests")
}

func TestReserveValidatesKeys(t *testing.T) {
	for name, tc := range map[string]struct {
		tx           InboxExecutor
		consumerName string
		eventKey     string
	}{
		"nil executor": {tx: nil, consumerName: "email-service-group", eventKey: "topic:evt-1"},
		"empty consumer": {tx: &fakeExecutor{}, consumerName: "", eventKey: "topic:evt-1"},
		"empty event key": {tx: &fakeExecutor{}, consumerName: "email-service-group", eventKey: ""},
	} {
		t.Run(name, func(t *testing.T) {
			if _, _, _, err := Reserve(context.Background(), tc.tx, tc.consumerName, tc.eventKey, "topic", 0, 1); !errors.Is(err, ErrInvalidInboxKey) {
				t.Fatalf("expected ErrInvalidInboxKey, got %v", err)
			}
		})
	}
}

func TestMarkProcessedValidatesKeys(t *testing.T) {
	if err := MarkProcessed(context.Background(), nil, "email-service-group", "topic:evt-1", 1); !errors.Is(err, ErrInvalidInboxKey) {
		t.Fatalf("expected ErrInvalidInboxKey, got %v", err)
	}
	if err := MarkProcessed(context.Background(), &fakeExecutor{}, "", "topic:evt-1", 1); !errors.Is(err, ErrInvalidInboxKey) {
		t.Fatalf("expected ErrInvalidInboxKey, got %v", err)
	}
}

func TestReleaseValidatesKeysAndRecordsError(t *testing.T) {
	if err := Release(context.Background(), nil, "email-service-group", "topic:evt-1", 1, nil); !errors.Is(err, ErrInvalidInboxKey) {
		t.Fatalf("expected ErrInvalidInboxKey, got %v", err)
	}

	exec := &fakeExecutor{}
	if err := Release(context.Background(), exec, "email-service-group", "topic:evt-1", 1, errors.New("smtp down")); err != nil {
		t.Fatalf("Release returned error: %v", err)
	}
	if exec.calls != 1 || len(exec.args) != 4 {
		t.Fatalf("expected one four-argument update, calls=%d args=%d", exec.calls, len(exec.args))
	}
	if !strings.Contains(exec.args[2].(string), "smtp down") {
		t.Fatalf("expected last_error to carry the processing error, got %q", exec.args[2])
	}
}

// TestReserveUsesLeaseFencedUpsert pins the reserve SQL to the lease-fenced
// upsert contract. The reserve path itself needs a real driver-backed row, so
// it is exercised by the integration suite; here the SQL constant is asserted.
func TestReserveUsesLeaseFencedUpsert(t *testing.T) {
	for _, needle := range []string{"consumer_inbox", "ON CONFLICT", "lease_until <= current_timestamp"} {
		if !strings.Contains(reserveConsumerInbox, needle) {
			t.Fatalf("reserve query does not contain %q: %s", needle, reserveConsumerInbox)
		}
	}
}

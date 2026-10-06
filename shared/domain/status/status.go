// Package status defines the canonical lifecycle (state machine) of financial
// records across the payment gateway domains. Every status transition that a
// client-facing command performs is validated here, so an invalid move (for
// example flipping a settled topup back to pending) fails fast with a
// 400 Bad Request instead of silently corrupting the ledger.
//
// The transition tables below are the single source of truth. Service code
// must call ValidateTransition instead of re-implementing inline checks, and
// any new transition must be added here first (see README.md for the rendered
// tables and the idempotency key format).
package status

import (
	"fmt"
	"sort"
	"strings"

	sharederrors "github.com/MamangRust/monolith-payment-gateway-shared/errors"
)

// Domain identifies a financial record type with its own lifecycle.
type Domain string

const (
	// DomainTopup covers balance-loading records (topups).
	DomainTopup Domain = "topup"
	// DomainTransfer covers peer-to-peer settlements (transfers).
	DomainTransfer Domain = "transfer"
	// DomainWithdraw covers outbound settlements (withdrawals).
	DomainWithdraw Domain = "withdraw"
	// DomainTransaction covers the merchant payment ledger (transactions).
	DomainTransaction Domain = "transaction"
	// DomainCardBilling covers recurring card billing cycles. The table is
	// defined for reference; wiring into the card service is still backlog
	// (see README.md).
	DomainCardBilling Domain = "card_billing"
)

// Canonical status values shared by the financial domains. Statuses are stored
// as short lowercase strings in the database.
const (
	Pending              = "pending"               // created, not yet settled
	Success              = "success"               // settled successfully
	Failed               = "failed"                // settlement failed
	CompensationRequired = "compensation_required" // automatic compensation failed; manual reconciliation required
	Authorized           = "authorized"            // transaction: hold placed on funds
	Captured             = "captured"              // transaction: hold settled to the merchant
	Voided               = "voided"                // transaction: hold released, no settlement
	Refunded             = "refunded"              // transaction: captured amount returned to the customer
	Unpaid               = "unpaid"                // card billing: cycle generated, not yet paid
	Paid                 = "paid"                  // card billing: cycle settled
	Completed            = "completed"             // card billing: cycle closed
)

// allowedTransitions maps domain -> from-status -> set of allowed to-statuses.
//
// Same-state transitions are permitted ONLY where explicitly listed. The
// `success -> success` entry exists for idempotent re-assertion in update
// flows (re-confirming an already-success record). Lifecycle operations that
// MUTATE balances (transaction Capture/Void/Refund) deliberately have NO
// same-state entry: calling them twice must be rejected, otherwise a retry
// would double-credit the merchant, double-release a hold, or double-refund.
//
// "failed" is reachable from terminal states as a SYSTEM-ONLY compensation:
// services call the status repo directly in their markAsFailed paths; client
// requests can never request a "failed" status through these tables.
var allowedTransitions = map[Domain]map[string]map[string]bool{
	DomainTopup: {
		Pending:              {Success: true, Failed: true, CompensationRequired: true},
		Failed:               {Pending: true, Success: true, CompensationRequired: true},
		Success:              {Success: true, Failed: true, CompensationRequired: true},
		CompensationRequired: {Failed: true, Pending: true, Success: true},
	},
	DomainTransfer: {
		Pending:              {Success: true, Failed: true, CompensationRequired: true},
		Failed:               {Pending: true, Success: true, CompensationRequired: true},
		Success:              {Success: true, Failed: true, CompensationRequired: true},
		CompensationRequired: {Failed: true, Pending: true, Success: true},
	},
	DomainWithdraw: {
		Pending:              {Success: true, Failed: true, CompensationRequired: true},
		Failed:               {Pending: true, Success: true, CompensationRequired: true},
		Success:              {Success: true, Failed: true, CompensationRequired: true},
		CompensationRequired: {Failed: true, Pending: true, Success: true},
	},
	DomainTransaction: {
		Pending:              {Authorized: true, Voided: true, Success: true, Failed: true, CompensationRequired: true},
		Authorized:           {Captured: true, Voided: true, Failed: true, CompensationRequired: true},
		Captured:             {Refunded: true, Failed: true, CompensationRequired: true},
		Success:              {Success: true, Failed: true, CompensationRequired: true},
		Failed:               {Pending: true, Success: true, CompensationRequired: true},
		CompensationRequired: {Failed: true, Pending: true, Success: true},
	},
	DomainCardBilling: {
		Pending:   {Unpaid: true, Paid: true, Completed: true, Failed: true},
		Unpaid:    {Paid: true, Completed: true, Failed: true},
		Paid:      {Completed: true},
		Completed: {Failed: true}, // reopened for correction
		Failed:    {Unpaid: true, Paid: true},
	},
}

// CanTransition reports whether moving a record of the given domain from one
// status to another is a valid lifecycle transition. Same-state moves are only
// allowed when the domain table lists them explicitly (see allowedTransitions).
func CanTransition(domain Domain, from, to string) bool {
	byFrom, ok := allowedTransitions[domain]
	if !ok {
		return false
	}
	toSet, ok := byFrom[from]
	if !ok {
		return false
	}
	return toSet[to]
}

// AllowedTransitions returns the sorted list of statuses reachable from the
// given state for a domain (excluding the always-allowed same-state move).
func AllowedTransitions(domain Domain, from string) []string {
	toSet := allowedTransitions[domain][from]
	out := make([]string, 0, len(toSet))
	for s := range toSet {
		out = append(out, s)
	}
	sort.Strings(out)
	return out
}

// ValidateTransition validates a status change and returns a 400 Bad Request
// error describing the allowed transitions when the move is not permitted.
// A missing target status is rejected up front.
func ValidateTransition(domain Domain, from, to string) error {
	if to == "" {
		return sharederrors.NewBadRequestError(fmt.Sprintf("status is required for %s status transition", domain))
	}
	if CanTransition(domain, from, to) {
		return nil
	}
	msg := fmt.Sprintf("invalid status transition for %s: %q -> %q", domain, from, to)
	if allowed := AllowedTransitions(domain, from); len(allowed) > 0 {
		msg += fmt.Sprintf(" (allowed from %q: %s)", from, strings.Join(allowed, ", "))
	}
	return sharederrors.NewBadRequestError(msg)
}

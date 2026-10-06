package kafka

import "testing"

func TestComputeRiskScoreUsesAuditScoringRules(t *testing.T) {
	tests := []struct {
		name   string
		amount int64
		mcc    string
		want   int
	}{
		{name: "low amount and unknown mcc", amount: 1_000_000, mcc: "0000", want: 0},
		{name: "medium amount", amount: 1_000_001, mcc: "0000", want: 5},
		{name: "large amount", amount: 5_000_001, mcc: "0000", want: 15},
		{name: "high amount and gambling", amount: 10_000_001, mcc: "7995", want: 70},
		{name: "crypto mcc", amount: 10_000_001, mcc: "6051", want: 65},
		{name: "money transfer mcc", amount: 10_000_001, mcc: "4829", want: 50},
		{name: "bars mcc", amount: 10_000_001, mcc: "5813", want: 45},
		{name: "pawn shop mcc", amount: 10_000_001, mcc: "5933", want: 45},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := computeRiskScore(tt.amount, tt.mcc); got != tt.want {
				t.Fatalf("computeRiskScore(%d, %q) = %d, want %d", tt.amount, tt.mcc, got, tt.want)
			}
		})
	}
}

func TestHighRiskThresholdIncludesScore70(t *testing.T) {
	if score := computeRiskScore(10_000_001, "7995"); score < 70 {
		t.Fatalf("expected audit maximum score to reach high-risk threshold, got %d", score)
	}
}

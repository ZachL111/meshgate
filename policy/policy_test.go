package policy

import "testing"

func TestFixtureDecisions(t *testing.T) {
	tests := []struct {
		name         string
		signal       Signal
		wantScore    int
		wantDecision string
	}{
		{name: "case_1", signal: Signal{Demand: 55, Capacity: 70, Latency: 20, Risk: 5, Weight: 12}, wantScore: 168, wantDecision: "accept"},
		{name: "case_2", signal: Signal{Demand: 74, Capacity: 86, Latency: 27, Risk: 18, Weight: 11}, wantScore: 152, wantDecision: "review"},
		{name: "case_3", signal: Signal{Demand: 67, Capacity: 95, Latency: 16, Risk: 8, Weight: 4}, wantScore: 181, wantDecision: "accept"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := Score(tc.signal); got != tc.wantScore {
				t.Fatalf("score = %d, want %d", got, tc.wantScore)
			}
			if got := Classify(tc.signal); got != tc.wantDecision {
				t.Fatalf("decision = %s, want %s", got, tc.wantDecision)
			}
		})
	}
}

package policy

import "testing"

func TestFixtureDecisions(t *testing.T) {
	signal := Signal{Demand: 55, Capacity: 70, Latency: 20, Risk: 5, Weight: 12}
	if got := Score(signal); got != 168 { t.Fatalf("score = %d", got) }
	if got := Classify(signal); got != "accept" { t.Fatalf("decision = %s", got) }
	signal := Signal{Demand: 74, Capacity: 86, Latency: 27, Risk: 18, Weight: 11}
	if got := Score(signal); got != 152 { t.Fatalf("score = %d", got) }
	if got := Classify(signal); got != "review" { t.Fatalf("decision = %s", got) }
	signal := Signal{Demand: 67, Capacity: 95, Latency: 16, Risk: 8, Weight: 4}
	if got := Score(signal); got != 181 { t.Fatalf("score = %d", got) }
	if got := Classify(signal); got != "accept" { t.Fatalf("decision = %s", got) }
}

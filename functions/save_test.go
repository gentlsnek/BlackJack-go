// go
package functions

import (
	"reflect"
	"sort"
	"testing"
)

// Test Deal function
func TestDeal(t *testing.T) {
	deck := []string{"A", "B", "C", "D", "E"}

	// Deal 3 cards
	hand, remaining := Deal(deck, 3)
	if len(hand) != 3 {
		t.Errorf("Expected hand of 3 cards, got %d", len(hand))
	}
	if len(remaining) != 2 {
		t.Errorf("Expected 2 cards remaining, got %d", len(remaining))
	}
	expectedHand := []string{"A", "B", "C"}
	if !reflect.DeepEqual(hand, expectedHand) {
		t.Errorf("Expected hand %v, got %v", expectedHand, hand)
	}

	// Deal more cards than deck size
	hand, remaining = Deal(deck, 10)
	if len(hand) != 5 {
		t.Errorf("Expected hand of 5 cards, got %d", len(hand))
	}
	if len(remaining) != 0 {
		t.Errorf("Expected 0 cards remaining, got %d", len(remaining))
	}

	// Deal zero cards
	hand, remaining = Deal(deck, 0)
	if len(hand) != 0 {
		t.Errorf("Expected hand of 0 cards, got %d", len(hand))
	}
	if len(remaining) != 5 {
		t.Errorf("Expected 5 cards remaining, got %d", len(remaining))
	}
}

// Test Shuffle function
func TestShuffle(t *testing.T) {
	original := []string{"A", "B", "C", "D", "E"}
	deck := make([]string, len(original))
	copy(deck, original)

	Shuffle(deck)

	// Check that deck has same elements as original
	origSorted := append([]string{}, original...)
	deckSorted := append([]string{}, deck...)
	sort.Strings(origSorted)
	sort.Strings(deckSorted)
	if !reflect.DeepEqual(origSorted, deckSorted) {
		t.Errorf("Shuffled deck has different elements: %v vs %v", deck, original)
	}

	// Check that order is changed (statistically, not always, so warn not fail)
	sameOrder := reflect.DeepEqual(deck, original)
	if sameOrder {
		t.Log("Warning: deck order did not change after shuffle (possible but unlikely)")
	}
}

// Note: Load() requires user input and file I/O, so is not tested here.
// To test Load, you would need to refactor it to accept dependencies or use interfaces for input/output.
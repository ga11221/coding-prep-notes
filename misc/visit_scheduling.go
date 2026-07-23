package main

import (
	"fmt"
	"math"
	"sort"
)

// ── Step 0: Restate + confirm ──────────────────────────────────────
// "So we have a list of booked visits — each has a patient ID, start,
// end. We need to place one new urgent visit of duration D somewhere
// in the existing schedule. There's a travel_time(a,b) function I can
// call. I'm looking for all valid start time windows. A few things I
// want to confirm:
//   - The booked visits are non-overlapping? Or should I handle that?
//   - The clinician starts at the first visit's location at its start
//     time — so no time before the first visit is available?
//   - Day ends after the last visit — no hard stop, so after-last is
//     open-ended?
//   - travel_time is symmetric? I'll assume yes, but I'll call it in
//     the right direction anyway."
//
// → We'll assume booked visits are valid and non-overlapping, and
//   that the day boundaries are as stated.

type Visit struct {
	PatientID string
	StartMin  int // minutes from midnight
	EndMin    int
}

type TimeWindow struct {
	StartMin int
	EndMin   int // math.MaxInt32 = unbounded (after last visit)
}

// ── Step 2: Outline approach as comments (talk through) ─────────────
// "Here's how I'm thinking about this:
//
//  1. If no booked visits → any start time works.
//  2. Sort by start time (don't trust the input to be ordered).
//  3. Before first visit: we agreed no slot — clinician starts at
//     first visit's location at its start time.
//  4. Between consecutive visits: for each gap, see if
//     travel(prev→new) + D + travel(new→next) fits.
//     If yes, valid start times are:
//       earliest = prev.EndMin + travel(prev→new)
//       latest   = next.StartMin - D - travel(new→next)
//  5. After last visit: start from last.EndMin + travel(last→new),
//     unbounded.
//
// Travel time is a function I'm calling — don't need to implement.
// Does that approach sound right?"

func FindValidSlots(booked []Visit, newPatient string, newDuration int, travel func(a, b string) int) []TimeWindow {
	// Step 1: empty case.
	if len(booked) == 0 {
		return []TimeWindow{{StartMin: 0, EndMin: math.MaxInt32}}
	}

	// Step 2: sort. Don't want to mutate the caller's slice.
	sorted := sortedByStart(booked)
	var windows []TimeWindow

	// Step 4: gaps between consecutive visits.
	for i := 0; i < len(sorted)-1; i++ {
		a, b := sorted[i], sorted[i+1]
		gap := b.StartMin - a.EndMin
		needed := travel(a.PatientID, newPatient) + newDuration + travel(newPatient, b.PatientID)
		if needed <= gap {
			windows = append(windows, TimeWindow{
				StartMin: a.EndMin + travel(a.PatientID, newPatient),
				EndMin:   b.StartMin - newDuration - travel(newPatient, b.PatientID),
			})
		}
	}

	// Step 5: after the last visit — open-ended.
	last := sorted[len(sorted)-1]
	startAfter := last.EndMin + travel(last.PatientID, newPatient)
	windows = append(windows, TimeWindow{StartMin: startAfter, EndMin: math.MaxInt32})

	// Return windows in chronological order.
	sort.Slice(windows, func(i, j int) bool {
		return windows[i].StartMin < windows[j].StartMin
	})
	return windows
}

func sortedByStart(visits []Visit) []Visit {
	sorted := make([]Visit, len(visits))
	copy(sorted, visits)
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].StartMin < sorted[j].StartMin
	})
	return sorted
}

func formatTime(minutes int) string {
	return fmt.Sprintf("%02d:%02d", minutes/60, minutes%60)
}

func main() {
	// ── Step 5: Trace through the given example ──────────────────────
	// "Booked: 9:00-9:30 (A), 11:00-11:30 (B).
	// travel A→new=15, new→B=20. D=30.
	//
	// Gap between A and B: 90 min. Needed: 15+30+20 = 65. Fits.
	//   earliest = 9:30 + 15 = 9:45
	//   latest   = 11:00 - 30 - 20 = 10:10
	//
	// After B: earliest = 11:30 + 20 = 11:50, unbounded.

	booked := []Visit{
		{PatientID: "A", StartMin: 9*60 + 0, EndMin: 9*60 + 30},
		{PatientID: "B", StartMin: 11*60 + 0, EndMin: 11*60 + 30},
	}

	travel := func(a, b string) int {
		m := map[string]map[string]int{
			"A":   {"B": 25, "new": 15},
			"B":   {"A": 25, "new": 20},
			"new": {"A": 15, "B": 20},
		}
		return m[a][b]
	}

	windows := FindValidSlots(booked, "new", 30, travel)
	fmt.Println("Valid start windows:")
	for _, w := range windows {
		start := formatTime(w.StartMin)
		if w.EndMin == math.MaxInt32 {
			fmt.Printf("  %s → end of day (or unbounded)\n", start)
		} else {
			fmt.Printf("  %s → %s\n", start, formatTime(w.EndMin))
		}
	}
	fmt.Println("Expected: 09:45 → 10:10, 11:50 onward")

	// ── Step 6: Edge cases ───────────────────────────────────────────
	// "Let me also walk through edge cases:
	//
	//   a. Empty schedule → should return a full-day window.
	//      (Already handled above.)
	//
	//   b. Single booked visit → only after-last window.
	//
	//   c. Gap too small → that gap is skipped.
	//
	//   d. Visits not sorted → still correct.
	//
	//   e. No valid slot anywhere → only the after-last window, which
	//      is always available."

	empty := FindValidSlots(nil, "new", 30, travel)
	fmt.Printf("\nEmpty booked: %d window(s) — starts at 00:00\n", len(empty))

	single := []Visit{
		{PatientID: "A", StartMin: 9*60 + 0, EndMin: 9*60 + 30},
	}
	one := FindValidSlots(single, "new", 30,
		func(a, b string) int {
			if a == b {
				return 0
			}
			return 15
		})
	fmt.Printf("Single visit: from %s onward\n", formatTime(one[0].StartMin))

	unsorted := []Visit{
		{PatientID: "B", StartMin: 11*60 + 0, EndMin: 11*60 + 30},
		{PatientID: "A", StartMin: 9*60 + 0, EndMin: 9*60 + 30},
	}
	unsortedResult := FindValidSlots(unsorted, "new", 30, travel)
	fmt.Printf("Unsorted input: still returns %d window(s)\n", len(unsortedResult))

	// ── Step 7: Time/space complexity ────────────────────────────────
	// "Time: O(n log n) due to sorting. The gap loop is O(n).
	// Space: O(n) for the sorted copy and the result slice.
	// Could be O(1) extra if we sorted in-place, but this is safer."
}

// ── Step 8: Discussion if time permits ─────────────────────────────
// "A few things I'd think about if we had more time:
//   - What if we could reorder existing visits? NP-hard (TSP variant).
//   - What if travel times vary by time of day (traffic)? Precompute
//     time-dependent shortest paths, but fundamentally the same loop.
//   - What if the new visit has a preferred time range? Filter windows.
//   - What if the clinician has a depot (starts at home)? Add a
//     travel(depot→first) and travel(last→depot) constraint, and a
//     before-first slot becomes available."

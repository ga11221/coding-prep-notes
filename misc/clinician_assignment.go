package main

import (
	"fmt"
	"math"
	"sort"
)

// ── Step 0: Restate + confirm ──────────────────────────────────────
// "So we have a set of clinicians, each with their booked visits for
// the day. A new urgent visit comes in with a skill requirement and
// duration. We need to find every clinician who can accommodate it
// and show the valid start time windows for each.
//
// Questions:
//   - Can the new visit start before a clinician's first visit?
//     → Same assumption as before: clinician starts at first visit's
//       location at its start time.
//   - Does every clinician have the same start-of-day constraint?
//     → Yes, each clinician starts at their own first visit.
//   - travel_time is callable per pair?
//     → Yes, and I can assume it works for any patient pair.
//   - If a clinician has no booked visits, they have the full day?
//     → Yes, but I'll clarify the day boundaries."

type Visit struct {
	PatientID string
	StartMin  int
	EndMin    int
}

type TimeWindow struct {
	StartMin int
	EndMin   int
}

type Clinician struct {
	ID      string
	Skills  []string
	Booked  []Visit
	DayEnd  int // minutes from midnight, default 24:00
}

type Placement struct {
	ClinicianID string
	Windows     []TimeWindow
}

// ── Step 2: Approach outline (talk through) ────────────────────────
// "Approach:
//
//  1. For each clinician, check if they have the required skill.
//     If not, skip.
//
//  2. Run the interval + travel insertion check from the previous
//     problem on this clinician's booked visits. Same logic:
//     - Sort booked visits by start time.
//     - Check gaps between consecutive visits.
//     - Check after the last visit (bounded by DayEnd if given).
//     - No slot before the first visit.
//
//  3. Collect all clinicians with at least one valid window.
//
// This is the same core algorithm composed over resources — just
// filtering by skill first. We're not doing anything fancy like
// load-balancing or preference matching unless asked."

func FindPlacements(
	clinicians []Clinician,
	newPatient string,
	newDuration int,
	requiredSkill string,
	travel func(a, b string) int,
) []Placement {
	var result []Placement

	for _, c := range clinicians {
		if !hasSkill(c.Skills, requiredSkill) {
			continue
		}

		windows := findSlots(c.Booked, newPatient, newDuration, c.DayEnd, travel)
		if len(windows) > 0 {
			result = append(result, Placement{ClinicianID: c.ID, Windows: windows})
		}
	}

	return result
}

func hasSkill(skills []string, target string) bool {
	for _, s := range skills {
		if s == target {
			return true
		}
	}
	return false
}

func findSlots(booked []Visit, newPatient string, newDuration int, dayEnd int, travel func(a, b string) int) []TimeWindow {
	if len(booked) == 0 {
		return []TimeWindow{{StartMin: 0, EndMin: dayEnd}}
	}

	sorted := sortedByStart(booked)
	var windows []TimeWindow

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

	last := sorted[len(sorted)-1]
	startAfter := last.EndMin + travel(last.PatientID, newPatient)
	if startAfter <= dayEnd && (startAfter+newDuration <= dayEnd || dayEnd == math.MaxInt32) {
		windows = append(windows, TimeWindow{StartMin: startAfter, EndMin: dayEnd})
	}

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
	if minutes == math.MaxInt32 {
		return "end of day"
	}
	return fmt.Sprintf("%02d:%02d", minutes/60, minutes%60)
}

// ── Step 5: Trace through example ──────────────────────────────────

func main() {
	travel := func(a, b string) int {
		m := map[string]map[string]int{
			"A":   {"B": 20, "C": 35, "new": 15},
			"B":   {"A": 20, "C": 15, "new": 25},
			"C":   {"A": 35, "B": 15, "new": 10},
			"new": {"A": 15, "B": 25, "C": 10},
		}
		return m[a][b]
	}

	clinicians := []Clinician{
		{
			ID:     "Nurse Jackie",
			Skills: []string{"RN", "wound-care"},
			Booked: []Visit{
				{PatientID: "A", StartMin: 9*60 + 0, EndMin: 9*60 + 30},
				{PatientID: "B", StartMin: 10*60 + 30, EndMin: 11*60 + 0},
			},
			DayEnd: 17 * 60,
		},
		{
			ID:     "PT Bob",
			Skills: []string{"PT"},
			Booked: []Visit{
				{PatientID: "C", StartMin: 10*60 + 0, EndMin: 10*60 + 45},
			},
			DayEnd: 17 * 60,
		},
		{
			ID:     "CNA Diane",
			Skills: []string{"CNA"},
			Booked: []Visit{
				{PatientID: "A", StartMin: 8*60 + 0, EndMin: 8*60 + 30},
				{PatientID: "B", StartMin: 12*60 + 0, EndMin: 13*60 + 0},
				{PatientID: "C", StartMin: 14*60 + 0, EndMin: 14*60 + 30},
			},
			DayEnd: 17 * 60,
		},
	}

	// "New visit: wound-care, 30 min. Only Nurse Jackie has wound-care."
	fmt.Println("=== wound-care visit (30 min) ===")
	placements := FindPlacements(clinicians, "new", 30, "wound-care", travel)
	for _, p := range placements {
		fmt.Printf("%s:\n", p.ClinicianID)
		for _, w := range p.Windows {
			fmt.Printf("  %s → %s\n", formatTime(w.StartMin), formatTime(w.EndMin))
		}
	}
	fmt.Println("Expected: only Nurse Jackie — after B (11:25 → 17:00, no between-visit slot: 60min gap < 70min needed)")

	// ── Step 6: Edge cases ───────────────────────────────────────────
	// "a. No clinician has the required skill → empty result."
	fmt.Println("\n=== PT visit (no PT booked above — wait, Bob is PT) ===")
	ptPlacements := FindPlacements(clinicians, "new", 30, "PT", travel)
	fmt.Printf("PT Bob has %d window(s)\n", len(ptPlacements[0].Windows))

	// "b. Clinician with empty schedule → full day window."
	fmt.Println("\n=== Empty schedule ===")
	emptyClinician := []Clinician{
		{ID: "Available RN", Skills: []string{"RN"}, DayEnd: 17 * 60},
	}
	emptyPlacements := FindPlacements(emptyClinician, "new", 30, "RN", travel)
	for _, p := range emptyPlacements {
		fmt.Printf("%s: %s → %s\n", p.ClinicianID, formatTime(p.Windows[0].StartMin), formatTime(p.Windows[0].EndMin))
	}

	// "c. After-last window clipped by DayEnd."
	tightEnd := []Clinician{
		{
			ID:     "Late RN",
			Skills: []string{"RN"},
			Booked: []Visit{
				{PatientID: "A", StartMin: 16*60 + 0, EndMin: 16*60 + 30},
			},
			DayEnd: 17 * 60,
		},
	}
	latePlacements := FindPlacements(tightEnd, "new", 45, "RN", travel)
	fmt.Printf("\nAfter-last clipped: ")
	if len(latePlacements) > 0 && len(latePlacements[0].Windows) > 0 {
		w := latePlacements[0].Windows[0]
		fmt.Printf("%s → %s (visit would end at %s, day ends at 17:00)\n",
			formatTime(w.StartMin), formatTime(w.EndMin),
			formatTime(w.StartMin+45))
	} else {
		fmt.Println("no valid slot (16:30 + 15 + 45 = 17:30 > 17:00 day end)")
	}

	// ── Step 7: Complexity ───────────────────────────────────────────
	// "Time: O(C * N log N) where C = clinicians, N = max booked per
	// clinician. Sorting per clinician dominates. Skill check is linear.
	// Space: O(N) per clinician for the sorted copy, plus result."
}

// ── Step 8: Discussion ─────────────────────────────────────────────
// "If I had more time:
//   - What if we need to pick the BEST clinician, not just all valid?
//     → Add scoring: earliest completion, least disruption, shortest
//       total travel, most even load distribution.
//   - What if clinicians have overlapping skills and we need to
//     balance? → Turns into an optimization problem. Greedy works
//     for feasibility but not optimality without more structure.
//   - What about multi-visit insertion? Same loop but consider each
//     new visit independently or batch them (NP-hard in general).
//   - What if travel times depend on clinician (different modes of
//     transport)? Pass travel as a method on Clinician or a 3-arg
//     function travel(clinician, a, b)."

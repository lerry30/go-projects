package timelogs

import (
	"time"
)

func ResolveNightShiftSequence(times []time.Time) []time.Time {
	seq := make([]time.Time, len(times))
	day := 1
	for i, tm := range times {
		resolved := tm.AddDate(0, 0, day)
		if i > 0 && !seq[i].After(seq[i-1]) {
			day++
			resolved = tm.AddDate(0, 0, day)
		}
		seq[i] = resolved
	}
	return seq
}
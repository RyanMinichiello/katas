package thief

func Steal(locker []string) int {
	var reps int

	for _, row := range locker {
		handAt := -1
		for i, character := range row {
			if character == '.' {
				continue
			}
			if character == '*' {
				if handAt == i {
					// already stole it
					continue
				}
				if i+1 < len(row) && row[i+1] == '*' {
					// double it up
					handAt = i+1
					reps = reps + 1
					continue
				}
				reps = reps + 1
			}
		}

	}
	return reps
}
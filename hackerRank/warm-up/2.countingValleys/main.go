package main

func main() {

}

func countingValleys(steps int32, path string) int32 {

	altitude := 0
	valleys := 0

	for _, step := range path {
		if step == 'U' {
			altitude++
		} else if step == 'D' {
			altitude--
		}

		// Check if we just came up to sea level
		if step == 'U' && altitude == 0 {
			valleys++
		}
	}

	return int32(valleys)

}

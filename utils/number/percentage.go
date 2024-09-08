package numberutils

func PercentageCalculator(total int32, percentage int32) int32 {
	if percentage > 100 {
		return 0
	}
	// Convert percentage to float32 to avoid integer division truncation
	return int32(float32(total) * (float32(percentage) / 100))
}

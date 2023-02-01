package maxestimate

func GetMaxEstimate(reps, weight float32) float32 {
	return weight / (1.0278 - 0.0278*reps)
}

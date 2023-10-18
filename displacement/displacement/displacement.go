package dispacement

func GenDisplaceFn(a, v0, s0 float64) func(t float64) float64 {
	return func(t float64) float64 {
		return (0.5*a*t+v0)*t + s0
	}
}
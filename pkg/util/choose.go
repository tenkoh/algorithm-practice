package util

func ChooseInt(a *int, b int, mode string) {
	switch mode {
	case "min":
		if *a > b {
			*a = b
		}
	case "max":
		if *a < b {
			*a = b
		}
	}
}

func ChooseFloat32(a *float32, b float32, mode string) {
	switch mode {
	case "min":
		if *a > b {
			*a = b
		}
	case "max":
		if *a < b {
			*a = b
		}
	}
}

func ChooseFloat64(a *float64, b float64, mode string) {
	switch mode {
	case "min":
		if *a > b {
			*a = b
		}
	case "max":
		if *a < b {
			*a = b
		}
	}
}

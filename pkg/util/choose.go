package util

func Chmin(p interface{}, v interface{}) {
	switch v.(type) {
	case int:
		a, ok := p.(*int)
		if !ok {
			return
		}
		if vv := v.(int); *a > vv {
			*a = vv
		}
	case float32:
		a, ok := p.(*float32)
		if !ok {
			return
		}
		if vv := v.(float32); *a > vv {
			*a = vv
		}
	case float64:
		a, ok := p.(*float64)
		if !ok {
			return
		}
		if vv := v.(float64); *a > vv {
			*a = vv
		}
	}
}

func Chmax(p interface{}, v interface{}) {
	switch v.(type) {
	case int:
		a, ok := p.(*int)
		if !ok {
			return
		}
		if vv := v.(int); *a < vv {
			*a = vv
		}
	case float32:
		a, ok := p.(*float32)
		if !ok {
			return
		}
		if vv := v.(float32); *a < vv {
			*a = vv
		}
	case float64:
		a, ok := p.(*float64)
		if !ok {
			return
		}
		if vv := v.(float64); *a < vv {
			*a = vv
		}
	}
}

package mapdiff

import "fmt"

// Result encapsulates the result of a comparison. It stores the result as well as the diff of the two objects.
type Result struct {
	Equal bool
	Diff  string
}

// Compare compares two json like objects and determines whether or not they are equal, and created a diff of both of
// them. As of right now objects being compared must have string keys and cannot be nested.
func Compare(x, y map[string]interface{}) Result {
	// TODO: Allow nested objects
	r := Result{
		Equal: true,
		Diff:  "",
	}

	for k, v := range x {
		if yval, yok := y[k]; !yok {
			r.Equal = false
			switch v.(type) {
			case string, rune:
				r.Diff += fmt.Sprintf("+%q: %q\n", k, v)
			default:
				r.Diff += fmt.Sprintf("+%q: %v\n", k, v)
			}
		} else if v != yval {
			r.Equal = false
			switch v.(type) {
			case string, rune:
				r.Diff += fmt.Sprintf("+%q: %v\n-%q: %q\n", k, v, k, yval)
			default:
				r.Diff += fmt.Sprintf("+%q: %v\n-%q: %v\n", k, v, k, yval)
			}
		} else {
			switch v.(type) {
			case string, rune:
				r.Diff += fmt.Sprintf("%q: %q\n", k, v)
			default:
				r.Diff += fmt.Sprintf("%q: %v\n", k, v)
			}
		}
	}

	for k, v := range y {
		if _, xok := x[k]; !xok {
			r.Equal = false
			switch v.(type) {
			case string, rune:
				r.Diff += fmt.Sprintf("-%q: %q\n", k, v)
			default:
				r.Diff += fmt.Sprintf("-%q: %v\n", k, v)
			}
		}
	}

	return r
}

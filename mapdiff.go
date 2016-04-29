package mapdiff

import "fmt"

type result struct {
	Equal bool
	Diff  string
}

func Compare(x, y map[string]interface{}) result {
	r := result{
		Equal: true,
		Diff:  "",
	}

	for k, v := range x {
		if yval, yok := y[k]; !yok {
			r.Equal = false
			r.Diff += fmt.Sprintf("+ %q: %v\n", k, v)
		} else if v != yval {
			r.Equal = false
			r.Diff += fmt.Sprintf("+ %q: %v\n- %v: %v", k, v, k, yval)
		} else {
			switch v.(type) {
			case string:
				r.Diff += fmt.Sprintf("%q: %q\n", k, v)
			default:
				r.Diff += fmt.Sprintf("%q: %v\n", k, v)
			}

		}
	}

	return r
}

// func makeDiffEntry(xk, yk string, xv, yv interface{}) (string, error) {
// 	e := fmt.Sprintf("+ %v: %v\n", xk, xv)
//
// 	return fmt.Sprintf("+ %v: %v\n")
// }

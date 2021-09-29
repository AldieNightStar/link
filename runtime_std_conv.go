package link

import "fmt"

func std_init_conv(r *RuntimeMemScope) {
	r.SetFunc("str", rt_std_str)
}

func rt_std_str(v []*Variable, r *RuntimeMemScope) *Variable {
	if len(v) < 1 {
		return NewVariable("")
	}
	f := fmt.Sprintf("%f", v[0].ReadNumber())
	return NewVariable(f)
}

package link

import "math"

func std_init_math(r *RuntimeMemScope) {
	r.SetFunc("add", rt_std_add)
	r.SetFunc("sub", rt_std_sub)
	r.SetFunc("div", rt_std_div)
	r.SetFunc("mul", rt_std_mul)
	r.SetFunc("mod", rt_std_mod)
	r.SetFunc("sqr", rt_std_sqr)
}

func rt_std_add(v []*Variable, r *RuntimeMemScope) *Variable {
	if len(v) < 2 {
		return nil
	}
	n1 := v[0].ReadNumber()
	n2 := v[1].ReadNumber()
	return NewVariable(n1 + n2)
}

func rt_std_sub(v []*Variable, r *RuntimeMemScope) *Variable {
	if len(v) < 2 {
		return nil
	}
	n1 := v[0].ReadNumber()
	n2 := v[1].ReadNumber()
	return NewVariable(n1 - n2)
}

func rt_std_div(v []*Variable, r *RuntimeMemScope) *Variable {
	if len(v) < 2 {
		return nil
	}
	n1 := v[0].ReadNumber()
	n2 := v[1].ReadNumber()
	return NewVariable(n1 / n2)
}

func rt_std_mul(v []*Variable, r *RuntimeMemScope) *Variable {
	if len(v) < 2 {
		return nil
	}
	n1 := v[0].ReadNumber()
	n2 := v[1].ReadNumber()
	return NewVariable(n1 * n2)
}

func rt_std_mod(v []*Variable, r *RuntimeMemScope) *Variable {
	if len(v) < 2 {
		return nil
	}
	n1 := int(v[0].ReadNumber())
	n2 := int(v[1].ReadNumber())
	return NewVariable(n1 % n2)
}

func rt_std_sqr(v []*Variable, r *RuntimeMemScope) *Variable {
	if len(v) < 2 {
		return nil
	}
	n1 := v[0].ReadNumber()
	return NewVariable(math.Sqrt(n1))
}

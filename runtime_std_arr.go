package link

func std_init_arr(r *RuntimeMemScope) {
	r.SetFunc("arr", rt_std_arr)
	r.SetFunc("arr-is", rt_std_is_arr)
	r.SetFunc("arr-get", rt_std_arr_get)
	r.SetFunc("arr-len", rt_std_arr_len)
}

func rt_std_arr(v []*Variable, r *RuntimeMemScope) *Variable {
	return NewVariable(v)
}

func rt_std_is_arr(v []*Variable, r *RuntimeMemScope) *Variable {
	if len(v) < 1 {
		return NewVariable(float64(0))
	}
	if _, ok := v[0].Value.([]*Variable); ok {
		return NewVariable(true)
	}
	return NewVariable(false)
}

func rt_std_arr_get(v []*Variable, r *RuntimeMemScope) *Variable {
	if len(v) < 2 {
		return NewVariable(float64(0))
	}
	if arr, ok := v[0].Value.([]*Variable); ok {
		elem := int(v[1].ReadNumber())
		if elem < 0 {
			elem = len(arr) - 1 + elem
			if elem < 0 {
				elem = 0
			}
		}
		if elem >= len(arr) {
			return nil
		}
		return arr[elem]
	}
	return nil
}

func rt_std_arr_len(v []*Variable, r *RuntimeMemScope) *Variable {
	if len(v) < 1 {
		return NewVariable(float64(0))
	}
	if arr, ok := v[0].Value.([]*Variable); ok {
		return NewVariable(float64(len(arr)))
	}
	return NewVariable(float64(0))
}

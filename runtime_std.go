package link

func std_init(r *RuntimeMemScope) {
	std_init_math(r)
	std_init_arr(r)
	std_init_conv(r)
}

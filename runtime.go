package link

import parser "github.com/AldieNightStar/goparser"

func run_code(cmds []*TokenCommand, r *RuntimeMemScope) *Variable {
	for _, cmd := range cmds {
		v := run_command(cmd, r)
		if v == nil {
			continue
		}
		if r := v.ReadReturnSpec(); r != nil {
			return r
		}
	}
	return nil
}

func run_command(cmd *TokenCommand, r *RuntimeMemScope) *Variable {
	val := r.GetVar(cmd.Name)
	if f, ok := val.(RuntimeFunc); ok {
		return f(run_args(cmd.Args, r), r)
	}
	return NewVariable(nil)
}

func run_arg(arg interface{}, r *RuntimeMemScope) *Variable {
	if t, ok := arg.(*TokenVariable); ok {
		return NewVariable(r.GetVar(t.Value))
	} else if t, ok := arg.(*parser.NumberToken); ok {
		return NewVariable(t.Value)
	} else if t, ok := arg.(*parser.StringToken); ok {
		return NewVariable(t.Value)
	} else if t, ok := arg.(*TokenCommand); ok {
		return run_command(t, r)
	} else if t, ok := arg.(*TokenCodeBlock); ok {
		return NewVariable(t)
	} else {
		return NewVariable(nil)
	}
}

func run_args(args []interface{}, r *RuntimeMemScope) []*Variable {
	arr := make([]*Variable, 0, 32)
	for _, arg := range args {
		arr = append(arr, run_arg(arg, r))
	}
	return arr
}

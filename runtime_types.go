package link

type Variable struct {
	Value interface{}
}

type SpecialSym struct {
	Value string
	Arg   interface{}
}

func NewVariable(val interface{}) *Variable {
	return &Variable{Value: val}
}

func NewVariableSpec(val string, arg interface{}) *Variable {
	return NewVariable(&SpecialSym{Value: val, Arg: arg})
}

func (v *Variable) ReadNumber() float64 {
	if n, ok := v.Value.(float64); ok {
		return n
	}
	return 0
}

func (v *Variable) ReadString() string {
	if s, ok := v.Value.(string); ok {
		return s
	} else if s, ok := v.Value.(*string); ok {
		return *s
	}
	return ""
}

func (v *Variable) ReadSpec() *SpecialSym {
	if s, ok := v.Value.(*SpecialSym); ok {
		return s
	}
	return nil
}

func (v *Variable) ReadReturnSpec() *Variable {
	if s := v.ReadSpec(); s != nil && s.Value == "return" {
		if v, ok := s.Arg.(*Variable); ok {
			return v
		}
	}
	return nil
}

func (v *Variable) SetReturnSpec(val *Variable) {
	v.Value = &SpecialSym{
		Value: "return",
		Arg:   val,
	}
}

func (v *Variable) ReadFunc() *RuntimeFunc {
	if f, ok := v.Value.(*RuntimeFunc); ok {
		return f
	}
	return nil
}

func (v *Variable) IsNil() bool {
	return v.Value == nil
}

func (v *Variable) ReadBool() bool {
	if b, ok := v.Value.(bool); ok {
		return b
	} else if b, ok := v.Value.(*bool); ok {
		return *b
	}
	return false
}

// =======================
// =======================
// =======================

type RuntimeMemScope struct {
	Mem    map[string]interface{}
	Parent *RuntimeMemScope
}

func NewRuntimeMemScope(parent *RuntimeMemScope) *RuntimeMemScope {
	r := &RuntimeMemScope{
		Mem:    make(map[string]interface{}),
		Parent: parent,
	}
	std_init(r)
	return r
}

func (r *RuntimeMemScope) GetVar(name string) interface{} {
	if m, ok := r.Mem[name]; ok {
		return m
	} else if r.Parent != nil {
		return r.Parent.GetVar(name)
	} else {
		return nil
	}
}

func (r *RuntimeMemScope) SetVar(name string, val interface{}) {
	// If parent has var - change it
	// If parent hasn't var - create local
	if r.Parent != nil {
		if _, ok := r.Parent.Mem[name]; ok {
			r.Parent.Mem[name] = val
			return
		}
	}
	r.Mem[name] = val
}

func (r *RuntimeMemScope) SetFunc(name string, f RuntimeFunc) {
	r.Mem[name] = f
}

// =======================
// =======================
// =======================

type RuntimeFunc func([]*Variable, *RuntimeMemScope) *Variable

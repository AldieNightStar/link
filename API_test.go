package link

import "testing"

func TestParse(t *testing.T) {
	src := "[print a][set b 22][count [add a 1]]"
	code := ParseCode(src)

	c1 := code[0]
	c2 := code[1]
	c3 := code[2]
	if c1.Name != "print" || c2.Name != "set" || c3.Name != "count" {
		t.Fatal("Commands not right")
	}
	if len(c1.Args) != 1 || len(c2.Args) != 2 || len(c3.Args) != 1 {
		t.Fatal(
			"Commands args len is not right: ",
			len(c1.Args),
			len(c2.Args),
			len(c3.Args),
		)
	}
}

func TestParseAndRunCode(t *testing.T) {
	src := "[set 'a' [add a 25]][set 'a' [add a -5]]"
	code := ParseCode(src)
	r := NewRuntimeMemScope(nil)

	r.SetFunc("add", func(v []*Variable, rms *RuntimeMemScope) *Variable {
		n1 := v[0].ReadNumber()
		n2 := v[1].ReadNumber()
		return NewVariable(n1 + n2)
	})

	r.SetFunc("set", func(v []*Variable, rms *RuntimeMemScope) *Variable {
		n := v[0].ReadString()
		rms.SetVar(n, v[1].Value)
		return nil
	})

	r.Mem["a"] = float64(10)

	RunCode(code, r)

	if r.Mem["a"] != float64(30) {
		t.Fatal("Number is not valid.", r.Mem["a"])
	}
}

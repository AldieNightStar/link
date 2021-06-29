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
	src := "[print a][print b][print [add a 1]]"
	code := ParseCode(src)
	r := NewRuntimeMemScope(nil)

	pcnt := 0
	var res float64

	r.SetFunc("print", func(v []*Variable, rms *RuntimeMemScope) *Variable {
		for _, s := range v {
			println(s.ReadString())
		}
		pcnt += 1
		return nil
	})

	r.SetFunc("add", func(v []*Variable, rms *RuntimeMemScope) *Variable {
		n1 := v[0].ReadNumber()
		n2 := v[1].ReadNumber()
		res = n1 + n2
		return NewVariable(n1 + n2)
	})

	r.Mem["a"] = float64(63)

	RunCode(code, r)

	if pcnt != 3 {
		t.Fatal("Prints count is not 3.", pcnt)
	}
	if res != 64 {
		t.Fatal("Result of adding is not 64.", res)
	}
}

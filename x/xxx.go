package main

import (
	"io/ioutil"
	"os"
	"strconv"

	"github.com/AldieNightStar/link"
)

func main() {
	f, _ := os.Open("code.txt")
	b, _ := ioutil.ReadAll(f)

	code := link.ParseCode(string(b))

	r := link.NewRuntimeMemScope(nil)

	r.SetFunc("print", func(v []*link.Variable, rms *link.RuntimeMemScope) *link.Variable {
		println(v[0].ReadString())
		return nil
	})

	r.SetFunc("rep", func(v []*link.Variable, rms *link.RuntimeMemScope) *link.Variable {
		name := v[0].ReadString()
		times := int(v[1].ReadNumber())
		code := v[2].Value.(*link.TokenCodeBlock)
		r2 := link.NewRuntimeMemScope(r)
		for i := 0; i < times; i++ {
			r2.Mem[name] = float64(i)
			link.RunCode(code.Commands, r2)
		}
		return nil
	})

	r.SetFunc("str", func(v []*link.Variable, rms *link.RuntimeMemScope) *link.Variable {
		return link.NewVariable(strconv.Itoa(int(v[0].ReadNumber())))
	})

	link.RunCode(code, r)
}

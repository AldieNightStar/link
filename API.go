package link

import _ "github.com/AldieNightStar/goparser"

func ParseCode(str string) []*TokenCommand {
	code := parser_code(str)
	return code
}

func RunCode(cmds []*TokenCommand, r *RuntimeMemScope) *Variable {
	return run_code(cmds, r)
}

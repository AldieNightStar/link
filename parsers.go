package link

import (
	"strings"

	parser "github.com/AldieNightStar/goparser"
)

func parser_code(s string) []*TokenCommand {
	iter := parser.Parse(s, []parser.Parser{
		parser_command,
	})
	arr := make([]*TokenCommand, 0, 32)
	for {
		res, _ := iter()

		if res == nil {
			break
		}

		if cmd, ok := res.Token.(*TokenCommand); ok {
			arr = append(arr, cmd)
		}
	}
	return arr
}

func parser_command(s string) *parser.Result {
	if s[0] != '[' {
		return nil
	}

	name, _ := parser.UntilOf(s[1:], []string{" ", "\t", "\r", "\n"})

	if len(name) < 1 {
		return nil
	}

	cnt := 2 + len(name)
	arr := make([]interface{}, 0, 32)

	iter := parser.Parse(s[cnt:], []parser.Parser{
		parser_argument,
		parser_end,
	})

	for {
		res, _ := iter()

		if t, ok := res.Token.(*tokEnd); ok {
			if t.Value == "]" {
				cnt += res.Count
				break
			}
		}
		if _, ok := res.Token.(*parser.UnknownToken); ok {
			continue
		}

		cnt += res.Count
		arr = append(arr, res.Token)
	}

	return &parser.Result{
		Token: &TokenCommand{
			Name: name,
			Args: arr,
		},
		Count: cnt,
	}
}

func parser_argument(s string) *parser.Result {
	return parser.ParseOne(s, []parser.Parser{
		parser.StringParser,
		parser.NumberParser,
		parser_command,
		parser_code_block,
		parser_variable,
	})
}

func parser_variable(s string) *parser.Result {
	sb := &strings.Builder{}
	for i := 0; i < len(s); i++ {
		c := s[i]

		if i == 0 {
			if strings.ContainsRune("0123456789.", rune(c)) {
				break
			}
		}

		if c == ' ' || c == '\t' || c == '\n' {
			break
		}
		if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (strings.ContainsRune("01234567890_-", rune(c))) {
			sb.WriteByte(c)
		} else {
			break
		}
	}

	if sb.Len() > 0 {
		return &parser.Result{
			Token: &TokenVariable{Value: sb.String()},
			Count: sb.Len(),
		}
	}

	return nil
}

func parser_end(s string) *parser.Result {
	if s[0] == ']' {
		return &parser.Result{Token: &tokEnd{Value: "]"}, Count: 1}
	} else if s[0] == '}' {
		return &parser.Result{Token: &tokEnd{Value: "}"}, Count: 1}
	}
	return nil
}

func parser_code_block(s string) *parser.Result {
	if s[0] != '{' {
		return nil
	}
	iter := parser.Parse(s, []parser.Parser{
		parser_command,
		parser_end,
	})

	codeBlock := &TokenCodeBlock{
		Commands: make([]*TokenCommand, 0, 32),
	}

	cnt := 1
	for {
		res, _ := iter()

		if t, ok := res.Token.(*tokEnd); ok {
			if t.Value == "}" {
				cnt += res.Count
				break
			}
		} else if cmd, ok := res.Token.(*TokenCommand); ok {
			cnt += res.Count
			codeBlock.Commands = append(codeBlock.Commands, cmd)
		}
	}

	return &parser.Result{
		Token: codeBlock,
		Count: cnt,
	}
}

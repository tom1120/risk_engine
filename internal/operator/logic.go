package operator

import (
	"fmt"
	"strings"
)

// evaluate 计算逻辑表达式的值
func EvaluateBoolExpr(expr string, variables map[string]bool) (bool, error) {

	// 将表达式拆分成一个个token
	tokens, err := splitExpression(expr)
	if err != nil {
		return false, err
	}

	// 开始执行逻辑运算
	stack := make([]bool, 0)
	opStack := make([]string, 0)
	for _, token := range tokens {
		switch token {
		case "&&":
			for len(opStack) > 0 && opStack[len(opStack)-1] == "!" {
				if len(stack) < 1 {
					return false, fmt.Errorf("invalid expression")
				}
				b := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				stack = append(stack, !b)
				opStack = opStack[:len(opStack)-1]
			}
			opStack = append(opStack, "&&")
		case "||":
			for len(opStack) > 0 && (opStack[len(opStack)-1] == "!" || opStack[len(opStack)-1] == "&&") {
				if len(stack) < 2 {
					return false, fmt.Errorf("invalid expression")
				}
				b1, b2 := stack[len(stack)-2], stack[len(stack)-1]
				stack = stack[:len(stack)-2]
				op := opStack[len(opStack)-1]
				opStack = opStack[:len(opStack)-1]
				stack = append(stack, evaluateOp(b1, b2, op))
			}
			opStack = append(opStack, "||")
		case "!":
			opStack = append(opStack, "!")
		case "(":
			opStack = append(opStack, "(")
		case ")":
			if len(opStack) < 1 {
				return false, fmt.Errorf("invalid expression")
			}
			for opStack[len(opStack)-1] != "(" {
				if len(stack) < 2 {
					return false, fmt.Errorf("invalid expression")
				}
				b1, b2 := stack[len(stack)-2], stack[len(stack)-1]
				stack = stack[:len(stack)-2]
				op := opStack[len(opStack)-1]
				opStack = opStack[:len(opStack)-1]
				stack = append(stack, evaluateOp(b1, b2, op))
				if len(opStack) == 0 {
					return false, fmt.Errorf("unmatched parentheses")
				}
			}
			opStack = opStack[:len(opStack)-1]
			if len(opStack) > 0 && opStack[len(opStack)-1] == "!" {
				if len(stack) < 1 {
					return false, fmt.Errorf("invalid expression")
				}
				b := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				stack = append(stack, !b)
				opStack = opStack[:len(opStack)-1]
			}
		default:
			if v, ok := variables[token]; ok {
				stack = append(stack, v)
				if len(opStack) > 0 && opStack[len(opStack)-1] == "!" {
					if len(stack) < 1 {
						return false, fmt.Errorf("invalid expression")
					}
					b := stack[len(stack)-1]
					stack = stack[:len(stack)-1]
					stack = append(stack, !b)
					opStack = opStack[:len(opStack)-1]
				}
			} else {
				return false, fmt.Errorf("unknown variable %s", token)
			}
		}
	}

	for len(opStack) > 0 {
		if len(stack) < 2 {
			return false, fmt.Errorf("invalid expression")
		}
		b1, b2 := stack[len(stack)-2], stack[len(stack)-1]
		stack = stack[:len(stack)-2]
		op := opStack[len(opStack)-1]
		opStack = opStack[:len(opStack)-1]
		stack = append(stack, evaluateOp(b1, b2, op))
	}

	if len(stack) != 1 {
		return false, fmt.Errorf("invalid expression")
	}
	return stack[0], nil
}

// evaluateOp 对两个 bool 值进行逻辑运算
func evaluateOp(b1, b2 bool, op string) bool {
	switch op {
	case "&&":
		return b1 && b2
	case "||":
		return b1 || b2
	default:
		panic("unsupported operator: " + op)
	}
}

// isValid 检查表达式是否合法
func isValid(expr string) bool {
	if len(expr) == 0 {
		return false
	}
	allowed := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!()+-_*%/|&,"
	stack := make([]rune, 0)
	for _, ch := range expr {
		if ch == '(' {
			stack = append(stack, ch)
		} else if ch == ')' {
			if len(stack) == 0 {
				return false
			}
			stack = stack[:len(stack)-1]
		} else if !strings.ContainsRune(allowed, ch) {
			return false
		}
	}
	return len(stack) == 0
}

// splitExpression 将表达式拆分为token
func splitExpression(expr string) ([]string, error) {
	expr = strings.ReplaceAll(expr, " ", "") // 去除空格
	if !isValid(expr) {
		return nil, fmt.Errorf("invalid expression")
	}
	tokens := make([]string, 0)
	buf := make([]rune, 0)

	for i := 0; i < len(expr); i++ {
		ch := rune(expr[i])
		if ch == '&' && i < len(expr)-1 && rune(expr[i+1]) == '&' {
			if len(buf) > 0 {
				tokens = append(tokens, string(buf))
				buf = []rune{}
			}
			tokens = append(tokens, "&&")
			i++
		} else if ch == '|' && i < len(expr)-1 && rune(expr[i+1]) == '|' {
			if len(buf) > 0 {
				tokens = append(tokens, string(buf))
				buf = []rune{}
			}
			tokens = append(tokens, "||")
			i++
		} else if ch == '!' || ch == '(' || ch == ')' {
			if len(buf) > 0 {
				tokens = append(tokens, string(buf))
				buf = []rune{}
			}
			tokens = append(tokens, string(ch))
		} else if ch == ',' {
			if len(buf) > 0 {
				tokens = append(tokens, string(buf))
				buf = []rune{}
			}
			tokens = append(tokens, string(ch))
		} else {
			buf = append(buf, ch)
		}
	}
	if len(buf) > 0 {
		tokens = append(tokens, string(buf))
	}
	return tokens, nil
}

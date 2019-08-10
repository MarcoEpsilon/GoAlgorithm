package apply

import (
	// we also can use 
	// stack "algorithm/linear-list/stack/seq"
	stack "algorithm/linear-list/stack/linked"
	"bytes"
	"errors"
	"strconv"
	"fmt"
)
// notice: suffix exp don't make sure the expression is legal
const (
	//raw text
	RawText = iota
	//number like 1,2,...
	RawNumber
	// + -
	OperandPrimary
	// * / %
	OperandMedium
	// (
	OperandLeftParentheses
	// )
	OperandRightParentheses
)

var (
	MissingOperandNumber = errors.New("MissingOperandNumber")
	MissingOperand = errors.New("MissingOperand")
	InvalidOperandNumber = errors.New("InvalidOperandNumber")
	InvalidOperand = errors.New("InvalidOperand")
	DivideZero = errors.New("DivideZero")
)
type SuffixExp = *suffixExp
type suffixExp struct {
	*bytes.Buffer
}
type OperandWithLevel struct {
	operand string
	level int
}
func NewSuffixExp(str string) (se SuffixExp) {
	se = &suffixExp {
		Buffer: new(bytes.Buffer),
	}
	se.parseSuffixExp(str)
	return se
}

func (se SuffixExp) popWithLevel(st stack.Stack,text string,level int) {
	switch level {
	case OperandMedium:
		se.WriteString(" ")
		if !st.IsEmpty() {
			ti, _ := st.Top()
			top := ti.(OperandWithLevel)
			if top.level == OperandMedium {
				se.WriteString(top.operand)
				_, _ = st.Pop()
				se.popWithLevel(st, text, level)
				return
			}
		}
		st.Push(OperandWithLevel {
			operand: text,
			level: level,
		})
	case OperandPrimary:
		se.WriteString(" ")
		if !st.IsEmpty() {
			ti, _ := st.Top()
			top := ti.(OperandWithLevel)
			if top.level == OperandPrimary || top.level == OperandMedium {
				se.WriteString(top.operand)
				_, _ = st.Pop()
				se.popWithLevel(st, text, level)
				return
			}
		}
		st.Push(OperandWithLevel {
			operand: text,
			level: level,
		})
	case OperandLeftParentheses:
		st.Push(OperandWithLevel {
			operand: text,
			level: level,
		})
	case OperandRightParentheses:
		se.WriteString(" ")
		for ; !st.IsEmpty(); {
			ti, _ := st.Top()
			top := ti.(OperandWithLevel)
			if top.level == OperandLeftParentheses {
				_, _ = st.Pop()
				return
			}
			se.WriteString(top.operand)
			_, _ = st.Pop()
		}
	case RawNumber:
		se.WriteString(text)
	// notice: may be need be modified
	case RawText:
		se.WriteString(text)
	default:
		// do nothing
	}
}

func (se SuffixExp) parseSuffixExp(exp string) {
	operands := stack.New()
	for _, r := range exp {
		s := string(r)
		switch s {
		case "*", "/", "%":
			se.popWithLevel(operands, s, OperandMedium)
		case "-", "+":
			se.popWithLevel(operands, s, OperandPrimary)
		case "(":
			se.popWithLevel(operands, s, OperandLeftParentheses)
		case ")":
			se.popWithLevel(operands, s, OperandRightParentheses)
		// now is ignore the number and dispatch to RawText
		default:
			se.popWithLevel(operands, s, RawText)
		}
	}
	for ; !operands.IsEmpty(); {
		ti, _ := operands.Pop()
		top := ti.(OperandWithLevel)
		se.WriteString(" ")
		se.WriteString(top.operand)
	}
	buff := bytes.TrimSpace(se.Bytes())
	se.Reset()
	se.Write(buff)
}

func eval(st stack.Stack, op string) (value float64, err error) {
	var left, right float64
	if !st.IsEmpty() {
		r, _ := st.Pop()
		right = r.(float64)
	} else {
		return 0.0, MissingOperandNumber
	}
	if !st.IsEmpty() {
		l, _ := st.Pop()
		left = l.(float64)
	} else {
		return 0.0, MissingOperandNumber
	}
	switch op {
	case "*":
		return left * right, nil
	case "/":
		if right == 0.0 {
			return 0.0, DivideZero
		}
		return left / right, nil
	case "%":
		return float64(int64(left) % int64(right)), nil
	case "+":
		return left + right, nil
	case "-":
		return left - right, nil
	default:
		return 0.0, InvalidOperand
	}
}
// now is supported by number
func (se SuffixExp) EvalSuffixExp() (value float64, err error) {
	numbers := stack.New()
	splited := bytes.Split(se.Bytes(), []byte(" "))
	for _ , operand := range splited {
		if bytes.ContainsAny(operand, "+-*/%") {
			v, err := eval(numbers, string(operand))
			if err != nil {
				return 0.0, err
			}
			numbers.Push(v)
		} else {
			v, err := strconv.ParseFloat(string(operand), 64)
			if err != nil {
				fmt.Println(err)
				fmt.Println(string(operand))
				return 0.0, InvalidOperandNumber
			}
			numbers.Push(v)
		}
	}
	if numbers.Length() != 1 {
		return 0.0, MissingOperand
	} else {
		v, _ := numbers.Pop()
		return v.(float64), nil
	}
}
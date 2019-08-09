package apply

import (
	// we also can use 
	// stack "algorithm/linear-list/stack/seq"
	stack "algorithm/linear-list/stack/linked"
	"bytes"
	"errors"
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
	MissingOperand = errors.New("MissOperand")
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
		se.WriteString(top.operand)
	}
}
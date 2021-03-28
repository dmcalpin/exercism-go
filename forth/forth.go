package forth

import (
	"container/list"
	"errors"
	"strconv"
	"strings"
)

type parseMode int

const (
	modeWordDefine parseMode = iota
	modeEvaluating parseMode = iota
)

func Forth(input []string) ([]int, error) {
	s := NewStack()
	s.mode = modeEvaluating

	for _, str := range input {
		tokens := strings.Split(str, " ")

		err := evaluateTokens(s, tokens)
		if err != nil {
			return nil, err
		}
	}

	return s.GetValues(), nil
}

func evaluateTokens(s *Stack, tokens []string) error {
	for i, token := range tokens {
		if s.mode == modeWordDefine {
			if token == ";" {
				s.mode = modeEvaluating
				s.customWord = ""
			} else if i == 1 {
				s.customWords[token] = []string{}
				s.customWord = token
			} else {
				s.customWords[s.customWord] = append(s.customWords[s.customWord], token)
			}
			continue
		}

		// evaluate custom commands
		customDef, ok := s.customWords[token]
		if ok {
			return evaluateTokens(s, customDef)
		}

		// evaluate numbers
		num, err := strconv.ParseInt(token, 10, 64)
		if err == nil {
			s.Push(int(num))
			continue
		}

		// evaluate built-in commands
		switch token {
		case "+":
			v1, v2, err := s.GetTwo()
			if err != nil {
				return err
			}
			sum := v1 + v2
			s.Push(sum)
		case "-":
			v1, v2, err := s.GetTwo()
			if err != nil {
				return err
			}
			diff := v1 - v2
			s.Push(diff)
		case "*":
			v1, v2, err := s.GetTwo()
			if err != nil {
				return err
			}
			mult := v1 * v2
			s.Push(mult)
		case "/":
			v1, v2, err := s.GetTwo()
			if err != nil {
				return err
			}
			if v2 == 0 {
				return errors.New("Division by zero")
			}
			mult := v1 / v2
			s.Push(mult)
		case "dup":
			v1, err := s.GetOne()
			if err != nil {
				return err
			}
			s.Push(v1)
			s.Push(v1)
		case "drop":
			_, err := s.GetOne()
			if err != nil {
				return err
			}
		case "swap":
			v1, v2, err := s.GetTwo()
			if err != nil {
				return err
			}
			s.Push(v2)
			s.Push(v1)
		case "over":
			v1, v2, err := s.GetTwo()
			if err != nil {
				return err
			}
			s.Push(v1)
			s.Push(v2)
			s.Push(v1)
		case ":":
			s.mode = modeWordDefine
		case ";":
			s.mode = modeEvaluating
			s.customWord = ""
		default:
			return errors.New("Undefined command")
		}

	}
	return nil
}

type Stack struct {
	lst         *list.List
	mode        parseMode
	customWord  string
	customWords map[string][]string
}

func NewStack() *Stack {
	s := Stack{
		lst:         list.New(),
		customWords: map[string][]string{},
	}
	return &s
}

func (s *Stack) Push(v int) *StackElement {
	elem := s.lst.PushFront(v)
	return &StackElement{
		Value:    elem.Value.(int),
		listElem: elem,
	}
}

func (s *Stack) Pop() *StackElement {
	elem := s.lst.Front()
	if elem != nil {
		s.lst.Remove(elem)
		return &StackElement{
			Value:    elem.Value.(int),
			listElem: elem,
		}
	} else {
		return nil
	}
}

func (s *Stack) Front() *StackElement {
	elem := s.lst.Front()
	if elem != nil {
		return &StackElement{
			Value:    elem.Value.(int),
			listElem: elem,
		}
	} else {
		return nil
	}
}

func (s *Stack) Size() int {
	return s.lst.Len()
}

func (s *Stack) Empty() bool {
	return s.Size() == 0
}

func (s *Stack) GetValues() []int {
	values := make([]int, s.Size())
	i := s.Size() - 1
	elem := s.Front()
	for elem != nil {
		values[i] = elem.Value
		elem = elem.Next()
		i--
	}
	return values
}

func (s *Stack) GetOne() (int, error) {
	if s.Front() == nil {
		return 0, errors.New("No value in stack")
	}
	v1 := s.Pop().Value
	return v1, nil
}

func (s *Stack) GetTwo() (int, int, error) {
	if s.Front() == nil || s.Front().Next() == nil {
		return 0, 0, errors.New("No value in stack")
	}
	v2 := s.Pop().Value
	v1 := s.Pop().Value
	return v1, v2, nil
}

type StackElement struct {
	Value    int
	listElem *list.Element
}

func (se *StackElement) Next() *StackElement {
	elem := se.listElem.Next()
	if elem != nil {
		return &StackElement{
			Value:    elem.Value.(int),
			listElem: elem,
		}
	} else {
		return nil
	}
}

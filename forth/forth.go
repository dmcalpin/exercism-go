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

		// define mode
		if tokens[0] == ":" {
			s.defineWord(tokens)
		} else {
			// evaluate mode
			err := evaluateTokens(s, tokens)
			if err != nil {
				return nil, err
			}
		}
	}

	return s.GetValues(), nil
}

func evaluateTokens(s *Stack, tokens []string) error {
	for _, token := range tokens {
		// evaluate numbers
		num, err := strconv.ParseInt(token, 10, 64)
		if err == nil {
			s.Push(int(num))
			continue
		}

		// evaluate custom commands
		customDef, ok := s.customWords[token]
		if ok {
			evaluateTokens(s, customDef)
			continue
		}

		// evaluate built-in commands
		cmd, ok := s.commands[token]
		if !ok {
			return errors.New("Unknown command")
		}
		err = cmd()
		if err != nil {
			return err
		}
	}
	return nil
}

type Stack struct {
	lst         *list.List
	mode        parseMode
	customWord  string
	customWords map[string][]string
	commands    map[string]func() error
}

func NewStack() *Stack {
	s := Stack{
		lst:         list.New(),
		customWords: map[string][]string{},
	}

	s.commands = map[string]func() error{
		"+": func() error {
			v1, v2, err := s.getTwo()
			if err != nil {
				return err
			}
			sum := v1 + v2
			s.Push(sum)
			return nil
		},
		"-": func() error {
			v1, v2, err := s.getTwo()
			if err != nil {
				return err
			}
			diff := v1 - v2
			s.Push(diff)
			return nil
		},
		"*": func() error {
			v1, v2, err := s.getTwo()
			if err != nil {
				return err
			}
			mult := v1 * v2
			s.Push(mult)
			return nil
		},
		"/": func() error {
			v1, v2, err := s.getTwo()
			if err != nil {
				return err
			}
			if v2 == 0 {
				return errors.New("Division by zero")
			}
			mult := v1 / v2
			s.Push(mult)
			return nil
		},
		"dup": func() error {
			v1, err := s.getOne()
			if err != nil {
				return err
			}
			s.Push(v1)
			s.Push(v1)
			return nil
		},
		"drop": func() error {
			_, err := s.getOne()
			if err != nil {
				return err
			}
			return nil
		},
		"swap": func() error {
			v1, v2, err := s.getTwo()
			if err != nil {
				return err
			}
			s.Push(v2)
			s.Push(v1)
			return nil
		},
		"over": func() error {
			v1, v2, err := s.getTwo()
			if err != nil {
				return err
			}
			s.Push(v1)
			s.Push(v2)
			s.Push(v1)
			return nil
		},
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

func (s *Stack) getOne() (int, error) {
	if s.Front() == nil {
		return 0, errors.New("No value in stack")
	}
	v1 := s.Pop().Value
	return v1, nil
}

func (s *Stack) getTwo() (int, int, error) {
	if s.Front() == nil || s.Front().Next() == nil {
		return 0, 0, errors.New("No value in stack")
	}
	v2 := s.Pop().Value
	v1 := s.Pop().Value
	return v1, v2, nil
}

func (s *Stack) defineWord(tokens []string) {
	for i, token := range tokens {
		if i == 0 || i == len(tokens)-1 {
			continue
		}

		if token == ";" { // close the mode
			s.mode = modeEvaluating
			s.customWord = ""
		} else if i == 1 { // initialize the word
			s.customWords[token] = []string{}
			s.customWord = token
		} else { // add to definition
			cmdVals, ok := s.customWords[token]
			if ok && len(cmdVals) == 1 {
				token = cmdVals[0]
				s.customWords[s.customWord] = append(s.customWords[s.customWord], token)
			} else {
				s.customWords[s.customWord] = append(s.customWords[s.customWord], token)
			}
		}
	}
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

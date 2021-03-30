package forth

import (
	"container/list"
	"errors"
	"strconv"
	"strings"
)

var (
	errDivisionByZero  = errors.New("Cannot divide by zero")
	errNoElem          = errors.New("No element on stack")
	errCommandNotFound = errors.New("Command not found")
	errOverride        = errors.New("Cannot override numbers")
)

func Forth(input []string) ([]int, error) {
	s := NewStack()

	var err error

	for _, str := range input {
		normalizedStr := strings.ToLower(str)
		tokens := strings.Split(normalizedStr, " ")

		if tokens[0] == ":" {
			err = s.defineWord(tokens)
		} else {
			err = s.evaluateTokens(tokens)
		}

		if err != nil {
			return nil, err
		}
	}

	return s.getValues(), nil
}

type Stack struct {
	*list.List
	customWords map[string][]string
}

func NewStack() *Stack {
	s := Stack{
		List:        list.New(),
		customWords: map[string][]string{},
	}
	return &s
}

func (s *Stack) Push(v int) *list.Element {
	return s.PushFront(v)
}

func (s *Stack) Pop() (*list.Element, error) {
	elem := s.Front()
	if elem != nil {
		s.Remove(elem)
		return elem, nil
	} else {
		return nil, errNoElem
	}
}

func (s *Stack) getValues() []int {
	values := make([]int, s.List.Len())
	i := s.List.Len() - 1
	elem := s.Front()
	for elem != nil {
		values[i] = elem.Value.(int)
		elem = elem.Next()
		i--
	}
	return values
}

// evaluateTokens pushes commands onto the stack
// depending on what type they are
func (s *Stack) evaluateTokens(tokens []string) error {
	for _, token := range tokens {
		// push numbers onto the stack
		num, err := strconv.ParseInt(token, 10, 64)
		if err == nil {
			s.Push(int(num))
			continue
		}

		// custom words are just a predefined
		// set of tokens, so we can just call
		// this function again on them
		customDef, ok := s.customWords[token]
		if ok {
			s.evaluateTokens(customDef)
			continue
		}

		// push predefined commands onto the stack
		err = s.runCommand(token)
		if err != nil {
			return err
		}
	}
	return nil
}

// defines a custom word. the slice starts with ":" and
// ends with ";" so these need to be removed. The second
// elem is the name of the custom word, everything
// that follows is the commands or numbers
func (s *Stack) defineWord(definition []string) error {
	customWord := definition[1]
	definition = definition[2 : len(definition)-1]
	cmds := make([]string, len(definition))

	// numbers not allowed
	_, err := strconv.ParseInt(customWord, 10, 64)
	if err == nil {
		return errOverride
	}

	for i, token := range definition {
		// use custom word def if found
		cmdVals, ok := s.customWords[token]
		if ok {
			token = cmdVals[0]
		}

		// add to definition
		cmds[i] = token
	}

	s.customWords[customWord] = cmds

	return nil
}

// runCommands runs predefined commands which may
// not be overridden
func (s *Stack) runCommand(cmd string) error {
	elem1, err := s.Pop() // this also handles "drop"
	if err != nil {
		return err
	}
	v1 := elem1.Value.(int)

	switch cmd {
	case "drop": // does nothing
	case "dup":
		s.Push(v1)
		s.Push(v1)
	default: // cmds which need 2 params
		elem2, err := s.Pop()
		if err != nil {
			return err
		}
		v2 := elem2.Value.(int)

		switch cmd {
		case "+":
			s.Push(v1 + v2)
		case "-":
			s.Push(v2 - v1)
		case "*":
			s.Push(v1 * v2)
		case "/":
			if v1 == 0 {
				return errDivisionByZero
			}
			s.Push(v2 / v1)
		case "swap":
			s.Push(v1)
			s.Push(v2)
		case "over":
			s.Push(v2)
			s.Push(v1)
			s.Push(v2)
		default:
			return errCommandNotFound
		}
	}

	return nil
}

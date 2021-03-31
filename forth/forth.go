package forth

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

var (
	ErrDivisionByZero  = errors.New("Cannot divide by zero")
	ErrNoElem          = errors.New("No element on stack")
	ErrCommandNotFound = errors.New("Command not found")
	ErrOverride        = errors.New("Cannot override numbers")
)

var numRegexp = regexp.MustCompile(`^\d$`)

func isNum(str string) bool {
	return numRegexp.MatchString(str)
}

func Forth(input []string) ([]int, error) {
	s := NewStack()

	var err error

	for _, str := range input {
		normalizedStr := strings.ToLower(str)
		tokens := strings.Fields(normalizedStr)

		if tokens[0] == ":" {
			err = s.defineWord(tokens)
		} else {
			err = s.evaluateTokens(tokens)
		}

		if err != nil {
			return nil, err
		}
	}

	return s.list, nil
}

type Stack struct {
	list        []int
	customWords map[string][]string
}

func NewStack() *Stack {
	s := Stack{
		list:        []int{},
		customWords: map[string][]string{},
	}
	return &s
}

func (s *Stack) Push(v ...int) {
	s.list = append(s.list, v...)
}

func (s *Stack) Pop() (int, error) {
	var elem int
	if len(s.list) != 0 {
		elem, s.list = s.list[len(s.list)-1], s.list[:len(s.list)-1]
		return elem, nil
	} else {
		return 0, ErrNoElem
	}
}

// evaluateTokens pushes commands onto the stack
// depending on what type they are
func (s *Stack) evaluateTokens(tokens []string) error {
	for _, token := range tokens {
		// push numbers onto the stack
		num, err := strconv.Atoi(token)
		if err == nil {
			s.Push(num)
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

	// numbers not allowed
	if isNum(customWord) {
		return ErrOverride
	}

	definition = definition[2 : len(definition)-1]
	cmds := make([]string, len(definition))

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
	v1, err := s.Pop() // this also handles "drop"
	if err != nil {
		return err
	}

	switch cmd {
	case "drop": // does nothing
	case "dup":
		s.Push(v1, v1)
	default: // cmds which need 2 params
		v2, err := s.Pop()
		if err != nil {
			return err
		}

		switch cmd {
		case "+":
			s.Push(v1 + v2)
		case "-":
			s.Push(v2 - v1)
		case "*":
			s.Push(v1 * v2)
		case "/":
			if v1 == 0 {
				return ErrDivisionByZero
			}
			s.Push(v2 / v1)
		case "swap":
			s.Push(v1, v2)
		case "over":
			s.Push(v2, v1, v2)
		default:
			return ErrCommandNotFound
		}
	}

	return nil
}

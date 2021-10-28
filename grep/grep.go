package grep

import (
	"errors"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Grep struct {
	Pattern             string
	SearchStr           string
	FlagLineNo          bool
	FlagFileName        bool
	FlagInverse         bool
	FlagCaseInsensitive bool
	FlagMatchEntireLine bool
	Filename            string
	Regex               *regexp.Regexp
	NumFiles            int
	Matches             []string
	CurrLine            int
	lastText            string
	matchString         func(string) bool
}

func Search(pattern string, flags []string, files []string) []string {
	config, err := NewGrepper(pattern, flags)
	if err != nil {
		log.Fatal(err)
	}

	config.NumFiles = len(files)

	for _, filename := range files {
		config.ScanFile(filename)
	}

	return config.Matches
}

func NewGrepper(
	pattern string,
	flags []string,
) (*Grep, error) {
	config := &Grep{
		Pattern:   pattern,
		SearchStr: pattern,
		Matches:   []string{},
	}

	for _, flag := range flags {
		switch flag {
		case "-n":
			config.FlagLineNo = true
		case "-i": // case insensitive
			config.FlagCaseInsensitive = true
		case "-l":
			config.FlagFileName = true
		case "-x": // match entire line
			config.FlagMatchEntireLine = true
		case "-v":
			config.FlagInverse = true
		default:
			return nil, errors.New("Unknown flag: " + flag)
		}
	}

	if config.FlagMatchEntireLine {
		if config.FlagCaseInsensitive {
			config.matchString = config.insensitiveFullMatch
		} else {
			config.matchString = config.exactMatch
		}
	} else {
		if config.FlagCaseInsensitive {
			config.matchString = config.insensitivePartialMatch
		} else {
			config.matchString = config.partialMatch
		}
	}

	return config, nil
}

func (config *Grep) ScanFile(filename string) {
	config.Filename = filename
	fileBytes, err := os.ReadFile(config.Filename)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(fileBytes), "\n")
	config.CurrLine = 1
	config.lastText = ""
	for _, text := range lines {
		if text == "" {
			continue
		}

		isMatch := config.matchString(text)

		// flagInverse causes isMatch to do
		// do the opposite.
		if config.FlagInverse == isMatch {
			config.CurrLine++
			continue
		}

		text = config.decorateText(text)
		if text != "" {
			config.lastText = text
			config.Matches = append(config.Matches, text)

			config.CurrLine++
		}
	}
}

func (g *Grep) decorateText(text string) string {
	if g.FlagFileName {
		// only add the file name once
		if g.Filename != g.lastText {
			return g.Filename
		}
		return ""
	}
	if g.FlagLineNo {
		text = strconv.Itoa(g.CurrLine) + ":" + text
	}
	if g.NumFiles > 1 {
		text = g.Filename + ":" + text
	}
	return text
}

func (g *Grep) insensitiveFullMatch(text string) bool {
	return strings.EqualFold(text, g.SearchStr)
}

func (g *Grep) exactMatch(text string) bool {
	return text == g.SearchStr
}

func (g *Grep) insensitivePartialMatch(text string) bool {
	return strings.Contains(strings.ToLower(text), strings.ToLower(g.SearchStr))
}

func (g *Grep) partialMatch(text string) bool {
	return strings.Contains(text, g.SearchStr)
}

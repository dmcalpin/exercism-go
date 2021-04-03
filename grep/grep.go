package grep

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

func Search(pattern string, flags []string, files []string) []string {
	matches := []string{}

	flagLineNo := false
	flagFileName := false
	flagInverse := false
	for _, flag := range flags {
		switch flag {
		case "-n":
			flagLineNo = true
		case "-i": // case insensitive
			pattern = "(?i)" + pattern
		case "-l":
			flagFileName = true
		case "-x": // match entire line
			pattern = "^" + pattern + "$"
		case "-v":
			flagInverse = true
		default:
			log.Fatal("Error: Unknown flag", flag)
			return nil
		}
	}

	re := regexp.MustCompile(pattern)

	numFiles := len(files)
	for _, fileName := range files {
		file, err := os.Open(fileName)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		line := 1
		lastText := ""
		for scanner.Scan() {
			text := scanner.Text()
			isMatch := re.MatchString(text)

			// flagInverse causes isMatch to do
			// do the opposite.
			if flagInverse == isMatch {
				line++
				continue
			}

			if flagFileName {
				// only add the file name once
				if fileName == lastText {
					continue
				}
				text = fileName
			} else {
				if flagLineNo {
					text = strconv.Itoa(line) + ":" + text
				}
				if numFiles > 1 {
					text = fileName + ":" + text
				}
			}

			lastText = text
			matches = append(matches, text)

			line++
		}
	}

	return matches
}

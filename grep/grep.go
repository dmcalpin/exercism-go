package grep

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func Search(pattern string, flags []string, files []string) []string {
	matches := []string{}

	flagLineNo := false
	flagFileName := false
	flagInverse := false
	for _, flag := range flags {
		if flag == "-n" {
			flagLineNo = true
		} else if flag == "-i" { // case incensitive
			pattern = "(?i)" + pattern
		} else if flag == "-l" {
			flagFileName = true
		} else if flag == "-x" { // match entire line
			pattern = "^" + pattern + "$"
		} else if flag == "-v" {
			flagInverse = true
		}
	}

	re := regexp.MustCompile(pattern)

	numFiles := len(files)
	for _, fileName := range files {
		file, err := os.Open(fileName)
		if err != nil {
			return []string{}
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		line := 1
		for scanner.Scan() {
			text := scanner.Text()
			isMatch := re.MatchString(text)
			if (!flagInverse && isMatch) || (flagInverse && !isMatch) {
				if flagFileName {
					// only add the file name once
					if len(matches) > 0 && fileName == matches[len(matches)-1] {
						continue
					}
					text = fileName
				} else {
					if flagLineNo {
						text = fmt.Sprintf("%d:%s", line, text)
					}
					if numFiles > 1 {
						text = fmt.Sprintf("%s:%s", fileName, text)
					}
				}

				matches = append(matches, text)
			}
			line++
		}

	}

	return matches
}

// Package tournament has a tally
// function
package tournament

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"sort"
	"strings"
)

type TeamRecord struct {
	Name          string
	MatchesPlayed int
	Wins          int
	Draws         int
	Losses        int
	Points        int
}

// Tally tallies up the scores
func Tally(r io.Reader, w io.Writer) error {
	tourneyInfo := map[string]*TeamRecord{}

	bytes, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	matches := strings.Split(strings.TrimSpace(string(bytes)), "\n")
	for _, match := range matches {
		if match == "" || match[0] == '#' {
			continue
		}
		matchInfo := strings.Split(match, ";")
		if len(matchInfo) != 3 {
			return errors.New("Bad format, need: home team;visiting team;outcome")
		}
		homeTeamName := matchInfo[0]
		visitingTeamName := matchInfo[1]
		matchResult := matchInfo[2]

		homeTeam, homeExists := tourneyInfo[homeTeamName]
		if homeExists == false {
			homeTeam = &TeamRecord{Name: homeTeamName}
			tourneyInfo[homeTeamName] = homeTeam
		}
		visitingTeam, visitingExists := tourneyInfo[visitingTeamName]
		if visitingExists == false {
			visitingTeam = &TeamRecord{Name: visitingTeamName}
			tourneyInfo[visitingTeamName] = visitingTeam
		}

		homeTeam.MatchesPlayed++
		visitingTeam.MatchesPlayed++

		switch matchResult {
		case "win":
			homeTeam.Wins++
			homeTeam.Points += 3
			visitingTeam.Losses++
		case "loss":
			homeTeam.Losses++
			visitingTeam.Wins++
			visitingTeam.Points += 3
		case "draw":
			homeTeam.Draws++
			homeTeam.Points++
			visitingTeam.Draws++
			visitingTeam.Points++
		default:
			return errors.New("invalid match result")
		}
	}

	var tourneySlice = make([]*TeamRecord, len(tourneyInfo))
	i := 0
	for _, result := range tourneyInfo {
		tourneySlice[i] = result
		i++
	}
	sort.SliceStable(tourneySlice, func(i, j int) bool {
		if tourneySlice[j].Points != tourneySlice[i].Points {
			return tourneySlice[j].Points < tourneySlice[i].Points
		} else if tourneySlice[i].Name < tourneySlice[j].Name {
			return true
		}
		return false
	})

	w.Write([]byte(formatResults(tourneySlice)))

	return nil
}

func formatResults(records []*TeamRecord) string {
	resultStr := fmt.Sprintf(
		"%-30s | %2s | %2s | %2s | %2s | %2s\n",
		"Team", "MP", "W", "D", "L", "P",
	)

	for _, record := range records {
		resultStr += fmt.Sprintf(
			"%-30s | %2d | %2d | %2d | %2d | %2d\n",
			record.Name,
			record.MatchesPlayed,
			record.Wins,
			record.Draws,
			record.Losses,
			record.Points,
		)
	}

	return resultStr
}

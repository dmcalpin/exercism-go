// Package tournament has a tally
// function
package tournament

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"sort"
	"strings"
)

// TeamRecord keeps track of a teams match
// results
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

	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		match := scanner.Text()
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

		err := addPoints(homeTeam, visitingTeam, matchResult)
		if err != nil {
			return err
		}

	}

	var tourneySlice = make([]*TeamRecord, len(tourneyInfo))
	i := 0
	for _, result := range tourneyInfo {
		tourneySlice[i] = result
		i++
	}

	// Sort by games points, then by name
	sort.SliceStable(tourneySlice, func(i, j int) bool {
		if tourneySlice[j].Points != tourneySlice[i].Points {
			return tourneySlice[j].Points < tourneySlice[i].Points
		} else if tourneySlice[i].Name < tourneySlice[j].Name {
			return true
		}
		return false
	})

	writeResults(w, tourneySlice)

	return nil
}

func addPoints(
	homeTeam *TeamRecord,
	visitingTeam *TeamRecord,
	matchResult string,
) error {
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
	return nil
}

func writeResults(w io.Writer, records []*TeamRecord) {
	result := bufio.NewWriter(w)
	_, err := result.WriteString(
		"Team                           | MP |  W |  D |  L |  P\n",
	)
	if err != nil {
		log.Fatal(err)
	}

	for _, record := range records {
		_, err = result.WriteString(fmt.Sprintf(
			"%-30s | %2d | %2d | %2d | %2d | %2d\n",
			record.Name,
			record.MatchesPlayed,
			record.Wins,
			record.Draws,
			record.Losses,
			record.Points,
		))
		if err != nil {
			log.Fatal(err)
		}
	}
	result.Flush()
}

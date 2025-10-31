package main

import (
	"fmt"
	"io"
	"os"
	"sort"
)

type Ranker interface {
	Ranking() []string
}

func RankPrinter(ranker Ranker, writer io.Writer) {
	for _, str := range ranker.Ranking() {
		io.WriteString(writer, str+"\n")
	}
}

type Team struct {
	Name    string
	Players []string
}

type League struct {
	Teams []Team
	Wins  map[string]int
}

func (l *League) MatchResult(teamA string, teamAScore int, teamB string, teamBScore int) {
	if teamAScore == teamBScore {
		return
	}

	winTeam := teamA
	if teamBScore > teamAScore {
		winTeam = teamB
	}

	l.Wins[winTeam]++
}

func (l *League) Ranking() []string {
	results := make([]string, 0, len(l.Teams))
	for _, t := range l.Teams {
		results = append(results, t.Name)
	}

	sort.Slice(results, func(i, j int) bool {
		return l.Wins[results[i]] > l.Wins[results[j]]
	})

	return results
}

func main() {
	teamA := Team{"Bull", []string{"Mike", "Joe"}}
	teamB := Team{"Lakers", []string{"Mike", "Joe"}}
	teamC := Team{"Rockets", []string{"Mike", "Joe"}}

	league := &League{
		Teams: []Team{teamA, teamB, teamC},
		Wins:  map[string]int{},
	}

	league.MatchResult(teamA.Name, 10, teamB.Name, 20)
	league.MatchResult(teamB.Name, 30, teamC.Name, 20)
	league.MatchResult(teamB.Name, 30, teamC.Name, 20)
	league.MatchResult(teamC.Name, 10, teamA.Name, 20)

	fmt.Println(league.Ranking())

	RankPrinter(league, os.Stdout)
}

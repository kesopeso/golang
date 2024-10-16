package main

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"sort"
)

func main() {
	league := LeagueFactory()
	league.SimulateMatches(100)
	RankPrinter(league, os.Stdout)
}

type Ranker interface {
	Ranking() []string
}

type Team struct {
	Name        string
	PlayerNames []string
}

type League struct {
	Teams map[string]Team
	Wins  map[string]int
}

func RankPrinter(ranker Ranker, writer io.Writer) {
	ranking := ranker.Ranking()
	for _, r := range ranking {
		io.WriteString(writer, r)
		writer.Write([]byte("\n"))
	}

}

func (l *League) SimulateMatches(numberOfMatches int) {
	for i := 0; i < numberOfMatches; i++ {
		teamOne, teamTwo := l.getTwoRandomTeams()
		scoreOne, scoreTwo := l.getTwoRandomScores()
		l.MatchResult(teamOne.Name, scoreOne, teamTwo.Name, scoreTwo)
	}
	fmt.Println("League wins", l.Wins)
}

func (l League) getTwoRandomTeams() (Team, Team) {
	teamsCount := len(l.Teams)

	i := rand.Intn(teamsCount)
	j := i
	for i == j {
		j = rand.Intn(teamsCount)
	}

	teamNames := make([]string, 0, teamsCount)
	for k := range l.Teams {
		teamNames = append(teamNames, k)
	}
	return l.Teams[teamNames[i]], l.Teams[teamNames[j]]
}

func (l League) getTwoRandomScores() (int, int) {
	scoreOne := 40 + rand.Intn(51)
	scoreTwo := 40 + rand.Intn(51)
	return scoreOne, scoreTwo
}

func (l *League) MatchResult(teamOne string, scoreOne int, teamTwo string, scoreTwo int) {
	if _, ok := l.Teams[teamOne]; !ok {
		return
	}
	if _, ok := l.Teams[teamTwo]; !ok {
		return
	}

	if scoreOne > scoreTwo {
		l.Wins[teamOne]++
	} else if scoreOne > scoreTwo {
		l.Wins[teamTwo]++
	}
}

func (l League) Ranking() []string {
	ranking := make([]string, 0, len(l.Teams))
	for k := range l.Teams {
		ranking = append(ranking, k)
	}

	sort.Slice(ranking, func(i, j int) bool {
		hasLessWins := l.Wins[ranking[i]] > l.Wins[ranking[j]]
		hasSameWins := l.Wins[ranking[i]] == l.Wins[ranking[j]]
		isAlphabeticallyLast := ranking[i] < ranking[j]
		return hasLessWins || (hasSameWins && isAlphabeticallyLast)
	})

	return ranking
}

func TeamFactory() map[string]Team {
	return map[string]Team{
		"Bulls":     {"Bulls", []string{"Vucevic", "Colby", "Lavine", "Derozan"}},
		"Pacers":    {"Pacers", []string{"Haliburton", "Hield", "Turner", "Tobin"}},
		"Cavaliers": {"Cavaliers", []string{"Garland", "Spida", "Supak", "Lavert"}},
		"Knicks":    {"Knicks", []string{"Hart", "Brunson", "Center", "Someone"}},
		"Mavericks": {"Mavericks", []string{"Doncic", "Irving", "Gafford", "Washington"}},
		"Celtics":   {"Celtics", []string{"Tatum", "Brown", "White", "Horford"}},
		"Lakers":    {"Lakers", []string{"LeBron", "AD", "Dinwiddie", "Bronny"}},
	}
}

func LeagueFactory() League {
	teams := TeamFactory()
	wins := map[string]int{}
	return League{Teams: teams, Wins: wins}
}

package tournament

import (
	"fmt"
	"io"
	"sort"
	"text/tabwriter"
)

// New creates a new struct for results of teams.
func (b *Board) New() *Board {
	m := make(map[string]*TeamStats)
	a := make([]string, 0)
	return &Board{m, a}
}

func (b *Board) addTeam(key string) {
	if _, ok := b.Teams[key]; ok {
		return
	}
	b.Teams[key] = &TeamStats{0, 0, 0, 0, 0}
	b.OrderTeams = append(b.OrderTeams, key)
}

func (b *Board) draw(key1 string, key2 string) {
	b.Teams[key1].D++
	b.Teams[key2].D++
}

func (b *Board) win(key1 string, key2 string) {
	b.Teams[key1].W++
	b.Teams[key2].L++
}

func (b *Board) loss(key1 string, key2 string) {
	b.win(key2, key1)
}

func (b *Board) addMatch(key1 string, key2 string) {
	b.addTeam(key1)
	b.addTeam(key2)
	b.Teams[key1].MP++
	b.Teams[key2].MP++
}

func (b *Board) calcPoints(key1 string, key2 string) {
	b.Teams[key1].P = b.Teams[key1].W*3 + b.Teams[key1].D
	b.Teams[key2].P = b.Teams[key2].W*3 + b.Teams[key2].D
}

// AddResult adds results for teams.
func (b *Board) AddResult(g Game) {
	b.addMatch(g.Team1, g.Team2)
	switch g.Result {
	case "draw":
		b.draw(g.Team1, g.Team2)
	case "win":
		b.win(g.Team1, g.Team2)
	case "loss":
		b.loss(g.Team1, g.Team2)
	}
	b.calcPoints(g.Team1, g.Team2)
}

// WriteBoard writes board with results of teams.
func (b *Board) WriteBoard(buffer io.Writer) {
	b.sort()
	const padding = 8
	w := tabwriter.NewWriter(buffer, 0, 0, padding, ' ', 0)
	fmt.Fprintln(w, "Team\t| MP |  W |  D |  L |  P")
	for _, v := range b.OrderTeams {
		fmt.Fprintf(w, "%s\t|  %d |  %d |  %d |  %d |  %d\n", v, b.Teams[v].MP, b.Teams[v].W, b.Teams[v].D, b.Teams[v].L, b.Teams[v].P)
	}
	w.Flush()
}

func (b *Board) compare(i, j int) bool {
	p1 := b.Teams[b.OrderTeams[i]]
	p2 := b.Teams[b.OrderTeams[j]]
	fiches1 := []int{p1.MP, p1.W, p1.D, p1.L, p1.P}
	fiches2 := []int{p2.MP, p2.W, p2.D, p2.L, p2.P}
	k := 0
	for k = len(fiches1) - 1; k >= 0; k-- {
		if fiches1[k] > fiches2[k] {
			return true
		}
		if fiches1[k] != fiches2[k] {
			break
		}
		if k == 0 {
			if b.OrderTeams[i] < b.OrderTeams[j] {
				return true
			}
		}
	}
	return false
}

func (b *Board) sort() {
	sort.Slice(b.OrderTeams, func(i, j int) bool { return b.compare(i, j) })
}

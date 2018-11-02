package tournament

import (
	"fmt"
	"io"
	"text/tabwriter"
)

func (b *Board) New() *Board {
	m := make(map[string]*TeamStats)
	a := make([]string, 0)
	return &Board{m, a}
}

func (b *Board) AddTeam(key string) {
	if _, ok := b.Teams[key]; ok {
		return
	}
	b.Teams[key] = &TeamStats{0, 0, 0, 0, 0}
	b.OrderTeams = append(b.OrderTeams, key)
}

func (b *Board) draw(key1 string, key2 string) {
	b.Teams[key1].D += 1
	b.Teams[key2].D += 1
}

func (b *Board) win(key1 string, key2 string) {
	b.Teams[key1].W += 1
	b.Teams[key2].L += 1
}

func (b *Board) loss(key1 string, key2 string) {
	b.win(key2, key1)
}

func (b *Board) addMatch(key1 string, key2 string) {
	b.AddTeam(key1)
	b.AddTeam(key2)
	b.Teams[key1].MP += 1
	b.Teams[key2].MP += 1
}

func (b *Board) calcPoints(key1 string, key2 string) {
	b.Teams[key1].P = b.Teams[key1].W*3 + b.Teams[key1].D
	b.Teams[key2].P = b.Teams[key2].W*3 + b.Teams[key2].D
}

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

func (b *Board) WriteBoard(buffer io.Writer) {
	const padding = 8
	w := tabwriter.NewWriter(buffer, 0, 0, padding, ' ', 0)
	fmt.Fprintln(w, "Team\t| MP |  W |  D |  L |  P")
	for _, v := range b.OrderTeams {
		fmt.Fprintf(w, "%s\t|  %d |  %d |  %d |  %d |  %d\n", v, b.Teams[v].MP, b.Teams[v].W, b.Teams[v].D, b.Teams[v].L, b.Teams[v].P)
	}
	w.Flush()
}

func (b *Board) Sort() {
	for i := 0; i < len(b.OrderTeams)-1; i++ {
		for j := 0; j < len(b.OrderTeams)-1; j++ {
			p1 := b.Teams[b.OrderTeams[j]]
			p2 := b.Teams[b.OrderTeams[j+1]]
			fiches1 := []int{p1.MP, p1.W, p1.D, p1.L, p1.P}
			fiches2 := []int{p2.MP, p2.W, p2.D, p2.L, p2.P}
			k := 0
			for k = len(fiches1) - 1; k > 0; k-- {
				if fiches1[k] < fiches2[k] {
					b.OrderTeams[j], b.OrderTeams[j+1] = b.OrderTeams[j+1], b.OrderTeams[j]
					continue
				}
				if fiches1[k] != fiches2[k] {
					break
				}
			}
			if k == 0 {
				if b.OrderTeams[j] > b.OrderTeams[j+1] {
					b.OrderTeams[j], b.OrderTeams[j+1] = b.OrderTeams[j+1], b.OrderTeams[j]
				}
			}
		}
	}
}

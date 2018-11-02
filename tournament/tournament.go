package tournament

import (
	"io"
	"fmt"
	"bufio"
	"strings"
	"text/tabwriter"
)

type Team struct {
	MP int
	W  int
	D  int
	L  int
	P  int
}

type T struct {
	Team1 string
	Team2 string
	Result string
}

func isValidStat(s []string) bool {
	for _, v := range s[:len(s)-1] {
		if len(v) < 1 {
			return false
		}
	}
	return true
}

func Tally(reader io.Reader, buffer io.Writer) error {
	scanner := bufio.NewScanner(reader)
	m := make(map[string]*Team)
	mOrder := make([]string, 0)
	for scanner.Scan() {
		stat := strings.Split(scanner.Text(), ";")
		if (len(stat) == 3 && isValidStat(stat)) {
			t := T{stat[0], stat[1], stat[2]}
			if _, ok := m[stat[0]]; !ok {
					m[t.Team1] = &Team{0, 0, 0, 0, 0}
					mOrder = append(mOrder, t.Team1)
			}
			if _, ok := m[stat[1]]; !ok {
					m[t.Team2] = &Team{0, 0, 0, 0, 0}
					mOrder = append(mOrder, t.Team2)
			}
			p1 := m[t.Team1]
			p2 := m[t.Team2]
			switch t.Result {
				case "draw":
					p1.D += 1
					p2.D += 1
				case "win":
					p1.W += 1
					p2.L += 1
				case "loss":
					p1.L += 1
					p2.W += 1
			}
			p1.MP += 1
			p2.MP += 1
			p1.P = p1.W * 3 + p1.D
			p2.P = p2.W * 3 + p2.D
			fmt.Printf("%v\n", t)
		}
	}
	fmt.Printf("%v\n", mOrder)
	for i := 0; i < len(mOrder) - 1; i++ {
		for j := 0; j < len(mOrder) - 1; j++ {
			p1 := m[mOrder[j]]
			p2 := m[mOrder[j+1]]
			if p1.P <= p2.P {
				mOrder[j],  mOrder[j+1] = mOrder[j+1], mOrder[j]
			}
		}
	}

	fmt.Printf("%v\n", mOrder)
	const padding = 8
	w := tabwriter.NewWriter(buffer, 0, 0, padding, ' ', 0)
	fmt.Fprintln(w, "Team\t| MP |  W |  D |  L |  P")
	for _, v := range mOrder {
		fmt.Fprintf(w,"%s\t|  %d |  %d |  %d |  %d |  %d\n", v, m[v].MP, m[v].W, m[v].D, m[v].L, m[v].P)
	}
	w.Flush()
	return nil
}

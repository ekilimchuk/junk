package tournament

import (
	"bufio"
	"io"
)

type Board struct {
	Teams      map[string]*TeamStats
	OrderTeams []string
}

type TeamStats struct {
	MP int
	W  int
	D  int
	L  int
	P  int
}

type Game struct {
	Team1  string
	Team2  string
	Result string
}

func Tally(reader io.Reader, buffer io.Writer) error {
	var board *Board
	board = board.New()
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		t, ok, err := GameParcer(scanner.Text())
		if err != nil {
			return err
		}
		if !ok {
			continue
		}
		board.AddResult(t)
	}
	board.Sort()
	board.WriteBoard(buffer)
	return nil
}

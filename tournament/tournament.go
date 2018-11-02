package tournament

import (
	"bufio"
	"io"
)

// Board is a struct with results of teams.
type Board struct {
	Teams      map[string]*TeamStats
	OrderTeams []string
}

// TeamStats is a struct with stats of a team.
type TeamStats struct {
	MP int // Matches Played
	W  int // Matches Won
	D  int // Matches Drawn (Tied)
	L  int // Matches Lost
	P  int // Points
}

// Game is a struct with a competition information.
type Game struct {
	Team1  string
	Team2  string
	Result string
}

// Tally counts points and writes a result board.
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
	board.WriteBoard(buffer)
	return nil
}

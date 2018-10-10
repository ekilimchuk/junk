package tournament

import (
	"io"
	"fmt"
//	"bytes"
//	"strings"
)

func Tally(reader io.Reader, buffer io.Writer) error {
	fmt.Printf("%d", len(reader))
	return nil
}

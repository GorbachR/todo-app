package utility

import (
	"errors"
	"strconv"
	"strings"
)

func ExtractIdUrl(input string, segmentPos int) (id int, err error) {
	segments := strings.Split(input, "/")

	if len(segments) > segmentPos || len(segments) < segmentPos {
		err = errors.New("Route doesn't match")
		return
	}

	id, err = strconv.Atoi(segments[segmentPos])
	return
}

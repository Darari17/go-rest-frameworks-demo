package helper

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
)

func ExtractPathID(r *http.Request) (int, error) {
	segments := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	if len(segments) < 1 {
		return 0, errors.New("invalid path")
	}

	idStr := segments[len(segments)-1]
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		return 0, errors.New("invalid ID: " + idStr)
	}

	return id, nil
}

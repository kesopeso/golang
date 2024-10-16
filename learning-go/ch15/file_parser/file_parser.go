package file_parser

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"strconv"
	"strings"
)

func ParseData(r io.Reader, maxLines int) ([]string, error) {
	s := bufio.NewScanner(r)
	if !s.Scan() {
		return nil, errors.New("empty")
	}
	countStr := s.Text()
	count, err := strconv.Atoi(countStr)
	if err != nil {
		return nil, err
	}

	if count > maxLines {
		return nil, errors.New("too many lines")
	}

	if count < 0 {
		return nil, errors.New("no negative numbers")
	}

	out := make([]string, 0, count)
	for i := 0; i < count; i++ {
		hasLine := s.Scan()
		if !hasLine {
			return nil, errors.New("too few lines")
		}
		line := strings.TrimSpace(s.Text())
		if len(line) == 0 {
			return nil, errors.New("no empty lines")
		}

		out = append(out, line)
	}
	return out, nil
}

func ParseDataFixed(r io.Reader) ([]string, error) {
	s := bufio.NewScanner(r)
	if !s.Scan() {
		return nil, errors.New("empty")
	}

	countStr := s.Text()
	count, err := strconv.Atoi(countStr)
	if err != nil {
		return nil, err
	}

	if count > 1000 {
		return nil, errors.New("too many")
	}

	if count < 0 {
		return nil, errors.New("no negative numbers")
	}

	out := make([]string, 0, count)
	for i := 0; i < count; i++ {
		hasLine := s.Scan()
		if !hasLine {
			return nil, errors.New("too few lines")
		}
		line := s.Text()
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			return nil, errors.New("blank line")
		}
		out = append(out, line)
	}
	return out, nil
}

func ToData(s []string) []byte {
	var b bytes.Buffer
	b.WriteString(strconv.Itoa(len(s)))
	b.WriteRune('\n')
	for _, v := range s {
		b.WriteString(v)
		b.WriteRune('\n')
	}
	return b.Bytes()
}

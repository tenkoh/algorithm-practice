package util

import (
	"bufio"
	"io"
	"strconv"
)

func NewScanner(r io.Reader) *bufio.Scanner {
	scan := bufio.NewScanner(r)
	scan.Split(bufio.ScanWords)
	return scan
}

func ScanInt(sc *bufio.Scanner) int {
	sc.Scan()
	i, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic("could not convert strint into int")
	}
	return i
}

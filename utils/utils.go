package utils

import (
	"bufio"
	"bytes"
	"os"
)

func GetFile(file string) string {

	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	wr := bytes.Buffer{}
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		wr.WriteString(sc.Text())
	}

	return wr.String()
}

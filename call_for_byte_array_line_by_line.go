package uhelpers

import (
	"bufio"
	"bytes"
)

func CallForByteArrayLineByLine(in []byte, f func(string, ...interface{}), prefix ...string) {
	bufReader := bufio.NewReader(bytes.NewReader(in))
	for {
		line, _, err := bufReader.ReadLine()
		if err != nil {
			break
		}

		if len(prefix) == 1 {
			f("%s%s", prefix[0], line)
		} else {
			f("%s", line)
		}
	}
}

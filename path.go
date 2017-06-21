package testutil

import (
	"bytes"
	"fmt"
	"os"
	"runtime"
	"strings"
)

func Getwd() string {
	if wd, err := os.Getwd(); err != nil {
		panic(err)
	} else {
		return wd
	}
}

func Where() string {
	_, file, line, _ := runtime.Caller(1)
	return fmt.Sprintf("%v:%v", file, line)
}

func EvalPath(d, p string) string {
	dirSlice := strings.Split(d, "/")
	pathSlice := strings.Split(p, "/")

	var dp = 0
	for i := range dirSlice {
		if dirSlice[i] != pathSlice[i] {
			dp = i
			break
		}
	}

	var buf = bytes.Buffer{}

	for i := 0; i < len(dirSlice)-dp; i++ {
		if i != 0 {
			buf.WriteString("/")
		}
		buf.WriteString("..")
	}

	for i := dp; i < len(pathSlice); i++ {
		buf.WriteString("/")
		buf.WriteString(pathSlice[i])
	}
	return buf.String()
}

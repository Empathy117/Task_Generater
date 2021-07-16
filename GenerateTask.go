package main

import (
	"io"
	"os"
)

func GenerateTask(test *Tests) {
	filename := "/home/empathy/in.txt"
	var f *os.File
	f, _ = os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC, 0600)
	_, _ = io.WriteString(f, test.Input)
}

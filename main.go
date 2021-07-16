package main

import (
	"bufio"
	"encoding/json"
	"io"
	"log"
	"net"
	"os/exec"
)

var Data Cases

type Input struct {
	Type string `json:"type,omitempty"`
}

type Output struct {
	Type string `json:"type,omitempty"`
}

type Tests struct {
	Input string `json:"input,omitempty"`
	Output string `json:"output,omitempty"`
}

type Java struct {
	MainClass string `json:"mainClass,omitempty"`
	TaskClass string `json:"taskClass,omitempty"`
}

type Batches struct {
	Id string `json:"id,omitempty"`
	Size int `json:"size,omitempty"`
}

type Languages struct {
	Java Java `json:"java,omitempty"`
}

type Cases struct {
	Name        string    `json:"name,omitempty"`
	Group       string    `json:"group,omitempty"`
	Url         string    `json:"url,omitempty"`
	Interactive bool      `json:"interactive,omitempty"`
	MemoryLimit int       `json:"memoryLimit,omitempty"`
	TimeLimit   int       `json:"timeLimit,omitempty"`
	TestCase    []Tests   `json:"tests"`
	TestType    string    `json:"testType,omitempty"`
	Input       Input     `json:"input,omitempty"`
	Output      Output    `json:"output,omitempty"`
	Langs       Languages `json:"languages"`
	Batch       Batches   `json:"batch"`
}

func handleConnection(conn net.Conn) {
	br := bufio.NewReader(conn)
	for{
		data, err := br.ReadString('\n')
		if err == io.EOF {
			err2 := json.Unmarshal([]byte(data), &Data)
			if err2 != nil {
				panic(err2)
			}
			cmd := exec.Command("rm", "*.t")
			_ = cmd.Run()
			GenerateTask(&Data.TestCase[0])
			break
		}
		log.Println(data)
	}
	_ = conn.Close()
}

func main() {
	ln, err := net.Listen("tcp", ":27121")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal("Some thing wrong: ", err)
		}
		go handleConnection(conn)
	}
}
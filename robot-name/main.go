package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"os"
	"sync"
	"strconv"
)

type Data struct {
	l sync.Mutex
	m map[string]bool
}

func (m *Data) saveMap() {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(m.m); err != nil {
		panic(err)
	}
	data := buf.Bytes()
	if err := ioutil.WriteFile("./data", data, 0777); err != nil {
		panic(err)
	}
}

func (m *Data) loadMap() {
	dat, err := ioutil.ReadFile("./data")
	if err != nil {
		panic(err)
	}
	dec := gob.NewDecoder(bytes.NewReader(dat))
	if err := dec.Decode(&m.m); err != nil {
		panic(err)
	}
}

func NewData() *Data {
	m := &Data{}
	m.l.Lock()
	if _, err := os.Stat("./data"); os.IsNotExist(err) {
		fmt.Printf("New\n")
		m.saveMap()
		m.l.Unlock()
		return m
	}
	fmt.Printf("Load\n")
	m.loadMap()
	m.l.Unlock()
	return m
}

func (m *Data) key(s string, b bool) {
	m.l.Lock()
	m.loadMap()
	m.m[s] = b
	m.saveMap()
	m.l.Unlock()
}

func (m *Data) AddKey(s string) {
	m.key(s, true)
}

func (m *Data) DelKey(s string) {
	m.key(s, true)
}

func main() {
	m := NewData()
	for i := 0; i < 10000; i++ {
		m.AddKey(strconv.Itoa(i))
		//fmt.Printf("File contents: %s\n", m.m)
		m.DelKey(strconv.Itoa(i))
		//fmt.Printf("File contents: %s\n", m.m)
	}
}

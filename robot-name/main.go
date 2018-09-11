package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"os"
	"sync"
)

type Data struct {
	l sync.Mutex
	m map[string]bool
}

func (m *Data) saveMap() {
	m.l.Lock()
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(m.m); err != nil {
		panic(err)
	}
	data := buf.Bytes()
	if err := ioutil.WriteFile("./data", data, 0777); err != nil {
		panic(err)
	}
	m.l.Unlock()
}

func (m *Data) loadMap() {
	m.l.Lock()
	dat, err := ioutil.ReadFile("./data")
	if err != nil {
		panic(err)
	}
	dec := gob.NewDecoder(bytes.NewReader(dat))
	if err := dec.Decode(&m.m); err != nil {
		panic(err)
	}
	m.l.Unlock()
}

func NewData() *Data {
	m := &Data{}
	if _, err := os.Stat("./data"); os.IsNotExist(err) {
		fmt.Printf("New\n")
		m.saveMap()
		return m
	}
	fmt.Printf("Load\n")
	m.loadMap()
	return m
}

func (m *Data) AddKey(s string) {
	m.l.Lock()
	m.m[s] = true
	m.l.Unlock()
}

func (m *Data) DelKey(s string) {
	m.l.Lock()
	m.m[s] = false
	m.l.Unlock()
}

func main() {
	m := NewData()
	m.loadMap()
	fmt.Printf("File contents: %s\n", m.m)
	m.AddKey("test")
	m.DelKey("test")
	m.AddKey("test2")
	m.saveMap()
	m.loadMap()
	fmt.Printf("File contents: %s\n", m.m)
}

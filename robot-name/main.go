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
	l sync.RWMutex
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
	m.l.RLock()
	dat, err := ioutil.ReadFile("./data")
	if err != nil {
		panic(err)
	}
	dec := gob.NewDecoder(bytes.NewReader(dat))
	if err := dec.Decode(&m.m); err != nil {
		panic(err)
	}
	m.l.RUnlock()
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

func main() {
	m := NewData()
	m.loadMap()
	fmt.Printf("File contents: %s\n", m.m)
	m.m = map[string]bool{"qwerty": true, "qwerty24": false}
	m.saveMap()
	m.loadMap()
	fmt.Printf("File contents: %s\n", m.m)
}

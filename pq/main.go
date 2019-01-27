package main

import (
	"fmt"
)

type list struct {
	element *element
	count   int
}

type element struct {
	power int
	name  string
	next  *element
}

func New() *list {
	return &list{element: &element{}, count: 0}
}

func (l *list) Insert(n int, s string) {
	fmt.Println()
	fmt.Printf("Add: %d -> %s\n", n, s)
	p := l.element
	if l.count == 0 {
		p.power = n
		p.name = s
		p.next = nil
		l.count++
		return
	}
	for {
		if p.power < n {
			old := &element{p.power, p.name, p.next}
			p.power = n
			p.name = s
			p.next = old
			l.count++
			return
		}
		if p.next == nil {
			p.next = &element{n, s, nil}
			l.count++
			return
		}
		p = p.next
	}
}

func (l *list) GetMax() {
	fmt.Println()
	if l.count == 0 {
		fmt.Println("Empty!")
		return
	}
	fmt.Printf("Max: %d -> %s\n", l.element.power, l.element.name)
	l.element = l.element.next
	l.count--
	return
}

func (l *list) Print() {
	fmt.Println()
	fmt.Printf("Count: %d\n", l.count)
	next := l.element
	for {
		if next == nil {
			break
		}
		fmt.Printf("%d -> %s\n", next.power, next.name)
		next = next.next
	}
}

func main() {
	pq := New()
	pq.Insert(0, "qwerty0")
	pq.Print()
	pq.Insert(3, "qwerty3")
	pq.Print()
	pq.Insert(2, "qwerty2")
	pq.Print()
	pq.Insert(3, "qwerty3")
	pq.Print()
	pq.Insert(10, "qwerty10")
	pq.Print()
	pq.GetMax()
	pq.Print()
	pq.GetMax()
	pq.Print()
}

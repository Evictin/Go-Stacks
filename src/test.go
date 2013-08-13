package main 

import (
	"fmt"
	"stack"
)

type Stuffer interface {
	Stuff()
}

type TypeA struct {
	a int
}

func (a TypeA) Stuff() {
	fmt.Println("A Stuff:", a.a)
}

type TypeB struct {
	Stuffer
}

func (b TypeB) Stuff() {
	fmt.Print("B Stuff:")
	b.Stuffer.Stuff()
}

func main() {
	fmt.Println("FILO Stack:\nPushing: \"Hello World\"")
	var s stack.Stack
	s.Push("Hello World")
	for i:=1; i<6; i++ {
		fmt.Println("Pushing:", i)
		s.Push(i)
	}
	
	for val, ok := s.Pop(); ok; val, ok = s.Pop() {
		fmt.Println("Popped:", val);
	}

	fmt.Println("\n===============================\nFIFO Queue:\nPushing: \"Hello World\"")
	var q stack.Queue
	q.Push("Hello World")
	for i:=1; i<6; i++ {
		fmt.Println("Pushing:", i)
		q.Push(i)
	}
	
	i := 0
	for val, ok := q.Pop(); ok; val, ok = q.Pop() {
		fmt.Println("Popped:", val);
		if i++; i==10 {
			break
		}
	}


	fmt.Println("\n===============================\nDecorator Test:")
	var c interface{}
	c = &TypeA{7}
	
	if d, ok := c.(Stuffer); ok {
		fmt.Println("c is a Stuffer")
		d.Stuff()
	} else {
		fmt.Println("c is NOT a Stuffer")
	}

	c = &TypeB{c.(Stuffer)}
	
	if d, ok := c.(Stuffer); ok {
		fmt.Println("c is a Stuffer")
		d.Stuff()
	} else {
		fmt.Println("c is NOT a Stuffer")
	}
}


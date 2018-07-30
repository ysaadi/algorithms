package main

import (
	"fmt"
  "math/rand"
  "math"
  "time"
)

type Node struct {
	Value int
}

// Stack is a basic LIFO stack that resizes as needed.
type Stack struct {
	nodes	[]*Node
	count	int
}

// Push adds a node to the stack.  Must manually grow splice if it isn't big enough.
func (s *Stack) Push(n *Node) {
	if s.count >= len(s.nodes) {
		nodes := make([]*Node, (len(s.nodes)+1)*2)
		copy(nodes, s.nodes)
		s.nodes = nodes
	}
	s.nodes[s.count] = n
	s.count++
}

// Pop removes and returns a node from the stack in last to first order.
func (s *Stack) Pop() *Node {
	if s.count == 0 {
		return nil
	}
	node := s.nodes[s.count-1]
	s.count--
	return node
}

// Queue is a basic FIFO queue based on a circular list that resizes as needed.
type Queue struct {
	nodes	[]*Node
	head	int
	tail	int
	count	int
}

// Push adds a node to the queue.
func (q *Queue) Push(n *Node) {
	if q.head == q.tail && q.count > 0 {
		nodes := make([]*Node, (len(q.nodes)+1)*2)
		copy(nodes, q.nodes[q.head:])
		copy(nodes[len(q.nodes)-q.head:], q.nodes[:q.head])
		q.head = 0
		q.tail = len(q.nodes)
		q.nodes = nodes
	}
	q.nodes[q.tail] = n
	q.tail = (q.tail + 1) % len(q.nodes)
	q.count++
}

// Pop removes and returns a node from the queue in first to last order.
func (q *Queue) Pop() *Node {
	if q.count == 0 {
		return nil
	}
	node := q.nodes[q.head]
	q.head = (q.head + 1) % len(q.nodes)
	q.count--
	return node
}

func main() {
  rand.Seed(time.Now().UTC().UnixNano())
	s := &Stack{nodes: make([]*Node, 0)}
  //parenRand := math.Round(rand.Float64() + .05)
  var parenIndex int
  var parenRand int
  for x:=0; x<20; x++ {
    parenRand = int(math.Round(rand.Float64()))
    //fmt.Printf("the Random num is %d\n", parenRand)
    if parenRand == 0 {
      s.Push(&Node{parenIndex})
      parenIndex+=1
      //fmt.Printf("the index is %d\n", parenIndex)
    } else{
      if s.count == 0{
        fmt.Printf("closed a Parenthesis without opening it at %d \n", parenIndex)
      } else{
      print(s.Pop().Value)
      parenIndex+=1
      }
    }
  }
  if s.count ==0{
    println("the parenthesis string was syntactically correct!")
  } else{
    fmt.Printf("A parenthesis was opened and not closed at %d", s.Pop().Value)
  }
  //
	// s.Push(&Node{1})
	// s.Push(&Node{2})
	// s.Push(&Node{3})
	// fmt.Printf("%v, %v, %v\n", s.Pop().Value, s.Pop().Value, s.Pop().Value)
  //
	// q := &Queue{nodes: make([]*Node, 3)}
	// q.Push(&Node{1})
	// q.Push(&Node{2})
	// q.Push(&Node{3})
	// fmt.Printf("%v, %v, %v\n", q.Pop().Value, q.Pop().Value, q.Pop().Value)
}

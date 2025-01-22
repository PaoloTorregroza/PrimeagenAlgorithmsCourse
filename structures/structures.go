package structures

import (
	"fmt"
	"unsafe"
)

func Array() {

	fmt.Println("ARRAYS")
	// Unbroken memory space

	// What are arrays?

	// Arrays are a continuous space in memory, that we can tell our code how to read.
	// EX: We can create a continuous space in memory of a total of 32 bits, and tell the compiler
	// we will read that memory in small spaces of 8 bits, each one containing a single value.
	spaceInMemory := make([]int8, 4) //32 bits continous memory space
	spaceInMemory[1] = 3
	fmt.Println(spaceInMemory)

	// This space in memory can be reinterpreted and readed 16bits at a time
	nowIts16 := (*[2]int16)(unsafe.Pointer(&spaceInMemory[0])) // We p
	nowIts16[1] = 0x6969                                       // this will write 16 bits of information (0x6969) after the first 16bits of the memory
	fmt.Println(spaceInMemory)

	// OUTPUT
	// [0 3 0 0]
	// [0 3 105 105]
	// 69 in hexadecimal is 105, we write 0x6969 because it is a number that
	// overflows the 8bit memory we configured in spaceInMemory

	// How do we get the value at an specific index?
	// We multiply the index by the bits offset (8 IE) and see what is at that point in memory
	//
	// How do we insert at an specific index?
	// We find the point in memory we want to write at, and write on the specified amount of bits our information
	//
	// How do we delete something?
	// Replacing the bits at an specific position with 0
}

// LINKED LISTS

// This can be known as a node based data structure, each node refers to the next one, in that way, data is
// organized in a non continuous way on memory, that means that adding and removing elements on this data structure
// is easier that in an array, but since memory is not continous it is going to take more time to read.
type Node[T any] struct {
	val  T
	next *Node[T]
	prev *Node[T]
}

type LinkedList[T any] struct {
	head *Node[T]
	tail *Node[T]
}

func (ll *LinkedList[T]) Add(val T) *Node[T] {
	l := ll.tail
	n := Node[T]{val: val, prev: l}

	if l != nil {
		l.next = &n
	}

	ll.tail = &n

	if ll.head == nil {
		ll.head = &n
	}

	return &n
}

func (ll *LinkedList[T]) Remove(node *Node[T]) {
	p := node.prev
	n := node.next

	p.next = n
	n.prev = p
	node.prev = nil
	node.next = nil
}

func (ll *LinkedList[T]) InsertAfter(node *Node[T], val T) *Node[T] {
	if node.next == nil {
		return ll.Add(val)
	}

	n := Node[T]{val: val, prev: node, next: node.next}
	node.next.prev = &n
	node.next = &n

	return &n
}

func (ll *LinkedList[T]) PrintAll() {
	current := ll.head
	for {
		if current == nil {
			break
		}

		fmt.Println(current)
		current = current.next
	}
}

func LinkedListShowcase() {
	myLl := LinkedList[int]{}

	nN := myLl.Add(23)

	myLl.Add(123)
	myLl.Add(69)
	toRemove := myLl.Add(123123123)
	myLl.InsertAfter(nN, 420)
	myLl.Add(1)

	myLl.Remove(toRemove)

	myLl.PrintAll()
}

// QUEUES

type QNode[T any] struct {
	value T
	next  *QNode[T]
}

type Queue[T any] struct {
	length int
	head   *QNode[T]
	tail   *QNode[T]
}

func (q *Queue[T]) Enqueue(item T) {
	q.length++

	n := QNode[T]{value: item}

	if q.length-1 == 0 {
		q.head = &n
		q.tail = &n
	} else {
		q.tail.next = &n
		q.tail = &n
	}
}

func (q *Queue[T]) Dequeue() T {
	if q.head == nil {
		var r T
		return r
	}

	q.length--

	h := q.head
	q.head = q.head.next

	// Non garbage collector languages should do a cleanup of h in here
	h.next = nil
	return h.value
}

func (q *Queue[T]) Peek() T {
	if q.head == nil {
		var r T
		return r
	}

	return q.head.value
}

func QueuesShowCase() {

}

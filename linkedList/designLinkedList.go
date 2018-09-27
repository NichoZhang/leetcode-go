package main

/**
* Your MyLinkedList object will be instantiated and called as such:
* obj := Constructor();
* param_1 := obj.Get(index);
* obj.AddAtHead(val);
* obj.AddAtTail(val);
* obj.AddAtIndex(index, val);
* obj.DeleteAtIndex(index);
*
* MyLinkedList linkedList = new MyLinkedList();
* linkedList.addAtHead(1);
* linkedList.addAtTail(3);
* linkedList.addAtIndex(1, 2);  // linked list becomes 1->2->3
* linkedList.get(1);            // returns 2
* linkedList.deleteAtIndex(1);  // now the linked list is 1->3
* linkedList.get(1);            // returns 3
 */

// MyLinkedList structure
type MyLinkedList struct {
	val  int
	next *MyLinkedList
}

// Constructor Initialize your data structure here.
func Constructor() MyLinkedList {
	return MyLinkedList{}
}

// Get the value of the index-th node in the linked list.
// If the index is ivalid, return -1.
func (m *MyLinkedList) Get(index int) int {
	if index < 0 {
		return -1
	}

	if m.next == nil {
		return -1
	}
	if index >= m.count() {
		return -1
	}

	//     doubleCheck := true
	for i := 0; i < index; i++ {
		m = m.next
	}

	return m.val
}

func (m *MyLinkedList) count() int {
	count := 0
	for m.next != nil {
		m = m.next
		count++
	}
	return count
}

// AddAtHead Add a node of value val before the first element of
// the linked list. After the insertion, the new node will be
// the first of the linked list.
func (m *MyLinkedList) AddAtHead(val int) {
	if m.next == nil {
		m.val = val
		m.next = &MyLinkedList{}
	} else {
		currNode := &MyLinkedList{
			val:  m.val,
			next: m.next,
		}
		m.val = val
		m.next = currNode
	}
}

// AddAtTail Add a node of value val to the last element of
// the linked list.
func (m *MyLinkedList) AddAtTail(val int) {
	for m.next != nil {
		m = m.next
	}
	m.next = new(MyLinkedList)
	m.val = val
}

// AddAtIndex Add a node of value before the index-th node in
// the linked list. If index equals to the length of linked list,
// the node will be appended to the end of linked list. If index
// is greater than the length, the node will not be inserted.
func (m *MyLinkedList) AddAtIndex(index int, val int) {
	if index == 0 {
		m.AddAtHead(val)
		return
	}

	if m.val == 0 && m.next == nil && index > 0 {
		return
	}

	for i := 0; i < index-1; i++ {
		m = m.next
	}
	currNode := &MyLinkedList{
		val:  val,
		next: m.next,
	}
	m.next = currNode
}

// DeleteAtIndex delete the index-th node in the linked list,
// if the index is valid.
func (m *MyLinkedList) DeleteAtIndex(index int) {
	if index >= m.count() {
		return
	}
	preNode := m
	nextNode := m
	for i := 0; i < index; i++ {
		nextNode = nextNode.next
	}
	for i := 0; i < index-1; i++ {
		preNode = preNode.next
	}
	preNode.next = nextNode.next
}

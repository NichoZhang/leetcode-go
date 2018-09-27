package main

import "testing"

func TestFunc1(t *testing.T) {
	linkedList := new(MyLinkedList)
	linkedList.AddAtHead(1)
	linkedList.AddAtTail(3)
	linkedList.AddAtIndex(1, 2)
	t.Log(linkedList.Get(1))
	linkedList.DeleteAtIndex(1)
	t.Log(linkedList.Get(1))
}

func TestFunc2(t *testing.T) {
	linkedList := new(MyLinkedList)
	linkedList.AddAtHead(1)
	linkedList.AddAtIndex(1, 2)
	t.Log(linkedList.Get(1))
	t.Log(linkedList.Get(0))
	t.Log(linkedList.Get(2))
}

func TestFunc3(t *testing.T) {
	linkedList := new(MyLinkedList)
	t.Log(linkedList.Get(0))
	linkedList.AddAtIndex(1, 2)
	t.Log(linkedList.Get(0))
	t.Log(linkedList.Get(1))
	linkedList.AddAtIndex(0, 1)
	t.Log(linkedList.Get(0))
	t.Log(linkedList.Get(1))
}

func TestFunc4(t *testing.T) {
	linkedList := new(MyLinkedList)
	linkedList.AddAtHead(56)
	linkedList.AddAtHead(41)
	linkedList.AddAtTail(98)
	linkedList.AddAtIndex(1, 33)
	linkedList.AddAtHead(72)
	linkedList.AddAtHead(52)
	linkedList.AddAtHead(89)
	linkedList.AddAtHead(0)
	linkedList.AddAtHead(98)
	linkedList.AddAtIndex(7, 97)
	linkedList.AddAtIndex(2, 51)
	linkedList.AddAtIndex(6, 49)

	linkedList.AddAtHead(72)
	linkedList.AddAtIndex(7, 8)
	linkedList.AddAtIndex(8, 58)
	linkedList.AddAtHead(10)
	linkedList.AddAtIndex(3, 6)
	linkedList.AddAtIndex(9, 61)
	linkedList.AddAtHead(63)

	linkedList.DeleteAtIndex(7)

	linkedList.AddAtIndex(16, 55)

	if linkedList.Get(4) != 6 {
		t.Fatal("wrong num")
	}
}

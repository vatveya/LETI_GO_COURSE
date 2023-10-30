package likedlist

import (
	"fmt"
	"strconv"
)

type node struct {
	val  int
	next *node
}

type linkedList struct {
	head *node
}

// ф-ция инициализации структуры с определенным кал-во элементов
func New(len int) *linkedList {
	newNode := &node{}
	newLinkedList := &linkedList{head: newNode}
	for i := 1; i < len; i++ {
		next := &node{}
		newNode.next = next
		newNode = next
		newNode.val = i
	}
	return newLinkedList
}

// размерность спсиска
func (l *linkedList) Size() int {
	newNode := l.head
	if l.head == nil {
		return 0
	}

	count := 1
	for newNode.next != nil {
		count += 1
		newNode = newNode.next
	}
	return count
}

// добавление значения в конец списка
func (l *linkedList) Add(val int) {
	newNode := l.head
	for newNode.next != nil {
		newNode = newNode.next
	}
	newNode.next = &node{val: val}
}

// удаление по заданной позиции
func (l *linkedList) DeleteFrom(position int) {
	newNode := l.head
	for i := 1; i < position; i++ {
		newNode = newNode.next
	}
	newNode.next = newNode.next.next
}

// обновление значение по позиции
func (l *linkedList) UpdateVal(val, position int) {
	newNode := l.head
	for i := 1; i < position; i++ {
		newNode = newNode.next
	}
	newNode.next.val = val
}

// получение значения по позиции
func (l *linkedList) GetVal(position int) int {
	newNode := l.head
	for i := 1; i < position; i++ {
		newNode = newNode.next
	}
	return newNode.next.val
}

// связный список из слайса
func LikedListFromSlice(s []int) *linkedList {
	newNode := &node{val: s[0]}
	newLinkedList := &linkedList{head: newNode}
	for _, i := range s[1:] {
		next := &node{}
		newNode.next = next
		newNode = next
		newNode.val = i

	}
	return newLinkedList
}

// вывод списка
func (l *linkedList) Print() {
	res := ""
	newNode := l.head
	for newNode.next != nil {
		res += strconv.Itoa(newNode.val) + " -> "
		newNode = newNode.next
	}
	res += strconv.Itoa(newNode.val)
	fmt.Println(res)
}

package singlylinkedlists

import (
	"errors"
)

var (
	errEmptyList = errors.New("list is empty")
	errIndexOutOfBounds = errors.New("index is out of bounds")
)

func New(values ...interface{}) *SinglyLinkedList {
	list := SinglyLinkedList{
		count: 0,
	}
	list.Add(values)
	return &list
}

// SLLNode singly linked list node
type SLLNode struct {
	Value interface{}
	Next *SLLNode
}

// SinglyLinkedList singly linked list structure
type SinglyLinkedList struct {
	count int
	head *SLLNode
	tail *SLLNode
}

// Size returns the count of the list
func (s *SinglyLinkedList) Size() int {
	return s.count
}

// IsEmpty checks if the list is empty 
func (s *SinglyLinkedList) IsEmpty() bool {
	return s.Size() == 0
}

// Clear clears list
func (s *SinglyLinkedList) Clear() {
	current := s.head
	if current != nil {
		tmp := current
		current = current.Next
		tmp.Next = nil
		tmp = nil
	}
	
	current = nil
	s.count = 0
	s.head = nil
	s.tail = nil
}

// GetHead returns head node of the list
func (s *SinglyLinkedList) GetHead() *SLLNode {
	return s.head
}

// GetLastValue returns last node value of the list
func (s *SinglyLinkedList) GetLastValue() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errEmptyList
	}
	return s.tail.Value, nil
}

// GetFirstValue returns head node value of the list
func (s *SinglyLinkedList) GetFirstValue() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errEmptyList
	}
	return s.head.Value, nil
}

// GetValues returns values
func (s *SinglyLinkedList) GetValues() []interface{} {
	values := make([]interface{}, 0, s.Size())
	current := s.head
	for current != nil {
		values = append(values, current.Value)
		current = current.Next
	}
	current = nil
	return values
}

// GetIndexOf returns the index of the first occurence
func (s *SinglyLinkedList) GetIndexOf(value interface{}) int {
	index := 0
	current := s.head
	for current != nil {
		if value == current.Value {
			return index
		}
		current = current.Next
		index++
	}
	current = nil
	return -1
}

// Add add to the list
func (s *SinglyLinkedList) Add(values ...interface{}) {
	for _, value := range values {
		s.InsertAt(s.Size(), value)
	}
}

// InsertAt insert value at specific index in the list
func (s *SinglyLinkedList) InsertAt(index int, value interface{}) error {
	if index < 0 || index > s.count {
		return errIndexOutOfBounds
	}

	element := &SLLNode{Value: value}
	if s.IsEmpty() {
		s.head = element
		s.tail = element
		s.count++
		element = nil
		return nil
	}

	if index == 0 {
		element.Next = s.head
		s.head = element
		s.count++
		element = nil
		return nil
	}

	if index == s.count {
		s.tail.Next = element
		s.tail = element
		s.count++
		element = nil
		return nil
	}

	current, err := s.getNode(index - 1)
	if err != nil {
		return err
	}

	element.Next = current.Next
	current.Next = element
	s.count++
	current = nil
	element = nil
	return nil
}

// RemoveAt remove value from the list at specific index
func (s *SinglyLinkedList) RemoveAt(index int) (interface{}, error) {
	if index < 0 || index > s.count {
		return nil, errIndexOutOfBounds
	}

	if s.count == 1 {
		value := s.head.Value
		s.head = nil
		s.tail = nil
		s.count--
		return value, nil
	}

	if index == 0 {
		current := s.head
		s.head = current.Next
		current.Next = nil
		value := current.Value
		s.count--
		current = nil
		return value, nil
	}

	current, err := s.getNode(index - 1)
	if err != nil {
		return nil, err
	}

	value := current.Next.Value
	current.Next = current.Next.Next
	if index == s.count {
		s.tail = current
	}
	s.count--
	current = nil
	return value, nil
}

// GetValueAt returns value at specific index in the list
func (s *SinglyLinkedList) GetValueAt(index int) (interface{}, error) {
	current, err := s.getNode(index)
	if err != nil {
		return nil, err
	}
	return current.Value, nil
}

// getNode returns node that at the specific index
func (s *SinglyLinkedList) getNode(index int) (*SLLNode, error) {
	if index < 0 || index > s.count {
		return nil, errIndexOutOfBounds
	}

	current := s.head
	for i := 0; i < index; i++ {
		current = current.Next
	}
	return current, nil
}

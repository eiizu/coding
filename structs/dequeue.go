package structs

type node struct {
	next *node
	prev *node
	data int
}

type Deque struct {
	head *node
	tail *node
}

func (dq *Deque) PushFront(data int) {
	if dq.head == nil {
		dq.head = &node{
			data: data,
		}
		dq.tail = dq.head
		return
	}
	tmp1 := dq.head
	tmp2 := &node{
		data: data,
		next: tmp1,
	}
	dq.head = tmp2
	return
}

func (dq *Deque) PushBack(data int) {
	if dq.tail == nil {
		dq.tail = &node{
			data: data,
		}
		dq.head = dq.tail
		return
	}
	tmp := dq.tail
	dq.tail.next = &node{
		data: data,
		prev: tmp,
	}
	dq.tail = dq.tail.next
	return
}

func (dq *Deque) PopFront() (int, error) {
	var data int
	if dq.head != nil {
		data = dq.head.data
		dq.head = dq.head.next
	}
	return data, nil
}

func (dq *Deque) PopBack() (int, error) {
	var data int
	if dq.tail != nil {
		data = dq.tail.data
		dq.tail = dq.tail.prev
	}
	return data, nil
}

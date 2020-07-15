package gstl

import "errors"

//queue 先进先出

type queue struct {
	valList []interface{}
	count   int
	size    int
}

func NewQueue(size int) *queue {
	return &queue{[]interface{}{}, 0, size}
}

//in val
func (q *queue) Push(v interface{}) error {
	if q.size <= q.count {
		return errors.New("out of queue size ")
	}
	q.valList = append(q.valList, v)
	q.count += 1
	return nil
}

//out val
func (q *queue) Pop() interface{} {
	if q.count <= 0 {
		return nil
	}
	v := q.valList[0]
	q.valList = q.valList[1:]
	q.count -= 1
	return v
}

func (q *queue) Length() int {
	return q.count
}

func (q *queue) IsEmpty() bool {
	return q.count == 0
}

func (q *queue) IsFull() bool {
	return q.count >= q.size
}

//获取队头
func (q *queue) GetFront() interface{} {
	if q.count <= 0 {
		return nil
	}
	v := q.valList[0]
	return v
}

//获取队尾
func (q *queue) GetBack() interface{} {
	if q.count <= 0 {
		return nil
	}
	v := q.valList[q.count-1]
	return v
}

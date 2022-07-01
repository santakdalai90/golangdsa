package queue

import (
    "errors"
    "fmt"
)

type Queue struct {
    Data []interface{}    // top towards the end
}

func New() *Queue {
    q := Queue {}

    q.Data = make([]interface{}, 0)

    return &q
}

func (q *Queue) Push(item interface{}) error {
    q.Data = append(q.Data, item)
    return nil
} 

func (q *Queue) Pop() (interface{}, error) {
    if q == nil || len(q.Data) == 0 {
        return nil, errors.New("empty queue, cannot pop")
    }
    front := q.Data[0]
    q.Data = q.Data[1:]
    return front, nil
}

func (q *Queue) Front() (interface{}, error) {
    if q == nil || len(q.Data) == 0 {
        return nil, errors.New("empty queue")
    }
    front := q.Data[0]
    return front, nil
}

func (q *Queue) Print() {
    if q == nil || len(q.Data) == 0 {
        fmt.Println("queue is empty")
        return
    }

    for _,data := range q.Data {
        fmt.Println(data)
    }
}
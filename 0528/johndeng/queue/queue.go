package queue

import "errors"

type Queue struct {
	bucket []string
	space  int
}

func (q *Queue) Push(data string) error {
	if q.space == 0 {
		return errors.New("Queue full!")
	}

	q.bucket = append(q.bucket, data)
	q.space--
	return nil
}

func (q *Queue) Pop() (string, error) {

	if len(q.bucket) == 0 {
		return "", errors.New("Empty queue")
	}

	data := q.bucket[0]

	q.bucket = q.bucket[1:len(q.bucket)]
	q.space++

	return data, nil

}

func (q *Queue) IsFull() bool {
	return q.space == 0
}

func NewQueue(len int) *Queue {
	return &Queue{
		space: len,
	}
}

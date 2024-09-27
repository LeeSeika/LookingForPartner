package localqueue

type Queue struct {
	inChan  chan interface{}
	outChan chan interface{}
	data    []interface{}
}

func NewQueue(chanCap, dataCap int) *Queue {
	q := &Queue{
		inChan:  make(chan interface{}, chanCap),
		outChan: make(chan interface{}, chanCap),
		data:    make([]interface{}, dataCap),
	}
	go q.transfer()

	return q
}

func (q *Queue) Push(val interface{}) {
	q.inChan <- val
}

func (q *Queue) Pop() interface{} {
	return <-q.outChan
}

func (q *Queue) popData() interface{} {
	if len(q.data) == 0 {
		return nil
	}

	val := q.data[0]
	q.data = q.data[1:]

	return val
}

func (q *Queue) outChanWrapper() chan interface{} {
	if len(q.data) == 0 {
		return nil
	}

	return q.outChan
}

func (q *Queue) transfer() {
	for {
		select {
		case val := <-q.inChan:
			q.data = append(q.data, val)
		case q.outChanWrapper() <- q.popData():
		}
	}
}

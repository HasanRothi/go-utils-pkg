package structsutils

func Enqueue(queue []uint64, element uint64) []uint64 {
	queue = append(queue, element) // Simply append to enqueue.
	return queue
}

func Dequeue(queue []uint64) (uint64, []uint64) {
	element := queue[0] // The first element is the one to be dequeued.
	if len(queue) == 1 {
		var tmp []uint64
		return element, tmp

	}
	return element, queue[1:] // Slice off the element once it is dequeued.
}

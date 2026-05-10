package deque

type Deque[T any] struct {
	items []T
}

func (d *Deque[T])PushBack(v T){
	d.items = append(d.items, v)
}

func (d *Deque[T])PushFront(v T){
	d.items = append([]T{v}, d.items...)
}

func (d *Deque[T])Size()int{
	return len(d.items)
}

func (d *Deque[T])At(index int)T{
	return d.items[index]
}

func New[T any ](values ...T)Deque[T]{
	return Deque[T]{
		items: values,
	}
}
package semaphore

type ticket struct{}

type Semaphore chan ticket

func New(capacity int) Semaphore {
	s := make(chan ticket, capacity)
	for i := 0; i < capacity; i++ {
		s <- ticket{}
	}
	return s
}

func (s Semaphore) P() {
	<-s
}

func (s Semaphore) V() {
	s <- ticket{}
}

func (s Semaphore) Empty() {
	for len(s) > 0 {
		<-s
	}
}

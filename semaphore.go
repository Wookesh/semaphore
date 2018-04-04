package semaphore

type ticket struct{}

type Semaphore chan ticket

func New(capacity int) Semaphore {
	c := make(chan ticket)
	for i := 0; i < capacity; i++ {
		c <- ticket{}
	}
	return c
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

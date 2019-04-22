package fanout

func Split(ch <-chan int, n int) []<-chan int {
	cs := make([]chan int, 0)

	for i := 0; i < n; i++ {
		cs = append(cs, make(chan int))
	}

	distributeChannels := func(ch <-chan int, cs []chan<- int) {
		defer func(cs []chan<- int) {
			for _, c := range cs {
				close(c)
			}
		}(cs)

		for {
			for _, c := range cs {
				select {
				case val, ok := <-ch:
					if !ok {
						return
					}
					c <- val
				}
			}
		}
	}

	go distributeChannels(ch, cs)

	return cs
}

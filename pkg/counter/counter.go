package counter

import "sync"

var (
	counter int
	m       sync.Mutex
)

func Get() int {
	m.Lock()
	defer m.Unlock()
	counter++
	return counter
}

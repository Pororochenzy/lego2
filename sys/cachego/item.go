package cachego

import "time"

type keyAndValue struct {
	key   string
	value interface{}
}
type Item struct {
	Object     interface{}
	Expiration int64
}

func (item Item) Expired() bool {
	if item.Expiration == 0 {
		return false
	}
	return time.Now().UnixNano() > item.Expiration
}

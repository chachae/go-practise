package singleton

import "sync"

type singleton3 struct {
}

var instance3 *singleton3
var lock = sync.Mutex{}

func getInstance3() *singleton3 {

	if instance3 == nil {
		lock.Lock()
		if instance3 == nil {
			instance3 = &singleton3{}
		}
		lock.Unlock()
	}

	return instance3

}

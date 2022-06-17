package singleton

/**
线程不安全
*/
type singleton1 struct {
}

var instance1 *singleton1

func getInstance() *singleton1 {
	if instance1 == nil {
		instance1 = &singleton1{}
	}
	return instance1
}

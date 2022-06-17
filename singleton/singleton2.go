package singleton

type singleton2 struct {
}

var instance2 *singleton2 = &singleton2{}

func getSingleton2() *singleton2 {
	return instance2
}

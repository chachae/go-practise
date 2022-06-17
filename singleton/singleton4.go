package singleton

import "sync"

/**
使用once.Do可以确保 ins 实例全局只被创建一次
once.Do 函数还可以确保当同时有多个创建动作时，只有一个创建动作在被执行。
*/
type singleton4 struct {
}

var instance4 *singleton4
var once sync.Once

func getInstance4() *singleton4 {
	once.Do(func() {
		instance4 = &singleton4{}
	})
	return instance4
}

package strategy

import (
	"fmt"
	"testing"
)

type IStrategy interface {
	calculate(a int, b int) int
}

type add struct {
}

func (*add) calculate(a int, b int) int {
	return a + b
}

type reduce struct {
}

func (*reduce) calculate(a int, b int) int {
	return a - b
}

type Operator struct {
	strategy IStrategy
}

// 设置策略
func (operator *Operator) setStrategy(strategy IStrategy) { operator.strategy = strategy }

func (operator *Operator) calculate(a, b int) int {
	return operator.strategy.calculate(a, b)
}

func TestStrategy(t *testing.T) {

	t.Run("testAdd", func(t *testing.T) {
		operator := Operator{}
		operator.setStrategy(&add{})
		fmt.Println(operator.calculate(100, 200))
	})

	t.Run("testReduce", func(t *testing.T) {
		operator := Operator{}
		operator.setStrategy(&reduce{})
		fmt.Println(operator.calculate(100, 200))
	})

}

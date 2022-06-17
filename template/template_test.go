package template

import (
	"fmt"
	"testing"
)

type Cooker interface {
	fire()
	cooke()
	closeFire()
}

type CookMenu struct {
}

func (CookMenu) fire() {
	fmt.Println("open fire...")
}

func (CookMenu) closeFire() {
	fmt.Println("close fire...")
}

func (CookMenu) cooke() {

}

func doCook(cooker Cooker) {
	cooker.fire()
	cooker.cooke()
	cooker.closeFire()
}

type cookTomato struct {
	CookMenu
}

func (*cookTomato) cooke() {
	fmt.Println("cooking tomato...")
}

type cookEge struct {
	CookMenu
}

func (*cookEge) cooke() {
	fmt.Println("cooking ege...")
}

func TestFunc(t *testing.T) {
	tomato := &cookTomato{}
	doCook(tomato)
	ege := &cookEge{}
	doCook(ege)
}

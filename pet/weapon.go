package pet

import (
	"fmt"
	"math/rand"
)

const (
	ATTR_CRIT = 1 // 暴击
	ATTR_FLASH = 2 // 闪避
	ATTR_REST = 3 // 休息
	ATTR_LOCK = 4 // 必中
	ATTR_CONT = 5 // 连续
	ATTR_DODGE = 6 // 闪避


	KIND_S = 0
	KIND_M = 1
	KIND_L = 2
	KIND_T = 3 // 投掷类
)

type attribute struct {
	Kind int
	Probability int
}

type weapon struct {
	Id int
	Name string
	MinAttack int
	MaxAttack int
	Attribute []attribute
	Probability int
	Level int
	Kind int
}



func (w *weapon)use(p1, p2 *Pet) string {
	hp := p2.hp
	attack := w.MinAttack + rand.Intn(w.MaxAttack - w.MinAttack)
	p2.hp -= attack
	return fmt.Sprintf(ATTACK_DESC, p1.name, p1.hp, w.Name, p2.name, hp, attack)
	
}
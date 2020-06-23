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

var Weapons = []weapon {
	{
		Id:0,
		Name: "平底锅",
		MinAttack: 18,
		MaxAttack: 22,
		Kind: KIND_S,
	},
	{
		Id:1,
		Name: "板砖",
		MinAttack: 3,
		MaxAttack: 8,
		Kind: KIND_M,
	},
	{
		Id:2,
		Name: "接力棒",
		MinAttack: 6,
		MaxAttack: 10,
		Kind: KIND_S,
	},
	{
		Id:3,
		Name: "汽水罐",
		MinAttack: 4,
		MaxAttack: 6,
		Kind: KIND_T,
	},

}


func (w *weapon)use(p1, p2 *pet) string {
	hp := p2.hp
	attack := w.MinAttack + rand.Intn(w.MaxAttack - w.MinAttack)
	p2.hp -= attack
	return fmt.Sprintf(ATTACK_DESC, p1.name, p1.hp, w.Name, p2.name, hp, attack)
	
}
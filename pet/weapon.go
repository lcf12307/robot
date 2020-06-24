package pet

import (
	"fmt"
	"math/rand"
)

type Attribute struct {
	name        string
	effect      func(p1, p2 *Pet) string
	Kind        int
	Probability int
}

type Weapon struct {
	Id          int
	Name        string
	MinAttack   int
	MaxAttack   int
	Attribute   []Attribute
	Probability int
	Level       int
	Kind        int
}

func (w *Weapon) use(p1, p2 *Pet) string {
	hp := p2.hp

	// todo 检查自己的状态
	// todo 检查敌人的状态
	// todo 检查武器效果
	// todo 校验是否闪避

	// 攻击对方，扣除血量
	attack := w.MinAttack + rand.Intn(w.MaxAttack-w.MinAttack)
	p2.hp -= attack
	return fmt.Sprintf(ATTACK_DESC, p1.name, p1.hp, w.Name, p2.name, hp, attack)

}

func (w *Weapon) crit(p1, p2 *Pet) string {
	w.MaxAttack = w.MaxAttack * 3 / 2
	return fmt.Sprintf(CRIT_DESC, p1.name)
}

func (w *Weapon) blood(p1, p2 *Pet) string {
	p1.status.Blood = attr{round: 1}
	return ""
}

func (w *Weapon) continuous(p1, p2 *Pet) string {
	p1.status.Cont = attr{
		round: 1,
		num:   100,
	}
	return ""
}

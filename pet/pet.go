package pet

import (
	"fmt"
	"math/rand"
)

type attr struct {
	num   int
	round int
}
type Status struct {
	IsStop    int  // 是否停止行动n回合
	IsReflect bool // 是否反弹伤害
	AttUp     int
	WeaUp     int
	SpeUp     int
	Fake      bool
	Reduc     int
	Recov     int
}
type Pet struct {
	name     string
	hp       int
	po       int
	ag       int
	sp       int
	level    int
	wps      []int
	sks      []int
	title    string
	addAttr  map[string]attr
	ranks    rank
}


func Adopt(name string) *Pet {
	nUser := &Pet{
		name:  name,
		hp:    50 + rand.Intn(50),
		po:    5 + rand.Intn(5),
		ag:    5 + rand.Intn(5),
		sp:    5 + rand.Intn(5),
		level: 5 + rand.Intn(5),
		wps:   nil,
		sks:   nil,
		title: "乐斗小菜",
	}
	for _, w := range Weapons {
		nUser.wps = append(nUser.wps, w.Id)
	}
	// todo 写入文件
	return nUser
}

func (p *Pet) LevelUp() {
	p.hp += rand.Intn(10)
	p.po += rand.Intn(3)
	p.ag += rand.Intn(3)
	p.sp += rand.Intn(3)
	p.level++
	// todo add title, 写入文件
}

func GroupFight(pets []*Pet) string {
	round := 1
	res := GROUP_FIGHT_TITLE
	for _, p := range pets {
		res += fmt.Sprintf(ATTR_DESC, p.name, p.hp, p.po, p.ag, p.sp)
	}
	// 目前是假设速度一直不变的情况下。
	pets = SortPetsBySp(pets)
	alive := len(pets)
	for {
		res += fmt.Sprintf(ROUND_TITLE, round)
		for i, p := range pets {
			// 跳过死人
			if p.hp <= 0 {
				continue
			}

			e := rand.Intn(len(pets))
			// 打到自己有概率加血或者停止
			if e == i {
				coin := rand.Intn(2)
				if coin == 1 {
					hp := pets[i].hp / 4
					pets[i].hp += hp
					res += fmt.Sprintf(HP_UP_DESC, p.name, hp)
				} else {
					res += fmt.Sprintf(TIRED_DESC, p.name)
				}
				continue
			}
			// 打到了尸体
			if pets[e].hp <= 0 {
				res += fmt.Sprintf(ATTACK_BODY_DESC, p.name)
				continue
			}
			w1 := p._chooseWeapon()
			res += w1.use(p, pets[e])
			if pets[e].hp <= 0 {
				res += fmt.Sprintf(BEAT_TITLE, p.name, pets[e].name)
				WriteRank(p, pets[e])
				alive --
				if alive == 1  {
					res += fmt.Sprintf(WIN_TITLE, p.name)
					return res
				}
			}
		}

		round ++
	}
}
func SortPetsBySp(pets []*Pet) []*Pet {
	l := len(pets)
	for i:=0; i<l; i++ {
		for j:=i; j<l; j++ {
			if pets[j].sp > pets[i].sp {
				pets[i], pets[j] = pets[j], pets[i]
			}
		}
	}
	return pets
}

func (p *Pet) Pk(ep *Pet) string {
	round := 1
	res := ""
	res += fmt.Sprintf(PK_TITLE, p.name, ep.name)
	res += fmt.Sprintf(ATTR_DESC, p.name, p.hp, p.po, p.ag, p.sp)
	res += fmt.Sprintf(ATTR_DESC, ep.name, ep.hp, ep.po, ep.ag, ep.sp)
	for {
		res += fmt.Sprintf(ROUND_TITLE, round)
		if p.sp > ep.sp {
			w1 := p._chooseWeapon()
			res += w1.use(p, ep)
			if ep.hp <= 0 {
				res += fmt.Sprintf(BEAT_TITLE, p.name, ep.name)
				WriteRank(p, ep)
				return res
			}

			w2 := ep._chooseWeapon()
			res += w2.use(ep, p)
			if p.hp <= 0 {
				res += fmt.Sprintf(BEAT_TITLE, ep.name, p.name)
				WriteRank(ep, p)
				return res
			}
		} else {
			w2 := ep._chooseWeapon()
			res += w2.use(ep, p)
			if p.hp <= 0 {
				res += fmt.Sprintf(BEAT_TITLE, ep.name, p.name)
				WriteRank(ep, p)
				return res
			}

			w1 := p._chooseWeapon()
			res += w1.use(p, ep)
			if ep.hp <= 0 {
				res += fmt.Sprintf(BEAT_TITLE, p.name, ep.name)
				WriteRank(p, ep)
				return res
			}
		}
		round++
	}
}

func (p *Pet) _chooseWeapon() *weapon {
	w := rand.Intn(len(p.wps))
	return &Weapons[w]
}




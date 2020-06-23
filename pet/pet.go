package pet

import (
	"fmt"
	"math/rand"
)

type attr struct {
	num int
	round int
}
type Status struct {
	IsStop int

}
type pet struct {
	name string
	hp int
	po int
	ag int
	sp int
	level int
	wps []int
	sks []int
	title string
	addAttr map[string]attr
}



func Adopt(name string) *pet {
	nUser := &pet{
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
	for _, w := range Weapons  {
		nUser.wps = append(nUser.wps, w.Id)
	}
	// todo 写入文件
	return nUser
}

func (p *pet)LevelUp() {
	p.hp += rand.Intn(10)
	p.po += rand.Intn(3)
	p.ag += rand.Intn(3)
	p.sp += rand.Intn(3)
	p.level ++
	// todo add title, 写入文件
}

func (p *pet)Pk(ep *pet) string {
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
				return res
			}

			w2 := ep._chooseWeapon()
			res += w2.use(ep, p)
			if p.hp <= 0 {
				res += fmt.Sprintf(BEAT_TITLE, ep.name, p.name)
				return res
			}
		} else {
			w2 := ep._chooseWeapon()
			res += w2.use(ep, p)
			if p.hp <= 0 {
				res += fmt.Sprintf(BEAT_TITLE, ep.name, p.name)
				return res
			}

			w1 := p._chooseWeapon()
			res += w1.use(p, ep)
			if ep.hp <= 0 {
				res += fmt.Sprintf(BEAT_TITLE, p.name, ep.name)
				return res
			}
		}
		round ++
	}
}

func (p *pet)_chooseWeapon() *weapon {
	w := rand.Intn(len(p.wps))
	return &Weapons[w]
}


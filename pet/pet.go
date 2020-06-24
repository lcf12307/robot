package pet

import (
	"fmt"
	"math/rand"
)

type attr struct {
	num   int
	round int
	from  string
}
type Status struct {
	IsStop    attr // 是否停止行动n回合
	IsReflect attr // 是否反弹伤害
	AttUp     attr // 提升攻击力
	WeaUp     attr // 提升力量
	SpeUp     attr // 提升速度
	Fake      attr // 装死
	Cont      attr //
	Recov     attr
	Blood     attr
}
type consumable struct {
	id     int
	status bool
}
type Pet struct {
	name   string
	hp     int
	po     int
	ag     int
	sp     int
	level  int
	wps    []consumable
	posks  []consumable
	pasks  []consumable
	title  string
	status Status
	ranks  rank
}

/**
 * @desc 获取一只新的pet
 */
func Adopt(name string) *Pet {
	nUser := &Pet{
		name:  name,
		hp:    30 + rand.Intn(30),
		po:    5 + rand.Intn(5),
		ag:    5 + rand.Intn(5),
		sp:    5 + rand.Intn(5),
		level: 5 + rand.Intn(5),
		wps:   nil,
		posks: nil,
		pasks: nil,
		title: "乐斗小菜",
	}
	for _, w := range Weapons {
		nUser.wps = append(nUser.wps, consumable{id: w.Id})
	}
	// todo 写入文件
	return nUser
}

/**
 * 升级
 */
func (p *Pet) LevelUp() {
	p.hp += rand.Intn(10)
	p.po += rand.Intn(3)
	p.ag += rand.Intn(3)
	p.sp += rand.Intn(3)
	p.level++
	// todo add title, 写入文件
}

/**
 * @desc 群架
 */
func GroupFight(pets []*Pet) string {
	round := 1

	var s string
	var killed bool
	res := GROUP_FIGHT_TITLE
	// 输出每个参赛选手文案
	for _, p := range pets {
		res += fmt.Sprintf(ATTR_DESC, p.name, p.hp, p.po, p.ag, p.sp)
	}
	// 目前是假设速度一直不变的情况下。
	pets = _sortPetsBySp(pets)
	alive := len(pets)
	for {
		res += fmt.Sprintf(ROUND_TITLE, round)
		for i, p := range pets {
			// 跳过死人
			if p.hp <= 0 {
				continue
			}
			// 随机匹配敌人
			e := rand.Intn(len(pets))
			// 打到自己有概率加血或者停止行动一回合
			if e == i {
				e = rand.Intn(len(pets))
				if e == i {
					coin := rand.Intn(2)
					if coin == 1 {
						hp := rand.Intn(60 - pets[i].hp)
						pets[i].hp += hp
						res += fmt.Sprintf(HP_UP_DESC, p.name, hp)
					} else {
						res += fmt.Sprintf(TIRED_DESC, p.name)
					}
					continue
				}
			}

			// 打到了尸体的概率太高了， 重新摇一下
			// 3个人比较适合这种方法， 但是更多人的话概率太高，需要调整
			// todo 如果将来开更多人的话， 死掉的就从数组中剔除，届时可以增加闪避
			if pets[e].hp <= 0 {
				e = rand.Intn(len(pets))
				if e == i || pets[e].hp <= 0 {
					res += fmt.Sprintf(ATTACK_BODY_DESC, p.name)
					continue
				}
			}

			s, killed = p._attack(pets[e])
			res += s
			if killed {
				alive--
				if alive <= 1 {
					res += fmt.Sprintf(WIN_TITLE, p.name)
					return res
				}
			}

		}

		round++
	}
}

/**
 * @desc 单挑
 */
func (p *Pet) Pk(ep *Pet) string {
	// round 当前回合数
	var s string
	var killed bool
	round := 1
	// res 结果文案 添加标题、每个人的属性
	res := ""
	res += fmt.Sprintf(PK_TITLE, p.name, ep.name)
	res += fmt.Sprintf(ATTR_DESC, p.name, p.hp, p.po, p.ag, p.sp)
	res += fmt.Sprintf(ATTR_DESC, ep.name, ep.hp, ep.po, ep.ag, ep.sp)
	for {
		// 记录回合数
		res += fmt.Sprintf(ROUND_TITLE, round)
		// 比较速度， 速度快的优先使用
		if p.sp > ep.sp {
			s, killed = p._attack(ep)
			res += s
			if killed {
				return res
			}
			s, killed = ep._attack(p)
			res += s
			if killed {
				return res
			}
		} else {
			s, killed = ep._attack(p)
			res += s
			if killed {
				return res
			}

			s, killed = p._attack(ep)
			res += s
			if killed {
				return res
			}

		}
		round++
	}
}

/**
 * @desc 选择武器
 */
func (p *Pet) _chooseWeapon() *Weapon {
	w := rand.Intn(len(p.wps))
	return &Weapons[w]
}

/**
 * @desc 按速度进行排序
 */
func _sortPetsBySp(pets []*Pet) []*Pet {
	l := len(pets)
	for i := 0; i < l; i++ {
		for j := i; j < l; j++ {
			if pets[j].sp > pets[i].sp {
				pets[i], pets[j] = pets[j], pets[i]
			}
		}
	}
	return pets
}

func (p *Pet) _attack(ep *Pet) (string, bool) {
	// todo 是否触发进攻者技能
	// todo 是否触发被进攻者技能
	// todo 决定使用武器还是使用技能

	// 选择自己的武器
	w1 := p._chooseWeapon()
	// 使用武器去攻击其他人
	s := w1.use(p, ep)
	var killed bool
	// 判断是否死亡
	if ep.hp <= 0 {
		// 记录最终数据、写入文件
		s += fmt.Sprintf(BEAT_TITLE, p.name, ep.name)
		WriteRank(p, ep)
		killed = true
	}
	return s, killed
}

// todo 开局前，检查被动技能， 并标记
// 使用责任链模式，依次调用对应的技能函数
func (p *Pet) _checkStatus() string {
	for _, skill := range p.pasks {
		Skills[skill.id].Effect(p, p)
	}
	return ""
}

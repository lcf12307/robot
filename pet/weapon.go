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

// todo 写入文件
var Weapons = []weapon {
	{
		Id:0,
		Name: "平底锅 :fried_egg: ",
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
	{
		Id:4,
		Name: "短剑",
		MinAttack: 3,
		MaxAttack: 8,
		Kind: KIND_S,
	},
	{
		Id:5,
		Name: "木剑",
		MinAttack: 10,
		MaxAttack: 25,
		Kind: KIND_S,
	},
	{
		Id:6,
		Name: "判官笔",
		MinAttack: 5,
		MaxAttack: 8,
		Kind: KIND_S,
	},
	{
		Id:7,
		Name: "流星球",
		MinAttack: 15,
		MaxAttack: 24,
		Kind: KIND_T,
	},
	{
		Id:8,
		Name: "老鼠 :rat:",
		MinAttack: 5,
		MaxAttack: 8,
		Kind: KIND_T,
	},
	{
		Id:9,
		Name: "小李飞刀",
		MinAttack: 5,
		MaxAttack: 10,
		Kind: KIND_T,
	},
	{
		Id:10,
		Name: "折凳",
		MinAttack: 11,
		MaxAttack: 13,
		Kind: KIND_M,
	},
	{
		Id:11,
		Name: "铁铲",
		MinAttack: 12,
		MaxAttack: 18,
		Kind: KIND_M,
	},
	{
		Id:12,
		Name: "环扣刀",
		MinAttack: 12,
		MaxAttack: 13,
		Kind: KIND_M,
	},
	{
		Id:13,
		Name: "红缨枪",
		MinAttack: 15,
		MaxAttack: 30,
		Kind: KIND_M,
	},
	{
		Id:14,
		Name: "双截棍",
		MinAttack: 9,
		MaxAttack: 13,
		Kind: KIND_M,
	},
	{
		Id:15,
		Name: "宽刃剑",
		MinAttack: 6,
		MaxAttack: 10,
		Kind: KIND_M,
	},
	{
		Id:16,
		Name: "幻影枪",
		MinAttack: 20,
		MaxAttack: 40,
		Kind: KIND_L,
	},
	{
		Id:17,
		Name: "木槌",
		MinAttack: 7,
		MaxAttack: 12,
		Kind: KIND_L,
	},
	{
		Id:18,
		Name: "棒球棒",
		MinAttack: 15,
		MaxAttack: 20,
		Kind: KIND_L,
	},
	{
		Id:19,
		Name: "狂魔镰",
		MinAttack: 15,
		MaxAttack: 25,
		Kind: KIND_L,
	},
	{
		Id:20,
		Name: "关刀",
		MinAttack: 20,
		MaxAttack: 35,
		Kind: KIND_L,
	},
	{
		Id:21,
		Name: "开山斧",
		MinAttack: 12,
		MaxAttack: 18,
		Kind: KIND_L,
	},
	{
		Id:22,
		Name: "充气锤子",
		MinAttack: 20,
		MaxAttack: 35,
		Kind: KIND_L,
	},
	{
		Id:23,
		Name: "三叉戟",
		MinAttack: 25,
		MaxAttack: 50,
		Kind: KIND_L,
	},
	{
		Id:24,
		Name: "青龙戟",
		MinAttack: 15,
		MaxAttack: 20,
		Kind: KIND_M,
	},
}


func (w *weapon)use(p1, p2 *pet) string {
	hp := p2.hp
	attack := w.MinAttack + rand.Intn(w.MaxAttack - w.MinAttack)
	p2.hp -= attack
	return fmt.Sprintf(ATTACK_DESC, p1.name, p1.hp, w.Name, p2.name, hp, attack)
	
}
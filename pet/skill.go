package pet

type Skill struct {
	Id     int
	Name   string
	Effect func(p1, p2 *Pet) string
	Kind   int
}

func Glue(p1, p2 *Pet) string {
	p2.status.IsStop = attr{
		num:   100,
		round: 3,
		from:  "glue",
	}
	return ""
}

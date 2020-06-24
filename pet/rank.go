package pet

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)
type rank struct {
	Name string
	Win int
	Lose int
	Point int
}

const filename  = "data/rank/%s"
func WriteRank(wp, lp *Pet)  {
	ranks := GetRank()
	var wFlag, lFlag bool
	for i, r := range ranks {
		// 胜利
		if r.Name == wp.name {
			ranks[i].Win ++
			ranks[i].Point ++
			wFlag = true
		}
		if r.Name == lp.name {
			ranks[i].Lose ++
			ranks[i].Point --
			lFlag = true
		}
	}
	if !wFlag {
		ranks = append(ranks, rank{
			Name:  wp.name,
			Win:   1,
			Lose:  0,
			Point: 1,
		})
	}
	if !lFlag {
		ranks = append(ranks, rank{
			Name:  lp.name,
			Win:   0,
			Lose:  1,
			Point: -1,
		})
	}
	ranks = SortRank(ranks)

	file := _getFileName()
	f, err := os.OpenFile(file, os.O_WRONLY, 0600)
	defer f.Close()
	if err != nil  {
		fmt.Println(err.Error())
	} else {
		rankJson, err := json.Marshal(ranks)

		_, err = f.Write(rankJson)
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	return
}

func GetRank() []rank {
	file := _getFileName()
	f, err := os.OpenFile(file, os.O_CREATE,0600)
	res := []rank{}
	defer f.Close()
	if err !=nil {
		fmt.Println(err.Error())
	} else {
		contentByte, err :=ioutil.ReadAll(f)
		if err != nil {
			fmt.Println(err.Error())
		}
		if len(contentByte) == 0 {
			return res
		}
		err  = json.Unmarshal(contentByte, &res)
		if err != nil {
			fmt.Println(err.Error())
		}
		return res
	}
	return res
}

func MyRank(s string) string {
	ranks := GetRank()
	for _, r := range ranks {
		if r.Name == s {
			return fmt.Sprintf(RANK_DESC, r.Name, r.Point, r.Win, r.Lose)
		}
	}
	return "你还没有对战信息，快去打一场吧！"
}

func SortRank(ranks []rank) []rank {
	l := len(ranks)
	for i:=0;i<l;i++ {
		for j:=i+1; j<l; j++  {
			if ranks[i].Point < ranks[j].Point {
				ranks[i], ranks[j] = ranks[j], ranks[i]
			}
		}
	}
	res := []rank{}
	for _, r := range ranks{
		if r.Name != "" {
			res = append(res, r)
		}
	}
	return res
}

func _getFileName() string {
	file := fmt.Sprintf(filename, time.Now().Format("2006-01-02"))
	return file
}

func RankResult() string {
	ranks := GetRank()
	res := fmt.Sprintf(RANK_TITLE, time.Now().Month(), time.Now().Day())
	for _, r := range ranks {
		res += fmt.Sprintf(RANK_DESC, r.Name, r.Point, r.Win, r.Lose)
	}
	return res
}
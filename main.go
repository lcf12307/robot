package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"math/rand"
	"net/http"
	"regexp"
	"robot/pet"
	"strings"
	"time"
)

const appKey = "0780a11ecdb6263d"
const youDongKey = "dc1a7874b6eda27445d4"
const youDOngSec = "a50daad969f8382ad39e"

func main() {
	r := gin.Default()
	r.POST("/test", func(c *gin.Context) {
		c.JSON(200, setOutPut("test", false, nil))
	})
	// 历史上的今天
	r.POST("/today", func(c *gin.Context) {
		res := getTodayInHistory()
		sss := setOutPut(res, false, nil)
		c.JSON(200, sss)
	})

	//问答机器人
	r.POST("/ask", func(c *gin.Context) {
		params := handleParam(c)
		ques := handleText(params["text"].(string))
		answ := getAnswer(ques)
		sss := setOutPut(answ, false, nil)
		c.JSON(200, sss)
	})
	// 知乎日报
	r.POST("zhihu", func(c *gin.Context) {
		text, attach := getZhihu()
		sss := setOutPut(text, false, attach)
		c.JSON(200, sss)
	})
	r.POST("ps", func(c *gin.Context) {
		params := handleParam(c)
		text := handleText(params["text"].(string))
		res := getPsResult(text)
		sss := setOutPut("", false, res)
		c.JSON(200, sss)
	})
	r.POST("dd", func(c *gin.Context) {
		params := handleParam(c)
		texts := strings.Split(handleText(params["text"].(string)), "@")
		if len(texts) == 0 {
			sss := setOutPut("发送文字有误，请重新发送", false, nil)
			c.JSON(200, sss)
			return
		}
		var res string
		// 根据分割后的字符来判断处理方式
		// 正常方式应该拆分多个接口比较好，为了方便机器人管理，写到一个接口里了
		switch texts[0] {
		// fight 功能
		case "":
			// 如果只有一个人的话 使用单挑功能
			if len(texts) == 2 {
				ename := strings.TrimSpace(texts[1])
				mname := params["user_name"].(string)

				//todo 每个人的属性，武器，都写入文件， 每次使用时读取
				// 目前每次对战都要重新领养
				p := pet.Adopt(mname)
				ep := pet.Adopt(ename)

				// 单挑函数
				res = p.Pk(ep)
				// 多人打架功能， 后期可以把单挑也封装进来。
			} else {
				mname := params["user_name"].(string)
				p := pet.Adopt(mname)
				pets := []*pet.Pet{
					p,
				}
				// 批量领取多人的武器
				// todo 每个人的属性都写入文件。
				for i := 1; i < len(texts); i++ {
					pets = append(pets, pet.Adopt(strings.TrimSpace(texts[i])))
				}
				// 群架函数
				res = pet.GroupFight(pets)
			}

			break
		// 获取rank榜
		case "rank":
			res = pet.RankResult()
			break
		// 获取自己的积分
		case "me":
			res = pet.MyRank(params["user_name"].(string))
			break
		}

		sss := setOutPut(res, false, nil)
		c.JSON(200, sss)
	})
	r.Run(":9999") // listen and serve on 0.0.0.0:8080
}

//处理字符串 去掉展示有问题的字符
func handleString(s string) string {
	s = strings.TrimSpace(s)
	s = strings.Replace(s, "</br>", "\n", -1)
	return s
}

// 处理json传的参数， 返回
func handleParam(c *gin.Context) map[string]interface{} {
	res := map[string]interface{}{}
	data, _ := ioutil.ReadAll(c.Request.Body)
	_ = json.Unmarshal(data, &res)
	return res
}

// 正则匹配出关键字之后的字符
func handleText(s string) string {
	r, _ := regexp.Compile(" .*")
	res := handleString(r.FindString(s))
	return res
}

// 设置输出的必要字段
func setOutPut(text string, isM bool, atts []map[string]interface{}) interface{} {
	res := map[string]interface{}{}

	if text != "" {
		res["text"] = text
	}

	res["markdown"] = isM

	if atts != nil {
		res["attachments"] = atts
	}
	return res
}

// 封装 get请求数据
func getResult(url string) map[string]interface{} {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)

	}

	defer resp.Body.Close()
	s, err := ioutil.ReadAll(resp.Body)

	// 处理返回数据
	var res map[string]interface{}
	err = json.Unmarshal(s, &res)
	return res
}

// 读取知乎日报信息
func getZhihu() (string, []map[string]interface{}) {
	url := "https://v1.alapi.cn/api/zhihu/latest"
	tmp := getResult(url)
	tmp2 := tmp["data"].(map[string]interface{})
	tmp3 := tmp2["stories"].([]interface{})
	i := rand.Intn(len(tmp3))
	tmp4 := tmp3[i].(map[string]interface{})
	attach := []map[string]interface{}{
		{
			"url":   tmp4["url"],
			"title": tmp4["title"],
		},
	}
	images := tmp4["images"].([]interface{})
	if len(images) > 0 {
		attach[0]["images"] = []map[string]interface{}{
			{
				"url":    images[0],
				"height": 150,
				"width":  240,
			},
		}
	}
	return tmp4["hint"].(string), attach
}

// 获取问答机器人
func getAnswer(ques string) string {
	url := "https://api.jisuapi.com/iqa/query?appkey=%s&question=%s"
	url = fmt.Sprintf(url, appKey, ques)
	tmp := getResult(url)
	tmp2 := tmp["result"].(map[string]interface{})
	return handleString(tmp2["content"].(string))
}

// 获取历史上的今天
func getTodayInHistory() string {
	url := "https://api.jisuapi.com/todayhistory/query?appkey=%s&month=%d&day=%d"
	url = fmt.Sprintf(url, appKey, time.Now().Month(), time.Now().Day())
	tmp := getResult(url)
	tmp2 := tmp["result"].([]interface{})
	tmp3 := tmp2[rand.Intn(len(tmp2))].(map[string]interface{})
	template := "%s年%s月%s日, %s"
	return fmt.Sprintf(template, tmp3["year"].(string), tmp3["month"].(string), tmp3["day"].(string), tmp3["title"].(string))
}

// 获取自动给图片增加文字结果
func getPsResult(s string) []map[string]interface{} {
	url := "http://api.guaqb.cn/v1/ps/?picurl=http://img.tukexw.com/img/9624c318fa9bfd4c.jpg&text=%s&key=%s&secret=%s&fontSize=25&circleSize=0&left=30&top=180"
	url = fmt.Sprintf(url, s, youDongKey, youDOngSec)
	return []map[string]interface{}{
		{
			"url":    url,
			"height": 150,
			"width":  240,
		},
	}
}

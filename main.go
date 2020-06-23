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
type apiResp struct {
	status int
	msg	string
	result []todayResp
}
type todayResp struct {
	title string
	year int
	month int
	day int
	content string
}
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
		sss  := setOutPut(answ, false, nil)
		c.JSON(200, sss)
	})
	// 知乎日报
	r.POST("zhihu", func(c *gin.Context) {
		text, attach :=getZhihu()
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
		ename := texts[1]
		mname := params["user_name"].(string)

		p := pet.Adopt(mname)
		ep := pet.Adopt(ename)
		res := p.Pk(ep)
		sss := setOutPut(res, false, nil)
		c.JSON(200, sss)
	})
	r.Run(":9999") // listen and serve on 0.0.0.0:8080
}
func handleString(s string) string {
	s = strings.Trim(s, " ")
	s = strings.Replace(s, "</br>", "\n", -1)
	return s
}
func handleParam(c *gin.Context) map[string]interface{} {
	res := map[string]interface{}{}
	data, _ := ioutil.ReadAll(c.Request.Body)
	_ = json.Unmarshal(data, &res)
	return res
}
func handleText(s string) string {
	r, _ := regexp.Compile(" .*")
	res := handleString(r.FindString(s))
	return res
}
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
func getResult(url string) map[string]interface{} {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)

	}

	defer resp.Body.Close()
	s,err:=ioutil.ReadAll(resp.Body)

	// 处理返回数据
	var res map[string]interface{}
	err = json.Unmarshal(s, &res)
	return res
}

func getZhihu() (string, []map[string]interface{}) {
	url := "https://v1.alapi.cn/api/zhihu/latest"
	tmp := getResult(url)
	tmp2 := tmp["data"].(map[string]interface{})
	tmp3 := tmp2["stories"].([]interface{})
	i := rand.Intn(len(tmp3))
	tmp4 := tmp3[i].(map[string]interface{})
	attach := []map[string]interface{}{
		{
			"url": tmp4["url"],
			"title": tmp4["title"],
		},
	}
	images := tmp4["images"].([]interface{})
	if len(images) > 0 {
		attach[0]["images"] = []map[string]interface{}{
			{
				"url": images[0],
				"height": 150,
				"width": 240,
			},
		}
	}
	return tmp4["hint"].(string), attach
}

func getAnswer(ques string) string {
	url := "https://api.jisuapi.com/iqa/query?appkey=%s&question=%s"
	url = fmt.Sprintf(url, appKey, ques)
	tmp := getResult(url)
	tmp2 := tmp["result"].(map[string]interface{})
	return handleString(tmp2["content"].(string))
}

func getTodayInHistory() string {
	url := "https://api.jisuapi.com/todayhistory/query?appkey=%s&month=%d&day=%d"
	url = fmt.Sprintf(url, appKey, time.Now().Month(), time.Now().Day())
	tmp := getResult(url)
	tmp2 := tmp["result"].([]interface{})
	tmp3 := tmp2[rand.Intn(len(tmp2))].(map[string]interface{})
	template := "%s年%s月%s日, %s"
	return fmt.Sprintf(template, tmp3["year"].(string), tmp3["month"].(string), tmp3["day"].(string), tmp3["title"].(string))
}

func getPsResult(s string) []map[string]interface{} {
	url := "http://api.guaqb.cn/v1/ps/?picurl=http://img.tukexw.com/img/9624c318fa9bfd4c.jpg&text=%s&key=%s&secret=%s&fontSize=25&circleSize=0&left=30&top=180"
	url = fmt.Sprintf(url, s, youDongKey, youDOngSec)
	return []map[string]interface{}{
		{
			"url": url,
			"height": 150,
			"width": 240,
		},
	}
}
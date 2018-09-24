package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//Video 视频
type Video struct {
	VideoURL    string `json:"video_url"`
	Duration    string `json:"duration"`
	ImageURL    string `json:"img_url"`
	ImageHeight int    `json:"img_height"`
	ImageWidth  int    `json:"img_width"`
	Size        string `json:"size"`
	Summary     string `json:"summary"`
}

//Image 图片
type Image struct {
	PicId     string `json:"pic_id"`
	ThumbURL  string `json:"des_url"`
	OriginURL string `json:"origin_url"`
	Width     int    `json:"origin_width"`
	Height    int    `json:"origin_height"`
	IsGif     bool   `json:"gif"`
}

//Blog 推送信息
type Blog struct {
	ID     string  `json:"mid"`
	Time   string  `json:"time"`
	Title  string  `json:"text"`
	Images []Image `json:"imgs"`
	Video  `json:"video"`
}

//XianZhiResponseData 鲜知返回的 Base Data
type XianZhiResponseData struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    struct {
		Cards []struct {
			MBLogs []Blog `json:"mblogs"`
		} `json:"cards"`
	} `json:"data"`
}

func main() {

	girlsParser()

}

//睡前一张美女图
func girlsParser() {

	fetchBlogsWithCardID("7554037350176769")
}

//头像壁纸看这里
func avatarWallpaperParser() {
	fetchBlogsWithCardID("7556011155036161")
}

//表情包，请自取
func facialBags() {
	fetchBlogsWithCardID("7554120768591873")
}

//今日最热足球资讯
func footballNews() {
	fetchBlogsWithCardID("7554897573091329")
}

//NBA资讯
func nbaNews() {
	fetchBlogsWithCardID("7556274885007361")
}

func cuteAnimals() {
	fetchBlogsWithCardID("")
}

//今日份沙雕
func todaySB() {
	fetchBlogsWithCardID("7555014120740865")
}

func fetchBlogsWithCardID(cardID string) {

	url := fmt.Sprintf("https://top.weibo.cn/xz/subject/show?ac=Wi-Fi&cardid=%v", cardID)
	res, err := http.Get(url)
	if err != nil {
		fmt.Printf("request error:%v", err)
	}

	bs, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		fmt.Printf("read data error:%v", err)
	}
	data := XianZhiResponseData{}
	jsonizationErr := json.Unmarshal(bs, &data)
	if jsonizationErr != nil {
		fmt.Printf("jsonization error:%v", jsonizationErr)
	}

	fmt.Printf("%+v", data.Data.Cards)

}

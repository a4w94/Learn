package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gocolly/colly"
)

var (
	userinfo       = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.159 Safari/537.36"
	visiturl       = "https://www.ambassador.com.tw/home/MovieList?Type=0"
	savefileroutie = "./testfile"
)

func main() {

	CheckDir(savefileroutie)

	//GetImg(savefileroutie + "/")
	c := colly.NewCollector() // 在colly中使用 Collector 這類物件 來做事情

	GetHtml(c)
	OnRequest(c)

	c.Visit(visiturl) // Visit 要放最後
}

func CheckResponse(c *colly.Collector) {

	//在發起請求之前，可以預先對Header的參數進行設定
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("1")
	})

	//如果在請求的時候發生錯誤
	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("2")
	})

	c.Head(visiturl)
	{
		fmt.Println("3")
	}
	//收到響應回復的時候
	c.OnResponse(func(r *colly.Response) {
		fmt.Println("4")
	})

	//收到的響應是HTML格式時（時間點比 OnResponse還晚），進行goquerySelector篩選
	c.OnHTML("body", func(e *colly.HTMLElement) {
		fmt.Println("5")
	})

	//收到的響應是XML格式時（時間點比 OnHTML還晚），進行xpathQuery篩選
	c.OnXML("//footer", func(e *colly.XMLElement) {
		fmt.Println("6")
	})

	//抓取網頁（與OnResponse相仿），但在最後才進行調用
	c.OnScraped(func(r *colly.Response) {
		fmt.Println("7")
	})

}

func OnRequest(c *colly.Collector) {
	c.OnRequest(func(r *colly.Request) { // iT邦幫忙需要寫這一段 User-Agent才給爬
		r.Headers.Set("User-Agent", userinfo)
	})
}

func GetAllHtml(c *colly.Collector) {
	c.OnResponse(func(r *colly.Response) { // 當Visit訪問網頁後，網頁響應(Response)時候執行的事情
		fmt.Println(string(r.Body)) // 返回的Response物件r.Body 是[]Byte格式，要再轉成字串
	})
}

var num int

func GetHtml(c *colly.Collector) {

	c.OnHTML(".cell", func(h *colly.HTMLElement) {
		// fmt.Println("1", h)
		// fmt.Println("2", h.Text)
		num++
		fmt.Println(num)
		h.ForEach("img[src]", downlodaImg)
		//GetImg(h.Text)
		//fmt.Println("tmp:", h.Text)
		// tmp := h.Attr("href")
		// fmt.Printf("tmp: %v\n", tmp)
	})
}
func downlodaImg(_ int, e *colly.HTMLElement) {
	srcRef := e.Attr("src")
	fullurl := srcRef
	res, _ := http.Get(fullurl)
	numstr := strconv.Itoa(num)
	savepath := savefileroutie + "/" + numstr + ".jpg"

	f, err := os.Create(savepath)
	if err != nil {
		panic(err)
	}
	io.Copy(f, res.Body)
	fmt.Println(srcRef)
}

func GetImg(imgUrl string) {
	//imgUrl := "https://www.ambassador.com.tw/assets/img/movies/EscapeRoomTournamentofChampions_180x270_Poster.jpg"

	//get data
	resp, err := http.Get(imgUrl)
	if err != nil {
		fmt.Println("fail")
		panic(err)
	} else {
		fmt.Println("獲取文件成功")
	}

	//defer resp.Body.Close()

	//創建一個文件 用於保存
	out, err := os.Create(savefileroutie + "/" + "testimg.jpg")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("創建文件成功")
	}

	//defer out.Close()

	_, err1 := io.Copy(out, resp.Body)
	if err != nil {
		panic(err1)
	} else {
		fmt.Println("文件複製成功")
	}
}

func CheckDir(path string) error {
	if _, err := os.Stat(path); err == nil {
		fmt.Println("資料夾已存在:", path)
		return nil
	} else {
		err := os.MkdirAll(path, 0711)
		if err != nil {
			return err
		} else {
			fmt.Println("創建成功:", path)
		}
	}
	_, err := os.Stat(path)

	filepath.Join()
	return err
}

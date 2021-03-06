package main

//test 111
//https://ithelp.ithome.com.tw/articles/10234820
import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"

	user "pack/src"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/yaml.v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	loginport = ":8000"
)

var openurl = "https://opendata.cwb.gov.tw/api/v1/rest/datastore/F-C0032-001?Authorization=CWB-096B3BC7-414C-4C11-A879-DBF28C469749&format=JSON"
var openurl2 = "https://api.openweathermap.org/data/2.5/onecall?lat=33.44&lon=-94.04&exclude=hourly,daily&appid=38eeafcc769d7625ba7de456d5f3eb5f"

//38eeafcc769d7625ba7de456d5f3eb5f

type Sport struct {
	ErrorCode  string       `json:"errorCode"`
	TotalCount int          `json:"totalCount"`
	Data       []DataDetail `json:"data"`
}

type DataDetail struct {
	ActivityNo              string `json:"activityNo"`
	ActivityKind            string `json:"activityKind"`
	ActivityType            string `json:"activityType"`
	ActivityName            string `json:"activityName"`
	ActivityCounty          string `json:"activityCounty"`
	ActivityContents        string `json:"activityContents"`
	ActivityDateBegin       string `json:"activityDateBegin"`
	ActivityDateEnd         string `json:"activityDateEnd"`
	ActivityOrganizer       string `json:"activityOrganizer"`
	ActivityContactName     string `json:"activityContactName"`
	ActivityContactMobileNo string `json:"activityContactMobileNo"`
	ActivityContactTelNo    string `json:"activityContactTelNo"`
	ActivityParticipants    string `json:"activityParticipants"`
	ActivityWebsit          string `json:"activityWebsit"`
}

type News struct {
	Status       string `json:"status"`
	TotalResults int    `json:"totalResults"`
	Articles     []struct {
		Source struct {
			Id   string `json:"id"`
			Name string `json:"name"`
		} `json:"source"`
		Author      string `json:"author"`
		Title       string `json:"title"`
		Description string `json:"description"`
		Url         string `json:"url"`
		UrltoImage  string `json:"urlToImage`
		PublishedAt string `json:"publishedAt"`
		Content     string `json:"content"`
	} `json:"articles"`
}

//https://newsapi.org/
func HandleGetData(c *gin.Context) {
	var sporturl = "https://newsapi.org/v2/everything?q=tesla&from=2021-07-31&sortBy=publishedAt&apiKey=5a88407f70554b379f9000506371942d"
	resq, err := http.Get(sporturl)
	if err != nil {
		panic(err)
	}
	body, _ := ioutil.ReadAll(resq.Body)

	var mapresult Sport
	err1 := json.Unmarshal(body, &mapresult)
	if err1 != nil {
		fmt.Println("jsontomap fail")
	}

	maxcount := mapresult.TotalCount
	arr := make([][]string, maxcount)
	returnmap := make(map[string]interface{})

	for k, v := range mapresult.Data {

		// fmt.Println(k)
		// fmt.Println(v)
		// fmt.Println()
		tmparr := []string{}
		key := reflect.TypeOf(v)    //??????
		value := reflect.ValueOf(v) //?????????
		for i := 0; i < key.NumField(); i++ {
			tmparr = append(tmparr, value.Field(i).String())
			//fmt.Println(value.Field(i).String())
		}
		arr[k] = tmparr
		//fmt.Println(tmp)

	}
	returnmap["Thead"] = []string{"??????", "Kind", "Type", "Name", "Country", "Contenes", "DateBegin", "DateEnd", "Organizer", "ContactName", "MobileNO", "ContactTelNo", "Participants", "Websit"}
	returnmap["Data"] = arr

	c.HTML(http.StatusOK, "index.html", returnmap)

}
func HandleGetDataTest() {
	var sporturl = "https://isports.sa.gov.tw/Api/Rest/V1/Activity.svc/GetActivityList?county=A&activityKind=1&paging=false"
	resq, err := http.Get(sporturl)
	if err != nil {
		panic(err)
	}
	body, _ := ioutil.ReadAll(resq.Body)
	//fmt.Println(string(body))

	var mapresult Sport
	err1 := json.Unmarshal(body, &mapresult)
	if err1 != nil {
		fmt.Println("jsontomap fail")
	}

	maxcount := mapresult.TotalCount
	arr := make([][]string, maxcount)
	returnmap := make(map[string]interface{})

	for k, v := range mapresult.Data {

		// fmt.Println(k)
		// fmt.Println(v)
		// fmt.Println()
		tmparr := []string{}
		key := reflect.TypeOf(v)    //??????
		value := reflect.ValueOf(v) //?????????
		for i := 0; i < key.NumField(); i++ {
			tmparr = append(tmparr, value.Field(i).String())
			//fmt.Println(value.Field(i).String())
		}
		arr[k] = tmparr
		//fmt.Println(tmp)

	}
	returnmap["Data"] = arr
	for i := 0; i < 10; i++ {
		//fmt.Println(returnmap["Data"][i])
	}

	//StructToMap(mapresult)
	//input2 := input["Data"]
	//test := StructToMap(input2)

	//fmt.Printf("%+v", input["Data"])

	//fmt.Println(test)

}

func HandleGetDataNewsTest() {
	var newsurl = "https://newsapi.org/v2/everything?q=bitcoin&apiKey=5a88407f70554b379f9000506371942d"
	resq, err := http.Get(newsurl)
	if err != nil {
		panic(err)
	}
	body, _ := ioutil.ReadAll(resq.Body)
	//fmt.Println(string(body))

	var mapresult News
	err1 := json.Unmarshal(body, &mapresult)
	if err1 != nil {
		fmt.Println("jsontomap fail")
	}
	//fmt.Println(string(body))

	//fmt.Println(mapresult.Articles[0])
	arr := make([][]string, mapresult.TotalResults)

	for k, v := range mapresult.Articles {

		arr[k] = append(arr[k], v.Source.Id, v.Source.Name, v.Author, v.Title, v.Description, v.Url, v.UrltoImage, v.Content)

	}

	//returnmapresult := make(map[string]interface{})

	//returnmapresult[]
	//fmt.Println(arr[1])

	// 	fmt.Println(k)
	// 	//fmt.Println(v)
	// 	// fmt.Println()
	// 	tmparr := []string{}
	// 	key := reflect.TypeOf(v)    //??????
	// 	value := reflect.ValueOf(v) //?????????
	// 	for i := 0; i < key.NumField(); i++ {
	// 		tmparr = append(tmparr, value.Field(i).String())
	// 		//fmt.Println(value.Field(i).String())
	// 	}
	// 	arr[k] = tmparr
	// 	//fmt.Println(tmp)

	// }

	//StructToMap(mapresult)
	//input2 := input["Data"]
	//test := StructToMap(input2)

	//fmt.Printf("%+v", input["Data"])

	//fmt.Println(test)

}
func NewsSearchPage(c *gin.Context) {
	c.HTML(http.StatusOK, "news.html", nil)

}
func HandleGetDataNews(c *gin.Context) {

	keyword, _ := c.GetPostForm("keyword")
	fmt.Println(keyword)
	url := fmt.Sprintf("https://newsapi.org/v2/everything?q=%s&from=%s&to=%s&sortBy=%s&apiKey=5a88407f70554b379f9000506371942d", keyword)
	fmt.Println(url)

	returnmap := NewDataGETAPI()

	c.HTML(http.StatusOK, "news.html", returnmap)

}

func NewDataGETAPI() map[string]interface{} {
	var newsurl = "https://newsapi.org/v2/everything?q=bitcoin&apiKey=5a88407f70554b379f9000506371942d"
	resq, err := http.Get(newsurl)
	if err != nil {
		panic(err)
	}
	body, _ := ioutil.ReadAll(resq.Body)
	//fmt.Println(string(body))

	var mapresult News
	err1 := json.Unmarshal(body, &mapresult)
	if err1 != nil {
		fmt.Println("jsontomap fail")
	}
	//fmt.Println(string(body))

	//fmt.Println(mapresult.Articles[0])
	arr := make([][]string, mapresult.TotalResults)

	for k, v := range mapresult.Articles {

		arr[k] = append(arr[k], v.Source.Id, v.Source.Name, v.Author, v.Title, v.Description, v.Url, v.UrltoImage, v.Content)

	}

	returnmap := make(map[string]interface{})
	returnmap["DetailData"] = arr
	returnmap["TableTitle"] = []string{"Source.Id", "Source.Name", "Author", "Title", "Description", "Url", "UrltoImage", "Content"}

	return returnmap
}

func StructToMap(obj Sport) map[string]interface{} {
	obj1 := reflect.TypeOf(obj)
	obj2 := reflect.ValueOf(obj)

	data := make(map[string]interface{})
	for i := 0; i < obj1.NumField(); i++ {
		fmt.Println(obj1.Field(i).Name)
		data[obj1.Field(i).Name] = obj2.Field(i).Interface()
	}
	return data
}

func main() {

	// apiserver := gin.Default()
	// apiserver.GET("/get", HandleGet)
	// apiserver.POST("/getdata", HandleGetData)
	// apiserver.GET("/first", test)
	// apiserver.Run(":8000")

	// var openurl = "https://ws.kinmen.gov.tw/001/Upload/0/relfile/0/0/843657b0-8f15-40bc-ba68-5dd5eec0cc5f.json"

	// gormtest()
	//HandleGetDataTest()
	//HandleGetDataNewsTest()
	i := 1
	if i == 1 {
		// mysqltest()
		// mytest.InitDB()
		// gormtest()

		server := gin.Default()
		//server.GET("/sport", HandleGetData)

		server.LoadHTMLGlob("template/html/*")        //??????html????????????
		server.Static("/assets", "./template/assets") //??????css????????????
		//server.GET("/test", test)
		server.GET("/news", NewsSearchPage)
		server.POST("/news", HandleGetDataNews)
		// server.GET("/login", LoginPage)
		// server.POST("/login", LoginAuth)
		server.Run(loginport)
	}
}

type IndexData struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

const (
	USERNAME = "root"
	PASSWORD = "a4w941207!!"
	NETWORK  = "tcp"
	SERVER   = "127.0.0.1"
	PORT     = 3306
	DATABASE = "UserLog"
)

type User struct {
	ID       int64  `json:"id" gorm:"primary_key;auto_increase'"`
	Username string `json:"username"`
	Password string `json:""`
}

func mysqltest() {
	conn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
	db, err := sql.Open("mysql", conn)
	if err != nil {
		fmt.Println("?????? MySQL ?????????????????????????????????", err)
		return
	}
	if err := db.Ping(); err != nil {
		fmt.Println("????????????????????????????????????", err.Error())
		return
	}
	defer db.Close()
}

func gormtest() {
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("?????? gorm ?????? DB ???????????????????????? " + err.Error())
	}
	fmt.Println("gorm connect success")

	if err := db.AutoMigrate(new(User)); err != nil {
		panic("database migrate failed, reason is" + err.Error())
	}

	user := &User{}
	user.Username = "test"
	user.Password = "test"

	if err := CreateUser(db, user); err != nil {
		panic("?????? user ?????????????????? " + err.Error())
	}
	if user, err := FindUser(db, 1); err != nil {
		panic("??????user??????????????????" + err.Error())
	} else {
		log.Println("????????????????????????", user)
	}

}

func CreateUser(db *gorm.DB, user *User) error {
	return db.Create(user).Error
}
func FindUser(db *gorm.DB, id int64) (*User, error) {
	user := new(User)
	user.ID = id
	err := db.First(&user).Error
	return user, err
}

func LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)

}
func LoginAuth(c *gin.Context) {

	var (
		username string
		password string
	)

	if in, isExist := c.GetPostForm("username"); isExist && in != "" {
		username = in
	} else {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"error": errors.New("???????????????????????????"),
		})
		return
	}

	if in, isExist := c.GetPostForm("password"); isExist && in != "" {
		password = in
	} else {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"error": errors.New("??????????????????"),
		})
		return
	}

	if err := user.Auth(username, password); err == nil {

		// c.HTML(http.StatusOK, "login.html", gin.H{
		// 	"success": "????????????",
		// })
		// data := IndexData{}
		// data.Title = "??????"
		// data.Content = "my firtst page"
		// c.HTML(http.StatusOK, "index.html", data)
		return

	} else {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{
			"error": err,
		})
		return
	}
}

type Info struct {
	Name string `yaml:"Name"`
	Age  int    `yaml:"Age"`
}

func ReadYaml() {
	var info Info
	config, err := ioutil.ReadFile("./info.yaml")
	if err != nil {
		panic(err)
	}

	err1 := yaml.Unmarshal(config, &info)
	if err1 != nil {
		panic(err)
	}
	fmt.Println(info)
}

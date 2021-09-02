package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/go-echarts/go-echarts/charts"
	"gopkg.in/yaml.v2"
)

func main() {
	ReadYaml()

	bar := charts.NewBar()
	bar.SetGlobalOptions(charts.TitleOpts{Title: "Bar-示例图"}, charts.ToolboxOpts{Show: true})
	//var nameItems = []string{"1", "2", "3"}
	var xarr = []int{-20, -15, -10, 0, 1}
	bar.AddXAxis(xarr).
		AddYAxis("商家A", []int{1, 2, 3, 4, 5}).
		AddYAxis("商家B", []int{1, 6, 7, 8, 9})
	f, err := os.Create("bar.html")
	if err != nil {
		log.Println(err)
	}
	bar.Render(f) // Render 可接收多个 io.writer 接口
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

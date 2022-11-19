package main

import (
	"fmt"
	"time"
)

type Movie struct {
	Name     string
	Language string
	Length   string
	Score    string
}

func main() {
	m := Movie{Name: "西线无战事", Language: "英语/德语", Length: "148分钟", Score: "8.8分"}
	fmt.Printf("请输入你的命令\n1.电影名字\n2.电影语言\n3.片长\n4.豆瓣评分\n5.观看电影\n6.退出程序")
	var option, x int
	for {
		_, err := fmt.Scanf("%d\n", &option)
		if err != nil {
			fmt.Println("格式错误")
			return
		}
		if option == 1 {
			fmt.Println(m.Name)
		} else if option == 2 {
			fmt.Println(m.Language)
		} else if option == 3 {
			fmt.Println(m.Length)
		} else if option == 4 {
			fmt.Println(m.Score)
		} else if option == 5 {
			go watch()
			go score()
			for {
				_, err := fmt.Scanf("%d\n", &x)
				if err != nil {
					fmt.Println("请重新输入")
					continue
				}
				if x <= 10 && x >= 6 {
					fmt.Println("感谢您的喜欢")
					fmt.Println("请重新输入你的命令")
					break
				} else if x < 6 && x >= 0 {
					fmt.Println("感谢评分，我们未来会做的更好")
					fmt.Println("请重新输入你的命令")
					break
				} else {
					fmt.Println("输入错误，请重新输入")
				}
			}
		} else if option == 6 {
			break
		} else {
			fmt.Println("无此功能，请重新输入")
		}
	}
}
func watch() {
	time.Sleep(1 * time.Second)
	fmt.Println("148分钟后……")
}
func score() {
	time.Sleep(2 * time.Second)
	fmt.Println("请输入你的评分（0~10分）")
}

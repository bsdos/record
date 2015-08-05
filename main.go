package main

import (
	"github.com/Unknwon/goconfig"
	"github.com/Unknwon/macaron"
	// "github.com/macaron-contrib/session"
	// "log"
	"fmt"
	"strconv"
)

func main() {
	m := macaron.Classic()
	m.Get("/", func() string {
		return "Hello world!"
	})
	// times, _ := conf.Int("app", "times")
	m.Get("/:username/:leixing", func(ctx *macaron.Context) string {
		// return "Hello " + ctx.Params(":name")
		//leixing = query,add,remove,stop
		leixing := ctx.Params(":leixing")
		username := ctx.Params(":username")
		conf, _ := goconfig.LoadConfigFile("data.ini")
		times, _ := conf.Int(username, "times")
		if leixing == "use" {
			if times > 0 {
				times -= 1
				stimes := strconv.Itoa(times)
				conf.SetValue(username, "times", stimes)
				goconfig.SaveConfigFile(conf, "data.ini")
				return "yes "
			} else {
				return "no"
			}
		} else if leixing == "query" {
			return strconv.Itoa(times)
		}

		return ""
	})
	m.Get("/:username/:leixing/:num/:password", func(ctx *macaron.Context) string {
		// return "Hello " + ctx.Params(":name")
		//leixing = query,add,remove,stop
		leixing := ctx.Params(":leixing")
		snum := ctx.Params(":num")
		username := ctx.Params(":username")
		password := ctx.Params(":password")
		conf, _ := goconfig.LoadConfigFile("data.ini")
		pwd, _ := conf.GetValue("app", "password")
		if password != pwd {
			return "sorry!"
		}
		fmt.Println(username)
		times, _ := conf.Int(username, "times")
		fmt.Println(times)
		num, _ := strconv.Atoi(snum)
		if leixing == "remove" {
			times = times - num
			fmt.Println(times)
			stimes := strconv.Itoa(times)
			if times > 0 {
				conf.SetValue(username, "times", stimes)
				goconfig.SaveConfigFile(conf, "data.ini")
			}
			return "now is " + stimes
		} else if leixing == "add" {
			// fmt.Println(num)
			times = times + num
			// fmt.Println(times)
			stimes := strconv.Itoa(times)
			conf.SetValue(username, "times", stimes)
			goconfig.SaveConfigFile(conf, "data.ini")
			return "now is " + stimes
		}
		return ""
	})

	m.Run()
}

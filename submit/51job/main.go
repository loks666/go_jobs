package main

import (
	"fmt"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
	"time"
)

func main() {
	// 启动 ChromeDriver
	var opts []selenium.ServiceOption
	service, err := selenium.NewChromeDriverService("path/to/chromedriver", 4444, opts...)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer service.Stop()

	// 连接到 Chrome
	caps := selenium.Capabilities{"browserName": "chrome"}
	chromeCaps := chrome.Capabilities{
		Path: "",
		Args: []string{
			"--start-maximized",
		},
	}
	caps.AddChrome(chromeCaps)
	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", 4444))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer wd.Quit()

	// 导航到页面
	err = wd.Get("https://www.example.com")
	if err != nil {
		fmt.Println(err)
		return
	}

	// 找到元素
	elem, err := wd.FindElement(selenium.ByCSSSelector, "#some-id")
	if err != nil {
		fmt.Println(err)
		return
	}

	// 操作元素
	err = elem.SendKeys("some text")
	if err != nil {
		fmt.Println(err)
		return
	}

	// 等待
	time.Sleep(5 * time.Second)
}

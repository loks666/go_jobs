package main

import (
	"github.com/tebeka/selenium"
)

type Job struct {
	Href      string
	JobName   string
	JobArea   string
	Salary    string
	Tag       string
	Recruiter string
}

var (
	EnableNotifications = true
	Page                = 1
	MaxPage             = 50
	LoginUrl            = "https://www.zhipin.com/web/user/?ka=header-login"
	BaseUrl             = "https://www.zhipin.com/web/geek/job?query=Java&city=101020100&page="
	SayHi               = "您好，我有7年的工作经验，有Java，Python，Golang，大模型的相关项目经验，希望应聘这个岗位，期待可以与您进一步沟通，谢谢！"
)

func main() {
	// Set up Selenium WebDriver.
	caps := selenium.Capabilities{"browserName": "chrome"}
	wd, err := selenium.NewRemote(caps, "")
	if err != nil {
		panic(err)
	}
	defer wd.Quit()

	// TODO: Implement the rest of the functionality.
}

func initDriver() {
	// TODO: Initialize the WebDriver.
}

func resumeSubmission(url string) int {
	// TODO: Implement the resume submission functionality.
	return 0
}

func login() {
	// TODO: Implement the login functionality.
}

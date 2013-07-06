package tests

import "github.com/robfig/revel"

type UserMailerTest struct {
	revel.TestSuite
}

func (t UserMailerTest) Before() { }
func (t UserMailerTest) After() { }

func (t UserMailerTest) TestSendReport() { }

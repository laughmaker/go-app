package log

import (
	"app/src/pkg/conf"
	"app/src/pkg/mail"
	"fmt"
	"log"
	"os"
	"runtime/debug"
	"strings"
	"time"
)

func Try() {
	fmt.Println("崩溃信息--------------------------")

	errs := recover()
	if errs == nil {
		return
	}

	f, err := os.OpenFile("runtime/logs/panic.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("logging.Setup err:%v", err)
	}
	defer f.Close()

	f.WriteString(fmt.Sprintf("[%s]", time.Now().Format("2006-01-02 15:04:05")))
	f.WriteString(fmt.Sprintf("\r\n%v\r\n", errs))
	f.WriteString("[call stack:] \r\n")
	f.WriteString(string(debug.Stack()))
	f.WriteString("\r\n")

	mail.Send(conf.App.LogMail, fmt.Sprintf("%v", errs), string(debug.Stack()), "")
}

func formatStack() string {
	stack := strings.Split(string(debug.Stack()), "	")
	var str string
	for idx, v := range stack {
		v = strings.ReplaceAll(v, "\r\n", " ")

		if idx == 0 {
			str += v + "\n"
			continue
		}

		if idx%2 == 0 {
			str += v + " \n "
		} else {
			str += v + "  ->  "
		}

		if idx == (len(stack) - 1) {
			str += "\r\n"
		}
	}

	return str
}

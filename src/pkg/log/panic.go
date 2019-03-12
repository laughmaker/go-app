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

	mail.Send(conf.App.LogMail, fmt.Sprintf("%v", errs), formatStack(), "")
}

func formatStack() string {
	stack := strings.Split(string(debug.Stack()), "\n")
	str := "<strong style=\"font-size:24px;\">TRACE</strong> <br>"
	for idx, v := range stack {
		if idx == 0 {
			str += v + "<br>"
			continue
		}

		if idx%2 == 0 {
			str += v + " </p> "
		} else {
			str += "<p style=\"padding:1px; margin:1px;\">" + v
			if idx < (len(stack) - 1) {
				str += "  <strong style=\"font-size:16px;\">-></strong> "
			}
		}

		if idx == (len(stack) - 1) {
			str += "\r\n"
		}
	}

	return str
}

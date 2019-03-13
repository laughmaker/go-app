package log

import (
	"app/src/pkg/conf"
	"app/src/pkg/mail"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"runtime/debug"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func Try(c *gin.Context) {
	fmt.Println("崩溃信息--------------------------")

	errs := recover()
	if errs == nil {
		return
	}

	write2File(errs, c)

	body := formatRequest(c) + formatStack()
	mail.Send(conf.App.LogMail, fmt.Sprintf("%v", errs), body, "")
}

func formatRequest(c *gin.Context) string {
	requestBody, _ := ioutil.ReadAll(c.Request.Body)
	header, _ := json.Marshal(c.Request.Header)
	params, _ := json.Marshal(c.Params)

	body := "<strong style=\"font-size:24px;\">BASIC</strong> <br>"
	body += "Method: " + c.Request.Method + "<br>"
	body += "Params: " + string(params) + "<br>"
	body += "RequestURI: " + c.Request.RequestURI + "<br>"
	body += "Host: " + c.Request.Host + "<br>"
	body += "ClientIp: " + c.ClientIP() + "<br>"
	body += "RemoteAddr: " + c.Request.RemoteAddr + "<br>"

	body += "<strong style=\"font-size:24px;\">HEADER</strong> <br>"
	body += string(header) + "<br>"

	body += "<strong style=\"font-size:24px;\">BODY</strong> <br>"
	body += string(requestBody) + "<br>"

	return body
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

func write2File(errs interface{}, c *gin.Context) {
	f, err := os.OpenFile("runtime/logs/panic.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("logging.Setup err:%v", err)
	}
	defer f.Close()

	f.WriteString(fmt.Sprintf("[%s]", time.Now().Format("2006-01-02 15:04:05")))
	f.WriteString(fmt.Sprintf("\r\n%v\r\n", errs))

	requestBody, _ := ioutil.ReadAll(c.Request.Body)
	header, _ := json.Marshal(c.Request.Header)
	params, _ := json.Marshal(c.Params)
	body := "method: " + c.Request.Method + "\r\n"
	body += "params: " + string(params) + "\r\n"
	body += "requestURI: " + c.Request.RequestURI + "\r\n"
	body += "host: " + c.Request.Host + "\r\n"
	body += "clientIp: " + c.ClientIP() + "\r\n"
	body += "remoteAddr: " + c.Request.RemoteAddr + "\r\n"
	body += "header:" + string(header) + "\r\n"
	body += "body:" + string(requestBody) + "\r\n"
	f.WriteString(body)

	f.WriteString("[call stack:] \r\n")
	f.WriteString(string(debug.Stack()))
	f.WriteString("\r\n")
}

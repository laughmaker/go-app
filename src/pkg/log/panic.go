package log

import (
	"fmt"
	"os"
	"runtime/debug"
)

func Try() {
	errs := recover()
	if errs == nil {
		return
	}

	exeName := os.Args[0]
	fmt.Println("exeName: ", exeName)

	pid := os.Getpid()
	fmt.Println("pid: ", pid)

	filename := "/Users/hzd/go/src/go-app/runtime/logs/panic.log"
	fmt.Println("dump to file", filename)

	f, err := os.Create(filename)
	if err != nil {
		return
	}
	defer f.Close()

	f.WriteString(fmt.Sprintf("%v\r\n", errs))
	f.WriteString("========\r\n")
	f.WriteString(string(debug.Stack()))
}

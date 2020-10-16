package main

import (
	"flag"
	"fmt"
	"mychat/connect"
	"mychat/site"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var module string
	flag.StringVar(&module, "module", "", "assign run module")
	flag.Parse()
	fmt.Println(fmt.Sprintf("start run %s module", module))
	switch module {
	case "site":
		site.New().Run()
	case "connect":
		connect.New().Run()
	default:
		fmt.Println("exiting,module param error!")
		return
	}

	fmt.Println(fmt.Sprintf("run %s module done", module))

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGTRAP)
	<-quit
	fmt.Println("Server exiting")
}

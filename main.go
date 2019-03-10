package main

import (
	"app/src/pkg/app"
	"app/src/pkg/conf"
	"app/src/pkg/db"
	"app/src/pkg/route"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func init() {
	app.Config()
}

// @title 接口文档
// @version 1.0
// @description 接口文档
// @termsOfService https://github.com/laughmaker/go-app
// @license.name MIT
// @license.url
func main() {
	run()

	db.Close()
}

func run() {
	gin := route.InitRoute()
	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", conf.Server.HttpPort),
		Handler:        gin,
		ReadTimeout:    conf.Server.ReadTimeout,
		WriteTimeout:   conf.Server.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Println("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	fmt.Println("shutdown server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		fmt.Printf("server shutdown:%s", err)
	}

	select {
	case <-ctx.Done():
		fmt.Println("timeout of 5 second!")
	}

	fmt.Println("server exiting!!")
}

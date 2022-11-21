package main

import (
	"fmt"
	"github.com/nats-io/stan.go"
	"log"
	"os"
	"wildberries/config"
)

func main() {
	conf, err := config.GetConfig() // передать переменную в функции
	if err != nil {
		log.Fatal("error during config downloading: ", err)
	}
	sc, err := stan.Connect(conf.StanClusterId, "clientID2")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err = sc.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()

	mod, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	err = sc.Publish(conf.Subject, mod)
	if err != nil {
		log.Println(err)
	}
}

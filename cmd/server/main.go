// cmd/server/main.go
package main

import (
	"fmt"
)

type People struct {
	Name string
}

func (p People) DoSomething() {
	// (p People), it's the method receiver
	p.Name = "S Darren"
	fmt.Printf("%s is do something\n", p.Name)
}

func (p *People) DoSomething2() {
	p.Name = "K Darren"
	fmt.Printf("%s is do something2\n", p.Name)
}

func main() {
	// configPath := flag.String("config", "etc/config.json", "path to config file")
	// flag.Parse()

	// cfg, err := config.Load(*configPath)
	// if err != nil {
	// 	log.Fatalf("failed to load config: %v", err)
	// }

	// fmt.Printf("Server will start on port %v\n", cfg.ServerPort)
	// TODO: 啟動 server 等
	p := People{Name: "Darren"}
	p.DoSomething()
	fmt.Println("should be darren. P NAME:", p.Name) // should be darren
	p.DoSomething2()
	fmt.Println("should be K darren. P NAME:", p.Name) // should be darren
}

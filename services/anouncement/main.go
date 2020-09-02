package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("Starting Anouncement service")
	var serv service
	err := serv.StartService()

	if err != nil {
		log.Fatalln(err)
	}

}

package main

import (
	"log"
	"os/exec"
	"time"
	"fmt"
)

func main() {

	cmd := exec.Command("cmd", "/c", "start", "powershell", "go", "run", "backup.go").Run()
	//go checkBackup()
	log.Print(cmd)

	counter := 0
	for i := 0; i < 5; i++ {
		counter +=1
		fmt.Println("The counter is", counter)
		time.Sleep(1*time.Second)
	}


	


}

// func sendData() {

// 	for {
// 	sendrate := 200 * time.Millisecond
// 	sender := make(chan int)
// 	message := 1

	
// 	fmt.Println(message)

	
// 	sender <- message
// 	time.Sleep(sendrate)
// 	}
// }
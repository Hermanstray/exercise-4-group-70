package main


import( "fmt"
	"time")

func main() {



	//go checkBackup();
	//go sendData()

	counter := 0
	for i := 0; i < 5; i++ {
		counter +=1
		fmt.Println("The counter is", counter)
		time.Sleep(1*time.Second)
	}

	

}




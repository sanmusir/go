package main
import(
	"fmt"
	"time"
)

func main(){
	done := make(chan bool,1)
	go func(){
		fmt.Println("star")
		time.Sleep(time.Second)
		done <- true
		fmt.Println("send")
	}()
	fmt.Println("waiting")
	<-done
	fmt.Println("end")
}

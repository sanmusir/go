package main
import (
	"confused"
	"fmt"
	"bufio"
    "os"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("输入base64编码: ")
	text, _ := reader.ReadString('\n')
	fmt.Println("解密结果: ")
	fmt.Println(confused.Decode(text))
}

package main
import (
	"fmt"
	"yar"
)

func main(){

	client ,err := yar.NewClientWithTcp("127.0.0.1",6789)

	if err != nil {

		fmt.Printf("err:%s",err)

	}

	var ret string

	err = client.Call("test",&ret)

	fmt.Println(ret)
	fmt.Println(err)

}


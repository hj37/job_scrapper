package main

import (
	"fmt"
	"time"
)

func main() {
	//goroutine 할려는 대상앞에 go만 적어주면 됨
	c := make(chan string) //make함수를 통해서 채널 변수를 추가할 수 있음
	people := [5]string{"nico", "flynn", "dal", "japanguy", "larry"}
	for _, person := range people {
		go isSexy(person, c)
	}
	for i := 0; i < len(people); i++ {
		fmt.Print("waiting for",i)
		fmt.Println(<-c)
	}

	//<-c 는 채널로부터 메세지를 가져오는거
}

func isSexy(person string, c chan string) {
	time.Sleep(time.Second * 10)
	c <- person + "is sexy"
}

func sexyCount(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "is sexy", i)
		time.Sleep(time.Second)
	}
}

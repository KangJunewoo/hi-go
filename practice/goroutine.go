package practice

import (
	"fmt"
	"time"
)

func goroutine() { // 메인함수는 고루틴을 기다리지 않음. 메인이 먼저 끝나면 남아있는 고루틴도 소멸.
	c := make(chan string)
	people := [2]string{"nico", "flynn"}
	for _, person := range people {
		go isSexy(person, c)
	}
	// result := <-c // 이 코드는 고루틴에서 메세지를 받을 때까지 기다림. 메시지큐가 이런 개념이려나?
	fmt.Println("waiting for msgs..")
	fmt.Println(<-c) // 이는 blocking operation이라고 할 수 있고,
	fmt.Println(<-c) // 채널수보다 많이 소환하면 데드락 에러 뜸.
	// for i := 0; i < len(people); i++ {
	// 	fmt.Println(<-c) // 이런 식으로 하면 사람이 몇 명이든 수월하게 받아내겠지.
	// }

}

func isSexy(person string, c chan string) {
	time.Sleep(time.Second * 5)
	c <- person + " is sexy"
}

package main

import (
	"fmt"
	"strings"
)

type person struct {
	name    string
	age     int
	favFood []string
}

func multiply(a, b int) int {
	return a * b
}

// 고는 함수 리턴값이 두개일 수 있다.
func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func lenAndUpper2(name string) (length int, uppercase string) {
	defer fmt.Println("I'm done")
	length = len(name)
	uppercase = strings.ToUpper(name)
	return

}

func canIDrink(age int) bool {
	switch koreanAge := age + 2; koreanAge {
	case 18:
		return false
	case 50:
		return true
	}
	// switch {
	// case age < 18:
	// 	return false
	// case age == 18:
	// 	return true
	// case age > 18:
	// 	return true
	// }
	return false
	// if koreanAge := age + 2; koreanAge < 18 {
	// 	return false
	// }
	// return true
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func superAdd(numbers ...int) int {
	total := 0

	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}

	for _, number := range numbers {
		total += number
	}
	// for index, number := range numbers {
	// 	fmt.Println(index, number)
	// }
	return total
}

func main() {
	const name string = "junu"
	var name2 string = "junu"
	// name2 := "junu"
	name2 = "merong"
	fmt.Println(name2)

	totalLength, upperName := lenAndUpper2(name)
	// totalLength, _ := lenAndUpper(name) 이런식으로 ignore 가능.
	fmt.Println(totalLength, upperName)
	repeatMe("nico", "lynn", "dal", "marl", "flynn")

	result := superAdd(1, 2, 3, 4, 5, 6)
	fmt.Println(result)
	fmt.Println(canIDrink(16))

	a := 2
	b := &a
	*b = 20
	// b := a
	// a = 10
	fmt.Println(a, b, &a, *b) // 오 포인터

	// names := [5]string{"nico", "lynn", "dal"}
	// names[3] = "alala"
	// names[4] = "alala"

	names := []string{"nico", "lynn", "dala"}
	names = append(names, "flynn")
	fmt.Println(names)

	nico := map[string]string{"name": "nico", "age": "12"}
	fmt.Println(nico)
	for key, value := range nico { // 무시하고 싶으면 언더바~
		fmt.Println(key, value)
	}
	favFood := []string{"kimchi", "ramen"}
	// lynn := person{"lynn", 18, favFood} // 좋은 방법은 아님.
	lynn := person{name: "nico", age: 18, favFood: favFood}
	fmt.Println(lynn)
}

package ifExercise

import "fmt"

//целое число, лежащее в диапазоне 1–999. Вывести его строкуописание вида «четное двузначное число», «нечетное трехзначное число»
//и т. д

func IfEx30(a int) {
	if a > 99 && a < 1000 {
		fmt.Println("Our number has three number")
	} else if a < 100 {
		fmt.Println("Our number has two number")
	} else if a%2 == 0 {
		fmt.Println("Our number is plus")
	} else if a == 0 {
		fmt.Println("Our number is zero")
	} else {
		fmt.Println("Our number is minus")
	}
}

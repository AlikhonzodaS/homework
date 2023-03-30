package ifExercise

import "fmt"

//Дано целое число. Вывести его строку-описание вида «отрицательное
//четное число», «нулевое число», «положительное нечетное число» и т. д

func IfEx(a int) {
	if a > 0 {
		fmt.Println("Our number is plus")
	} else if a == 0 {
		fmt.Println("0")
	} else {
		fmt.Println("Our number is minus")
	}
}

//var a int
//fmt.Scan(&a)
//ifExercise.IfEx(a)

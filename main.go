package main

import (
	"fmt"
	// "sort"
)

func sum(nums []int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

func filter(slice []int) []int {
	for i, _ := range slice {
		if slice[i]%2 == 0 {
			fmt.Println(slice[i])
		}
	}
	return slice
}

func remove(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}

func main() {

	var num int = 123
	fmt.Println(num)
	// Задача 3.4:
	// Напиши функцию, которая фильтрует срез и возвращает только чётные числа

	slice := []int{1, 4, 8, 11, 15, 18, 22, 9}
	remove(slice, 0)
	fmt.Println(slice)
	// a := slice
	// filter(a)
	// Задача 3.3:
	// Напиши функцию, которая принимает срез чисел и возвращает сумму всех элементов

	// slice := []int{1, 2, 3, 4}
	// fmt.Println(sum(slice))

	// Задача 3.4:
	// Напиши функцию, которая фильтрует срез и возвращает только чётные числа

	// Заданиие 1
	// fmt.Println("Привет, Telegram-бот !")

	// Заданиие 2
	// var name string = "Bob"
	// age := 25
	// fmt.Printf("Имя : %s, Возраст: %d\n", name, age)

	// Заданиие 2.2
	// var name string
	// var age uint8
	// fmt.Print("What is your name ?")
	// fmt.Scan(&name)
	// fmt.Print("How old are you &")
	// fmt.Scan(&age)
	// fmt.Printf("Hellow " + name + " ! Your age is " + fmt.Sprint(age) + " .")

	// Заданиие 3
	// const pi float64 = 3.1415
	// var R int = 12
	// res := (float64(R*R) * pi) / 4
	// fmt.Printf("S = " + fmt.Sprint(res))

	// Заданиие 4
	// var age uint8
	// fmt.Printf("How is you age ?")
	// fmt.Scan(&age)
	// if age >= 18 {
	// 	fmt.Println("Your age is " + fmt.Sprint(age) + " !")
	// 	fmt.Println("Ты совершеннолетний ! ")
	// } else if age < 18 {
	// 	fmt.Println("Your age is " + fmt.Sprint(age) + " !")
	// 	fmt.Println("Ты не достиг совершеннолетия !")
	// } else {
	// 	fmt.Println("Err!")
	// }
	// Заданиие 4.2
	// var num int32
	// fmt.Printf("select some number")
	// fmt.Scan(&num)
	// if num > 0 {
	// 	fmt.Printf("Число положительное и равно : " + fmt.Sprint(num))
	// } else if num < 0 {
	// 	fmt.Printf("Число отрицательное и равно : " + fmt.Sprint(num))
	// } else {
	// 	fmt.Printf("Число равно " + fmt.Sprint(num))
	// }
	//Задание 5
	// for i := 0; i < 5; i++ {
	// 	fmt.Println(i)
	// }
	//Задание 5.2
	// var N int
	// fmt.Println("Введите диапозон N")
	// fmt.Scan(&N)
	// for i := 0; i <= N; i++ {
	// 	fmt.Println(i)
	// }
	//Задание 5.3
	// var i, N int
	// fmt.Print("Введите начальное число i :")
	// fmt.Scan(&i)
	// fmt.Print("Введите начальное число N :")
	// fmt.Scan(&N)

	// sum := 0

	// for num := i; num <= N; num++ {
	// 	sum += num
	// }
	// fmt.Println("Сумма чисел от ", i, "до", N, "равна", sum)
	// Задание 6
	// nums := []int{1, 2, 3}
	// nums = append(nums, 4, 5)
	// fmt.Println(nums)
	// Задание 6.2
	// nums := []int{1, 2, 3, 4, 5}
	// nums = append(nums, 6)
	// fmt.Println(nums)
	// fmt.Println(len(nums))
	// fmt.Println(cap(nums))

	////////День 1.5
	// 1
	// var num int
	// fmt.Print("Введите число ")
	// fmt.Scan(&num)
	// if num > 0 {
	// 	fmt.Println("положительное")
	// } else if num < 0 {
	// 	fmt.Println("отрицательное")
	// } else {
	// 	fmt.Println("ноль")
	// }
	// 1.2
	// var num int
	// fmt.Printf("Введите число ")
	// fmt.Scan(&num)
	// switch {
	// case num%2 == 0:
	// 	fmt.Println("Чётное")
	// default:
	// 	fmt.Println("Нечётное")
	// }
	// 1.3
	// var num uint
	// fmt.Println("Введите число от 1 до 7 ")
	// fmt.Scan(&num)
	// switch {
	// case num == 1:
	// 	fmt.Println("Понедельник")
	// case num == 2:
	// 	fmt.Println("Вторник")
	// case num == 3:
	// 	fmt.Println("Среда")
	// case num == 4:
	// 	fmt.Println("Четверг")
	// case num == 5:
	// 	fmt.Println("Пятница")
	// case num == 6:
	// 	fmt.Println("Суббота")
	// case num == 7:
	// 	fmt.Println("Воскресенье")
	// default:
	// 	fmt.Println("Введите корректное число")
	// }
	// if num == 1 {
	// 	fmt.Println("Понедельник")
	// } else if num == 2 {
	// 	fmt.Println("Вторник")
	// } else if num == 3 {
	// 	fmt.Println("Среда")
	// } else if num == 4 {
	// 	fmt.Println("Четверг")
	// } else if num == 5 {
	// 	fmt.Println("Пятница")
	// } else if num == 6 {
	// 	fmt.Println("Суббота")
	// } else if num == 7 {
	// 	fmt.Println("Воскресенье")
	// } else {
	// 	fmt.Println("Введите корректное число")
	// }
	// 2.1
	// var N int
	// fmt.Println("Select the N")
	// fmt.Scan(&N)
	// for i := 1; i <= N; i++ {
	// 	if i%3 == 0 {
	// 		fmt.Println(i)
	// 	}
	// }

	//2.2
	// var N int
	// fmt.Print("Need to find factorial a number : select the N ")
	// fmt.Scan(&N)
	// fac := 1
	// for i := fac; i <= N; i++ {
	// 	fac *= i
	// 	fmt.Println(fac)
	// }

	//2.3
	// words := []string{"Hello", "word", "Go"}
	// for i, word := range words {
	// 	fmt.Printf("Index %d : %s\n", i, word)
	// }

	//2.3.2
	// nums := []int{1, 2, 3, 4}
	// for i, num := range nums {
	// 	fmt.Printf("Index %d : %v\n", i, num)
	// }
	//2.4
	// var a, b int = 1, 9
	// res := a * b
	// for i := a; i <= b; i++ {
	// 	res = i * b
	// 	for j := 1; j <= b; j++ {
	// 		res = j * i
	// 		fmt.Printf("%d * %d = %d\n", i, j, res)
	// 	}
	// 	fmt.Printf("%d * %d = %d\n", i, b, res)
	// }
	// Массивы и Срезы
	//3.1
	//  Создай массив из 5 целых чисел. Используй индексы для доступа к элементам

	// var numbers [5]int = [5]int{10, 20, 30, 40, 50}
	// fmt.Println("Index 1 :", numbers[0])
	// fmt.Println("Index 2 :", numbers[1])
	// fmt.Println("Index 3 :", numbers[2])
	// fmt.Println("Index 4 :", numbers[3])
	// fmt.Println("Index 5 :", numbers[4])

	// numbers[1] = 25
	// fmt.Println("Index 2 have been changed :", numbers[1])

	// fmt.Println("All array :", numbers)
	// fmt.Println("All element's on array :")
	// for i := 0; i < len(numbers); i++ {
	// 	fmt.Printf("numbers[%d] = %d\n", i, numbers[i])
	// }
	// Задача 3.2:
	// Создай срез, добавь в него 10 чисел, выведи его длину и вместимость
	// nums := []int{1, 2, 3, 50, 60}
	// nums = append(nums, 4, 5, 6, 7, 8, 9, 10, 20, 30, 40)
	// fmt.Println(nums)
	// fmt.Println(len(nums))
	// fmt.Println(cap(nums))

	// slice := []int{1}
	// sort.Ints(slice)
	// slice = append(slice, 2, 3, 4, 5, 6, 7, 8, 9, 10, 20)
	// fmt.Println(slice)
	// fmt.Println(len(slice))
	// fmt.Println(cap(slice))

}

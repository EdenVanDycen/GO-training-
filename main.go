package main

import "fmt"

func main() {
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
	nums := []int{1, 2, 3}
	nums = append(nums, 4, 5)
	fmt.Println(nums)
}

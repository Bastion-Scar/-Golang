package main

import (
	"fmt"
	"regexp"
)

func phone(input string) {
	re, err := regexp.Compile("^\\+\\d{11}$")
	if err != nil {
		fmt.Println(err)
	}
	result := re.FindAllString(input, -1)
	if len(result) == 0 {
		fmt.Println("Ошибка, некорректно введен телефон")
	} else {
		fmt.Println(result)
	}
}

func email(input string) {
	re, err := regexp.Compile("^[\\w._%+-]+@[\\w._-]+[a-zA-Z]{2,5}")
	if err != nil {
		fmt.Println(err)
	}
	result := re.FindAllString(input, -1)
	if len(result) == 0 {
		fmt.Println("Ошибка, некорректный email")
	} else {
		fmt.Println(result)
	}

}

func password(input string) {
	re, err := regexp.Compile("^[a-zA-Z0-9]{8}$") //Хотел сделать прикольней, но голова загружена из за завтрашнего зачета...
	if err != nil {
		fmt.Println(err)
	}
	result := re.FindAllString(input, -1)
	if len(result) == 0 {
		fmt.Println("Ошибка, пароль слабый")
	} else {
		fmt.Println(result)
	}
}

func main() {
	phone("+79041782718") //Тут пройдет
	phone("+7904178271")  //Ошибка

	email("chema2007ivan@mail.ru") //Пройдет
	email("chema2007ivanmail.ru")  //Ошибка

	password("01915W85") //Пройдет
	password("0195")     //Ошибка
}

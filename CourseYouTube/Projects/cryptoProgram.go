package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
)

type User struct {
	id     int
	name   string
	age    int
	wallet string
}

func CreateUser(name string, age int) User {
	id := rand.IntN(100)
	switch {
	case age < 18:
		fmt.Printf("%v,Account is not created, your age is lower 18\n", name)
		return User{}
	default:
		fmt.Printf("%s, account with id %d created\n", name, id)
		return User{id: id, name: name, age: age, wallet: ""}
	}

}
func CreateWallet(user *User) {
	user.wallet = fmt.Sprintf("WlT-%d-%s", user.id, user.name)
	fmt.Println("✅ Кошелек успешно создан!")
	fmt.Printf("Адрес вашего кошелька: %s\n", user.wallet)
}

func main() {
	person := CreateUser("Gopher", 18)
	if person.wallet == "" {
		fmt.Println("Хотите создать кошелек для вашего аккаунта?")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		if input == "\n" {
			CreateWallet(&person)
			fmt.Printf("Ваши данные: \n %v", person)
		} else {
			fmt.Printf("Ваши данные: \n %v", person)
		}
	}
}

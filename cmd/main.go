package main

import (
	"fmt"
	"github.com/shuhrat-shokirov/wallet/pkg/wallet"
)

func main(){
	svc := &wallet.Service{}
	account, err := svc.RegisterAccount("+992000000000")
	if err != nil{
		fmt.Println(err)
		return
	}

	err = svc.Deposit(account.ID, 10)
	if err != nil{
		switch err{
		case wallet.ErrAmountMustBePositive:
			fmt.Println("Сумма должна быть полажительной")
		case wallet.ErrAccountNotFound:
			fmt.Println("Аккаунт пользователя не найдено")
		}
		return
	}

	fmt.Println(account.Balance)
}
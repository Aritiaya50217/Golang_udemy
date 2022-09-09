package main

import "fmt"

type user interface {
	GetAccount()
}

type cart interface {
	user
	SaveItem()
}

type wallet interface {
	user
	GetBalance()
}

type order interface {
	cart
	wallet
	Checkout()
}

type userImpl struct{}

func (u userImpl) GetAccount() {
	fmt.Println("User-GetAccount")
}

type cartImpl struct {
	user
}

func (c cartImpl) SaveItem() {
	fmt.Println("Cart-SaveItem")
}

type walletImpl struct {
	user
}

func (w walletImpl) GetBalance() {
	fmt.Println("Wallet-GetBalance")
}

type orderImpl struct {
	cart
	wallet
}

func main() {
	user := userImpl{}
	user.GetAccount()
	order := orderImpl{cartImpl{user}, walletImpl{user}}
	order.GetBalance()
	order.SaveItem()

	// duplicate method
	fmt.Println("==== Duplicate Method ====")
	fmt.Print("order.cart.GetAccount : ")
	order.cart.GetAccount()
	fmt.Print("order.wallet.GetAccount : ")
	order.wallet.GetAccount()
}

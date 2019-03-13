package main

import (
	"fmt"
)

type Payer interface {
	Pay(int) error
}

type Replenisher interface {
	Replenish(int) error
}

type Wallet struct {
	Cash int
}

func (w *Wallet) Replenish(amount int) error {
	if amount < 0 {
		return fmt.Errorf("amount is negative")
	}
	w.Cash = w.Cash + amount
	return nil
}
func (w *Wallet) Pay(amount int) error {
	if w.Cash < amount {
		return fmt.Errorf("Не хватает денег в кошельке")
	}
	w.Cash -= amount
	return nil
}

func Buy(p Payer) {
	err := p.Pay(10)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Спасибо за покупку через %T\n%T\n", p, &p)
}

func Fill(r Replenisher) {
	err := r.Replenish(110)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Спасибо за пополнение через %T\n%T \n", r, &r)
}

func main() {
	myWallet := &Wallet{Cash: 100}
	Buy(myWallet)
	Fill(myWallet)
}

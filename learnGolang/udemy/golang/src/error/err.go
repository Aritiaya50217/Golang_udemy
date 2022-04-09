package main

import (
	"errors"
	"fmt"
	"log"
	"time"
)

const dbReady = false
const balance = 200
const numberToSuccess = 10

func connectDb(nTry int) error {
	if nTry == numberToSuccess {
		return nil
	}
	return errors.New("busy")
}
func waitForDatabase() error {
	timeout := 3 * time.Second
	deadline := time.Now().Add(timeout)

	for tries := 0; time.Now().Before(deadline); tries++ {
		err := connectDb(tries)
		if err == nil {
			return nil
		}
		log.Printf("database is not responding (%v); retrying", err)
		time.Sleep(time.Second << uint(tries))
	}
	return fmt.Errorf("waitForDatabase : timeout %v", timeout)
}

func getBalance() (int64, error) {
	if !dbReady {
		// retry
		//log.Println("retrying...")
		if err := waitForDatabase(); err != nil {
			return 0, fmt.Errorf("getBalance: %v", err)
		}

	}
	return balance, nil
}

func errorPropagation(amount int64) (int64, error) {
	balance, err := getBalance()
	if err != nil {
		return 0, fmt.Errorf("withdraw: %v", err)
	}

	if amount > balance {

		return 0, errors.New("errorPropagation : insufficient found")
	}
	return amount, nil
}

func main() {
	log.SetFlags(0)
	amount2, err := errorPropagation(200)
	if err != nil {
		log.Fatal("main :", err)
	}
	fmt.Println("Please collect your money :", amount2)

}

/*
You are given an m x n integer grid accounts where accounts[i][j] is the amount of money the i​​​​​​​​​​​th​​​​ customer has in the j​​​​​​​​​​​th​​​​ bank.
Return the wealth that the richest customer has.
A customer's wealth is the amount of money they have in all their bank accounts.
The richest customer is the customer that has the maximum wealth.
*/

package RichestCustomer

import "sync"

func maximumWealth(accounts [][]int) int {
	maxSum := 0
	ch := make(chan int, len(accounts))
	wg := sync.WaitGroup{}

	for _, account := range accounts {
		account := account

		wg.Add(1)
		go func() {
			defer wg.Done()

			acctSum := 0
			for _, v := range account {
				acctSum += v
			}
			ch <- acctSum
		}()
	}

	wg.Wait()
	close(ch)
	for acctSum := range ch {
		if acctSum > maxSum {
			maxSum = acctSum
		}
	}

	return maxSum
}

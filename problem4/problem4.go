package problem4

import (
	"time"

	"github.com/ilopshire/eulerGo/util"
)

func Solve() {
	start := time.Now()

	var solution int
	pq := new(productQueue)
	pq.init(100, 999)
	for {
		value := pq.pop()
		if isPalindrome(value) {
			solution = value
			break
		}
	}

	util.PrintSolution(solution, time.Since(start))
}

//productProvider
type productProvider struct {
	multiplier  int
	staticValue int
	nextProduct int
}

//calculate populates the nextProduct field
func (p *productProvider) calculate() {
	p.nextProduct = p.multiplier * p.staticValue
}

//pop gets the next product then recalculates nextProduct
func (p *productProvider) pop() int {
	product := p.nextProduct
	p.multiplier--
	p.calculate()

	return product
}

//productQueue stores a sorted slice of productProviders
type productQueue struct {
	productProviders []productProvider
	min, max         int
}

//init sets up the productQueue with a min and max value
func (q *productQueue) init(min, max int) {
	q.productProviders = make([]productProvider, max-min+1)
	q.min = min
	q.max = max

	//this will be sorted
	for index := range q.productProviders {
		q.productProviders[index].multiplier = max
		q.productProviders[index].staticValue = min + index
		q.productProviders[index].calculate()
	}
}

//pop gets the next largest integer
func (q *productQueue) pop() int {
	value := q.productProviders[q.max-q.min].pop()

	//keep the list sorted
	offset := 0
	for q.productProviders[q.max-q.min-offset].nextProduct < q.productProviders[q.max-q.min-offset-1].nextProduct {
		//swap down the list
		temp := q.productProviders[q.max-q.min-offset]
		q.productProviders[q.max-q.min-offset] = q.productProviders[q.max-q.min-offset-1]
		q.productProviders[q.max-q.min-offset-1] = temp
		offset++
	}

	return value
}

func isPalindrome(num int) bool {
	if num > 0 && num < 10 {
		return true
	}

	rev := 0
	newNum := num

	for newNum > 0 {
		rev = rev*10 + newNum%10
		newNum /= 10
	}

	return num == rev
}

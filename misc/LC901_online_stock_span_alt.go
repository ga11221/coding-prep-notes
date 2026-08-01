package main

// arrays/strings:
// sliding window
// pointer arithmetic
// bit array/buckets
// sort first
// dp
// working backwards
// maps/sets (for strings)
// queues (for strings)

// n^2
// keep position/order of price in list
// as prices come in, store in a second, sorted list along with indexes
// use binsearch to find lrg=largest(small) in sorted list
// if lrg=-1, return 1
// nxtLrg = idx/pos(sortedList[lrg+1])
// use binsearch to find nxtLrg=idx of largest(small)/next largest price, ret i - nextLrg
type PricePair struct {
	price, positionInStream int
}
type StockSpanner struct {
	sortedPrices     []PricePair
	positionInStream int
}

func Constructor() StockSpanner {
	return StockSpanner{
		sortedPrices: []PricePair{},
	}
}

func (this *StockSpanner) Next(price int) int {
	// insert in sorted order into this.sortedPrices:
	pricePair := PricePair{
		price:            price,
		positionInStream: this.positionInStream,
	}
	this.insert(pricePair)
	// incr positionInStream
	this.positionInStream++
	// binsearch for nextLrg pricePair
	var nextLrg PricePair = this.findNextLrg(price)
	return pricePair.positionInStream - nextLrg.positionInStream
}

func (this *StockSpanner) insert(pricePair PricePair) {

}

func (this *StockSpanner) findNextLrg(pricePair PricePair) {

}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */

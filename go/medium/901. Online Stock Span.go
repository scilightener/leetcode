package medium

type DataItem struct {
	price int
	ans   int
}

type StockSpanner []DataItem

func Constructor() StockSpanner {
	return make([]DataItem, 0)
}

func (this *StockSpanner) Next(price int) int {
	ans := 1
	for len(*this) > 0 && (*this)[len(*this)-1].price <= price {
		ans += (*this)[len(*this)-1].ans
		*this = (*this)[:len(*this)-1]
	}

	*this = append(*this, DataItem{price, ans})
	return ans
}

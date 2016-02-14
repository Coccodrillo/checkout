package checkout

import "errors"

type ItemForSale interface {
	GetCode() string
	GetPrice() float64
}

type Item struct {
	code  string
	price float64
}

func (i Item) GetCode() string {
	return i.code
}

func (i Item) GetPrice() float64 {
	return i.price
}

func getItemByCode(code string) (item ItemForSale, err error) {
	// some sort of db call
	items := []ItemForSale{
		Item{"VOUCHER", 5.00},
		Item{"TSHIRT", 20.00},
		Item{"MUG", 7.50},
	}
	for k := range items {
		if items[k].GetCode() == code {
			return items[k], nil
		}
	}
	return item, errors.New("Not found")
}

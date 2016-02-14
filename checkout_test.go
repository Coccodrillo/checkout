package checkout

import (
	"sync"
	"testing"
)

func TestExamples(t *testing.T) {
	examples := []struct {
		items  []string
		result string
	}{
		{[]string{"VOUCHER", "TSHIRT", "MUG"}, "32.50€"},
		// contrary to the examples given, the correct value for this one is 30.00€, since 25.00€ would imply buy one get one, not buy two get one
		{[]string{"VOUCHER", "TSHIRT", "VOUCHER"}, "30.00€"},
		{[]string{"TSHIRT", "TSHIRT", "TSHIRT", "VOUCHER", "TSHIRT"}, "81.00€"},
		{[]string{"VOUCHER", "TSHIRT", "VOUCHER", "VOUCHER", "MUG", "TSHIRT", "TSHIRT"}, "74.50€"},
	}
	pricingRules := []PricingRule{
		BuyTwoGetOneFree{[]string{"VOUCHER"}},
		BulkOneOff{[]string{"TSHIRT"}},
	}
	for _, example := range examples {
		for i := 0; i < 10000; i++ {
			if result := testAddConcurentlyAndprintTotalForCart(example.items, pricingRules); result != example.result {
				t.Errorf("total not correct, received %s, expected %s", result, example.result)
			}
		}
	}
}

func testAddConcurentlyAndprintTotalForCart(items []string, pricingRules []PricingRule) string {
	var wg sync.WaitGroup
	co := NewCheckout(pricingRules)
	queue := make(chan bool, len(items))
	wg.Add(len(items))
	for _, item := range items {
		go func(item string) {
			defer wg.Done()
			co.Scan(item)
			queue <- true
		}(item)
	}
	wg.Wait()
	price := co.GetTotal()
	return price
}

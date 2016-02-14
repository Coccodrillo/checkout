package checkout

import (
	"fmt"
	"sync"
)

func NewCheckout(rules []PricingRule) *CheckOut {
	return &CheckOut{
		pricingRules: rules,
	}
}

type CheckOut struct {
	pricingRules []PricingRule
	items        []ItemForSale
	total        float64
	sync.RWMutex
}

func (c *CheckOut) Scan(name string) {
	item, err := getItemByCode(name)
	if err == nil {
		c.Lock()
		defer c.Unlock()
		c.items = append(c.items, item)
	}
}

func (c *CheckOut) GetTotal() string {
	c.total = 0 // zero it to prevent mistakes when calling it more times
	total := make(chan float64)
	for i := 0; i < len(c.items); i++ {
		go func(total chan float64, item ItemForSale) {
			total <- item.GetPrice()
		}(total, c.items[i])
	}
	for i := 0; i < len(c.pricingRules); i++ {
		go func(total chan float64, rule PricingRule) {
			total <- rule.Calculate(c.items)
		}(total, c.pricingRules[i])
	}
	for i := 0; i < len(c.items)+len(c.pricingRules); i++ {
		c.total += <-total
	}
	return fmt.Sprintf("%.2f€", c.total)
}

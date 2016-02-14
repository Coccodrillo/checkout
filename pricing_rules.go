package checkout

type PricingRule interface {
	GetApplicableCodes() []string
	Calculate([]ItemForSale) float64
}

type BuyTwoGetOneFree struct {
	AppliesTo []string
}

func (b BuyTwoGetOneFree) GetApplicableCodes() []string {
	return b.AppliesTo
}

func (b BuyTwoGetOneFree) Calculate(a []ItemForSale) (sumChange float64) {
	forCodes := b.GetApplicableCodes()
	for _, code := range forCodes {
		var counter int
		item, err := getItemByCode(code)
		if err == nil {
			for j := range a {
				if code == a[j].GetCode() {
					counter++
				}
			}
			sumChange -= float64(counter/3) * item.GetPrice()
		}
	}
	return sumChange
}

type BulkOneOff struct {
	AppliesTo []string
}

func (b BulkOneOff) GetApplicableCodes() []string {
	return b.AppliesTo
}

func (b BulkOneOff) Calculate(a []ItemForSale) (sumChange float64) {
	forCodes := b.GetApplicableCodes()
	for _, code := range forCodes {
		var counter int
		for j := range a {
			if code == a[j].GetCode() {
				counter++
			}
		}
		if counter > 2 {
			sumChange -= float64(counter)
		}
	}
	return sumChange
}

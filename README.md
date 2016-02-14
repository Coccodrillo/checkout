## This is a thread safe checkout example

This tool assumes you are working in a standard Go workspace, as described in
http://golang.org/doc/code.html. The project was build and tested using Go version 1.5.2, but should work in any version 1.1+.

There are no outside dependencies, so the steps to set up the project are the following:

## Install

```console
$ go get github.com/coccodrillo/checkout
```

## Usage

```golang
package main

import (
	"fmt"
	"github.com/coccodrillo/checkout"
)

func main() {
	examples := []struct {
		items  []string
		result string
	}{
		{[]string{"VOUCHER", "TSHIRT", "MUG"}, "32.50€"},
		{[]string{"VOUCHER", "TSHIRT", "VOUCHER"}, "30.00€"},
		{[]string{"TSHIRT", "TSHIRT", "TSHIRT", "VOUCHER", "TSHIRT"}, "81.00€"},
		{[]string{"VOUCHER", "TSHIRT", "VOUCHER", "VOUCHER", "MUG", "TSHIRT", "TSHIRT"}, "74.50€"},
	}
	pricingRules := []checkout.PricingRule{
		checkout.BuyTwoGetOneFree{[]string{"VOUCHER"}},
		checkout.BulkOneOff{[]string{"TSHIRT"}},
	}
	for _, example := range examples {
		co := checkout.NewCheckout(pricingRules)
		for _, itemCode := range example.items {
			co.Scan(itemCode)
		}
		fmt.Println(co.GetTotal() == example.result)
	}
}

```

The project can be tested using `go test` and built using `go build`. Tests were written using only Go testing package so there are no requirements there. The tests are repeated 10k times each as a simple proof of concurrent access.
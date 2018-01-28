package cart

import "testing"

func TestAddAndGetProducts(t *testing.T) {
	c := New()
	c.Add("apple1")
	c.Add("orange")

	products := c.GetAll()
	if len(products) != 2 {
		t.Fatalf("wrong producrs number(%d)", len(products))
	}
	if products[0] != "apple" && products[1] != "apple" {
		t.Error("There is not apple")
		t.Log("cart : ", products)
	}
}

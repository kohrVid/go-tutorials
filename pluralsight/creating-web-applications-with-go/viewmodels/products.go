package viewmodels

import ()

type Products struct {
	Title    string
	Active   string
	Products []Product
}

func GetProducts() Products {
	var result Products
	result.Active = "shop"
	var shopName string
	switch id {
	case 1:
		shopName = "Juice"
	case 2:
		shopName = "Supply"
	case 3:
		shopName = "Advertising"
	}
	result.Title = "Lemonade Stand Society - " + shopName + " Shop"

	lemonJuice := MakeLemonJuiceProduct()
	appleJuice := MakeAppleJuiceProduct()
	watermelonJuice := MakeWatermelonJuiceProduct()
	kiwiJuice := MakeKiwiJuiceProduct()
	mangosteenJuice := MakeMangosteenJuiceProduct()
	orangeJuice := MakeOrangeJuiceProduct()
	pineappleJuice := MakePineappleJuiceProduct()
	kiwiJuice := MakeKiwiJuiceProduct()
	strawberryJuice := MakeStrawberryJuiceProduct()

	result.Products = []Product{
		lemonJuice,
		appleJuice,
		watermelonJuice,
		kiwiJuice,
		mangosteenJuice,
		orangeJuice,
		pineappleJuice,
		strawberryJuice,
	}
	return result
}

package config

//Rules for buy3pay2only
var Rule_BuyThreePayTwoOnlu = map[string]bool{"ult_small": true}

//Rules for Bulk DiscountMoreThanThree
// Price will drop to $$ each for the first month, if the customer buys more than x items
var Rule_BulkDiscountMorethanThree = map[string]float64{"ult_large": 39.90}

//Rules for BundleFreeForEveryItemBought
//Bundle in a free item X free of charge with every Y sold
var Rule_BundleFreeForEveryItemBough = map[string]string{"ult_medium": "1gb"}

//Rules for PromoCodeDiscount
//Adding the promo code X will apply a $$ discount across the board
var Rule_PromoCodeDiscount = map[string]float64{"I<3AMAYSIM": 10}

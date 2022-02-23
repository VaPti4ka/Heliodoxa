package utils

type Category interface {
	String() string
}

type SpendCategory string

const (
	OneTime   = SpendCategory("One-Time")  // OneTime spend is base category for something like taxi, lunch, etc
	Repeating = SpendCategory("Repeating") // Repeating spent is base category for something like real estate, subscription, etc
)

// String return string value of Category
func (s SpendCategory) String() string {
	return string(s)
}

type DepositCategory string

const (
	Salary = DepositCategory("Salary") // Salary is just salary
	Credit = DepositCategory("Credit") // Credit may be bank-credit, or friend-loan or something else
)

// String return string value of Category
func (d DepositCategory) String() string {
	return string(d)
}

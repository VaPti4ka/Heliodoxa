package expenses

import (
	"time"

	"github.com/VaPti4ka/Heliodoxa/internal/pkg/utils"
)

type BaseSpend interface {
	CreateNewSpend() (int64, error)
	UpdateSpend() (Spend, error)
	ReadSpend(id int64) (Spend, error)
	DeleteSpend() error
}

// Spend - base type of money Spend
type Spend struct {
	ID       int64               // ID - uniq Spend id
	Title    string              // Title - Spend title
	Amount   string              // Amount - Spend amount
	Date     time.Time           // Date - date and time of Spend
	Category utils.SpendCategory // Category - custom category of Spend. Use utils.SpendCategory
}

// CreateNewSpend create new entry in DB
// TODO implement me
func (s Spend) CreateNewSpend() (int64, error) {
	return 0, nil
}

// ReadSpend get information about spend entry from DB and return Spend struct
// TODO implement me
func (s Spend) ReadSpend(id int64) (Spend, error) {
	var resultSpend Spend

	return resultSpend, nil
}

// UpdateSpend find spend entry in BD by ID and reWrite info from Deposit to DB
// TODO implement me
func (s Spend) UpdateSpend() (Spend, error) {
	return Spend{}, nil
}

// DeleteSpend remove spend entry from DB
// TODO implement me
func (s Spend) DeleteSpend() error {
	return nil
}

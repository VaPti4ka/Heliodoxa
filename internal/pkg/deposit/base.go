// Package deposit is
package deposit

import (
	"time"

	"github.com/VaPti4ka/Heliodoxa/internal/pkg/utils"
)

type BaseDeposit interface {
	CreateNewDeposit() (int64, error)
	UpdateDeposit() (Deposit, error)
	ReadDeposit(id int64) (Deposit, error)
	DeleteDeposit() error
}

// Deposit - base type of money deposit
type Deposit struct {
	ID       int64               // ID - uniq Deposit id
	Title    string              // Title - Deposit title
	Amount   string              // Amount - Deposit amount
	Date     time.Time           // Date - date and time of Deposit
	Category utils.SpendCategory // Category - custom category of Deposit. Use utils.SpendCategory
}

// CreateNewDeposit create new entry in DB.
// Return ID of created entry
// TODO implement me
func (d Deposit) CreateNewDeposit() (int64, error) {
	return 0, nil
}

// ReadDeposit get information about deposit entry from DB
// Return Deposit struct
// TODO implement me
func (d Deposit) ReadDeposit(id int64) (Deposit, error) {
	var resultDeposit Deposit

	return resultDeposit, nil
}

// UpdateDeposit find deposit entry in BD by ID and reWrite info from Deposit to DB
// TODO implement me
func (d Deposit) UpdateDeposit() (Deposit, error) {
	return Deposit{}, nil
}

// DeleteDeposit remove info about deposit from DB
// TODO implement me
func (d Deposit) DeleteDeposit() error {
	return nil
}

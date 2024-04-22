package surveys

import (
	"github.com/carterjackson/ranked-pick-api/internal/common"
)

type CreateParams struct {
	Title       string
	State       string
	Visibility  string
	Description string
}

func Create(ctx *common.Context, iparams interface{}) error {
	// params := iparams.(CreateParams)

	return nil
}

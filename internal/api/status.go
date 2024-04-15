package api

import "github.com/carterjackson/ranked-pick-api/internal/common"

func StatusHandler(ctx *common.Context) (interface{}, error) {
	return "ready", nil
}

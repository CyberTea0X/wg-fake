package app

import (
	"context"

	"github.com/lastbyte32/wg-fake/internal/sender"
)

type wizard interface {
	Magic() error
}

func Run(ctx context.Context, serverAddr string, localPort uint) error {
	wiz, err := sender.New(ctx, serverAddr, localPort)
	if err != nil {
		return err
	}
	return wiz.Magic()
}

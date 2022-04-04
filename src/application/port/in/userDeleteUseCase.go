package in

import (
	"context"
)

type DeleteUseCase interface {
	Delete(context.Context, string) error
}

package db

import (
	"context"
)

type Storer interface {
	ListTarget(context.Context) ([]Target, error)
	CreateTarget(context.Context, Target) (Target,error)

}

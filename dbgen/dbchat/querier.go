// Code generated by sqlc. DO NOT EDIT.

package dbchat

import (
	"context"
)

type Querier interface {
	AddMessage(ctx context.Context, arg AddMessageParams) (Message, error)
	GetMessages(ctx context.Context) ([]Message, error)
}

var _ Querier = (*Queries)(nil)
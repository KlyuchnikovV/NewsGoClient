package ui

import (
	"context"
	"main/models"
)

type MenuElement struct {
	Key     rune
	Header  string
	Handler func(ctx context.Context, client models.RssClient) error
}

func FindMenu(m []MenuElement, key rune) (*MenuElement, bool) {
	for i := range m {
		if m[i].Key == key {
			return &m[i], true
		}
	}
	return nil, false
}

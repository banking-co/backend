package dto

import (
	"gorm.io/gorm"
	"rabotyaga-go-backend/models"
	"rabotyaga-go-backend/types"
	"time"
)

type Item struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`

	Type   types.ItemType   `json:"type"`
	Rarity types.ItemRarity `json:"rarity"`

	Cost int `json:"cost"`

	IsBuyable  bool `json:"isBuyable"`
	IsSellable bool `json:"isSellable"`

	CreatedAt time.Time      `json:"createdAt,omitempty"`
	UpdatedAt time.Time      `json:"updatedAt,omitempty"`
	DeletedAt gorm.DeletedAt `json:"deletedAt,omitempty"`
}

func ItemWrap(i *models.Item) *Item {
	if i == nil {
		return nil
	}

	return &Item{
		ID:   i.ID,
		Name: i.Name,

		Type:   i.Type,
		Rarity: i.Rarity,

		Cost: i.Cost,

		IsBuyable:  i.IsBuyable,
		IsSellable: i.IsSellable,

		CreatedAt: i.CreatedAt,
		UpdatedAt: i.UpdatedAt,
		DeletedAt: i.DeletedAt,
	}
}

func ItemsWrap(is []*models.Item) []*Item {
	var items = make([]*Item, 0, len(is))
	for _, i := range is {
		if i != nil {
			items = append(items, ItemWrap(i))
		}
	}

	return items
}

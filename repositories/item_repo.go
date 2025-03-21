package repositories

import (
	"errors"
	"gin-fleamarket/models"
)

// repositoriesのinterface
// repositoriesが満たすべきメソッドを定義する
type IItemRepository interface {
	FindAll() (*[]models.Item, error)
	FindById(itemId uint) (*models.Item, error)
}

type ItemMemoryRepository struct {
	items []models.Item
}

// factory関数
func NewItemMemoryRepository(items []models.Item) IItemRepository {
	return &ItemMemoryRepository{items: items}
}

func (r *ItemMemoryRepository) FindAll() (*[]models.Item, error) {
	return &r.items, nil
}

func (r *ItemMemoryRepository) FindById(itemId uint) (*models.Item, error) {
	for _, v := range r.items {
		if v.Id == itemId {
			return &v, nil
		}
	}
	return nil, errors.New("Item not found")
}

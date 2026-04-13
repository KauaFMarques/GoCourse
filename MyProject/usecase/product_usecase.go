package usecase

import (
	"meu-projeto/model"
	"meu-projeto/repository"
)

// A interface PRECISA começar com letra maiúscula para ser exportada
type ProductUseCase interface {
	GetProducts() ([]model.Product, error)
	CreateProduct(product model.Product) (model.Product, error)
	GetProductById(id_product int) (*model.Product, error)
}

type productUseCase struct {
	repository repository.ProductRepository
}

// Recomendo retornar a interface ou o ponteiro da struct
func NewProductUsecase(repo repository.ProductRepository) ProductUseCase {
	return &productUseCase{
		repository: repo,
	}
}

func (pu *productUseCase) GetProducts() ([]model.Product, error) {
	return pu.repository.GetProducts()
}

func (pu *productUseCase) CreateProduct(product model.Product) (model.Product, error) {
	productId, err := pu.repository.CreateProduct(product)
	if err != nil {
		return model.Product{}, err // Retorna struct vazia se der erro
	}

	product.ID = productId
	return product, nil // Retorna o produto completo com o ID do banco
}

func (pu *productUseCase) GetProductById(id_product int) (*model.Product, error) {
	return pu.repository.GetProductById(id_product)
}

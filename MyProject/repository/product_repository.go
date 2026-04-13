package repository

import (
	"database/sql"
	"fmt"
	"meu-projeto/model"
)

type ProductRepository struct {
	connection *sql.DB
}

func NewProductRepository(connection *sql.DB) ProductRepository {
	return ProductRepository{
		connection: connection,
	}
}

func (pr *ProductRepository) GetProducts() ([]model.Product, error) {
	query := "SELECT id_product, name, price FROM products"
	rows, err := pr.connection.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close() // Isso garante que feche mesmo se der erro no loop

	var productList []model.Product
	var productObj model.Product

	for rows.Next() {
		err = rows.Scan(
			&productObj.ID,
			&productObj.Name,
			&productObj.Price)

		if err != nil {
			fmt.Println(err)
			return []model.Product{}, err
		}
		productList = append(productList, productObj)
	}

	rows.Close()

	return productList, nil
}

func (pr *ProductRepository) CreateProduct(product model.Product) (int, error) {
	var id int // <--- Você precisa declarar a variável antes de usar no .Scan()

	// Ajustei o SQL para bater com a tabela 'products' que criamos
	query, err := pr.connection.Prepare("INSERT INTO products" +
		"(name, price) " +
		"VALUES ($1, $2) RETURNING id_product") // Use o nome correto da coluna

	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	defer query.Close() // Use defer para fechar o statement com segurança

	// Aqui a variável 'id' agora existe e pode receber o valor
	err = query.QueryRow(product.Name, product.Price).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	return id, nil
}

// ADICIONE o * antes de model.Product para permitir o retorno de nil
func (pr *ProductRepository) GetProductById(id_product int) (*model.Product, error) {
	// Use nomes de colunas específicos para evitar problemas de Scan
	query, err := pr.connection.Prepare("SELECT id_product, name, price FROM products WHERE id_product = $1")
	if err != nil {
		return nil, err // Agora funciona porque o retorno aceita ponteiro
	}
	defer query.Close()

	var produto model.Product
	err = query.QueryRow(id_product).Scan(
		&produto.ID,
		&produto.Name,
		&produto.Price,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Produto não encontrado
		}
		return nil, err
	}

	return &produto, nil // Retorna o ENDEREÇO do produto
}

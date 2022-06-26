package repository

import (
	"database/sql"
	"log"
	"tech-test-azura-lab/models"
)

type ProductRepo struct {
	DB *sql.DB
}

func NewProductRepostory(db *sql.DB) *ProductRepo {
	return &ProductRepo{
		DB: db,
	}
}

func (r *ProductRepo) GetAllProduct() ([]models.Product, error) {
	sqlstmt := `SELECT id, nama, harga, rating FROM product`
	result := []models.Product{}

	rows, err := r.DB.Query(sqlstmt)

	if err != nil {
		log.Fatalln("Error : ", err.Error())
		if err == sql.ErrNoRows {
			return []models.Product{}, nil
		}
		return nil, err
	}

	for rows.Next() {
		var p models.Product = models.Product{}
		err = rows.Scan(&p.ID, &p.Nama, &p.Harga, &p.Rating)
		if err != nil {
			log.Println(err.Error())
		}
		result = append(result, p)
	}
	return result, nil
}

func (r *ProductRepo) GetProductDetail(id int) (*models.Product, error) {

	sqlstmt := `SELECT * FROM product WHERE ID = $1`
	var p models.Product

	row := r.DB.QueryRow(sqlstmt, id)

	err := row.Scan(&p.ID, &p.Nama, &p.Harga, &p.Rating, &p.Likes, &p.Deskripsi)

	if err != nil {
		log.Println("Error : ", err.Error())
		return nil, err
	}

	return &p, nil
}

func (r *ProductRepo) CreateProduct(product models.Product) error {
	sqlstmt := `INSERT INTO product (nama, harga, rating, likes, deskripsi) VALUES ($1, $2, $3, $4, $5)`

	_, err := r.DB.Exec(sqlstmt, product.Nama, product.Harga, product.Rating, product.Likes, product.Deskripsi)

	if err != nil {
		log.Println("Error When Inserting to The Table : ", err.Error())
		return err
	}

	return nil
}

func (r *ProductRepo) UpdateProduct(product models.Product) error {
	sqlstmt := `
  UPDATE product
  SET
    nama = $1,
    harga = $2,
    rating = $3,
    likes = $4,
    deskripsi = $5
  WHERE id = $6;`

	_, err := r.DB.Exec(sqlstmt, product.Nama, product.Harga, product.Rating, product.Likes, product.Deskripsi, product.ID)

	if err != nil {
		log.Println("Error When Updating a Row : ", err.Error())
		return err
	}

	return nil
}

func (r *ProductRepo) DeleteProduct(id int) error {
	sqlstmt := `DELETE FROM product WHERE id = $1`

	_, err := r.DB.Exec(sqlstmt, id)

	if err != nil {
		log.Println("Error When Deleting in a Row : ", err.Error())
		return err
	}
	return nil
}

package repo

import "fmt"

type Product struct {
	ID          int     `json:"id"` //add tag
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgURL      string  `json:"imageUrl"`
}

type ProductRepo interface {
	Create(p Product) (*Product, error)
	Get(id int) (*Product, error)
	List() ([]*Product, error)
	Update(id int, p Product) (*Product, error)
	Delete(id int) error
}

type productRepo struct {
	productList []*Product
}

func NewProductRepo() ProductRepo {
	repo := &productRepo{}

	generateInitialProducts(repo)

	return repo
}

func (r productRepo) Create(p Product) (*Product, error) {

	p.ID = len(r.productList) + 1
	r.productList = append(r.productList, &p)

	return &p, nil
}
func (r productRepo) Get(id int) (*Product, error) {
	for _, p := range r.productList {
		if p.ID == id {
			return p, nil
		}
	}
	return nil, fmt.Errorf("product not found")
}
func (r productRepo) List() ([]*Product, error) {
	var products []*Product
	for _, p := range r.productList {
		products = append(products, p)
	}
	return products, nil
}
func (r productRepo) Update(id int, p Product) (*Product, error) {
	for i, product := range r.productList {
		if product.ID == id {
			r.productList[i] = &p
			return &p, nil
		}
	}
	return nil, fmt.Errorf("product not found")
}
func (r productRepo) Delete(id int) error {
	for i, product := range r.productList {
		if product.ID == id {
			r.productList = append(r.productList[:i], r.productList[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("product not found")
}

func generateInitialProducts(r *productRepo) {
	prd1 := &Product{
		ID:          1,
		Title:       "Orange",
		Description: "This is orange.",
		Price:       150.00,
		ImgURL:      "https://encrypted-tbn1.gstatic.com/images?q=tbn:ANd9GcRwJf4xPOftZLLgkWjr2eMMumu9XuRdKiGc1eZXFku9OmA4lSymIPBm_vU0bFIM_BjpYOsI_pX7O9c9sRUxfkXxE0N1x_bWSERw7SXHP3A",
	}

	prd2 := &Product{
		ID:          2,
		Title:       "Apple",
		Description: "This is apple.",
		Price:       120.00,
		ImgURL:      "https://encrypted-tbn3.gstatic.com/images?q=tbn:ANd9GcSbWqLu3A4mBXzZNi1POVqPrGV4UAuPw3Bl9RwTDjJqjCQrMiid5VgbDGHbAIA0rcFkAWdpjba5sqFdLdIvRXZ27_H_dPxXDpebJuqHiHQ",
	}

	prd3 := &Product{
		ID:          3,
		Title:       "Banana",
		Description: "This is banana.",
		Price:       40.00,
		ImgURL:      "https://www.allrecipes.com/thmb/lc7nSL9L5zMHXz9t6PMAVm9biNM=/1500x0/filters:no_upscale():max_bytes(150000):strip_icc()/ar-new-banana-adobe-ar-2x1-917fdde58d194b529b41042ebff1c031.jpg",
	}

	r.productList = append(r.productList, prd1)
	r.productList = append(r.productList, prd2)
	r.productList = append(r.productList, prd3)

}

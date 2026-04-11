package database

type Product struct {
	ID          int     `json:"id"` //add tag
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgURL      string  `json:"imageUrl"`
}

var productList []Product

func Store(p Product) Product {
	p.ID = len(productList) + 1
	productList = append(productList, p)
	return p
}

func List() []Product {
	return productList
}

func GetById(pId int) *Product {
	for _, product := range productList {
		if product.ID == pId {
			return &product
		}
	}
	return nil
}

func Update(p Product) {
	for idx, product := range productList {
		if product.ID == p.ID {
			productList[idx] = p
		}
	}
}

func Delete(pId int) {
	var tempList []Product
	for _, product := range productList {
		if product.ID != pId {
			tempList = append(tempList, product)
			// tempList[idx] = product

		}
	}
	productList = tempList
	// productList = append(productList, tempList...)
}

func init() {

	prd1 := Product{
		ID:          1,
		Title:       "Orange",
		Description: "This is orange.",
		Price:       150.00,
		ImgURL:      "https://encrypted-tbn1.gstatic.com/images?q=tbn:ANd9GcRwJf4xPOftZLLgkWjr2eMMumu9XuRdKiGc1eZXFku9OmA4lSymIPBm_vU0bFIM_BjpYOsI_pX7O9c9sRUxfkXxE0N1x_bWSERw7SXHP3A",
	}
	prd2 := Product{
		ID:          2,
		Title:       "Apple",
		Description: "This is apple.",
		Price:       120.00,
		ImgURL:      "https://encrypted-tbn3.gstatic.com/images?q=tbn:ANd9GcSbWqLu3A4mBXzZNi1POVqPrGV4UAuPw3Bl9RwTDjJqjCQrMiid5VgbDGHbAIA0rcFkAWdpjba5sqFdLdIvRXZ27_H_dPxXDpebJuqHiHQ",
	}
	prd3 := Product{
		ID:          3,
		Title:       "Banana",
		Description: "This is banana.",
		Price:       40.00,
		ImgURL:      "https://www.allrecipes.com/thmb/lc7nSL9L5zMHXz9t6PMAVm9biNM=/1500x0/filters:no_upscale():max_bytes(150000):strip_icc()/ar-new-banana-adobe-ar-2x1-917fdde58d194b529b41042ebff1c031.jpg",
	}
	prd4 := Product{
		ID:          4,
		Title:       "Grapes",
		Description: "This is grapes.",
		Price:       140.00,
		ImgURL:      "https://d2t8nl1y0ie1km.cloudfront.net/images/6706bbb12b01cd80b29a0488_Grapes-Red-N-500gm_1.webp",
	}
	prd5 := Product{
		ID:          5,
		Title:       "Mango",
		Description: "This is mango.",
		Price:       80.00,
		ImgURL:      "https://belescooverseas.com/wp-content/uploads/2024/04/Mango.jpeg",
	}

	productList = append(productList, prd1)
	productList = append(productList, prd2)
	productList = append(productList, prd3)
	productList = append(productList, prd4)
	productList = append(productList, prd5)
}
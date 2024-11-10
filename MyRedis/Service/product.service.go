package service

import repository "github.com/paramet02/Repository"

type proService struct {
	proRepo repository.ProductRepo
}

func NewproService(proRepo repository.ProductRepo) ProductService {
	return proService{proRepo}
}

func (s proService) CreateProduct(product Product) error {
	// แปลงจาก service.Product ไปเป็น repository.product โดยใช้ NewProduct
	// เพราะ service.Product ไม่สามารถใส่ค่าลงใน repository.product โดยตรงได้ 
	// จึงต้องใช้ฟังก์ชัน NewProduct ในการแปลงค่าจาก service ไปเป็น repository
	repoProduct := repository.NewProduct(product.Name, product.Quantity)

	// เรียกใช้ CreateProduct ของ repository เพื่อสร้างข้อมูลใหม่ในฐานข้อมูล
	return s.proRepo.CreateProduct(repoProduct)
}

func (s proService) UpdateProduct(product Product) error {
	// แปลงจาก service.Product ไปเป็น repository.product โดยใช้ NewUpdateProduct
	// เพราะ service.Product ไม่สามารถใส่ค่าลงใน repository.product โดยตรงได้
	// จึงต้องใช้ฟังก์ชัน NewUpdateProduct เพื่อแปลงค่าจาก service ไปเป็น repository
	repoProduct := repository.NewUpdateProduct(product.ID, product.Name, product.Quantity)

	// เรียกใช้ UpdateProduct ของ repository เพื่ออัปเดตข้อมูลในฐานข้อมูล
	err := s.proRepo.UpdateProduct(repoProduct)
	if err != nil {
		return err // ถ้ามีข้อผิดพลาดเกิดขึ้นใน repository ให้ส่งกลับ
	}
	return nil // ถ้าไม่มีข้อผิดพลาดให้คืนค่า nil
}

func (s proService) GetProducts() (pro []Product, err error) {
	// เรียกใช้ GetProducts ของ repository เพื่อดึงข้อมูลทั้งหมด
	products, err := s.proRepo.GetProducts()
	if err != nil {
		return nil, err
	}

	// แปลงข้อมูลจาก repository.product ไปเป็น service.Product 
	// เพื่อให้สามารถส่งกลับไปยัง client ได้
	for _, p := range products {
		pro = append(pro, Product{
			ID:       p.ID,
			Name:     p.Name,
			Quantity: p.Quantity,
		})
	}

	return pro, nil
}

func (s proService) GetProduct(name string) (pro []Product, err error) {
	// เรียกใช้ GetProduct ของ repository เพื่อค้นหาสินค้าตามชื่อ
	newproduct, err := s.proRepo.GetProduct(name)
	if err != nil {
		return nil, err
	}

	// แปลงข้อมูลจาก repository.product ไปเป็น service.Product 
	// เพื่อให้สามารถส่งกลับไปยัง client ได้
	for _, p := range newproduct {
		pro = append(pro, Product{
			ID:       p.ID,
			Name:     p.Name,
			Quantity: p.Quantity,
		})
	}

	return pro, nil
}
package sellers

// Modelo de sellers
type Seller struct {
	ID           int    `json:"id"`
	Cid          int    `json:"cid"`
	Company_name string `json:"company_name"`
	Adress       string `json:"adress"`
	Telephone    int    `json:"telephone"`
}

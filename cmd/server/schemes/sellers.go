package schemes

type Seller struct {
	ID           int    `json:"id"`
	Cid          int    `json:"cid"`
	Company_name string `json:"company_name"`
	Address      string `json:"adress"`
	Telephone    int    `json:"telephone"`
}

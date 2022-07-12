package schemas

type Seller struct {
	ID           int    `json:"id"`
	Cid          int    `json:"cid"`
	Company_name string `json:"company_name"`
	Address      string `json:"address"`
	Telephone    string `json:"telephone"`
}

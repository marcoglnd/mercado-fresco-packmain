package domain

type Carrier struct {
	ID          int64  `json:"id"`
	Cid         string `json:"cid" binding:"required"`
	CompanyName string `json:"company_name" binding:"required"`
	Address     string `json:"address" binding:"required"`
	Telephone   string `json:"telephone" binding:"required"`
	LocalityId  int64  `json:"locality_id" binding:"required"`
}

type CarrierReport struct {
	LocalityId    int64  `json:"locality_id"`
	LocalityName  string `json:"locality_name"`
	CarriersCount int64  `json:"carriers_count"`
}

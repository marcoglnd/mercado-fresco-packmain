package domain

type CreateCarrierInput struct {
	Cid         string `json:"cid" binding:"required"`
	CompanyName string `json:"company_name" binding:"required"`
	Address     string `json:"address" binding:"required"`
	Telephone   string `json:"telephone" binding:"required"`
	LocalityId  int64  `json:"locality_id" binding:"required"`
}

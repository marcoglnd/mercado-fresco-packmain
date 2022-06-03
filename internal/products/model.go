package products

type Product struct {
	Id                             int     `json:"id"`
	Description                    string  `json:"description"`
	ExpirationRate                 int     `json:"expiration_rate"`
	FreezingRate                   int     `json:"freezing_rate"`
	Height                         float64 `json:"height"`
	Length                         float64 `json:"length"`
	NetWeight                      float64 `json:"netweight"`
	ProductCode                    string  `json:"product_code"`
	RecommendedFreezingTemperature float64 `json:"recommended_freezing_temperature"`
	Width                          float64 `json:"width"`
	ProductTypeId                 int     `json:"product_type_id"`
	SellerId                       int     `json:"seller_id"`
}

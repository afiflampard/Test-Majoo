package entities

type OrderRequest struct {
	IDPembeli uint `json:"id_pembeli"`
	IDProduct uint `json:"id_product"`
	Jumlah    uint `json:"jumlah"`
}

package entities

type OrderRequest struct {
	IDPembeli uint `json:"id_pembeli"`
	IDProduct uint `json:"id_product"`
	Jumlah    uint `json:"jumlah"`
}
type RequestProduct struct {
	NamaProduk string `json:"nama_product"`
	Harga      uint   `json:"harga"`
	Stock      uint   `json:"stock"`
}

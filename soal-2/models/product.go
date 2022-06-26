package models

type Product struct {
	ID        uint   `json:"id" db:"id"`
	Nama      string `json:"nama" db:"nama"`
	Harga     uint   `json:"harga" db:"harga"`
	Rating    uint   `json:"rating" db:"rating"`
	Likes     uint   `json:"likes,omitempty" db:"likes"`
	Deskripsi string `json:"desc,omitempty" db:"deskripsi"`
}

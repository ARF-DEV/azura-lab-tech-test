package main

import (
	"fmt"
	"sort"
)

type Produk struct {
	Nama   string
	Harga  int
	Rating float32
	Likes  uint
}

func (p Produk) String() string {
	return fmt.Sprintf("%v | %v | %v | %v\n", p.Nama, p.Harga, p.Rating, p.Likes)
}

type compFunc func(p1, p2 Produk) bool //Comparison Function

type ProdukSorter struct {
	ProdukList   []Produk
	CompFuncList []compFunc
}

func OrderBy(lessFunctions ...compFunc) *ProdukSorter {
	return &ProdukSorter{
		CompFuncList: lessFunctions,
	}
}

func (ps *ProdukSorter) Sort(slice []Produk) {
	ps.ProdukList = slice
	sort.Sort(ps)
}

func (ps *ProdukSorter) Len() int {
	return len(ps.ProdukList)
}

func (ps *ProdukSorter) Swap(i, j int) {
	ps.ProdukList[i], ps.ProdukList[j] = ps.ProdukList[j], ps.ProdukList[i]
}

func (ps *ProdukSorter) Less(i, j int) bool {
	p, q := ps.ProdukList[i], ps.ProdukList[j]

	for _, comp := range ps.CompFuncList {
		if comp(p, q) {
			return true
		} else if comp(q, p) {
			return false
		}
	}

	return ps.CompFuncList[len(ps.CompFuncList)-1](p, q)
}

func main() {
	listOfProduk := []Produk{
		{"Indomie", 3000, 5, 150},
		{"Laptop", 4000000, 4.5, 123},
		{"Aqua", 3000, 4, 250},
		{"Smart TV", 4000000, 4.5, 42},
		{"Headphone", 4000000, 3.5, 90},
		{"Very Smart TV", 4000000, 3.5, 87},
	}

	Harga := func(i, j Produk) bool {
		return i.Harga > j.Harga
	}

	Rating := func(i, j Produk) bool {
		return i.Rating > j.Rating
	}

	Like := func(i, j Produk) bool {
		return i.Likes > j.Likes
	}

	OrderBy(Harga, Rating, Like).Sort(listOfProduk)

	fmt.Println(listOfProduk)
}

package identity

import "fmt"

type Person struct {
	absen     int
	nama      string
	alamat    string
	pekerjaan string
	alasan    string
}

var Identity = []Person{
	{absen: 1, nama: "Taufik", alamat: "Jakarta", pekerjaan: "Backend Engineer", alasan: "Lorem ipsum dolor sit amet"},
	{absen: 2, nama: "Hidayah", alamat: "Bandung", pekerjaan: "Backend Engineer", alasan: "consectetur, adipisci velit"},
	{absen: 3, nama: "Muhammad", alamat: "Jogjakarta", pekerjaan: "Android Engineer", alasan: "sed quia non numquam eius modi tempora incidunt"},
	{absen: 4, nama: "Temmy", alamat: "Pontianak", pekerjaan: "Frontend Engineer", alasan: "ut labore et dolore magnam aliquam quaerat voluptatem"},
	{absen: 5, nama: "Alex", alamat: "Aceh", pekerjaan: "Frontend Engineer", alasan: "Ut enim ad minima veniam"},
}

func (identity Person) ShowData() {
	fmt.Printf("Nama 		: %v\n", identity.nama)
	fmt.Printf("Alamat		: %v\n", identity.alamat)
	fmt.Printf("Pekerjaan	: %v\n", identity.pekerjaan)
	fmt.Printf("Alasan		: %v\n", identity.alasan)
}

package main

import "fmt"

type Kendaraan struct {
	Suara string
}

func (k *Kendaraan) Akselerasi() {
	fmt.Println(k.Suara)
}

type Sepeda struct {
	Kendaraan
	Rantai string
}

func (s *Sepeda) Akselerasi() {
	s.Rantai = "Perlu perbaikan"
	fmt.Println(s.Suara)
}

type Mobil struct {
	Kendaraan
	Bensin string
}

func (m *Mobil) Akselerasi() {
	m.Bensin = "Kosong"
	fmt.Println(m.Suara)
}

func main() {
	kendaraan1 := Kendaraan{Suara: "Broom"}
	kendaraan1.Akselerasi()

	sepeda1 := Sepeda{Kendaraan: Kendaraan{Suara: "Swoosh"}, Rantai: "Normal"}
	sepeda1.Akselerasi()
	fmt.Println(sepeda1.Rantai)

	mobil1 := Mobil{Kendaraan: Kendaraan{Suara: "Vroom"}, Bensin: "Penuh"}
	mobil1.Akselerasi()
	fmt.Println(mobil1.Bensin)
}

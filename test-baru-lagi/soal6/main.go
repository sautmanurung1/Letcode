package main

import "fmt"

func main() {
	people := map[string]int{"A": 100, "C": 100, "B": 50, "D": 50}
	bags := map[string]int{"A": 25, "C": 25}

	// Menyebrangi sungai dengan cara berulang sampai semua orang sudah menyebrang
	for len(people) > 0 {
		boat := []string{}
		weight := 0

		// Menambahkan orang ke perahu sampai mencapai kapasitas maksimum
		for person, personWeight := range people {
			if personWeight+weight <= 100 {
				boat = append(boat, person)
				weight += personWeight

				// Jika orang tersebut membawa tas, maka menambahkan berat tas ke dalam perahu
				if bags[person] > 0 {
					weight += bags[person]
				}
			}
		}

		// Menghapus orang dan tas yang menyeberang dari daftar orang dan tas
		for _, person := range boat {
			delete(people, person)

			// Jika orang membawa tas, maka menghapus tas dari daftar tas
			if bags[person] > 0 {
				delete(bags, person)
			}

			// Mencetak pesan bahwa orang berhasil menyeberang dengan atau tanpa tas
			if person == "B" || person == "D" {
				fmt.Printf("%s menyebrang tanpa tas\n", person)
			} else {
				fmt.Printf("%s menyebrang dengan tas berat %d kg\n", person, bags[person])
			}
		}

		// Mencetak pesan bahwa perahu sudah menyebrang dan jumlah orang serta tas yang tersisa
		fmt.Printf("Perahu menyebrang dengan %d kg beban, dan %d orang serta %d tas tersisa di sisi awal sungai\n", weight, len(people), len(bags))
	}

	// Mencetak pesan bahwa semua orang dan barang berhasil menyebrang
	fmt.Println("Semua orang dan barang berhasil menyebrang")
}

package main
import (
	"fmt"
	"log"
	"os"
	"strconv"
)

type biodata struct {
	Nama      string
	Alamat    string
	Pekertaan string
	Alasan    string
}

func main() {
	var allbio = []biodata{
		{
			Nama:      "Asep",
			Alamat:    "Cirebon",
			Pekertaan: "Developer",
			Alasan:    "menambah wawasan ",
		},
		{
			Nama:      "Adi",
			Alamat:    "Cirebon",
			Pekertaan: "Developer",
			Alasan:    "menambah wawasan ",
		},
		{
			Nama:      "Bagas",
			Alamat:    "Cirebon",
			Pekertaan: "Developer",
			Alasan:    "menambah wawasan ",
		},
		{
			Nama:      "Agung",
			Alamat:    "Cirebon",
			Pekertaan: "Developer",
			Alasan:    "menambah wawasan ",
		},
	}
	args := os.Args[1:2]      
	idx := args[0]                
	ind, err := strconv.Atoi(idx)

	if err != nil {
		log.Fatal(err)
	}

	if ind < 4 {
		fmt.Printf("Nama = %s", allbio[ind].Nama)
		fmt.Println("")
		fmt.Printf("Alamat = %s", allbio[ind].Alamat)
		fmt.Println("")
		fmt.Printf("Pekerjaan = %s", allbio[ind].Pekertaan)
		fmt.Println("")
		fmt.Printf("Alasan = %s", allbio[ind].Alasan)
	} else {
		fmt.Println("Tidak ada data ")
	}
}
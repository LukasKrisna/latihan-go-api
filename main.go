package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type ToDo struct {
	Kegiatan string `json:"kegiatan"`
	Waktu    string `json:"waktu"`
}

type JSONResponse struct {
	Code    int    `json:"code"`
	Success bool   `json:"success"`
	Message string `json:"message"`
	// Data    []ToDo `json:"data"`
	Data interface{} `json:"data"`
}

func main() {
	daftarKegiatan := []ToDo{}
	daftarKegiatan = append(daftarKegiatan, ToDo{"Mengikuti kelas GoAPI H1", "2021-11-22"})
	daftarKegiatan = append(daftarKegiatan, ToDo{"Mengikuti kelas GoAPI H2", "2021-11-23"})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			w.Header().Add("Content-Type", "application/json")

			// res := JSONResponse{
			// 	http.StatusOK,
			// 	true,
			// 	"Uji Coba GET method pada Postman",
			// 	[]ToDo{},
			// }
			// resJSON, err := json.Marshal(res)
			// if err != nil {
			// 	http.Error(w, "Terjadi kesalahan", http.StatusInternalServerError)
			// }
			// w.Write(resJSON)

			res := JSONResponse{
				http.StatusOK,
				true,
				"Berhasil mendapatkan daftar aktivitas",
				daftarKegiatan,
			}

			resJSON, err := json.Marshal(res)
			if err != nil {
				fmt.Println("Terjadi kesalahan")
				http.Error(w, "Terjadi kesalahan", http.StatusInternalServerError)
				return
			}

			w.Write(resJSON)

		} else if r.Method == "POST" {
			w.Header().Add("Content-Type", "application/json")
			jsonDecode := json.NewDecoder(r.Body)
			aktivitasBaru := ToDo{}
			res := JSONResponse{}

			if err := jsonDecode.Decode(&aktivitasBaru); err != nil {
				fmt.Println("Terjadi kesalahan")
				http.Error(w, "Terjadi kesalahan", http.StatusInternalServerError)
				return
			}

			res.Code = http.StatusCreated
			res.Success = true
			res.Message = "Berhasil menambahkan data"
			res.Data = aktivitasBaru

			resJSON, err := json.Marshal(res)
			if err != nil {
				fmt.Println("Terjadi kesalahan")
				http.Error(w, "Terjadi kesalahan saat ubah JSON", http.StatusInternalServerError)
				return
			}

			w.Write(resJSON)

		}
	})

	fmt.Println("Listening on: 8080 ....")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

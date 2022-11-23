package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type apiResults struct{
	Count int `json:"count"`
	Results []struct{
		Name string `json:"name"`
		RotationPeriod string `json:"rotation_period"`
		OrbitalPeriod string `json:"orbital_period"`
		Diameter string `json:"diameter"`
		Climate string `json:"climate"`
		Gravity string `json:"gravity"`
		Terrain string `json:"terrain"`
		SurfaceWater string `json:"surface_water"`
		Population string `json:"population"`
	} `json:"results"`
}

func main() {
	var results apiResults
  response, err :=	http.Get("https://swapi-trybe.herokuapp.com/api/planets/?page=1")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	body, _ := io.ReadAll(response.Body)
	json.Unmarshal(body, &results)
	fmt.Println("resultados", results)

	http.HandleFunc("/test", func(res http.ResponseWriter, req *http.Request) {
		parsed, _ := json.Marshal(results)
		res.Write(parsed)
	})
	http.ListenAndServe(":9090", nil)
}

// func handler(res http.ResponseWriter, req *http.Request) {
	
// }

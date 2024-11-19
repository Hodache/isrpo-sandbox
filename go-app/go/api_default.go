/*
 * Дендрарий
 *
 * Учет деревьев, кустарников, лиан
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
)

var plants = []Plant{
	{
		Id:           0,
		Species:      "Сосна обыкновенная",
		PlantingYear: 2002,
	},
	{
		Id:           1,
		Species:      "Сосна обыкновенная",
		PlantingYear: 2004,
	},
	{
		Id:           2,
		Species:      "Сосна кедровая сибирская",
		PlantingYear: 2003,
	},
	{
		Id:           3,
		Species:      "Сосна кедровая сибирская",
		PlantingYear: 2011,
	},
	{
		Id:           4,
		Species:      "Дуб крупноплодный",
		PlantingYear: 2005,
	},
	{
		Id:           5,
		Species:      "Дуб монгольский",
		PlantingYear: 2005,
	},
	{
		Id:           6,
		Species:      "Дуб каштанолистный",
		PlantingYear: 2015,
	},
	{
		Id:           7,
		Species:      "Берёза пушистая",
		PlantingYear: 2007,
	},
	{
		Id:           8,
		Species:      "Берёза кустарниковая",
		PlantingYear: 2009,
	},
	{
		Id:           9,
		Species:      "Шиповник красно-бурый",
		PlantingYear: 2017,
	},
	{
		Id:           10,
		Species:      "Шиповник блестящий",
		PlantingYear: 2020,
	},
	{
		Id:           11,
		Species:      "Шиповник сизый",
		PlantingYear: 2022,
	},
}
var idCounter int = 12

func CreatePlant(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	var plant Plant

	err := json.NewDecoder(r.Body).Decode(&plant)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else if plant.Species == "" {
		http.Error(w, "Вид растения должен быть указан", http.StatusBadRequest)
		return
	} else if plant.PlantingYear == 0 {
		http.Error(w, "Год посадки должен быть указан", http.StatusBadRequest)
		return
	}

	plant.Id = idCounter
	idCounter++
	plants = append(plants, plant)
	w.WriteHeader(200)
}

func DeletePlant(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for i, plant := range plants {
		if plant.Id == id {
			plants = append(plants[:i], plants[i+1:]...)
			w.WriteHeader(200)
			return
		}
	}

	http.Error(w, "Растение с таким id не найдено", 404)
}

func GetPlantById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for _, plant := range plants {
		if plant.Id == id {
			json.NewEncoder(w).Encode(plant)
			return
		}
	}

	http.Error(w, "Растение с таким id не найдено", 404)
}

func GetPlants(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// Генерация и логирование ошибок
	errorChance := rand.Intn(100)
	if errorChance < 3 {
		logrus.Error("Произошла ошибка с шансом 3%, код ответа 500")
		http.Error(w, "Ошибка сервера (3%)", http.StatusInternalServerError)
		return
	} else if errorChance < 10 {
		logrus.Error("Произошла ошибка с шансом 7%, код ответа 502")
		http.Error(w, "Ошибка сервера (7%)", http.StatusBadGateway)
		return
	}

	json.NewEncoder(w).Encode(plants)
}

func UpdatePlant(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var newPlant Plant

	bodyError := json.NewDecoder(r.Body).Decode(&newPlant)
	if bodyError != nil {
		http.Error(w, bodyError.Error(), http.StatusBadRequest)
		return
	}

	for i, plant := range plants {
		if plant.Id == id {
			plants[i].PlantingYear = newPlant.PlantingYear
			plants[i].Species = newPlant.Species
			w.WriteHeader(200)
			return
		}
	}

	http.Error(w, "Растение с таким id не найдено", 404)
}

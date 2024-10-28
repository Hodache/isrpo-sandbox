/*
 * Дендрарий
 *
 * Учет деревьев, кустарников, лиан
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type Plant struct {
	Id int `json:"id,omitempty"`
	// Биологический вид
	Species string `json:"species"`
	// Год посадки растения
	PlantingYear int `json:"planting_year"`
}
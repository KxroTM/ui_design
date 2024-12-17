package app

import (
	"database/sql"
	"fmt"
	"math/rand"
	"strconv"

	"github.com/go-faker/faker/v4"
	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

func CreateDb() {
	_, err := Db.Exec(`
		CREATE TABLE IF NOT EXISTS products (
			id INTEGER PRIMARY KEY,
			name TEXT,
			description TEXT,
			price INTEGER,
			image_url TEXT
		)
	`)

	if err != nil {
		panic(err)
	}

}

func FakerProducts() {
	for i := 0; i < 600; i++ {
		product := Product{
			Name:        faker.Word(),
			Description: faker.Sentence(),
			Price:       generateRandomPrice(10, 200),
			Image_Url:   generateRandomImageURL(300, 300),
		}
		_, err := Db.Exec(`
		INSERT INTO products (name, description, price, image_url)
		VALUES (?, ?, ?, ?)
	`, product.Name, product.Description, product.Price, product.Image_Url)

		if err != nil {
			panic(err)
		}
	}
}

func generateRandomPrice(min, max float64) string {
	return strconv.FormatFloat(min+rand.Float64()*(max-min), 'f', 2, 32)
}

func generateRandomImageURL(width, height int) string {
	randomValue := rand.Intn(1000000)
	return fmt.Sprintf("https://picsum.photos/%d/%d?random=%d", width, height, randomValue)
}

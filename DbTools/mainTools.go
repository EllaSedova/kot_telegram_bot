package DbTools

import (
	"database/sql"
	"fmt"
	"math/rand"

	"github.com/BurntSushi/toml"
	_ "github.com/lib/pq"
)

type Config struct {
	Bot struct {
		Token string
	}
	Database struct {
		Port     int
		Host     string
		User     string
		Name     string
		Password string
	}
}

func MakeConnection() *sql.DB {
	var c Config
	flConfigPath := "config.toml"
	if _, err := toml.DecodeFile(flConfigPath, &c); err != nil {
		panic(err)
	}
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
		c.Database.Host, c.Database.Port, c.Database.User, c.Database.Password, c.Database.Name)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	return db
}
func AddCatPhoto(message string, msgType string) {
	fmt.Println("photo", message, msgType)
	db := MakeConnection()
	result, err := db.Exec("insert into cat_photo (photo, breed) values ($1, $2)", message, msgType)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
func GetCatPhoto(cat_breed string) string {
	var text = "no"
	db := MakeConnection()
	rows, err := db.Query("select distinct photo from cat_photo where breed = $1", cat_breed)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	ids := []string{}

	for rows.Next() {
		photo := ""
		err := rows.Scan(&photo)
		if err != nil {
			fmt.Println(err)
			continue
		}
		ids = append(ids, photo)
	}
	var row = ""
	if len(ids) != 0 {
		row = ids[rand.Intn(len(ids))]
	} else {
		row = text
	}
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(row)
	return row
}

package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

type App struct {
	db *pgx.Conn
}

func NewApp(conn *pgx.Conn) *App {
	return &App{db: conn}
}

type Image struct {
	Title   string `json:"title"`
	AltText string `json:"alt_text"`
	URL     string `json:"url"`
}

func (app App) GetImages(w http.ResponseWriter, r *http.Request) {

	// images := []Image{
	// 	{
	// 		Title:   "Sunset",
	// 		AltText: "Clouds at sunset",
	// 		URL:     "https://images.unsplash.com/photo-1506815444479-bfdb1e96c566?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1000&q=80",
	// 	},
	// 	{
	// 		Title:   "Mountain",
	// 		AltText: "A mountain at sunset",
	// 		URL:     "https://images.unsplash.com/photo-1540979388789-6cee28a1cdc9?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1000&q=80",
	// 	},
	// }

	// get images

	rows, err := app.db.Query(context.Background(), "SELECT title, url, alt_text FROM public.images")
	if err != nil {
		log.Printf("query error: %s", err.Error())
		return
	}

	defer rows.Close()

	var images []Image

	for rows.Next() {
		var url, alt_text, title string
		err := rows.Scan(&title, &url, &alt_text)

		if err != nil {
			log.Printf("query error: %s", err.Error())
			return
		}
		image := Image{Title: title, URL: url, AltText: alt_text}

		images = append(images, image)

	}

	var resp []byte

	ind := r.URL.Query().Get("indent")
	if ind != "" {
		i, err := strconv.Atoi(ind)
		if err != nil {
			log.Fatal(err)
		}
		indent := strings.Repeat(" ", i)
		resp, err = json.MarshalIndent(images, "", indent)
	} else {
		resp, err = json.Marshal(images)
		if err != nil {
			log.Fatal("unable to marshal json")
		}
	}
	w.Write(resp)
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error loading .env file %v\n", err)
	}

	// connect to database
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("error connecting to database: %v\n", err)
	}
	defer conn.Close(context.Background())

	app := NewApp(conn)

	http.HandleFunc("GET /images.json", app.GetImages)
	log.Println("starting server on 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

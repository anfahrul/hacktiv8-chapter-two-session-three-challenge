package main

import (
	"hacktiv8-chapter-two-session-three-challenge/database"
	routers "hacktiv8-chapter-two-session-three-challenge/routes"

	_ "github.com/lib/pq"
)

func main() {
	db := database.GetConnection()
	defer db.Close()

	const PORT = ":8080"

	routers.StartServer().Run(PORT)
}

package main

import (
	"encoding/csv"
	"os"
)

func main() {
	writer := csv.NewWriter(os.Stdout)

	_ = writer.Write([]string{"syauqi", "damario", "djohan"})
	_ = writer.Write([]string{"akane", "miyabi", "hikari"})
	_ = writer.Write([]string{"hideo", "takaya", "shin"})

	writer.Flush()
}
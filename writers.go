package nflscores

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

// WriteFile writes to a given filename
func WriteFile(fileName string) {
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	WriteCSV(file)
	file.Close()
	return
}

// WriteCSV scrapes games and writes to given io.Writer
func WriteCSV(output io.Writer) {
	w := csv.NewWriter(output)

	if err := w.Write(GetGameHeaders()); err != nil {
		log.Fatalln("error writing record to csv:", err)
	}

	// Scrape the available years
	years := ScrapeYears()

	// Loop through each year
	for _, year := range years {
		// Scrape the available weeks and week names in a given year
		weekIDs, weekNames := ScrapeWeeks(year)

		// Loop through each week in a given year
		for index := range weekIDs {
			// Scrape all game data from a given week in a given year
			newGames := ScrapeGames(year, weekIDs[index], weekNames[index])

			// Append each new game to the running slice of games
			for _, newGame := range newGames {
				if err := w.Write(newGame.ToRow()); err != nil {
					log.Fatalln("error writing record to csv:", err)
				}
			} // Append
		} // Loop through weeks
	} // Loop through years

	// Write any buffered data to the underlying writer (standard output).
	w.Flush()

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}

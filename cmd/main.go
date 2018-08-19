package main

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/ocelotsloth/nflscores"
)

func main() {
	w := csv.NewWriter(os.Stdout)

	if err := w.Write(nflscores.GetGameHeaders()); err != nil {
		log.Fatalln("error writing record to csv:", err)
	}

	// Scrape the available years
	years := nflscores.ScrapeYears()

	// Loop through each year
	for _, year := range years {
		// Scrape the available weeks and week names in a given year
		weekIDs, weekNames := nflscores.ScrapeWeeks(year)

		// Loop through each week in a given year
		for index := range weekIDs {
			// Scrape all game data from a given week in a given year
			newGames := nflscores.ScrapeGames(year, weekIDs[index], weekNames[index])

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

} // main()

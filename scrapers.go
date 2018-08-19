package nflscores

import (
	"fmt"

	"github.com/gocolly/colly"
)

// ScrapeYears returns a list of all Football seasons available.
func ScrapeYears() []string {
	var years []string

	c := colly.NewCollector(
		colly.AllowedDomains("www.nfl.com"),
	)

	c.OnHTML(".year-selector a[href]", func(e *colly.HTMLElement) {
		years = append(years, e.Text)
	})

	// Start scraping on https://hackerspaces.org
	c.Visit("https://www.nfl.com/scores/")

	return years
}

// ScrapeWeeks returns a map of week IDs to human names for a given year.
func ScrapeWeeks(year string) ([]string, []string) {
	var weekIDs, weekNames []string

	c := colly.NewCollector(
		colly.AllowedDomains("www.nfl.com"),
	)

	c.OnHTML(".pre-season-games a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// Removes /scores/YYYY/ from the link text and stores the result:
		weekIDs = append(weekIDs, link[13:])
		weekNames = append(weekNames, fmt.Sprintf("Pre Season %s", e.Text))
	})

	c.OnHTML(".reg-season-games a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// Removes /scores/YYYY/ from the link text and stores the result:
		weekIDs = append(weekIDs, link[13:])
		weekNames = append(weekNames, fmt.Sprintf("Regular Season %s", e.Text))
	})

	c.OnHTML(".post-season-games a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// Removes /scores/YYYY/ from the link text and stores the result:
		weekIDs = append(weekIDs, link[13:])
		weekNames = append(weekNames, fmt.Sprintf("Pre Season %s", e.Text))
	})

	// Start scraping
	c.Visit(fmt.Sprintf("https://www.nfl.com/scores/%s/PRE0", year))

	return weekIDs, weekNames
}

// ScrapeGames returns an array of all games from a given week.
func ScrapeGames(year string, weekID string, weekName string) []Game {

	var games []Game

	c := colly.NewCollector(
		colly.AllowedDomains("www.nfl.com"),
	)

	c.OnHTML(".scorebox-wrapper", func(e *colly.HTMLElement) {
		var newGame Game

		newGame.date = e.ChildText(".date")
		newGame.weekID = weekID
		newGame.weekName = weekName
		newGame.year = year

		newGame.homeName = e.ChildText(".home-team .team-data .team-info .team-name a[href]")
		newGame.homeQ1 = e.ChildText(".home-team .team-data .quarters-score .first-qt")
		newGame.homeQ2 = e.ChildText(".home-team .team-data .quarters-score .second-qt")
		newGame.homeQ3 = e.ChildText(".home-team .team-data .quarters-score .third-qt")
		newGame.homeQ4 = e.ChildText(".home-team .team-data .quarters-score .fourth-qt")
		newGame.homeOT = e.ChildText(".home-team .team-data .quarters-score .ot-qt")
		newGame.homeFinal = e.ChildText(".home-team .team-data .total-score")

		newGame.awayName = e.ChildText(".away-team .team-data .team-info .team-name a[href]")
		newGame.awayQ1 = e.ChildText(".away-team .team-data .quarters-score .first-qt")
		newGame.awayQ2 = e.ChildText(".away-team .team-data .quarters-score .second-qt")
		newGame.awayQ3 = e.ChildText(".away-team .team-data .quarters-score .third-qt")
		newGame.awayQ4 = e.ChildText(".away-team .team-data .quarters-score .fourth-qt")
		newGame.awayOT = e.ChildText(".away-team .team-data .quarters-score .ot-qt")
		newGame.awayFinal = e.ChildText(".away-team .team-data .total-score")

		games = append(games, newGame)

	})

	c.Visit(fmt.Sprintf("https://www.nfl.com/scores/%s/%s", year, weekID))

	return games
}

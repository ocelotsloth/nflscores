package nflscores

// Game stores all the information about a game.
type Game struct {
	date      string `csv:"date"`
	weekID    string `csv:"weekID"`
	weekName  string `csv:"weekName"`
	year      string `csv:"year"`
	homeName  string `csv:"homeName"`
	homeQ1    string `csv:"homeQ1"`
	homeQ2    string `csv:"homeQ2"`
	homeQ3    string `csv:"homeQ3"`
	homeQ4    string `csv:"homeQ4"`
	homeOT    string `csv:"homeOT"`
	homeFinal string `csv:"homeFinal"`
	awayName  string `csv:"awayName"`
	awayQ1    string `csv:"awayQ1"`
	awayQ2    string `csv:"awayQ2"`
	awayQ3    string `csv:"awayQ3"`
	awayQ4    string `csv:"awayQ4"`
	awayOT    string `csv:"awayOT"`
	awayFinal string `csv:"awayFinal"`
}

// ToRow returns input as row of strings
func (g Game) ToRow() []string {
	return []string{
		g.date,
		g.weekID,
		g.weekName,
		g.year,
		g.homeName,
		g.homeQ1,
		g.homeQ2,
		g.homeQ3,
		g.homeQ4,
		g.homeOT,
		g.homeFinal,
		g.awayName,
		g.awayQ1,
		g.awayQ2,
		g.awayQ3,
		g.awayQ4,
		g.awayOT,
		g.awayFinal,
	}
}

// GetGameHeaders returns a row of string headers
func GetGameHeaders() []string {
	return []string{
		"date",
		"weekID",
		"weekName",
		"year",
		"homeName",
		"homeQ1",
		"homeQ2",
		"homeQ3",
		"homeQ4",
		"homeOT",
		"homeFinal",
		"awayName",
		"awayQ1",
		"awayQ2",
		"awayQ3",
		"awayQ4",
		"awayOT",
		"awayFinal",
	}
}

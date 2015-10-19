package theme

type Theme struct {
	Name        string
	Description string
	Homepage    string
	Tags        []string
	License     string
	MinVersion  string

	Author struct {
		Name     string
		Homepage string
	}
}

// Package spiderstub is a stub for spider package
// Implements same interface, but returns hardcoded values
package spiderstub

// Bot - implements Scanner interface
type Bot struct{}

// New - creates and returns a new spider bot
func New() *Bot {
	c := Bot{}
	return &c
}

// Scan осуществляет рекурсивный обход ссылок сайта, указанного в URL,
// с учётом глубины перехода по ссылкам, переданной в depth.
func (s *Bot) Scan(url string) (data map[string]string, err error) {
	stubdata := map[string]string{
		"https://katata.games":          "Katata Games | A family game studio",
		"https://katata.games/contact/": "Contact Katata | Katata Games",
		"https://katata.games/top-run":  "Top Run | Katata Games",
	}

	return stubdata, nil
}

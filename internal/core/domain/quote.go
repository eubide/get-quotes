package domain

// Quote represents a quote entity in the domain
type Quote struct {
	Text string
}

// NewQuote creates a new Quote instance
func NewQuote(text string) *Quote {
	return &Quote{
		Text: text,
	}
}

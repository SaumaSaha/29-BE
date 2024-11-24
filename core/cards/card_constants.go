package cards

//go:generate enumer -type Suit,Value -output card_constants_generated.go -linecomment -json
type Suit uint

const (
	_        Suit = iota
	Hearts         // Hearts
	Diamonds       // Diamonds
	Clubs          // Clubs
	Spades         // Spades
)

type Value uint

const (
	_     Value = iota
	Ace          // Ace
	Two          // Two
	Three        // Three
	Four         // Four
	Five         // Five
	Six          // Six
	Seven        // Seven
	Eight        // Eight
	Nine         // Nine
	Ten          // Ten
	Jack         // Jack
	Queen        // Queen
	King         // King
)

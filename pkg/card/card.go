package card

var FamilySpade = "spade"
var FamilyHeart = "heart"
var FamilyDiamond = "diamond"
var FamilyClub = "club"

var CardValueAce = "A"
var CardValueTwo = "2"
var CardValueThree = "3"
var CardValueFour = "4"
var CardValueFive = "5"
var CardValueSix = "6"
var CardValueSeven = "7"
var CardValueEight = "8"
var CardValueNine = "9"
var CardValueTen = "10"
var CardValueJack = "J"
var CardValueQueen = "Q"
var CardValueKing = "K"

type Card struct {
    Image string
    Family string
    Value string
}

var deck []Card = []Card {
    {Image: "🂡", Family: FamilySpade, Value: CardValueAce},
    {Image: "🂢", Family: FamilySpade, Value: CardValueTwo},
    {Image: "🂣", Family: FamilySpade, Value: CardValueThree},
    {Image: "🂤", Family: FamilySpade, Value: CardValueFour},
    {Image: "🂥", Family: FamilySpade, Value: CardValueFive},
    {Image: "🂦", Family: FamilySpade, Value: CardValueSix},
    {Image: "🂧", Family: FamilySpade, Value: CardValueSeven},
    {Image: "🂨", Family: FamilySpade, Value: CardValueEight},
    {Image: "🂩", Family: FamilySpade, Value: CardValueNine},
    {Image: "🂪", Family: FamilySpade, Value: CardValueTen},
    {Image: "🂫", Family: FamilySpade, Value: CardValueJack},
    {Image: "🂭", Family: FamilySpade, Value: CardValueQueen},
    {Image: "🂮", Family: FamilySpade, Value: CardValueKing},

    {Image: "🂱", Family: FamilyHeart, Value: CardValueAce},
    {Image: "🂲", Family: FamilyHeart, Value: CardValueTwo},
    {Image: "🂳", Family: FamilyHeart, Value: CardValueThree},
    {Image: "🂴", Family: FamilyHeart, Value: CardValueFour},
    {Image: "🂵", Family: FamilyHeart, Value: CardValueFive},
    {Image: "🂶", Family: FamilyHeart, Value: CardValueSix},
    {Image: "🂷", Family: FamilyHeart, Value: CardValueSeven},
    {Image: "🂸", Family: FamilyHeart, Value: CardValueEight},
    {Image: "🂹", Family: FamilyHeart, Value: CardValueNine},
    {Image: "🂺", Family: FamilyHeart, Value: CardValueTen},
    {Image: "🂻", Family: FamilyHeart, Value: CardValueJack},
    {Image: "🂽", Family: FamilyHeart, Value: CardValueQueen},
    {Image: "🂾", Family: FamilyHeart, Value: CardValueKing},

    {Image: "🃁", Family: FamilyDiamond, Value: CardValueAce},
    {Image: "🃂", Family: FamilyDiamond, Value: CardValueTwo},
    {Image: "🃃", Family: FamilyDiamond, Value: CardValueThree},
    {Image: "🃄", Family: FamilyDiamond, Value: CardValueFour},
    {Image: "🃅", Family: FamilyDiamond, Value: CardValueFive},
    {Image: "🃆", Family: FamilyDiamond, Value: CardValueSix},
    {Image: "🃇", Family: FamilyDiamond, Value: CardValueSeven},
    {Image: "🃈", Family: FamilyDiamond, Value: CardValueEight},
    {Image: "🃉", Family: FamilyDiamond, Value: CardValueNine},
    {Image: "🃊", Family: FamilyDiamond, Value: CardValueTen},
    {Image: "🃋", Family: FamilyDiamond, Value: CardValueJack},
    {Image: "🃍", Family: FamilyDiamond, Value: CardValueQueen},
    {Image: "🃎", Family: FamilyDiamond, Value: CardValueKing},

    {Image: "🃑", Family: FamilyClub, Value: CardValueAce},
    {Image: "🃒", Family: FamilyClub, Value: CardValueTwo},
    {Image: "🃓", Family: FamilyClub, Value: CardValueThree},
    {Image: "🃔", Family: FamilyClub, Value: CardValueFour},
    {Image: "🃕", Family: FamilyClub, Value: CardValueFive},
    {Image: "🃖", Family: FamilyClub, Value: CardValueSix},
    {Image: "🃗", Family: FamilyClub, Value: CardValueSeven},
    {Image: "🃘", Family: FamilyClub, Value: CardValueEight},
    {Image: "🃘", Family: FamilyClub, Value: CardValueNine},
    {Image: "🃘", Family: FamilyClub, Value: CardValueTen},
    {Image: "🃛", Family: FamilyClub, Value: CardValueJack},
    {Image: "🃝", Family: FamilyClub, Value: CardValueQueen},
    {Image: "🃞", Family: FamilyClub, Value: CardValueKing},
}

func NewDecks(count int) []Card {
    var cards []Card
    for i := 0; i < count; i++ {
        cards = append(cards, deck...)
    }

    return cards
}

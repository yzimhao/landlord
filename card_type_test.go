package landlord

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCardType(t *testing.T) {

	data := map[CardType][]Hand{
		danpai: []Hand{
			//黑桃3
			[]Card{
				parseCard(SuitSpade, "3"),
			},
		},
		duizi: []Hand{
			//对3
			[]Card{
				parseCard(SuitSpade, "3"),
				parseCard(SuitHeart, "3"),
			},
		},
		sanbudai: []Hand{
			//三个3
			[]Card{
				parseCard(SuitSpade, "3"),
				parseCard(SuitHeart, "3"),
				parseCard(SuitClub, "3"),
			},
		},
		sandaiyi: []Hand{
			//三个3带4
			[]Card{
				parseCard(SuitSpade, "3"),
				parseCard(SuitHeart, "3"),
				parseCard(SuitClub, "3"),
				parseCard(SuitClub, "4"),
			},
		},
		sandaier: []Hand{
			//三个3带对4
			[]Card{
				parseCard(SuitSpade, "3"),
				parseCard(SuitHeart, "3"),
				parseCard(SuitClub, "3"),
				parseCard(SuitClub, "4"),
				parseCard(SuitSpade, "4"),
			},
		},

		shunzi: []Hand{
			//34567
			[]Card{
				parseCard(SuitSpade, "3"),
				parseCard(SuitSpade, "4"),
				parseCard(SuitSpade, "5"),
				parseCard(SuitSpade, "6"),
				parseCard(SuitSpade, "7"),
			},

			//45678910jqka
			[]Card{
				parseCard(SuitSpade, "4"),
				parseCard(SuitSpade, "5"),
				parseCard(SuitSpade, "6"),
				parseCard(SuitSpade, "7"),
				parseCard(SuitSpade, "8"),
				parseCard(SuitSpade, "9"),
				parseCard(SuitSpade, "10"),
				parseCard(SuitSpade, "J"),
				parseCard(SuitSpade, "Q"),
				parseCard(SuitSpade, "K"),
				parseCard(SuitSpade, "A"),
			},
		},

		liandui: []Hand{
			[]Card{
				parseCard(SuitSpade, "4"),
				parseCard(SuitHeart, "4"),
				parseCard(SuitSpade, "5"),
				parseCard(SuitHeart, "5"),
				parseCard(SuitSpade, "6"),
				parseCard(SuitHeart, "6"),
			},
			[]Card{
				parseCard(SuitSpade, "4"),
				parseCard(SuitHeart, "4"),
				parseCard(SuitSpade, "5"),
				parseCard(SuitHeart, "5"),
				parseCard(SuitSpade, "6"),
				parseCard(SuitHeart, "6"),
				parseCard(SuitSpade, "7"),
				parseCard(SuitHeart, "7"),
			},
		},

		feiji: []Hand{
			[]Card{
				parseCard(SuitSpade, "4"),
				parseCard(SuitHeart, "4"),
				parseCard(SuitClub, "4"),
				parseCard(SuitSpade, "5"),
				parseCard(SuitHeart, "5"),
				parseCard(SuitClub, "5"),
			},
		},

		// feijidaidan: []Hand{
		// 	[]Card{
		// 		parseCard(SuitSpade, "4"),
		// 		parseCard(SuitHeart, "4"),
		// 		parseCard(SuitClub, "4"),
		// 		parseCard(SuitClub, "3"),

		// 		parseCard(SuitSpade, "5"),
		// 		parseCard(SuitHeart, "5"),
		// 		parseCard(SuitClub, "5"),
		// 		parseCard(SuitClub, "7"),
		// 	},
		// },

		// feijidaidui: []Hand{
		// 	[]Card{
		// 		parseCard(SuitSpade, "4"),
		// 		parseCard(SuitHeart, "4"),
		// 		parseCard(SuitClub, "4"),
		// 		parseCard(SuitHeart, "8"),
		// 		parseCard(SuitClub, "8"),

		// 		parseCard(SuitSpade, "5"),
		// 		parseCard(SuitHeart, "5"),
		// 		parseCard(SuitClub, "5"),

		// 		parseCard(SuitHeart, "7"),
		// 		parseCard(SuitClub, "7"),
		// 	},
		// },
	}

	for k, v := range data {
		Convey(k.String(), t, func() {
			for _, cards := range v {
				tt := CheckCardsType(cards)
				So(tt, ShouldEqual, k)
			}
		})
	}

}

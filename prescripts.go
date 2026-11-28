package main

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"math/rand/v2"
	"unicode"
)

var singles = []string{
	"Walk up to a random stranger and hand them a blank canvas",
	"Within three days, knit a scarf with a butterfly pattern",
	"See green from a white wall",
	"When your eyes meet another person’s, nod at them",
	"Clap without a sound, slash without a voice",
	"Return to your home this instant. You may leave once a dog barks in front of your house one time",
	"Neatly clip the nails of the sixty-second person you come across",
	"Do not utter a single word for five days",
	"Stand on any three-way intersection at 3:38 tomorrow, look to the east, and wave seven time.",
	fmt.Sprintf("Recite the first %d values of pie", rand.IntN( (30-20) + 20 )),
	"Initiate a game of Never Have I Ever with the first five people you encounter. When one folds a finger, break it",
	"With a sharpened blade, cry out like the scent of fallen leaves and the roar of a waterfall",
	"Breath of the other side. Folded wound. Distant sound. Open gates",
	"Cut off the right arm of every 25 year old living within a 5 mile vicinity approximate to your location by the end of the day",
	"Behead seven dwellers of the Backstreets who resist when a blade touches their throat",
}

var timeConditions = []string{
	"At precisely 14:00",
	"The next time the clock strikes 7:16",
	"Tonight",
	"This evening",
	"During the next full moon",
	"Tomorrow",
	"As soon as the sun rises",
	"As soon as the sun falls",
	"The next time it rains",
	"The next time it snows",
	"When the world comes to the end",
	"On the morning after receiving the Prescript",
}

var actions = []string {
	"drink of cup of water",
	"jump into a fire",
	"throw a baseball at someone's window",
	"go purchase a new book",
	"ask the next person you see for directions",
	"eat a bowl of chicken",
	"stare at the eyes of the next person you see for 30 minutes",
	"drink 1 gallon of milk",
	"destroy life",
	"eat a plate full of 20 cookies",
	"throw a party at your home",
	"burn your house down",
	"buy any item at any store",
	"put a coin into a piggy bank",
	"climb a tree",
	"plant a seed",
	"drink a bottle of water",
	"sit down for 10 hours",
	"walk backwards for the rest of the day",
	"cheat death",
	"cheat life",
	"walk to the nearest city with one leg",
	"cut off all your nails",
	"pat the back of the first person you see",
	"drink the life out of a painting",
	"skin alive the 5th person you see on a Tuesday",
	"drive into a lamppost",
	"drive into a lamppost or something solid to draw attention and damage the car",
	"kill 8 billion people immediately as soon as you can",
	"walk out into the middle of an intersection at 5:47 pm",
	"love the universe",
	"take on a new name",
	"Kill all the doors of the ceiling",
	"eat your vocal cords",
	"drink three cups of water",
	"spin a wheel and throw a cake at the person determined by the result",
	"perform a dance that consists of no more than 8 rotations",
}

var conditions = []string {
	"while holding a glass of water",
	"without speaking a single word",
	"without taking a single step",
}

var markers = []string{
	"then",
	"right after that",
	"and then",
	"afterwards",
	"and once you do",
	"wait 20 minutes and then",
}

var postScripts = []string{
	"Make sure to lock your door afterwords",
	"Fail this prescript on purpose",
	"Once you are done, take a shower",
	"No time limit",
	"Reflect on what you have done",
	"Disobey your next prescript",
	"Once this prescript is done, exchange this prescript with another Index member",
	"This directive is absolute",
}

type Prescript struct {
	Prescript string `json:"prescript"`
}

type NumGenerator struct {
	rng  *rand.Rand
}

func (n *NumGenerator) RandRange(start int, end int) int {
	if start > end  {
		return 0
	}

	if n.rng == nil {
		return rand.IntN((end+start) + start)
	} else {
		return n.rng.IntN((end+start) + start)
	}
}

func (n *NumGenerator) IntN(end int) int {
	if end < 0 {
		return 0
	}

	if n.rng == nil {
		return rand.IntN(end)
	} else {
		return n.rng.IntN(end)
	}
}

func CapitalizeFirst(s string) string {
	if len(s) == 0 {
		return ""
	}

	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])

	return string(runes)
}

func GeneratePrescript(seed string) Prescript {
	var randSource *rand.Rand
	if len(seed) > 0 {
		sum := sha256.Sum256([]byte(seed))	
		firstSeed := binary.BigEndian.Uint64(sum[0:16])
		secondSeed := binary.BigEndian.Uint64(sum[16:32])
		pcg := rand.NewPCG(firstSeed, secondSeed)
		randSource = rand.New(pcg)
	}

	rng := NumGenerator{randSource}

	structureID := rng.IntN(5)

	var prescriptTxt string;
	switch structureID {
	case 0:
		idx1 := rng.IntN(len((timeConditions)))
		timeCon := timeConditions[idx1]

		idx2 := rng.IntN(len(actions))
		action := actions[idx2]

		prescriptTxt = timeCon + ", " + action + "."
	case 1:
		idx1 := rng.IntN(len((actions)))
		actions1 := actions[idx1] 

		idx2 := rng.IntN(len(markers))
		marker := markers[idx2] 
		
		idx3 := rng.IntN(len(actions) - 1)
		
		if idx3 >= idx1 {
			idx3++
		}

		actions2 := actions[idx3]

		prescriptTxt =  CapitalizeFirst(actions1) + ", " + marker + ", " + actions2
	case 2:
		idx1 := rng.IntN(len(actions))
		action1 := actions[idx1]
		
		idx2 := rng.IntN(len(markers))
		marker := markers[idx2]

		idx3 := rng.IntN(len(actions) - 1)

		if idx3 >= idx1 {
			idx3++
		}

		action2 := actions[idx3]

		idx4 := rng.IntN(len(postScripts))
		postScript := postScripts[idx4]

		prescriptTxt = CapitalizeFirst(action1) + ", " + marker + ", " + action2 + ". " + postScript
	case 3:
		idx1 := rng.IntN(len(timeConditions))
		timeCon := timeConditions[idx1]

		idx2 := rng.IntN(len(actions))
		action := actions[idx2]

		idx3 := rng.IntN(len(postScripts))
		postScript := postScripts[idx3]

		prescriptTxt = timeCon + ", " + action + ". " + postScript + "."
	case 4:
		prescriptTxt = singles[rng.IntN(len(singles))] + "."
	}

	return Prescript{prescriptTxt}
}

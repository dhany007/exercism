package dndcharacter

import (
	"math"
	"math/rand"
	"sort"
)

type Character struct {
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
	Hitpoints    int
}

// Modifier calculates the ability modifier for a given ability score
func Modifier(score int) int {
	modifier := math.Floor((float64(score) - 10) / 2)
	return int(modifier)
}

// Ability uses randomness to generate the score for an ability
func Ability() int {
	ability := 0
	var scores []int

	for i := 0; i < 4; i++ {
		score := rand.Intn(6) + 1 // random 1-6
		scores = append(scores, score)
	}

	sort.Slice(scores, func(i, j int) bool {
		return i > j
	})

	for i := 1; i < len(scores); i++ {
		ability += scores[i]
	}

	return ability
}

// GenerateCharacter creates a new Character with random scores for abilities
func GenerateCharacter() Character {
	var character Character

	character.Strength = Ability()
	character.Dexterity = Ability()
	character.Constitution = Ability()
	character.Intelligence = Ability()
	character.Wisdom = Ability()
	character.Charisma = Ability()
	character.Hitpoints = 10 + Modifier(character.Constitution)

	return character
}

package main

import (
	"math/rand"

	"github.com/meyskens/wwg-welcome/gopherize"
)

var needs = map[string]bool{"Body": true, "Eyes": true}

func createForWord(word string) (string, error) {
	rand.Seed(nameToSeed(word))

	categories, _ := gopherize.MapAllCategories()

	gopher := gopherize.NewGopher()
	for _, b := range buildOrder {
		category := categories[b]

		if rand.Intn(3) != 1 && !needs[category.Name] {
			continue
		}
		image := category.Images[rand.Intn(len(category.Images))]
		gopher.SetImage(image.ID)
	}

	return gopher.GetImageURL()
}

func nameToSeed(name string) int64 {
	bytes := []byte(name)

	var i int64
	for _, b := range bytes {
		i += int64(b)
	}

	return i
}

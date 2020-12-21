package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

type Recipe struct {
	Name        string
	MustContain AllergenSet
	Ingredients IngredientSet
}

type Ingredient struct {
	Name string
}

type Allergen struct {
	Name                string
	PossibleIngredients IngredientSet
}

type Input struct {
	Ingredients map[string]*Ingredient
	Recipes map[string]*Recipe
	Allergens map[string]*Allergen
}

func NewInput() Input {
	return Input{
		Ingredients: make(map[string]*Ingredient),
		Recipes: make(map[string]*Recipe),
		Allergens: make(map[string]*Allergen),
	}
}

func (in Input) GetOrCreateIngredient(name string) *Ingredient {
	if in.Ingredients[name] == nil {
		in.Ingredients[name] = &Ingredient{
			Name: name,
		}
	}
	return in.Ingredients[name]
}

func (in Input) GetOrCreateRecipe(name string) *Recipe {
	if in.Recipes[name] == nil {
		in.Recipes[name] = &Recipe{
			Name: name,
			MustContain: make(AllergenSet),
			Ingredients: make(IngredientSet),
		}
	}
	return in.Recipes[name]
}

func (in Input) GetOrCreateAllergen(name string) *Allergen {
	if in.Allergens[name] == nil {
		in.Allergens[name] = &Allergen{
			Name:                name,
		}
	}
	return in.Allergens[name]
}


type AllergenSet map[*Allergen]struct{}

func (as AllergenSet) Copy() AllergenSet {
	c := make(AllergenSet)
	for k, v := range as {
		c[k] = v
	}
	return c
}

func (as AllergenSet) Without(allergen *Allergen) AllergenSet {
	c := as.Copy()
	delete(c, allergen)
	return c
}


type IngredientSet map[*Ingredient]struct{}

func (is IngredientSet) Contains(ingredient *Ingredient) bool {
	_, ok := is[ingredient]
	return ok
}

func (is IngredientSet) Copy() IngredientSet {
	c := make(IngredientSet)
	for k, v := range is {
		c[k] = v
	}
	return c
}

func (is IngredientSet) Without(ingredient *Ingredient) IngredientSet {
	c := is.Copy()
	delete(c, ingredient)
	return c
}

func (is IngredientSet) Intersect(other IngredientSet) IngredientSet {
	c := make(IngredientSet)
	for ingredient := range is {
		if is.Contains(ingredient) && other.Contains(ingredient) {
			c[ingredient] = struct{}{}
		}
	}
	return c
}


type IngredientAllergyMap map[*Ingredient]*Allergen

func (iam IngredientAllergyMap) Copy() IngredientAllergyMap {
	c := make(IngredientAllergyMap)
	for k, v := range iam {
		c[k] = v
	}
	return c
}

func (iam IngredientAllergyMap) With(ingredient *Ingredient, allergen *Allergen) IngredientAllergyMap {
	c := iam.Copy()
	c[ingredient] = allergen
	return c
}

func ReadInput(filename string) Input {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	input := NewInput()
	recipeName := 'A'
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line ," (contains ")

		recipe := input.GetOrCreateRecipe(fmt.Sprintf("%c", recipeName))

		ingredientNames := strings.Split(parts[0], " ")
		for _, ingredientName := range ingredientNames {
			ingredient := input.GetOrCreateIngredient(ingredientName)
			recipe.Ingredients[ingredient] = struct{}{}
		}

		allergenNames := strings.Split(strings.TrimRight(parts[1], ")"), " ")
		for _, allergenName_ := range allergenNames {
			allergenName := strings.Trim(allergenName_, " ,)")
			allergen := input.GetOrCreateAllergen(allergenName)
			recipe.MustContain[allergen] = struct{}{}
		}

		recipeName++
	}
	return input
}

func Solve(in Input, knownAllergens IngredientAllergyMap, unknownAllergens AllergenSet, unknownIngredients IngredientSet) map[*Ingredient]*Allergen {
	if len(unknownAllergens) == 0 {
		return knownAllergens
	}
	for allergen := range unknownAllergens {
		for ingredient := range unknownIngredients.Intersect(allergen.PossibleIngredients) {
			result := Solve(in, knownAllergens.With(ingredient, allergen), unknownAllergens.Without(allergen), unknownIngredients.Without(ingredient))
			if result != nil {
				return result
			}
		}
	}

	return nil
}

// Each allergen must be in the intersection of the ingredients where that allergen is listed.
func Preprocess(input Input) {
	for _, allergen := range input.Allergens {
		for _, recipe := range input.Recipes {
			if _, recipeHasThisAllergen := recipe.MustContain[allergen]; !recipeHasThisAllergen {
				continue
			}
			if allergen.PossibleIngredients == nil {
				allergen.PossibleIngredients = recipe.Ingredients.Copy()
			}
			allergen.PossibleIngredients = allergen.PossibleIngredients.Intersect(recipe.Ingredients)
		}
	}
}

func Part1(input Input) IngredientAllergyMap {
	knownAllergens := make(map[*Ingredient]*Allergen)

	unknownAllergens := make(AllergenSet)
	for _, allergen := range input.Allergens {
		unknownAllergens[allergen] = struct{}{}
	}

	unknownIngredients := make(IngredientSet)
	for _, ingredient := range input.Ingredients {
		unknownIngredients[ingredient] = struct{}{}
	}

	allergenFor := Solve(input, knownAllergens, unknownAllergens, unknownIngredients)

	num := 0
	for _, recipe := range input.Recipes {
		for ingredient := range recipe.Ingredients {
			if allergenFor[ingredient] == nil {
				num++
			}
		}
	}
	fmt.Println("Part 1:", num)
	return allergenFor
}

func Part2(allergensForIngredient IngredientAllergyMap) {
	var ingredientsSortedByAllergen []*Ingredient
	for ingredient := range allergensForIngredient {
		ingredientsSortedByAllergen = append(ingredientsSortedByAllergen, ingredient)
	}
	sort.Slice(ingredientsSortedByAllergen, func(i,j int)bool {
		return strings.Compare(allergensForIngredient[ingredientsSortedByAllergen[i]].Name, allergensForIngredient[ingredientsSortedByAllergen[j]].Name) == -1
	})
	var ingredientNames []string
	for _, ingredient := range ingredientsSortedByAllergen {
		ingredientNames = append(ingredientNames, ingredient.Name)
	}
	fmt.Println("Part 2:", strings.Join(ingredientNames, ","))
}


func main() {
	input := ReadInput("input.txt")
	Preprocess(input)
	allergensForIngredient := Part1(input)
	Part2(allergensForIngredient)
}
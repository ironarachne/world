package drink

import (
	"context"
	"fmt"
	"github.com/ironarachne/world/pkg/pattern"

	"github.com/ironarachne/world/pkg/language"
	"github.com/ironarachne/world/pkg/random"
	"github.com/ironarachne/world/pkg/resource"
)

const drinkingCultureStyleError = "failed to generate drinking culture style: %w"
const drinkingSocialRuleError = "failed to generate random drinking social rules: %w"
const toastingRuleError = "failed to generate random toasting rules: %w"

// Style is a type of drinking culture
type Style struct {
	TheToast           string          `json:"the_toast"`
	ToastingRules      []string        `json:"toasting_rules"`
	UniqueDrinkPattern pattern.Pattern `json:"unique_drink_pattern"`
	SocialRules        []string        `json:"social_rules"`
}

// Generate procedurally generates a drinking style based on a set of available resources
func Generate(ctx context.Context, lang language.Language, resources []resource.Resource) (Style, error) {
	unique, err := generateUniqueDrinkPattern(ctx, lang, resources)
	if err != nil {
		err = fmt.Errorf(drinkingCultureStyleError, err)
		return Style{}, err
	}

	socialRules, err := getSocialRules(ctx)
	if err != nil {
		err = fmt.Errorf(drinkingCultureStyleError, err)
		return Style{}, err
	}

	toastingRules, err := getToastingRules(ctx)
	if err != nil {
		err = fmt.Errorf(drinkingCultureStyleError, err)
		return Style{}, err
	}

	theToast, err := lang.NewWord(ctx)
	if err != nil {
		err = fmt.Errorf(drinkingCultureStyleError, err)
		return Style{}, err
	}

	style := Style{
		TheToast:           theToast,
		ToastingRules:      toastingRules,
		UniqueDrinkPattern: unique,
		SocialRules:        socialRules,
	}
	return style, nil
}

func getSocialRules(ctx context.Context) ([]string, error) {
	rules := []string{}

	allAges := []string{
		"children may not drink alcohol",
		"drinking alcohol is a mark of passage into adulthood",
		"children are allowed small amounts of alcohol",
	}

	age, err := random.String(ctx, allAges)
	if err != nil {
		err = fmt.Errorf(drinkingSocialRuleError, err)
		return []string{}, err
	}

	rules = append(rules, age)

	allClasses := []string{
		"people of all social statuses drink the same things",
		"there are marked differences between the drinks of different social statuses",
		"people of lower social status are not permitted to drink alcohol",
		"there are no particular rules about which social classes drink what types of alcohol",
	}

	class, err := random.String(ctx, allClasses)
	if err != nil {
		err = fmt.Errorf(drinkingSocialRuleError, err)
		return []string{}, err
	}

	rules = append(rules, class)

	allSpeeds := []string{
		"drinking slowly is frowned upon",
		"drinking slowly is considered good manners",
		"finishing your cup in one gulp is a mark of youth",
		"finishing your cup in one gulp is for celebrations only",
	}

	speed, err := random.String(ctx, allSpeeds)
	if err != nil {
		err = fmt.Errorf(drinkingSocialRuleError, err)
		return []string{}, err
	}

	rules = append(rules, speed)

	allPours := []string{
		"the young refill the cups of the old",
		"when you don't want any more, you are supposed to turn your empty cup over",
		"the host fills the cups of their guests",
		"the host never fills their own cup",
	}

	pour, err := random.String(ctx, allPours)
	if err != nil {
		err = fmt.Errorf(drinkingSocialRuleError, err)
		return []string{}, err
	}

	rules = append(rules, pour)

	return rules, nil
}

func getToastingRules(ctx context.Context) ([]string, error) {
	var rules []string

	allLead := []string{
		"the youngest leads the toast",
		"the oldest leads the toast",
		"the one who travelled the furthest leads the toast",
		"the host leads the toast",
	}

	lead, err := random.String(ctx, allLead)
	if err != nil {
		err = fmt.Errorf(toastingRuleError, err)
		return []string{}, err
	}

	rules = append(rules, lead)

	allFrequencies := []string{
		"there is rarely more than one toast in a gathering",
		"each person is expected to make a toast at a gathering",
		"there are no expectations regarding the frequency of toasts",
		"each adult is expected to make a toast at a gathering",
	}

	freq, err := random.String(ctx, allFrequencies)
	if err != nil {
		err = fmt.Errorf(toastingRuleError, err)
		return []string{}, err
	}

	rules = append(rules, freq)

	allMotions := []string{
		"all participants raise their cup",
		"the leader of the toast raises their cup",
		"everyone at the table tips their cup towards the eldest",
		"everyone making the toast spills a little of their drink as a show of cordiality",
	}

	motion, err := random.String(ctx, allMotions)
	if err != nil {
		err = fmt.Errorf(toastingRuleError, err)
		return []string{}, err
	}

	rules = append(rules, motion)

	return rules, nil
}

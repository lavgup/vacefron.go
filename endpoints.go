package vacefron

import (
	"net/url"
	"strconv"
)

func BatmanSlap(text1, text2 string, batman, robin url.URL) (interface{}, error) {
	return makeRequest("batmanslap", url.Values{
		"text1":  {text1},
		"text2":  {text2},
		"batman": {batman.String()},
		"robin":  {robin.String()},
	})
}

func CarReverse(text string) (interface{}, error) {
	return makeRequest("carreverse", url.Values{
		"text": {text},
	})
}

func DistractedBoyfriend(boyfriend, woman, girlfriend url.URL) (interface{}, error) {
	return makeRequest("distractedbf", url.Values{
		"boyfriend":  {boyfriend.String()},
		"woman":      {woman.String()},
		"girlfriend": {girlfriend.String()},
	})
}

func DockOfShame(image url.URL) (interface{}, error) {
	return makeRequest("dockofshame", url.Values{
		"user": {image.String()},
	})
}

func Drip(image url.URL) (interface{}, error) {
	return makeRequest("drip", url.Values{
		"user": {image.String()},
	})
}

func Ejected(name string, wasImposter bool, color string) (interface{}, error) {
	return makeRequest("ejected", url.Values{
		"name":     {name},
		"imposter": {strconv.FormatBool(wasImposter)},
		"crewmate": {color},
	})
}

func EmergencyMeeting(text string) (interface{}, error) {
	return makeRequest("emergencymeeting", url.Values{
		"text": {text},
	})
}

func FirstTime(image url.URL) (interface{}, error) {
	return makeRequest("firsttime", url.Values{
		"user": {image.String()},
	})
}

func Grave(image url.URL) (interface{}, error) {
	return makeRequest("grave", url.Values{
		"user": {image.String()},
	})
}

func IAmSpeed(image url.URL) (interface{}, error) {
	return makeRequest("iamspeed", url.Values{
		"user": {image.String()},
	})
}

func ICanMilkYou(face, cow url.URL) (interface{}, error) {
	return makeRequest("icanmilkyou", url.Values{
		"user1": {face.String()},
		"user2": {cow.String()},
	})
}

func Heaven(image url.URL) (interface{}, error) {
	return makeRequest("heaven", url.Values{
		"user": {image.String()},
	})
}

func Npc(text1, text2 string) (interface{}, error) {
	return makeRequest("npc", url.Values{
		"text1": {text1},
		"text2": {text2},
	})
}

func RankCard(
	username string,
	avatar,
	customBackground url.URL,
	level,
	rank,
	currentXp,
	nextLevelXp,
	previousLevelXp int,
	xpColor string,
	isBoosting,
	circleAvatar bool,
) (interface{}, error) {
	return makeRequest("rankcard", url.Values{
		"username":        {username},
		"avatar":          {avatar.String()},
		"custombg":        {customBackground.String()},
		"level":           {strconv.Itoa(level)},
		"rank":            {strconv.Itoa(rank)},
		"currentxp":       {strconv.Itoa(currentXp)},
		"nextlevelxp":     {strconv.Itoa(nextLevelXp)},
		"previouslevelxp": {strconv.Itoa(previousLevelXp)},
		"xpcolor":         {xpColor},
		"isboosting":      {strconv.FormatBool(isBoosting)},
		"circleavatar":    {strconv.FormatBool(circleAvatar)},
	})
}

func Stonks(image url.URL, notStonks bool) (interface{}, error) {
	return makeRequest("stonks", url.Values{
		"user":      {image.String()},
		"notstonks": {strconv.FormatBool(notStonks)},
	})
}

func TableFlip(image url.URL) (interface{}, error) {
	return makeRequest("tableflip", url.Values{
		"user": {image.String()},
	})
}

func Water(text string) (interface{}, error) {
	return makeRequest("water", url.Values{
		"text": {text},
	})
}

func Wide(image url.URL) (interface{}, error) {
	return makeRequest("wide", url.Values{
		"image": {image.String()},
	})
}

func Wolverine(image url.URL) (interface{}, error) {
	return makeRequest("wolverine", url.Values{
		"user": {image.String()},
	})
}

func WomanYellingAtCat(woman, cat url.URL) (interface{}, error) {
	return makeRequest("womanyellingatcat", url.Values{
		"woman": {woman.String()},
		"cat":   {cat.String()},
	})
}

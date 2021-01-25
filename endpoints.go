package vacefron

import (
	"net/url"
	"strconv"
)

// BatmanSlap returns a "batman slap" image
func BatmanSlap(text1, text2 string, batman, robin url.URL) (interface{}, error) {
	return makeRequest("batmanslap", url.Values{
		"text1":  {text1},
		"text2":  {text2},
		"batman": {batman.String()},
		"robin":  {robin.String()},
	})
}

// CarReverse returns a "car reverse" image
func CarReverse(text string) (interface{}, error) {
	return makeRequest("carreverse", url.Values{
		"text": {text},
	})
}

// DistractedBoyfriend returns a "distracted boyfriend" image
func DistractedBoyfriend(boyfriend, woman, girlfriend url.URL) (interface{}, error) {
	return makeRequest("distractedbf", url.Values{
		"boyfriend":  {boyfriend.String()},
		"woman":      {woman.String()},
		"girlfriend": {girlfriend.String()},
	})
}

// DockOfShame returns a "dock of shame" image
func DockOfShame(image url.URL) (interface{}, error) {
	return makeRequest("dockofshame", url.Values{
		"user": {image.String()},
	})
}

// Drip returns a "drip" image
func Drip(image url.URL) (interface{}, error) {
	return makeRequest("drip", url.Values{
		"user": {image.String()},
	})
}

// Ejected returns an "ejected" image
func Ejected(name string, wasImposter bool, color string) (interface{}, error) {
	return makeRequest("ejected", url.Values{
		"name":     {name},
		"imposter": {strconv.FormatBool(wasImposter)},
		"crewmate": {color},
	})
}

// EmergencyMeeting returns an "emergency meeting" image
func EmergencyMeeting(text string) (interface{}, error) {
	return makeRequest("emergencymeeting", url.Values{
		"text": {text},
	})
}

// FirstTime returns a "first time" image
func FirstTime(image url.URL) (interface{}, error) {
	return makeRequest("firsttime", url.Values{
		"user": {image.String()},
	})
}

// Grave returns a "grave" image
func Grave(image url.URL) (interface{}, error) {
	return makeRequest("grave", url.Values{
		"user": {image.String()},
	})
}

// IAmSpeed returns an "I am speed" image
func IAmSpeed(image url.URL) (interface{}, error) {
	return makeRequest("iamspeed", url.Values{
		"user": {image.String()},
	})
}

// ICanMilkYou returns an "I can milk you" image
func ICanMilkYou(face, cow url.URL) (interface{}, error) {
	return makeRequest("icanmilkyou", url.Values{
		"user1": {face.String()},
		"user2": {cow.String()},
	})
}

// Heaven returns a "heaven" image
func Heaven(image url.URL) (interface{}, error) {
	return makeRequest("heaven", url.Values{
		"user": {image.String()},
	})
}

// Npc returns a "npc" image
func Npc(text1, text2 string) (interface{}, error) {
	return makeRequest("npc", url.Values{
		"text1": {text1},
		"text2": {text2},
	})
}

// RankCard returns a "rank card" image
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

// Stonks returns a "stonks" image
func Stonks(image url.URL, notStonks bool) (interface{}, error) {
	return makeRequest("stonks", url.Values{
		"user":      {image.String()},
		"notstonks": {strconv.FormatBool(notStonks)},
	})
}

// TableFlip returns a "table flip" image
func TableFlip(image url.URL) (interface{}, error) {
	return makeRequest("tableflip", url.Values{
		"user": {image.String()},
	})
}

// Water returns a "water image"
func Water(text string) (interface{}, error) {
	return makeRequest("water", url.Values{
		"text": {text},
	})
}

// Wide returns a widened version of an image
func Wide(image url.URL) (interface{}, error) {
	return makeRequest("wide", url.Values{
		"image": {image.String()},
	})
}

// Wolverine returns a "wolverine" image
func Wolverine(image url.URL) (interface{}, error) {
	return makeRequest("wolverine", url.Values{
		"user": {image.String()},
	})
}

// WomanYellingAtCat returns a "woman yelling at cat" image
func WomanYellingAtCat(woman, cat url.URL) (interface{}, error) {
	return makeRequest("womanyellingatcat", url.Values{
		"woman": {woman.String()},
		"cat":   {cat.String()},
	})
}

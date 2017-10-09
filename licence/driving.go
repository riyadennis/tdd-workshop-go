package licence

import (
	"time"
	"strings"
	"fmt"
	"math/rand"
)

type Person struct {
	name        string
	dateOfBirth string
}
type LicenseNumberGenerator struct {
	person             Person
	licenseNumber RandomNumberGenerator
}
type RandomNumberGenerator interface {
	// returns a string as number can be 0
	createRandomNumber(numberOfDigits int) string
}

func (l LicenseNumberGenerator) createLicenceNumber(numberOfDigits int) (string) {
	return l.person.createRandomNumber(numberOfDigits)
}

func (p Person) createRandomNumber(numberOfDigits int) string {
	var fl, dob string
	var randomNum int

	fl = getFirstLettersFromName(p.name)
	dob = generateDigitsFromDate(p.dateOfBirth)
	randomNum = createRandomDigits(100, 999)

	str := fmt.Sprintf("%s%d%s", fl, randomNum, dob)
	if len(str) == numberOfDigits {
		return str
	}
	return ""
}

func getFirstLettersFromName(name string) (string) {
	names := strings.Split(name, " ")
	var f string
	for _, n := range names {
		f += string([]rune(n)[0])
	}
	return f
}

func createRandomDigits(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}
func generateDigitsFromDate(date string) (string) {
	return date
}

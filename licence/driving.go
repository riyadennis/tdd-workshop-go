package licence

import (
	"time"
	"strings"
	"fmt"
	"math/rand"
)

type Person struct {
	name        string
	dateOfBirth time.Time
}
type LicenseNumberGenerator struct {
	person        Person
	licenseNumber RandomNumberGenerator
}
type RandomNumberGenerator interface {
	// returns a string as number can be 0
	CreateRandomNumber(numberOfDigits int) string
}

func (l LicenseNumberGenerator) CreateLicenceNumber(numberOfDigits int) (string) {
	return l.person.createRandomNumber(numberOfDigits)
}

func (p Person) createRandomNumber(numberOfDigits int) string {
	var fl string
	var dob string
	var randomNum int

	fl = p.getFirstLettersFromName()
	dob = generateDigitsFromDate(p.dateOfBirth)
	randomNum = CreateRandomDigits(10000, 99999)
	str := fmt.Sprintf("%s%d%s", fl, randomNum, dob)
	if len(str) == numberOfDigits {
		return str
	}
	return ""
}

func (p Person) getFirstLettersFromName() (string) {
	names := strings.Split(p.name, " ")
	var f string
	for _, n := range names {
		f += string([]rune(n)[0])
	}
	return f
}

func CreateRandomDigits(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}
func generateDigitsFromDate(date time.Time) (string) {
	dates := strings.Split(date.String(), "-")
	return dates[0]+dates[1]
}

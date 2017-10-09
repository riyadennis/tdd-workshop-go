package licence
import (
	"testing"
	"strings"
	"time"
)
// Please implement a scruct that can produce driving licence numbers using TDD.
//
// A drving licence number it produced by taking the applicants initials, date of birth (YYYYMMDD)
// and a 3 digit random number.
//
// i.e
//     Mark David Bradley DOB: 12/05/1997 driving licence number could be MDB19970512999
//     Harry Jim James Smith DOB 09/10/1985 driving licence number could be HJJS19851009999
//     Jane Bond 01/01/2001 driving licence number could be JB20010101999
//
// You will need to create a stub test double of the RandomNumberGenerator interface to ensure
// tests always pass

func TestCreateRandomNumberContainsNameLetters(t *testing.T){
	p := new(Person)
	p.dateOfBirth = time.Date(1997, time.May, 12, 00, 00, 00, 000, time.UTC)
	p.name = "Mark David Bradley"
	l := new(LicenseNumberGenerator)
	l.person = *p

	randomString := p.getFirstLettersFromName()
	if strings.Contains(randomString,  "MDB") == false {
		t.Error("No string returned")
	}
}

func TestCreateRandomNumberTest(t *testing.T){
	p := new(Person)
	p.dateOfBirth = time.Date(1997, time.November, 9, 00, 00, 00, 000, time.UTC)
	p.name = "Mark David Bradley"

	l := new(LicenseNumberGenerator)
	l.person = *p

	randomInteger := CreateRandomDigits(10000, 99999)

	if randomInteger < 10000 && randomInteger > 99999 {
		t.Error("Invalid random number generated")
	}
}

func TestCreateStringFromDOB(t *testing.T){
	p := new(Person)
	p.dateOfBirth = time.Date(1985, time.May, 12, 00, 00, 00, 000, time.UTC)
	p.name = "Mark David Bradley"

	l := new(LicenseNumberGenerator)
	l.person = *p

	randomString := l.CreateLicenceNumber(14)

	if len(randomString) != 14 {
		t.Error("No random number attached")
	}
}
func TestCreateStringForDifferentPerson(t *testing.T){
	p := new(Person)
	p.dateOfBirth = time.Date(1997, time.May, 12, 00, 00, 00, 000, time.UTC)
	p.name = "Harry Jim James Smith"

	l := new(LicenseNumberGenerator)
	l.person = *p

	randomString := l.CreateLicenceNumber(15)

	if len(randomString) != 15 {
		t.Error("No random number attached")
	}
}
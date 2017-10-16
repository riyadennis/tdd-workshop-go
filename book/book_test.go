package book

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/pkg/errors"
)
type bookRepositoryMock struct{
	mock.Mock
}

func (br bookRepositoryMock) findByISBN(isbn string)(Book, error) {
	args := br.Called(isbn)
	return args.Get(0).(Book), args.Error(1)
}

func TestBookServiceReturnsValidJSONWhenBookFound(t *testing.T) {
	drNoJSON := `{"isbn":"0099576929","ean":"1357924680","title":"Dr No","author":"Ian Fleming","price":3.99,"year_published":1993}`
	casinoRoyaleJSON := `{"isbn":"0230037496","ean":"2468013579","title":"Casino Royale","author":"Ian Fleming","price":4.99,"year_published":1990}`

	drNo := NewBook("0099576929", "1357924680", "Dr No", "Ian Fleming", 3.99, 1993)
	casinoRoyale := NewBook("0230037496", "2468013579", "Casino Royale", "Ian Fleming", 4.99, 1990)

	bookRepo := new(bookRepositoryMock)
	bookRepo.On("findByISBN", "0099576929").Return(drNo, nil)

	bookRepo.On("findByISBN", "0230037496").Return(casinoRoyale, nil)

	bookService := BookService{bookRepo}

	assert.Equal(t, drNoJSON, bookService.getBook("0099576929"))
	assert.Equal(t, casinoRoyaleJSON, bookService.getBook("0230037496"))
}

func TestBookServiceReutrnsValidJsonWhenBookNotFound(t *testing.T) {
	errorMoonrakerJSON := `{"status":404,"code":1,"title":"Book not found","detail":"Could not find book with ISBN: 0099576872"}`
	errorGoldfingerJSON := `{"status":404,"code":1,"title":"Book not found","detail":"Could not find book with ISBN: 1784871095"}`

	bookRepo := new(bookRepositoryMock)
	bookService := BookService{bookRepo}

	bookRepo.On("findByISBN", "0099576872").Return(Book{}, errors.New("NO BOOK"))
	bookRepo.On("findByISBN", "1784871095").Return(Book{}, errors.New("NO BOOK"))

	assert.Equal(t, errorMoonrakerJSON, bookService.getBook("0099576872"))
	assert.Equal(t, errorGoldfingerJSON, bookService.getBook("1784871095"))
}

func TestBookServiceReturnValidJsonWhenBookRepositoryUnavaiable(t *testing.T) {
	errorRepositoryJSON := `{"status":500,"code":2,"title":"Repository unavailable","detail":"Service unavailable, try again later"}`

	bookRepo := new(bookRepositoryMock)
	bookService := BookService{bookRepo}

	bookRepo.On("findByISBN","009957702X").Return(Book{}, errors.New("Service unavailable, try again later"))
	assert.Equal(t, errorRepositoryJSON, bookService.getBook("009957702X"))

}

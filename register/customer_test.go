package register

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"fmt"
)

type emailSenderMock struct {
	mock.Mock
}

func (es emailSenderMock) SendEmail(toEmail, toName, subject, message string) {
	es.Called(toEmail,toName,  subject, message)
}
func TestCustomerIsEmailedWhenRegistrationIsSuccessful(t *testing.T) {
	spyEmailSender := new(emailSenderMock)
	spyEmailSender.On("SendEmail", "mark@testing.com", "Mark", "Welcome", "You have registered")
	customerRegister := NewCustomerRegister(spyEmailSender)

	customer, _ := customerRegister.registerCustomer("Mark", "Bradley", "mark@testing.com")
	spyEmailSender.AssertExpectations(t)

	assert.Equal(t, "Mark", customer.FirstName)
	assert.Equal(t, "Bradley", customer.LastName)
	assert.Equal(t, "mark@testing.com", customer.Email)

}

func TestCustomerIsNotEmailsWhenEmailAddressIsInvalid(t *testing.T) {
	spyEmailSender := new(emailSenderMock)
	customerRegister := NewCustomerRegister(spyEmailSender)

	spyEmailSender.AssertNotCalled(t, "SendEmail", "mark@testingcom", "Mark", "Welcome", "Invalid email")
	_, err := customerRegister.registerCustomer("Mark", "Bradley", "mark@testingcom")
	fmt.Print(err)
	assert.Error(t, err)
}

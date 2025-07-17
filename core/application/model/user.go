package model

import (
	"errors"
	_ "github.com/jbub/banking/iban"
	"net/mail"
	"strings"
)

type PersonalInfo struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
}

type AddressInfo struct {
	Street      string `json:"street"`
	HouseNumber string `json:"house_number"`
	ZipCode     string `json:"zip_code"`
	City        string `json:"city"`
}

type PaymentInfo struct {
	AccountOwner string `json:"account_owner"`
	IBAN         string `json:"iban"`
}

// Validate checks if all fields are valid
func (p *PersonalInfo) Validate() error {
	if strings.TrimSpace(p.Firstname) == "" {
		return errors.New("firstname is required")
	}
	if strings.TrimSpace(p.Lastname) == "" {
		return errors.New("lastname is required")
	}
	if strings.TrimSpace(p.Email) == "" {
		return errors.New("email is required")
	}

	if _, err := mail.ParseAddress(p.Email); err != nil {
		return errors.New("invalid email format")
	}
	return nil
}

func (a *AddressInfo) Validate() error {
	if strings.TrimSpace(a.Street) == "" {
		return errors.New("street is required")
	}
	if strings.TrimSpace(a.HouseNumber) == "" {
		return errors.New("house number is required")
	}
	if strings.TrimSpace(a.City) == "" {
		return errors.New("city is required")
	}
	if strings.TrimSpace(a.ZipCode) == "" {
		return errors.New("zipcode is required")
	}

	return nil
}

func (p *PaymentInfo) Validate() error {
	if strings.TrimSpace(p.AccountOwner) == "" {
		return errors.New("account owner is required")
	}
	if strings.TrimSpace(p.IBAN) == "" {
		return errors.New("iban is required")
	}

	return nil
}

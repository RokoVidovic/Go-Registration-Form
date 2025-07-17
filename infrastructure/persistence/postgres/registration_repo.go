package postgres

import (
	"Day_2_Go_Registration_Form/core/application/model"
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"io"
)

type PostgresRegistrationRepo struct {
	DB *pgx.Conn
}

func (r *PostgresRegistrationRepo) SavePersonalInfo(ctx context.Context, data model.PersonalInfo) error {
	_, err := r.DB.Exec(ctx, `
        INSERT INTO registrations_temp (first_name, last_name, email)
        VALUES ($1, $2, $3)
    `, data.Firstname, data.Lastname, data.Email)
	return err
}

func (r *PostgresRegistrationRepo) SaveAddressInfo(ctx context.Context, data model.AddressInfo) error {
	_, err := r.DB.Exec(ctx, `
        INSERT INTO registrations_temp (address_street,address_house_number,address_zip_code,address_city)
        VALUES ($1, $2, $3, $4)
    `, data.Street, data.HouseNumber, data.ZipCode, data.City)
	return err
}

func (r *PostgresRegistrationRepo) SavePaymentInfo(ctx context.Context, data model.PaymentInfo) error {
	_, err := r.DB.Exec(ctx, `
        INSERT INTO registrations_temp (payment_owner,payment_iban)
        VALUES ($1, $2)
    `, data.AccountOwner, data.IBAN)
	return err
}

func (r *PostgresRegistrationRepo) SetUUID(ctx context.Context) error {
	_, err := r.DB.Exec(ctx, `
        INSERT INTO registrations_temp (user_id)
        VALUES ($1)
    `, sessionId())
	return err
}

func sessionId() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	fmt.Println(base64.URLEncoding.EncodeToString(b))
	//return base64.URLEncoding.EncodeToString(b)
	fmt.Println(uuid.New().String())
	return uuid.New().String()
}

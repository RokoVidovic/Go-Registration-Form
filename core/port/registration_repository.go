package port

import (
	"Day_2_Go_Registration_Form/core/application/model"
	"context"
)

type RegistrationRepository interface {
	SavePersonalInfo(ctx context.Context, data model.PersonalInfo) error
	SaveAddressInfo(ctx context.Context, data model.AddressInfo) error
	SavePaymentInfo(ctx context.Context, data model.PaymentInfo) error

	SetUUID(ctx context.Context) error
}

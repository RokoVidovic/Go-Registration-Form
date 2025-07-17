package service

import (
	"Day_2_Go_Registration_Form/core/application/model"
	"Day_2_Go_Registration_Form/core/port"
	"context"
)

type RegistrationService struct {
	Repo port.RegistrationRepository
}

func (s *RegistrationService) SubmitStep1(ctx context.Context, data model.PersonalInfo) error {
	// validate input here if needed
	return s.Repo.SavePersonalInfo(ctx, data)
}

func (s *RegistrationService) SubmitStep2(ctx context.Context, data model.AddressInfo) error {
	return s.Repo.SaveAddressInfo(ctx, data)
}

func (s *RegistrationService) SubmitStep3(ctx context.Context, data model.PaymentInfo) error {
	return s.Repo.SavePaymentInfo(ctx, data)
}

func (s *RegistrationService) SetUUID(ctx context.Context) error {
	return s.Repo.SetUUID(ctx)
}

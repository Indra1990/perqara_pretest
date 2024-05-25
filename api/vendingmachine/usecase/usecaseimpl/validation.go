package usecaseimpl

import (
	"perqara_testing/api/vendingmachine/usecase"

	"github.com/go-ozzo/ozzo-validation/v4/is"

	validation "github.com/go-ozzo/ozzo-validation"
)

func itemCreateValidation(cmd usecase.VendingMachineRequestCommand) (err error) {
	errValidStruct := validation.ValidateStruct(&cmd,
		validation.Field(&cmd.Item, validation.Required, validation.Length(3, 255)),
		validation.Field(&cmd.Price, validation.Required, is.Digit),
	)

	if errValidStruct != nil {
		err = errValidStruct
		return
	}

	return
}

func itemUpdateValidation(cmd usecase.VendingMachineUpdateCommand) (err error) {
	errValidStruct := validation.ValidateStruct(&cmd,
		validation.Field(&cmd.Item, validation.Required, validation.Length(3, 255)),
		validation.Field(&cmd.Price, validation.Required, is.Digit),
	)

	if errValidStruct != nil {
		err = errValidStruct
		return
	}

	return
}

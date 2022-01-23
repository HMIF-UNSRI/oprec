package usecase

import (
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/go-playground/validator/v10"
	"oprec/domain"
	"oprec/pkg/helper"
)

type participantUsecaseImpl struct {
	repository domain.ParticipantRepository
	validate   *validator.Validate
}

func NewParticipantUsecaseImpl(repository domain.ParticipantRepository, validate *validator.Validate) *participantUsecaseImpl {
	return &participantUsecaseImpl{repository: repository, validate: validate}
}

func (usecase participantUsecaseImpl) Register(ctx context.Context, payload domain.ParticipantPayload) {
	err := usecase.validate.Struct(&payload)
	helper.PanicIfErr(err)

	// Validate class
	isFound := false
	for _, class := range domain.ParticipantClass {
		if class == payload.Class {
			isFound = true
		}
	}
	if !isFound {
		panic(domain.NewBadRequestError("class is not valid"))
	}

	// Validate division
	isDivision1Found := false
	isDivision2Found := false
	for key, values := range domain.Roles {
		for _, value := range values {
			if payload.Division1.Name == fmt.Sprintf("%s - %s", key, value) {
				isDivision1Found = true
			}
			if payload.Division2.Name == fmt.Sprintf("%s - %s", key, value) {
				isDivision2Found = true
			}
		}

		if isDivision1Found && isDivision2Found {
			break
		}
	}

	if !(isDivision1Found && isDivision2Found) {
		panic(domain.NewBadRequestError("division is not valid"))
	}

	// Search is participant already exist by nim and email
	statement, args, err := usecase.repository.SelectBuilder().
		Limit(1).Where(sq.Eq{"nim": payload.NIM, "email": payload.Email}).ToSql()
	helper.PanicIfErr(err)
	_, err = usecase.repository.FindOneBy(ctx, statement, args)
	if err == nil {
		panic(domain.NewAlreadyExistError("participant is already exist by nim or email"))
	}

	// Save data
	usecase.repository.Save(ctx, payload.FillForNewRecord())
}

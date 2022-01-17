package usecase

import (
	"context"
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
	isFound := false
	for _, class := range domain.ParticipantClass {
		if class == payload.Class {
			isFound = true
		}
	}
	if !isFound {
		panic(domain.NewBadRequestError("class is not valid"))
	}

	// Search is participant already exist by nim and email
	statement, args, err := usecase.repository.SelectBuilder().
		Limit(1).Where(sq.Eq{"nim": payload.NIM, "email": payload.Email}).ToSql()
	helper.PanicIfErr(err)
	_, err = usecase.repository.FindOneBy(ctx, statement, args)
	if err == nil {
		panic(domain.NewAlreadyExistError("nim or email is already exist"))
	}

	// Save data
	usecase.repository.Save(ctx, payload.FillForNewRecord())
}

func (usecase participantUsecaseImpl) GenerateForm(ctx context.Context, participant domain.Participant) {
	panic("implement me")
}

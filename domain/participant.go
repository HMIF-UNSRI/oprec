package domain

import (
	"context"
	sq "github.com/Masterminds/squirrel"
)

var ParticipantClass = [...]string{"Regular A", "Regular B", "Regular C", "Bilingual A", "Bilingual B"}

type Participant struct {
	ID             int
	Name           string
	NIM            string
	Force          int
	Class          string
	Domicile       string
	Address        string
	Email          string
	WhatsappNumber string
	LineID         string
}

type ParticipantPayload struct {
	Name           string `validate:"required,lt=100"`
	NIM            string `validate:"required,gt=10,lt=18"`
	Force          int    `validate:"required,number,min=2020,max=2021"`
	Class          string `validate:"required"`
	Domicile       string `validate:"required,lt=200"`
	Address        string `validate:"required,lt=700"`
	Email          string `validate:"required,email,lt=255"`
	WhatsappNumber string `validate:"required,lt=25"`
	LineID         string `validate:"required,lt=30"`
}

func (p Participant) AsPayload() ParticipantPayload {
	return ParticipantPayload{
		Name:           p.Name,
		NIM:            p.NIM,
		Force:          p.Force,
		Class:          p.Class,
		Domicile:       p.Domicile,
		Address:        p.Address,
		Email:          p.Email,
		WhatsappNumber: p.WhatsappNumber,
		LineID:         p.LineID,
	}
}

func (p ParticipantPayload) FillForNewRecord() Participant {
	return Participant{
		Name:           p.Name,
		NIM:            p.NIM,
		Force:          p.Force,
		Class:          p.Class,
		Domicile:       p.Domicile,
		Address:        p.Address,
		Email:          p.Email,
		WhatsappNumber: p.WhatsappNumber,
		LineID:         p.LineID,
	}
}

type ParticipantRepository interface {
	Save(ctx context.Context, participant Participant)
	FindOneBy(ctx context.Context, statement string, args []interface{}) (Participant, error)
	FindAllBy(ctx context.Context, statement string, args []interface{}) []Participant
	SelectBuilder() sq.SelectBuilder
}

type ParticipantUsecase interface {
	Register(ctx context.Context, payload ParticipantPayload)
	GenerateForm(ctx context.Context, participant Participant)
}

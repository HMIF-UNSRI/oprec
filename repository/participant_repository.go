package repository

import (
	"context"
	"database/sql"
	"errors"
	sq "github.com/Masterminds/squirrel"
	"oprec/domain"
	"oprec/pkg/helper"
)



type participantRepositoryImpl struct {
	db *sql.DB
}

func NewParticipantRepositoryImpl(db *sql.DB) *participantRepositoryImpl {
	return &participantRepositoryImpl{db: db}
}

func (repository participantRepositoryImpl) Save(ctx context.Context, participant domain.Participant) {
	statement := "INSERT INTO participants(name, nim, `force`, class, domicile, address, email, whatsapp_number, line_id) VALUE (?, ?, ?, ?, ?, ?, ?, ?, ?);"
	_, err := repository.db.ExecContext(ctx, statement, &participant.Name, &participant.NIM, &participant.Force,
		&participant.Class, &participant.Domicile, &participant.Address, &participant.Email,
		&participant.WhatsappNumber, &participant.LineID)
	helper.PanicIfErr(err)
}

func (repository participantRepositoryImpl) FindOneBy(ctx context.Context, statement string, args []interface{}) (domain.Participant, error) {
	rows, err := repository.db.QueryContext(ctx, statement, args...)
	helper.PanicIfErr(err)
	defer rows.Close()

	participant := domain.Participant{}
	if rows.Next() {
		err := rows.Scan(&participant.ID, &participant.Name, &participant.NIM, &participant.Force,
			&participant.Class, &participant.Domicile, &participant.Address, &participant.Email,
			&participant.WhatsappNumber, &participant.LineID)
		helper.PanicIfErr(err)

		return participant, nil
	} else {
		return participant, errors.New("participant is not found")
	}
}

func (repository participantRepositoryImpl) FindAllBy(ctx context.Context, statement string, args []interface{}) []domain.Participant {
	rows, err := repository.db.QueryContext(ctx, statement, args...)
	helper.PanicIfErr(err)
	defer rows.Close()

	var participants []domain.Participant
	for rows.Next() {
		participant := domain.Participant{}

		err := rows.Scan(&participant.ID, &participant.Name, &participant.NIM, &participant.Force,
			&participant.Class, &participant.Domicile, &participant.Address, &participant.Email,
			&participant.WhatsappNumber, &participant.LineID)
		helper.PanicIfErr(err)

		participants = append(participants, participant)
	}

	return participants
}

func (repository participantRepositoryImpl) SelectBuilder() sq.SelectBuilder {
	return sq.Select("id", "name", "nim","`force`","class","domicile",
		"address","email","whatsapp_number","line_id").From("participants")
}

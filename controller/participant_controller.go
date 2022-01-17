package controller

import (
	"net/http"
	"oprec/domain"
	"oprec/pkg/helper"
	"oprec/view"
	"strconv"
)

type participantController struct {
	participantUsecase domain.ParticipantUsecase
}

func RegisterParticipantController(mux *http.ServeMux, participantUsecase domain.ParticipantUsecase) {
	handler := &participantController{participantUsecase: participantUsecase}

	mux.HandleFunc("/register", handler.Register)
}

func (controller participantController) Register(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		writer.WriteHeader(http.StatusOK)
		view.Templates.ExecuteTemplate(writer, "participant_form.gohtml", map[string]interface{}{
			"Title": "Register Form",
			"Class": domain.ParticipantClass,
		})
	} else if request.Method == http.MethodPost {
		// Submit
		err := request.ParseForm()
		helper.PanicIfErr(err)

		force, err := strconv.Atoi(request.PostForm.Get("force"))
		if err != nil {
			panic(domain.NewBadRequestError("force is not number"))
		}

		payload := domain.ParticipantPayload{
			Name:           request.PostForm.Get("name"),
			NIM:            request.PostForm.Get("nim"),
			Force:          force,
			Class:          request.PostForm.Get("class"),
			Domicile:       request.PostForm.Get("domicile"),
			Address:        request.PostForm.Get("address"),
			Email:          request.PostForm.Get("email"),
			WhatsappNumber: request.PostForm.Get("whatsapp_number"),
			LineID:         request.PostForm.Get("line_id"),
		}
		controller.participantUsecase.Register(request.Context(), payload)
		writer.WriteHeader(http.StatusOK)
		http.Redirect(writer, request, "/register", http.StatusTemporaryRedirect)
	} else {
		//TODO: Method not allowed
	}
}

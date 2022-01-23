package controller

import (
	"github.com/google/uuid"
	"io"
	"net/http"
	"oprec/domain"
	"oprec/pkg/helper"
	"oprec/view"
	"os"
	"path/filepath"
	"strconv"
)

type participantController struct {
	participantUsecase domain.ParticipantUsecase
}

func RegisterParticipantController(mux *http.ServeMux, participantUsecase domain.ParticipantUsecase) {
	handler := &participantController{participantUsecase: participantUsecase}

	mux.HandleFunc("/daftar", handler.Register)
	mux.HandleFunc("/daftar/berhasil", handler.RegisterSuccess)
}

func (controller participantController) Register(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		writer.WriteHeader(http.StatusOK)
		view.Templates.ExecuteTemplate(writer, "participant_form.gohtml", map[string]interface{}{
			"Title": "Formulir Pendaftaran",
			"Class": domain.ParticipantClass,
			"Roles": domain.Roles,
		})
	} else if request.Method == http.MethodPost {
		// KPM
		err := request.ParseMultipartForm(domain.MaxImageSize)
		helper.PanicIfErr(err)

		// Save kpm
		file, fileHeader, err := request.FormFile("kpm")
		helper.PanicIfErr(err)
		defer file.Close()

		// Validation
		if fileHeader.Filename == "" {
			panic(domain.NewBadRequestError("kpm is empty"))
		}

		if fileHeader.Size > domain.MaxImageSize {
			panic(domain.NewBadRequestError("kpm size exceeds the capacity of 4MB"))
		}

		extension := filepath.Ext(fileHeader.Filename)
		if extension != ".jpg" && extension != ".jpeg" && extension != ".png" && extension != ".pdf" {
			panic(domain.NewBadRequestError("only jpg, jpeg, png, and pdf formats are accepted"))
		}

		// Generate new kpm name
		kpmName := uuid.NewString() + extension
		fileDestination, err := os.Create("uploads/" + kpmName)
		helper.PanicIfErr(err)

		_, err = io.Copy(fileDestination, file)
		helper.PanicIfErr(err)

		// Form text data
		force, err := strconv.Atoi(request.PostForm.Get("force"))
		if err != nil {
			panic(domain.NewBadRequestError("force is not number"))
		}

		division1 := domain.Division{
			Name:   request.PostForm.Get("division1_name"),
			Reason: request.PostForm.Get("division1_reason"),
		}
		division2 := domain.Division{
			Name:   request.PostForm.Get("division2_name"),
			Reason: request.PostForm.Get("division2_reason"),
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
			MainReason:     request.PostForm.Get("main_reason"),
			Division1:      division1,
			Division2:      division2,
			KPMFileName:    kpmName,
		}

		controller.participantUsecase.Register(request.Context(), payload)
		writer.WriteHeader(http.StatusOK)

		http.Redirect(writer, request, domain.BaseUrl+"/daftar/berhasil", http.StatusMovedPermanently)
	} else {
	}
}

func (controller participantController) RegisterSuccess(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		writer.WriteHeader(http.StatusOK)
		view.Templates.ExecuteTemplate(writer, "participant_success.gohtml", map[string]interface{}{
			"Title": "Berhasil Mendaftar",
		})
	} else {

	}
}
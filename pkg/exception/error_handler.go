package exception

import (
	"github.com/go-playground/validator/v10"
	"net/http"
	"oprec/domain"
	"oprec/view"
)

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err interface{}) {
	if alreadyExistError(writer, request, err) {
		return
	}

	if badRequestError(writer, request, err) {
		return
	}

	if validationErrors(writer, request, err) {
		return
	}

	internalServerError(writer, request, err)
}

func validationErrors(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		writer.WriteHeader(http.StatusBadRequest)
		view.Templates.ExecuteTemplate(writer, "participant_form.gohtml", map[string]interface{}{
			"Title":   "Register Form",
			"Class":   domain.ParticipantClass,
			"Roles":   domain.Roles,
			"Message": exception.Error(),
		})
		return true
	} else {
		return false
	}
}

func badRequestError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(domain.BadRequestError)
	if ok {
		writer.WriteHeader(http.StatusBadRequest)
		view.Templates.ExecuteTemplate(writer, "participant_form.gohtml", map[string]interface{}{
			"Title":   "Register Form",
			"Class":   domain.ParticipantClass,
			"Roles":   domain.Roles,
			"Message": exception.Error,
		})
		return true
	} else {
		return false
	}
}

//func notFoundError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
//	exception, ok := err.(domain.NotFoundError)
//	if ok {
//
//		return true
//	} else {
//		return false
//	}
//}

func alreadyExistError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(domain.AlreadyExistError)
	if ok {
		writer.WriteHeader(http.StatusBadRequest)
		view.Templates.ExecuteTemplate(writer, "participant_form.gohtml", map[string]interface{}{
			"Title":   "Register Form",
			"Class":   domain.ParticipantClass,
			"Roles":   domain.Roles,
			"Message": exception.Error,
		})
		return true
	} else {
		return false
	}
}

func internalServerError(writer http.ResponseWriter, request *http.Request, err interface{}) {
	view.Templates.ExecuteTemplate(writer, "template_error.gohtml", map[string]interface{}{
		"Title":   "Register Form",
		"Code":    500,
		"Message": err,
	})
}

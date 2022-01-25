package exception

import (
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"net/http"
	"oprec/domain"
	"oprec/view"
)

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err interface{}, logger *logrus.Logger) {
	if alreadyExistError(writer, request, err, logger) {
		return
	}

	if badRequestError(writer, request, err, logger) {
		return
	}

	if validationErrors(writer, request, err, logger) {
		return
	}

	internalServerError(writer, request, err, logger)
}

func validationErrors(writer http.ResponseWriter, request *http.Request, err interface{}, logger *logrus.Logger) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		view.Templates.ExecuteTemplate(writer, "template_error.gohtml", map[string]interface{}{
			"Title":   "400 - Bad Request",
			"Message": exception.Error(),
		})

		logger.WithFields(logrus.Fields{
			"status": 400,
			"method": request.Method,
			"path":   request.URL.EscapedPath(),
			"error":  exception.Error(),
		}).Errorln()
		return true
	} else {
		return false
	}
}

func badRequestError(writer http.ResponseWriter, request *http.Request, err interface{}, logger *logrus.Logger) bool {
	exception, ok := err.(domain.BadRequestError)
	if ok {
		view.Templates.ExecuteTemplate(writer, "template_error.gohtml", map[string]interface{}{
			"Title":   "400 - Bad Request",
			"Message": exception.Error,
		})
		logger.WithFields(logrus.Fields{
			"status": 400,
			"method": request.Method,
			"path":   request.URL.EscapedPath(),
			"error":  exception.Error,
		}).Errorln()
		return true
	} else {
		return false
	}
}

//func notFoundError(writer http.ResponseWriter, request *http.Request, err interface{}, logger *logrus.Logger) bool {
//	exception, ok := err.(domain.NotFoundError)
//	if ok {
//
//		return true
//	} else {
//		return false
//	}
//}

func alreadyExistError(writer http.ResponseWriter, request *http.Request, err interface{}, logger *logrus.Logger) bool {
	exception, ok := err.(domain.AlreadyExistError)
	if ok {
		view.Templates.ExecuteTemplate(writer, "template_error.gohtml", map[string]interface{}{
			"Title":   "400 - Bad Request",
			"Message": exception.Error,
		})
		logger.WithFields(logrus.Fields{
			"status": 400,
			"method": request.Method,
			"path":   request.URL.EscapedPath(),
			"error":  exception.Error,
		}).Errorln()
		return true
	} else {
		return false
	}
}

func internalServerError(writer http.ResponseWriter, request *http.Request, err interface{}, logger *logrus.Logger) {
	view.Templates.ExecuteTemplate(writer, "template_error.gohtml", map[string]interface{}{
		"Title":   "500 - Internal Server Error",
		"Message": err,
	})
	logger.WithFields(logrus.Fields{
		"status": 500,
		"method": request.Method,
		"path":   request.URL.EscapedPath(),
		"error":  err,
	}).Errorln()
}

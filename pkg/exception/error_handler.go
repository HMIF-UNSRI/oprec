package exception

//
//func ErrorHandler(writer http.ResponseWriter, request *http.Request, err interface{}) {
//
//	if notFoundError(writer, request, err) {
//		return
//	}
//
//	if alreadyExistError(writer, request, err) {
//		return
//	}
//
//	if badRequestError(writer, request, err) {
//		return
//	}
//
//	if validationErrors(writer, request, err) {
//		return
//	}
//
//	internalServerError(writer, request, err)
//}
//
//func validationErrors(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
//	exception, ok := err.(validator.ValidationErrors)
//	if ok {
//
//	} else {
//		return false
//	}
//}
//
//func badRequestError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
//	exception, ok := err.(domain.BadRequestError)
//	if ok {
//
//		return true
//	} else {
//		return false
//	}
//}
//
//func notFoundError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
//	exception, ok := err.(domain.NotFoundError)
//	if ok {
//
//		return true
//	} else {
//		return false
//	}
//}
//
//func alreadyExistError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
//	exception, ok := err.(domain.AlreadyExistError)
//	if ok {
//		return true
//	} else {
//		return false
//	}
//}
//
//func internalServerError(writer http.ResponseWriter, request *http.Request, err interface{}) {
//	writer.Header().Set("Content-Type", "application/json")
//	writer.WriteHeader(http.StatusInternalServerError)
//}

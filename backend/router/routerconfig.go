package router

import (
	"net/http"
)

const (
	//ContentTypenHeader is the http header key for Content-Type
	ContentTypenHeader string = "Content-Type"
	//ContentTypenHeaderValue is the http header value for Content-Type
	ContentTypenHeaderValue string = "application/json"
	//AccessControlAllowOriginHeader is the http header key for Access-Control-Allow-Origin
	AccessControlAllowOriginHeader string = "Access-Control-Allow-Origin"
	//AccessControlAllowOriginHeaderValue is the http header value for Access-Control-Allow-Origin
	AccessControlAllowOriginHeaderValue string = "*"
)

//SetHtppWriterHeaders set the defualt http writer headers.
func SetHtppWriterHeaders(writer http.ResponseWriter) http.ResponseWriter {
	writer.Header().Set(ContentTypenHeader, ContentTypenHeaderValue)
	writer.Header().Set(AccessControlAllowOriginHeader, AccessControlAllowOriginHeaderValue)
	return writer
}

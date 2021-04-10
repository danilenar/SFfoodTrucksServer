// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// GetFoodtrucksAndLocationLocationOKCode is the HTTP code returned for type GetFoodtrucksAndLocationLocationOK
const GetFoodtrucksAndLocationLocationOKCode int = 200

/*GetFoodtrucksAndLocationLocationOK OK

swagger:response getFoodtrucksAndLocationLocationOK
*/
type GetFoodtrucksAndLocationLocationOK struct {
}

// NewGetFoodtrucksAndLocationLocationOK creates GetFoodtrucksAndLocationLocationOK with default headers values
func NewGetFoodtrucksAndLocationLocationOK() *GetFoodtrucksAndLocationLocationOK {

	return &GetFoodtrucksAndLocationLocationOK{}
}

// WriteResponse to the client
func (o *GetFoodtrucksAndLocationLocationOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}
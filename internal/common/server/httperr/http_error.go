package errors

import (
	"net/http"

	"encoding/json"
	"errors"
	"github.com/Avielyo10/edge-api/internal/edge/domain/image"
	"github.com/go-chi/render"
	"gorm.io/gorm"
)

// Error gets a error string from an APIError
func (e APIError) Error() string        { return e.message }
func (e APIError) Code() int            { return e.code }
func (e *APIError) SetMessage(t string) { e.message = t }

type APIError struct {
	message string
	code    int
}

// NewInternalServerError creates a new InternalServerError
func NewInternalServerError() APIError {
	return APIError{
		message: "Internal Server Error",
		code:    http.StatusInternalServerError,
	}
}

// NewBadRequest creates a new BadRequest
func NewBadRequest(message string) APIError {
	return APIError{
		message: errors.New("Bad Request: " + message).Error(),
		code:    http.StatusBadRequest,
	}
}

// NewNotFound creates a new NotFound
func NewNotFound(message string) APIError {
	return APIError{
		message: errors.New("Not Found: " + message).Error(),
		code:    http.StatusNotFound,
	}
}

// HandleImageErrors handles errors from the image domain
func HandleImageErrors(w http.ResponseWriter, r *http.Request, err error) {
	switch err {
	case image.ErrImageNotFound:
		render.Status(r, NewNotFound(err.Error()).Code())
		render.JSON(w, r, err)
	case image.ErrAlreadyBuilding, image.ErrEmptyContext,
		image.ErrInvalidStatus, image.ErrInvalidVersion, image.ErrInvalidUser,
		image.ErrInvalidOutputType:
		render.Status(r, NewBadRequest(err.Error()).Code())
		render.JSON(w, r, NewBadRequest(err.Error()))
	case gorm.ErrRecordNotFound:
		render.Status(r, NewNotFound(err.Error()).Code())
		render.JSON(w, r, NewNotFound(err.Error()))
	default:
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, NewInternalServerError())
	}
}

// UnmarshalJSON is a custom unmarshaller for APIError
func (e *APIError) UnmarshalJSON(data []byte) error {
	var err struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}
	if err := json.Unmarshal(data, &err); err != nil {
		return err
	}
	e.message = err.Message
	e.code = err.Code
	return nil
}

// MarshalJSON is a custom marshaller for APIError
func (e APIError) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}{
		Code:    e.Code(),
		Message: e.Error(),
	})
}

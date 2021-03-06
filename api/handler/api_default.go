/*
 * go-architecture-sample
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package handler

import (
	"encoding/json"
	"net/http"
	"strings"
)

// DefaultApiController binds http requests to an api service and writes the service results to the http response
type DefaultApiController struct {
	service      DefaultApiServicer
	errorHandler ErrorHandler
}

// DefaultApiOption for how the controller is set up.
type DefaultApiOption func(*DefaultApiController)

// WithDefaultApiErrorHandler inject ErrorHandler into controller
func WithDefaultApiErrorHandler(h ErrorHandler) DefaultApiOption {
	return func(c *DefaultApiController) {
		c.errorHandler = h
	}
}

// NewDefaultApiController creates a default api controller
func NewDefaultApiController(s DefaultApiServicer, opts ...DefaultApiOption) Router {
	controller := &DefaultApiController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all of the api route for the DefaultApiController
func (c *DefaultApiController) Routes() Routes {
	return Routes{
		{
			"UserDelete",
			strings.ToUpper("Delete"),
			"/user",
			c.UserDelete,
		},
		{
			"UserListPost",
			strings.ToUpper("Post"),
			"/user/list",
			c.UserListPost,
		},
		{
			"UserPost",
			strings.ToUpper("Post"),
			"/user",
			c.UserPost,
		},
	}
}

// UserDelete -
func (c *DefaultApiController) UserDelete(w http.ResponseWriter, r *http.Request) {
	inlineObject1Param := InlineObject1{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&inlineObject1Param); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertInlineObject1Required(inlineObject1Param); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.UserDelete(r.Context(), inlineObject1Param)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// UserListPost -
func (c *DefaultApiController) UserListPost(w http.ResponseWriter, r *http.Request) {
	inlineObject2Param := InlineObject2{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&inlineObject2Param); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertInlineObject2Required(inlineObject2Param); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.UserListPost(r.Context(), inlineObject2Param)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// UserPost -
func (c *DefaultApiController) UserPost(w http.ResponseWriter, r *http.Request) {
	inlineObjectParam := InlineObject{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&inlineObjectParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertInlineObjectRequired(inlineObjectParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.UserPost(r.Context(), inlineObjectParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

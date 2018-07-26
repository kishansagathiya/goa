// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// cars HTTP client types
//
// Command:
// $ goa gen goa.design/goa/examples/streaming/design -o
// $(GOPATH)/src/goa.design/goa/examples/streaming

package client

import (
	goa "goa.design/goa"
	carssvc "goa.design/goa/examples/streaming/gen/cars"
	carssvcviews "goa.design/goa/examples/streaming/gen/cars/views"
)

// AddStreamingBody is the type of the "cars" service "add" endpoint HTTP
// request body.
type AddStreamingBody struct {
	// Car to add.
	Car *CarStreamingBody `form:"car,omitempty" json:"car,omitempty" xml:"car,omitempty"`
}

// UpdateStreamingBody is the type of the "cars" service "update" endpoint HTTP
// request body.
type UpdateStreamingBody []*CarStreamingBody

// ListResponseBody is the type of the "cars" service "list" endpoint HTTP
// response body.
type ListResponseBody struct {
	// The make of the car
	Make *string `form:"make,omitempty" json:"make,omitempty" xml:"make,omitempty"`
	// The car model
	Model *string `form:"model,omitempty" json:"model,omitempty" xml:"model,omitempty"`
	// The car body style
	BodyStyle *string `form:"body_style,omitempty" json:"body_style,omitempty" xml:"body_style,omitempty"`
}

// AddResponseBody is the type of the "cars" service "add" endpoint HTTP
// response body.
type AddResponseBody []*StoredCarResponseBody

// UpdateResponseBody is the type of the "cars" service "update" endpoint HTTP
// response body.
type UpdateResponseBody []*StoredCarResponseBody

// LoginUnauthorizedResponseBody is the type of the "cars" service "login"
// endpoint HTTP response body for the "unauthorized" error.
type LoginUnauthorizedResponseBody string

// ListInvalidScopesResponseBody is the type of the "cars" service "list"
// endpoint HTTP response body for the "invalid-scopes" error.
type ListInvalidScopesResponseBody string

// ListUnauthorizedResponseBody is the type of the "cars" service "list"
// endpoint HTTP response body for the "unauthorized" error.
type ListUnauthorizedResponseBody string

// AddInvalidScopesResponseBody is the type of the "cars" service "add"
// endpoint HTTP response body for the "invalid-scopes" error.
type AddInvalidScopesResponseBody string

// AddUnauthorizedResponseBody is the type of the "cars" service "add" endpoint
// HTTP response body for the "unauthorized" error.
type AddUnauthorizedResponseBody string

// UpdateInvalidScopesResponseBody is the type of the "cars" service "update"
// endpoint HTTP response body for the "invalid-scopes" error.
type UpdateInvalidScopesResponseBody string

// UpdateUnauthorizedResponseBody is the type of the "cars" service "update"
// endpoint HTTP response body for the "unauthorized" error.
type UpdateUnauthorizedResponseBody string

// CarStreamingBody is used to define fields on request body types.
type CarStreamingBody struct {
	// The make of the car
	Make string `form:"make" json:"make" xml:"make"`
	// The car model
	Model string `form:"model" json:"model" xml:"model"`
	// The car body style
	BodyStyle string `form:"body_style" json:"body_style" xml:"body_style"`
}

// StoredCarResponseBody is used to define fields on response body types.
type StoredCarResponseBody struct {
	// The make of the car
	Make *string `form:"make,omitempty" json:"make,omitempty" xml:"make,omitempty"`
	// The car model
	Model *string `form:"model,omitempty" json:"model,omitempty" xml:"model,omitempty"`
	// The car body style
	BodyStyle *string `form:"body_style,omitempty" json:"body_style,omitempty" xml:"body_style,omitempty"`
}

// NewAddStreamingBody builds the HTTP request body from the payload of the
// "add" endpoint of the "cars" service.
func NewAddStreamingBody(p *carssvc.AddStreamingPayload) *AddStreamingBody {
	body := &AddStreamingBody{}
	if p.Car != nil {
		body.Car = marshalCarToCarStreamingBody(p.Car)
	}
	return body
}

// NewUpdateStreamingBody builds the HTTP request body from the payload of the
// "update" endpoint of the "cars" service.
func NewUpdateStreamingBody(p []*carssvc.Car) UpdateStreamingBody {
	body := make([]*CarStreamingBody, len(p))
	for i, val := range p {
		body[i] = &CarStreamingBody{
			Make:      val.Make,
			Model:     val.Model,
			BodyStyle: val.BodyStyle,
		}
	}
	return body
}

// NewLoginUnauthorized builds a cars service login endpoint unauthorized error.
func NewLoginUnauthorized(body LoginUnauthorizedResponseBody) carssvc.Unauthorized {
	v := carssvc.Unauthorized(body)
	return v
}

// NewListStoredCarOK builds a "cars" service "list" endpoint result from a
// HTTP "OK" response.
func NewListStoredCarOK(body *ListResponseBody) *carssvcviews.StoredCarView {
	v := &carssvcviews.StoredCarView{
		Make:      body.Make,
		Model:     body.Model,
		BodyStyle: body.BodyStyle,
	}
	return v
}

// NewListInvalidScopes builds a cars service list endpoint invalid-scopes
// error.
func NewListInvalidScopes(body ListInvalidScopesResponseBody) carssvc.InvalidScopes {
	v := carssvc.InvalidScopes(body)
	return v
}

// NewListUnauthorized builds a cars service list endpoint unauthorized error.
func NewListUnauthorized(body ListUnauthorizedResponseBody) carssvc.Unauthorized {
	v := carssvc.Unauthorized(body)
	return v
}

// NewAddStoredCarCollectionCreated builds a "cars" service "add" endpoint
// result from a HTTP "Created" response.
func NewAddStoredCarCollectionCreated(body AddResponseBody) carssvcviews.StoredCarCollectionView {
	v := make([]*carssvcviews.StoredCarView, len(body))
	for i, val := range body {
		v[i] = &carssvcviews.StoredCarView{
			Make:      val.Make,
			Model:     val.Model,
			BodyStyle: val.BodyStyle,
		}
	}
	return v
}

// NewAddInvalidScopes builds a cars service add endpoint invalid-scopes error.
func NewAddInvalidScopes(body AddInvalidScopesResponseBody) carssvc.InvalidScopes {
	v := carssvc.InvalidScopes(body)
	return v
}

// NewAddUnauthorized builds a cars service add endpoint unauthorized error.
func NewAddUnauthorized(body AddUnauthorizedResponseBody) carssvc.Unauthorized {
	v := carssvc.Unauthorized(body)
	return v
}

// NewUpdateStoredCarCollectionOK builds a "cars" service "update" endpoint
// result from a HTTP "OK" response.
func NewUpdateStoredCarCollectionOK(body UpdateResponseBody) carssvcviews.StoredCarCollectionView {
	v := make([]*carssvcviews.StoredCarView, len(body))
	for i, val := range body {
		v[i] = &carssvcviews.StoredCarView{
			Make:      val.Make,
			Model:     val.Model,
			BodyStyle: val.BodyStyle,
		}
	}
	return v
}

// NewUpdateInvalidScopes builds a cars service update endpoint invalid-scopes
// error.
func NewUpdateInvalidScopes(body UpdateInvalidScopesResponseBody) carssvc.InvalidScopes {
	v := carssvc.InvalidScopes(body)
	return v
}

// NewUpdateUnauthorized builds a cars service update endpoint unauthorized
// error.
func NewUpdateUnauthorized(body UpdateUnauthorizedResponseBody) carssvc.Unauthorized {
	v := carssvc.Unauthorized(body)
	return v
}

// Validate runs the validations defined on StoredCarResponseBody
func (body *StoredCarResponseBody) Validate() (err error) {
	if body.Make == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("make", "body"))
	}
	if body.Model == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("model", "body"))
	}
	if body.BodyStyle == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("body_style", "body"))
	}
	return
}

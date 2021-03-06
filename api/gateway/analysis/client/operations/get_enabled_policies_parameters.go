// Code generated by go-swagger; DO NOT EDIT.

// Panther is a scalable, powerful, cloud-native SIEM written in Golang/React.
// Copyright (C) 2020 Panther Labs Inc
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.
//

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetEnabledPoliciesParams creates a new GetEnabledPoliciesParams object
// with the default values initialized.
func NewGetEnabledPoliciesParams() *GetEnabledPoliciesParams {
	var ()
	return &GetEnabledPoliciesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetEnabledPoliciesParamsWithTimeout creates a new GetEnabledPoliciesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetEnabledPoliciesParamsWithTimeout(timeout time.Duration) *GetEnabledPoliciesParams {
	var ()
	return &GetEnabledPoliciesParams{

		timeout: timeout,
	}
}

// NewGetEnabledPoliciesParamsWithContext creates a new GetEnabledPoliciesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetEnabledPoliciesParamsWithContext(ctx context.Context) *GetEnabledPoliciesParams {
	var ()
	return &GetEnabledPoliciesParams{

		Context: ctx,
	}
}

// NewGetEnabledPoliciesParamsWithHTTPClient creates a new GetEnabledPoliciesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetEnabledPoliciesParamsWithHTTPClient(client *http.Client) *GetEnabledPoliciesParams {
	var ()
	return &GetEnabledPoliciesParams{
		HTTPClient: client,
	}
}

/*GetEnabledPoliciesParams contains all the parameters to send to the API endpoint
for the get enabled policies operation typically these are written to a http.Request
*/
type GetEnabledPoliciesParams struct {

	/*Type
	  Type of analysis logic to retrieve

	*/
	Type string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get enabled policies params
func (o *GetEnabledPoliciesParams) WithTimeout(timeout time.Duration) *GetEnabledPoliciesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get enabled policies params
func (o *GetEnabledPoliciesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get enabled policies params
func (o *GetEnabledPoliciesParams) WithContext(ctx context.Context) *GetEnabledPoliciesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get enabled policies params
func (o *GetEnabledPoliciesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get enabled policies params
func (o *GetEnabledPoliciesParams) WithHTTPClient(client *http.Client) *GetEnabledPoliciesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get enabled policies params
func (o *GetEnabledPoliciesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithType adds the typeVar to the get enabled policies params
func (o *GetEnabledPoliciesParams) WithType(typeVar string) *GetEnabledPoliciesParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the get enabled policies params
func (o *GetEnabledPoliciesParams) SetType(typeVar string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *GetEnabledPoliciesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param type
	qrType := o.Type
	qType := qrType
	if qType != "" {
		if err := r.SetQueryParam("type", qType); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

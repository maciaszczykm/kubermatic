// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteNodeForClusterLegacyParams creates a new DeleteNodeForClusterLegacyParams object
// with the default values initialized.
func NewDeleteNodeForClusterLegacyParams() *DeleteNodeForClusterLegacyParams {
	var ()
	return &DeleteNodeForClusterLegacyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteNodeForClusterLegacyParamsWithTimeout creates a new DeleteNodeForClusterLegacyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteNodeForClusterLegacyParamsWithTimeout(timeout time.Duration) *DeleteNodeForClusterLegacyParams {
	var ()
	return &DeleteNodeForClusterLegacyParams{

		timeout: timeout,
	}
}

// NewDeleteNodeForClusterLegacyParamsWithContext creates a new DeleteNodeForClusterLegacyParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteNodeForClusterLegacyParamsWithContext(ctx context.Context) *DeleteNodeForClusterLegacyParams {
	var ()
	return &DeleteNodeForClusterLegacyParams{

		Context: ctx,
	}
}

// NewDeleteNodeForClusterLegacyParamsWithHTTPClient creates a new DeleteNodeForClusterLegacyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteNodeForClusterLegacyParamsWithHTTPClient(client *http.Client) *DeleteNodeForClusterLegacyParams {
	var ()
	return &DeleteNodeForClusterLegacyParams{
		HTTPClient: client,
	}
}

/*DeleteNodeForClusterLegacyParams contains all the parameters to send to the API endpoint
for the delete node for cluster legacy operation typically these are written to a http.Request
*/
type DeleteNodeForClusterLegacyParams struct {

	/*ClusterID*/
	ClusterID string
	/*Dc*/
	Dc string
	/*NodeID*/
	NodeID string
	/*ProjectID*/
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete node for cluster legacy params
func (o *DeleteNodeForClusterLegacyParams) WithTimeout(timeout time.Duration) *DeleteNodeForClusterLegacyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete node for cluster legacy params
func (o *DeleteNodeForClusterLegacyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete node for cluster legacy params
func (o *DeleteNodeForClusterLegacyParams) WithContext(ctx context.Context) *DeleteNodeForClusterLegacyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete node for cluster legacy params
func (o *DeleteNodeForClusterLegacyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete node for cluster legacy params
func (o *DeleteNodeForClusterLegacyParams) WithHTTPClient(client *http.Client) *DeleteNodeForClusterLegacyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete node for cluster legacy params
func (o *DeleteNodeForClusterLegacyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the delete node for cluster legacy params
func (o *DeleteNodeForClusterLegacyParams) WithClusterID(clusterID string) *DeleteNodeForClusterLegacyParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the delete node for cluster legacy params
func (o *DeleteNodeForClusterLegacyParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithDc adds the dc to the delete node for cluster legacy params
func (o *DeleteNodeForClusterLegacyParams) WithDc(dc string) *DeleteNodeForClusterLegacyParams {
	o.SetDc(dc)
	return o
}

// SetDc adds the dc to the delete node for cluster legacy params
func (o *DeleteNodeForClusterLegacyParams) SetDc(dc string) {
	o.Dc = dc
}

// WithNodeID adds the nodeID to the delete node for cluster legacy params
func (o *DeleteNodeForClusterLegacyParams) WithNodeID(nodeID string) *DeleteNodeForClusterLegacyParams {
	o.SetNodeID(nodeID)
	return o
}

// SetNodeID adds the nodeId to the delete node for cluster legacy params
func (o *DeleteNodeForClusterLegacyParams) SetNodeID(nodeID string) {
	o.NodeID = nodeID
}

// WithProjectID adds the projectID to the delete node for cluster legacy params
func (o *DeleteNodeForClusterLegacyParams) WithProjectID(projectID string) *DeleteNodeForClusterLegacyParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the delete node for cluster legacy params
func (o *DeleteNodeForClusterLegacyParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteNodeForClusterLegacyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
		return err
	}

	// path param dc
	if err := r.SetPathParam("dc", o.Dc); err != nil {
		return err
	}

	// path param node_id
	if err := r.SetPathParam("node_id", o.NodeID); err != nil {
		return err
	}

	// path param project_id
	if err := r.SetPathParam("project_id", o.ProjectID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
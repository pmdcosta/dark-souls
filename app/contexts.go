// Code generated by goagen v1.2.0-dirty, DO NOT EDIT.
//
// API "souls": Application Contexts
//
// Command:
// $ goagen
// --design=github.com/pmdcosta/dark-souls/design
// --out=$(GOPATH)/src/github.com/pmdcosta/dark-souls
// --version=v1.2.0-dirty

package app

import (
	"context"
	"github.com/goadesign/goa"
	"net/http"
)

// LoadSavesContext provides the saves load action context.
type LoadSavesContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewLoadSavesContext parses the incoming request URL and body, performs validations and creates the
// context used by the saves controller load action.
func NewLoadSavesContext(ctx context.Context, r *http.Request, service *goa.Service) (*LoadSavesContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := LoadSavesContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *LoadSavesContext) OK(resp []byte) error {
	ctx.ResponseData.Header().Set("Content-Type", "text/plain")
	ctx.ResponseData.WriteHeader(200)
	_, err := ctx.ResponseData.Write(resp)
	return err
}

// SaveSavesContext provides the saves save action context.
type SaveSavesContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewSaveSavesContext parses the incoming request URL and body, performs validations and creates the
// context used by the saves controller save action.
func NewSaveSavesContext(ctx context.Context, r *http.Request, service *goa.Service) (*SaveSavesContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := SaveSavesContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *SaveSavesContext) OK(resp []byte) error {
	ctx.ResponseData.Header().Set("Content-Type", "text/plain")
	ctx.ResponseData.WriteHeader(200)
	_, err := ctx.ResponseData.Write(resp)
	return err
}

// StartSavesContext provides the saves start action context.
type StartSavesContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewStartSavesContext parses the incoming request URL and body, performs validations and creates the
// context used by the saves controller start action.
func NewStartSavesContext(ctx context.Context, r *http.Request, service *goa.Service) (*StartSavesContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := StartSavesContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *StartSavesContext) OK(resp []byte) error {
	ctx.ResponseData.Header().Set("Content-Type", "text/plain")
	ctx.ResponseData.WriteHeader(200)
	_, err := ctx.ResponseData.Write(resp)
	return err
}

// StopSavesContext provides the saves stop action context.
type StopSavesContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewStopSavesContext parses the incoming request URL and body, performs validations and creates the
// context used by the saves controller stop action.
func NewStopSavesContext(ctx context.Context, r *http.Request, service *goa.Service) (*StopSavesContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := StopSavesContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *StopSavesContext) OK(resp []byte) error {
	ctx.ResponseData.Header().Set("Content-Type", "text/plain")
	ctx.ResponseData.WriteHeader(200)
	_, err := ctx.ResponseData.Write(resp)
	return err
}

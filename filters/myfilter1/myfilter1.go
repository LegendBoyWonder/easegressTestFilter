package myfilter1

import (
	egCtx "github.com/megaease/easegress/v2/pkg/context"
	"github.com/megaease/easegress/v2/pkg/filters"
)

// import github.com/megaease/easegress/v2/pkg/logger

const (
	// resultMyFilter1Fail is the result when MyFilter1.Handle(ctx) fails. You can add more results here, please remember to update Kind.Results
	resultMyFilter1Fail = "MyFilter1Fail"
)

// Kind is the kind of the %s filter. It is used to register and create a filter instance.
var Kind = &filters.Kind{
	CreateInstance: func(spec filters.Spec) filters.Filter {
		return &MyFilter1{spec: spec.(*Spec)}
	},
	DefaultSpec: func() filters.Spec {
		return &Spec{}
	},
	Description: "MyFilter1 is a custom filter",
	Name:        "MyFilter1",
	Results:     []string{resultMyFilter1Fail},
}

// MyFilter1 is a custom filter
type MyFilter1 struct {
	spec *Spec

	/*
	   your code here
	*/
}

// Spec is the spec of MyFilter1
type Spec struct {
	filters.BaseSpec `json:",inline"`

	/*
	   your code here
	*/
}

var _ filters.Filter = (*MyFilter1)(nil)
var _ filters.Spec = (*Spec)(nil)

// Validate validates Spec
func (s *Spec) Validate() error {
	/*
	   your code here
	*/
	return nil
}

// Name returns the name of the filter.
func (m *MyFilter1) Name() string {
	return m.spec.Name()
}

// Kind returns the kind of the filter.
func (m *MyFilter1) Kind() *filters.Kind {
	return Kind
}

// Spec returns the spec of the filter.
func (m *MyFilter1) Spec() filters.Spec {
	return m.spec
}

// Init initializes the filter.
func (m *MyFilter1) Init() {
	/*
	   your code here
	*/
}

// Inherit inherits previous filter.
func (m *MyFilter1) Inherit(previousGeneration filters.Filter) {
	// Pipeline will close the previous filter automatically.
	/*
	   your code here
	*/
}

// Handle handles the request.
func (m *MyFilter1) Handle(ctx *egCtx.Context) string {
	// hint: req := ctx.GetRequest().(*httpprot.Request)
	// empty string means success
	/*
	   your code here
	*/
	return ""
}

// Status returns the status of the filter.
func (m *MyFilter1) Status() interface{} {
	/*
	   your code here
	*/
	return nil
}

// Close closes the filter.
func (m *MyFilter1) Close() {
	/*
	   your code here
	*/
}

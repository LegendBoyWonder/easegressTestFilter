/*
Code generated by egbuilder. DO NOT EDIT.
*/

package registry

import (
	"github.com/megaease/easegress/v2/pkg/filters"
	"github.com/my/repo/filters/myfilter1"
)

// import github.com/megaease/easegress/v2/pkg/logger

func init() {
	// register filter MyFilter1
	filters.Register(myfilter1.Kind)

}

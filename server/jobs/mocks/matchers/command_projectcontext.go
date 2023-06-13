// Code generated by pegomock. DO NOT EDIT.
package matchers

import (
	pegomock "github.com/petergtz/pegomock/v3"
	"reflect"

	command "github.com/runatlantis/atlantis/server/events/command"
)

func AnyCommandProjectContext() command.ProjectContext {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(command.ProjectContext))(nil)).Elem()))
	var nullValue command.ProjectContext
	return nullValue
}

func EqCommandProjectContext(value command.ProjectContext) command.ProjectContext {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue command.ProjectContext
	return nullValue
}

func NotEqCommandProjectContext(value command.ProjectContext) command.ProjectContext {
	pegomock.RegisterMatcher(&pegomock.NotEqMatcher{Value: value})
	var nullValue command.ProjectContext
	return nullValue
}

func CommandProjectContextThat(matcher pegomock.ArgumentMatcher) command.ProjectContext {
	pegomock.RegisterMatcher(matcher)
	var nullValue command.ProjectContext
	return nullValue
}

// Code generated by pegomock. DO NOT EDIT.
package matchers

import (
	pegomock "github.com/petergtz/pegomock/v3"
	"reflect"

	slack "github.com/slack-go/slack"
)

func AnyPtrToSlackAuthTestResponse() *slack.AuthTestResponse {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(*slack.AuthTestResponse))(nil)).Elem()))
	var nullValue *slack.AuthTestResponse
	return nullValue
}

func EqPtrToSlackAuthTestResponse(value *slack.AuthTestResponse) *slack.AuthTestResponse {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue *slack.AuthTestResponse
	return nullValue
}

func NotEqPtrToSlackAuthTestResponse(value *slack.AuthTestResponse) *slack.AuthTestResponse {
	pegomock.RegisterMatcher(&pegomock.NotEqMatcher{Value: value})
	var nullValue *slack.AuthTestResponse
	return nullValue
}

func PtrToSlackAuthTestResponseThat(matcher pegomock.ArgumentMatcher) *slack.AuthTestResponse {
	pegomock.RegisterMatcher(matcher)
	var nullValue *slack.AuthTestResponse
	return nullValue
}

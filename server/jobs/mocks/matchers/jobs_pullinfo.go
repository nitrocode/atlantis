// Code generated by pegomock. DO NOT EDIT.
package matchers

import (
	pegomock "github.com/petergtz/pegomock/v3"
	"reflect"

	jobs "github.com/runatlantis/atlantis/server/jobs"
)

func AnyJobsPullInfo() jobs.PullInfo {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(jobs.PullInfo))(nil)).Elem()))
	var nullValue jobs.PullInfo
	return nullValue
}

func EqJobsPullInfo(value jobs.PullInfo) jobs.PullInfo {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue jobs.PullInfo
	return nullValue
}

func NotEqJobsPullInfo(value jobs.PullInfo) jobs.PullInfo {
	pegomock.RegisterMatcher(&pegomock.NotEqMatcher{Value: value})
	var nullValue jobs.PullInfo
	return nullValue
}

func JobsPullInfoThat(matcher pegomock.ArgumentMatcher) jobs.PullInfo {
	pegomock.RegisterMatcher(matcher)
	var nullValue jobs.PullInfo
	return nullValue
}

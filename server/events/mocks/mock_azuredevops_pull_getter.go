// Code generated by pegomock. DO NOT EDIT.
// Source: github.com/runatlantis/atlantis/server/events (interfaces: AzureDevopsPullGetter)

package mocks

import (
	azuredevops "github.com/mcdafydd/go-azuredevops/azuredevops"
	pegomock "github.com/petergtz/pegomock/v3"
	models "github.com/runatlantis/atlantis/server/events/models"
	"reflect"
	"time"
)

type MockAzureDevopsPullGetter struct {
	fail func(message string, callerSkip ...int)
}

func NewMockAzureDevopsPullGetter(options ...pegomock.Option) *MockAzureDevopsPullGetter {
	mock := &MockAzureDevopsPullGetter{}
	for _, option := range options {
		option.Apply(mock)
	}
	return mock
}

func (mock *MockAzureDevopsPullGetter) SetFailHandler(fh pegomock.FailHandler) { mock.fail = fh }
func (mock *MockAzureDevopsPullGetter) FailHandler() pegomock.FailHandler      { return mock.fail }

func (mock *MockAzureDevopsPullGetter) GetPullRequest(_param0 models.Repo, _param1 int) (*azuredevops.GitPullRequest, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockAzureDevopsPullGetter().")
	}
	params := []pegomock.Param{_param0, _param1}
	result := pegomock.GetGenericMockFrom(mock).Invoke("GetPullRequest", params, []reflect.Type{reflect.TypeOf((**azuredevops.GitPullRequest)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 *azuredevops.GitPullRequest
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(*azuredevops.GitPullRequest)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockAzureDevopsPullGetter) VerifyWasCalledOnce() *VerifierMockAzureDevopsPullGetter {
	return &VerifierMockAzureDevopsPullGetter{
		mock:                   mock,
		invocationCountMatcher: pegomock.Times(1),
	}
}

func (mock *MockAzureDevopsPullGetter) VerifyWasCalled(invocationCountMatcher pegomock.InvocationCountMatcher) *VerifierMockAzureDevopsPullGetter {
	return &VerifierMockAzureDevopsPullGetter{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
	}
}

func (mock *MockAzureDevopsPullGetter) VerifyWasCalledInOrder(invocationCountMatcher pegomock.InvocationCountMatcher, inOrderContext *pegomock.InOrderContext) *VerifierMockAzureDevopsPullGetter {
	return &VerifierMockAzureDevopsPullGetter{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		inOrderContext:         inOrderContext,
	}
}

func (mock *MockAzureDevopsPullGetter) VerifyWasCalledEventually(invocationCountMatcher pegomock.InvocationCountMatcher, timeout time.Duration) *VerifierMockAzureDevopsPullGetter {
	return &VerifierMockAzureDevopsPullGetter{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		timeout:                timeout,
	}
}

type VerifierMockAzureDevopsPullGetter struct {
	mock                   *MockAzureDevopsPullGetter
	invocationCountMatcher pegomock.InvocationCountMatcher
	inOrderContext         *pegomock.InOrderContext
	timeout                time.Duration
}

func (verifier *VerifierMockAzureDevopsPullGetter) GetPullRequest(_param0 models.Repo, _param1 int) *MockAzureDevopsPullGetter_GetPullRequest_OngoingVerification {
	params := []pegomock.Param{_param0, _param1}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "GetPullRequest", params, verifier.timeout)
	return &MockAzureDevopsPullGetter_GetPullRequest_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockAzureDevopsPullGetter_GetPullRequest_OngoingVerification struct {
	mock              *MockAzureDevopsPullGetter
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockAzureDevopsPullGetter_GetPullRequest_OngoingVerification) GetCapturedArguments() (models.Repo, int) {
	_param0, _param1 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1], _param1[len(_param1)-1]
}

func (c *MockAzureDevopsPullGetter_GetPullRequest_OngoingVerification) GetAllCapturedArguments() (_param0 []models.Repo, _param1 []int) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]models.Repo, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(models.Repo)
		}
		_param1 = make([]int, len(c.methodInvocations))
		for u, param := range params[1] {
			_param1[u] = param.(int)
		}
	}
	return
}

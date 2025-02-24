// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package sdk

import (
	"sync"
)

// Ensure, that TaskHandlerMock does implement TaskHandler.
// If this is not the case, regenerate this file with moq.
var _ TaskHandler = &TaskHandlerMock{}

// TaskHandlerMock is a mock implementation of TaskHandler.
//
// 	func TestSomethingThatUsesTaskHandler(t *testing.T) {
//
// 		// make and configure a mocked TaskHandler
// 		mockedTaskHandler := &TaskHandlerMock{
// 			ExecuteFunc: func(keptnHandle IKeptn, event KeptnEvent) (interface{}, *Error) {
// 				panic("mock out the Execute method")
// 			},
// 		}
//
// 		// use mockedTaskHandler in code that requires TaskHandler
// 		// and then make assertions.
//
// 	}
type TaskHandlerMock struct {
	// ExecuteFunc mocks the Execute method.
	ExecuteFunc func(keptnHandle IKeptn, event KeptnEvent) (interface{}, *Error)

	// calls tracks calls to the methods.
	calls struct {
		// Execute holds details about calls to the Execute method.
		Execute []struct {
			// KeptnHandle is the keptnHandle argument value.
			KeptnHandle IKeptn
			// Event is the event argument value.
			Event KeptnEvent
		}
	}
	lockExecute sync.RWMutex
}

// Execute calls ExecuteFunc.
func (mock *TaskHandlerMock) Execute(keptnHandle IKeptn, event KeptnEvent) (interface{}, *Error) {
	if mock.ExecuteFunc == nil {
		panic("TaskHandlerMock.ExecuteFunc: method is nil but TaskHandler.Execute was just called")
	}
	callInfo := struct {
		KeptnHandle IKeptn
		Event       KeptnEvent
	}{
		KeptnHandle: keptnHandle,
		Event:       event,
	}
	mock.lockExecute.Lock()
	mock.calls.Execute = append(mock.calls.Execute, callInfo)
	mock.lockExecute.Unlock()
	return mock.ExecuteFunc(keptnHandle, event)
}

// ExecuteCalls gets all the calls that were made to Execute.
// Check the length with:
//     len(mockedTaskHandler.ExecuteCalls())
func (mock *TaskHandlerMock) ExecuteCalls() []struct {
	KeptnHandle IKeptn
	Event       KeptnEvent
} {
	var calls []struct {
		KeptnHandle IKeptn
		Event       KeptnEvent
	}
	mock.lockExecute.RLock()
	calls = mock.calls.Execute
	mock.lockExecute.RUnlock()
	return calls
}

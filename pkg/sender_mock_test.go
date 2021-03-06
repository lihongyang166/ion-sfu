// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package sfu

import (
	"github.com/pion/rtp"
	"sync"
)

// Ensure, that SenderMock does implement Sender.
// If this is not the case, regenerate this file with moq.
var _ Sender = &SenderMock{}

// SenderMock is a mock implementation of Sender.
//
//     func TestSomethingThatUsesSender(t *testing.T) {
//
//         // make and configure a mocked Sender
//         mockedSender := &SenderMock{
//             CloseFunc: func()  {
// 	               panic("mock out the Close method")
//             },
//             CurrentSpatialLayerFunc: func() uint8 {
// 	               panic("mock out the CurrentSpatialLayer method")
//             },
//             IDFunc: func() string {
// 	               panic("mock out the ID method")
//             },
//             OnCloseHandlerFunc: func(fn func())  {
// 	               panic("mock out the OnCloseHandler method")
//             },
//             SwitchSpatialLayerFunc: func(layer uint8)  {
// 	               panic("mock out the SwitchSpatialLayer method")
//             },
//             SwitchTemporalLayerFunc: func(layer uint8)  {
// 	               panic("mock out the SwitchTemporalLayer method")
//             },
//             WriteRTPFunc: func(in1 *rtp.Packet)  {
// 	               panic("mock out the WriteRTP method")
//             },
//             statsFunc: func() string {
// 	               panic("mock out the stats method")
//             },
//         }
//
//         // use mockedSender in code that requires Sender
//         // and then make assertions.
//
//     }
type SenderMock struct {
	// CloseFunc mocks the Close method.
	CloseFunc func()

	// CurrentSpatialLayerFunc mocks the CurrentSpatialLayer method.
	CurrentSpatialLayerFunc func() uint8

	// IDFunc mocks the ID method.
	IDFunc func() string

	// OnCloseHandlerFunc mocks the OnCloseHandler method.
	OnCloseHandlerFunc func(fn func())

	// SwitchSpatialLayerFunc mocks the SwitchSpatialLayer method.
	SwitchSpatialLayerFunc func(layer uint8)

	// SwitchTemporalLayerFunc mocks the SwitchTemporalLayer method.
	SwitchTemporalLayerFunc func(layer uint8)

	// WriteRTPFunc mocks the WriteRTP method.
	WriteRTPFunc func(in1 *rtp.Packet)

	// statsFunc mocks the stats method.
	statsFunc func() string

	// calls tracks calls to the methods.
	calls struct {
		// Close holds details about calls to the Close method.
		Close []struct {
		}
		// CurrentSpatialLayer holds details about calls to the CurrentSpatialLayer method.
		CurrentSpatialLayer []struct {
		}
		// ID holds details about calls to the ID method.
		ID []struct {
		}
		// OnCloseHandler holds details about calls to the OnCloseHandler method.
		OnCloseHandler []struct {
			// Fn is the fn argument value.
			Fn func()
		}
		// SwitchSpatialLayer holds details about calls to the SwitchSpatialLayer method.
		SwitchSpatialLayer []struct {
			// Layer is the layer argument value.
			Layer uint8
		}
		// SwitchTemporalLayer holds details about calls to the SwitchTemporalLayer method.
		SwitchTemporalLayer []struct {
			// Layer is the layer argument value.
			Layer uint8
		}
		// WriteRTP holds details about calls to the WriteRTP method.
		WriteRTP []struct {
			// In1 is the in1 argument value.
			In1 *rtp.Packet
		}
		// stats holds details about calls to the stats method.
		stats []struct {
		}
	}
	lockClose               sync.RWMutex
	lockCurrentSpatialLayer sync.RWMutex
	lockID                  sync.RWMutex
	lockOnCloseHandler      sync.RWMutex
	lockSwitchSpatialLayer  sync.RWMutex
	lockSwitchTemporalLayer sync.RWMutex
	lockWriteRTP            sync.RWMutex
	lockstats               sync.RWMutex
}

// Close calls CloseFunc.
func (mock *SenderMock) Close() {
	if mock.CloseFunc == nil {
		panic("SenderMock.CloseFunc: method is nil but Sender.Close was just called")
	}
	callInfo := struct {
	}{}
	mock.lockClose.Lock()
	mock.calls.Close = append(mock.calls.Close, callInfo)
	mock.lockClose.Unlock()
	mock.CloseFunc()
}

// CloseCalls gets all the calls that were made to Close.
// Check the length with:
//     len(mockedSender.CloseCalls())
func (mock *SenderMock) CloseCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockClose.RLock()
	calls = mock.calls.Close
	mock.lockClose.RUnlock()
	return calls
}

// CurrentSpatialLayer calls CurrentSpatialLayerFunc.
func (mock *SenderMock) CurrentSpatialLayer() uint8 {
	if mock.CurrentSpatialLayerFunc == nil {
		panic("SenderMock.CurrentSpatialLayerFunc: method is nil but Sender.CurrentSpatialLayer was just called")
	}
	callInfo := struct {
	}{}
	mock.lockCurrentSpatialLayer.Lock()
	mock.calls.CurrentSpatialLayer = append(mock.calls.CurrentSpatialLayer, callInfo)
	mock.lockCurrentSpatialLayer.Unlock()
	return mock.CurrentSpatialLayerFunc()
}

// CurrentSpatialLayerCalls gets all the calls that were made to CurrentSpatialLayer.
// Check the length with:
//     len(mockedSender.CurrentSpatialLayerCalls())
func (mock *SenderMock) CurrentSpatialLayerCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockCurrentSpatialLayer.RLock()
	calls = mock.calls.CurrentSpatialLayer
	mock.lockCurrentSpatialLayer.RUnlock()
	return calls
}

// ID calls IDFunc.
func (mock *SenderMock) ID() string {
	if mock.IDFunc == nil {
		panic("SenderMock.IDFunc: method is nil but Sender.ID was just called")
	}
	callInfo := struct {
	}{}
	mock.lockID.Lock()
	mock.calls.ID = append(mock.calls.ID, callInfo)
	mock.lockID.Unlock()
	return mock.IDFunc()
}

// IDCalls gets all the calls that were made to ID.
// Check the length with:
//     len(mockedSender.IDCalls())
func (mock *SenderMock) IDCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockID.RLock()
	calls = mock.calls.ID
	mock.lockID.RUnlock()
	return calls
}

// OnCloseHandler calls OnCloseHandlerFunc.
func (mock *SenderMock) OnCloseHandler(fn func()) {
	if mock.OnCloseHandlerFunc == nil {
		panic("SenderMock.OnCloseHandlerFunc: method is nil but Sender.OnCloseHandler was just called")
	}
	callInfo := struct {
		Fn func()
	}{
		Fn: fn,
	}
	mock.lockOnCloseHandler.Lock()
	mock.calls.OnCloseHandler = append(mock.calls.OnCloseHandler, callInfo)
	mock.lockOnCloseHandler.Unlock()
	mock.OnCloseHandlerFunc(fn)
}

// OnCloseHandlerCalls gets all the calls that were made to OnCloseHandler.
// Check the length with:
//     len(mockedSender.OnCloseHandlerCalls())
func (mock *SenderMock) OnCloseHandlerCalls() []struct {
	Fn func()
} {
	var calls []struct {
		Fn func()
	}
	mock.lockOnCloseHandler.RLock()
	calls = mock.calls.OnCloseHandler
	mock.lockOnCloseHandler.RUnlock()
	return calls
}

// SwitchSpatialLayer calls SwitchSpatialLayerFunc.
func (mock *SenderMock) SwitchSpatialLayer(layer uint8) {
	if mock.SwitchSpatialLayerFunc == nil {
		panic("SenderMock.SwitchSpatialLayerFunc: method is nil but Sender.SwitchSpatialLayer was just called")
	}
	callInfo := struct {
		Layer uint8
	}{
		Layer: layer,
	}
	mock.lockSwitchSpatialLayer.Lock()
	mock.calls.SwitchSpatialLayer = append(mock.calls.SwitchSpatialLayer, callInfo)
	mock.lockSwitchSpatialLayer.Unlock()
	mock.SwitchSpatialLayerFunc(layer)
}

// SwitchSpatialLayerCalls gets all the calls that were made to SwitchSpatialLayer.
// Check the length with:
//     len(mockedSender.SwitchSpatialLayerCalls())
func (mock *SenderMock) SwitchSpatialLayerCalls() []struct {
	Layer uint8
} {
	var calls []struct {
		Layer uint8
	}
	mock.lockSwitchSpatialLayer.RLock()
	calls = mock.calls.SwitchSpatialLayer
	mock.lockSwitchSpatialLayer.RUnlock()
	return calls
}

// SwitchTemporalLayer calls SwitchTemporalLayerFunc.
func (mock *SenderMock) SwitchTemporalLayer(layer uint8) {
	if mock.SwitchTemporalLayerFunc == nil {
		panic("SenderMock.SwitchTemporalLayerFunc: method is nil but Sender.SwitchTemporalLayer was just called")
	}
	callInfo := struct {
		Layer uint8
	}{
		Layer: layer,
	}
	mock.lockSwitchTemporalLayer.Lock()
	mock.calls.SwitchTemporalLayer = append(mock.calls.SwitchTemporalLayer, callInfo)
	mock.lockSwitchTemporalLayer.Unlock()
	mock.SwitchTemporalLayerFunc(layer)
}

// SwitchTemporalLayerCalls gets all the calls that were made to SwitchTemporalLayer.
// Check the length with:
//     len(mockedSender.SwitchTemporalLayerCalls())
func (mock *SenderMock) SwitchTemporalLayerCalls() []struct {
	Layer uint8
} {
	var calls []struct {
		Layer uint8
	}
	mock.lockSwitchTemporalLayer.RLock()
	calls = mock.calls.SwitchTemporalLayer
	mock.lockSwitchTemporalLayer.RUnlock()
	return calls
}

// WriteRTP calls WriteRTPFunc.
func (mock *SenderMock) WriteRTP(in1 *rtp.Packet) {
	if mock.WriteRTPFunc == nil {
		panic("SenderMock.WriteRTPFunc: method is nil but Sender.WriteRTP was just called")
	}
	callInfo := struct {
		In1 *rtp.Packet
	}{
		In1: in1,
	}
	mock.lockWriteRTP.Lock()
	mock.calls.WriteRTP = append(mock.calls.WriteRTP, callInfo)
	mock.lockWriteRTP.Unlock()
	mock.WriteRTPFunc(in1)
}

// WriteRTPCalls gets all the calls that were made to WriteRTP.
// Check the length with:
//     len(mockedSender.WriteRTPCalls())
func (mock *SenderMock) WriteRTPCalls() []struct {
	In1 *rtp.Packet
} {
	var calls []struct {
		In1 *rtp.Packet
	}
	mock.lockWriteRTP.RLock()
	calls = mock.calls.WriteRTP
	mock.lockWriteRTP.RUnlock()
	return calls
}

// stats calls statsFunc.
func (mock *SenderMock) stats() string {
	if mock.statsFunc == nil {
		panic("SenderMock.statsFunc: method is nil but Sender.stats was just called")
	}
	callInfo := struct {
	}{}
	mock.lockstats.Lock()
	mock.calls.stats = append(mock.calls.stats, callInfo)
	mock.lockstats.Unlock()
	return mock.statsFunc()
}

// statsCalls gets all the calls that were made to stats.
// Check the length with:
//     len(mockedSender.statsCalls())
func (mock *SenderMock) statsCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockstats.RLock()
	calls = mock.calls.stats
	mock.lockstats.RUnlock()
	return calls
}

// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package sfu

import (
	"github.com/pion/rtcp"
	"github.com/pion/rtp"
	"github.com/pion/webrtc/v3"
	"sync"
)

// Ensure, that ReceiverMock does implement Receiver.
// If this is not the case, regenerate this file with moq.
var _ Receiver = &ReceiverMock{}

// ReceiverMock is a mock implementation of Receiver.
//
//     func TestSomethingThatUsesReceiver(t *testing.T) {
//
//         // make and configure a mocked Receiver
//         mockedReceiver := &ReceiverMock{
//             AddSenderFunc: func(sender Sender)  {
// 	               panic("mock out the AddSender method")
//             },
//             CloseFunc: func()  {
// 	               panic("mock out the Close method")
//             },
//             DeleteSenderFunc: func(pid string)  {
// 	               panic("mock out the DeleteSender method")
//             },
//             GetPacketFunc: func(sn uint16) *rtp.Packet {
// 	               panic("mock out the GetPacket method")
//             },
//             OnCloseHandlerFunc: func(fn func())  {
// 	               panic("mock out the OnCloseHandler method")
//             },
//             ReadRTCPFunc: func() chan rtcp.Packet {
// 	               panic("mock out the ReadRTCP method")
//             },
//             SpatialLayerFunc: func() uint8 {
// 	               panic("mock out the SpatialLayer method")
//             },
//             TrackFunc: func() *webrtc.Track {
// 	               panic("mock out the Track method")
//             },
//             WriteRTCPFunc: func(in1 rtcp.Packet) error {
// 	               panic("mock out the WriteRTCP method")
//             },
//             statsFunc: func() string {
// 	               panic("mock out the stats method")
//             },
//         }
//
//         // use mockedReceiver in code that requires Receiver
//         // and then make assertions.
//
//     }
type ReceiverMock struct {
	// AddSenderFunc mocks the AddSender method.
	AddSenderFunc func(sender Sender)

	// CloseFunc mocks the Close method.
	CloseFunc func()

	// DeleteSenderFunc mocks the DeleteSender method.
	DeleteSenderFunc func(pid string)

	// GetPacketFunc mocks the GetPacket method.
	GetPacketFunc func(sn uint16) *rtp.Packet

	// OnCloseHandlerFunc mocks the OnCloseHandler method.
	OnCloseHandlerFunc func(fn func())

	// ReadRTCPFunc mocks the ReadRTCP method.
	ReadRTCPFunc func() chan rtcp.Packet

	// SpatialLayerFunc mocks the SpatialLayer method.
	SpatialLayerFunc func() uint8

	// TrackFunc mocks the Track method.
	TrackFunc func() *webrtc.Track

	// WriteRTCPFunc mocks the WriteRTCP method.
	WriteRTCPFunc func(in1 rtcp.Packet) error

	// statsFunc mocks the stats method.
	statsFunc func() string

	// calls tracks calls to the methods.
	calls struct {
		// AddSender holds details about calls to the AddSender method.
		AddSender []struct {
			// Sender is the sender argument value.
			Sender Sender
		}
		// Close holds details about calls to the Close method.
		Close []struct {
		}
		// DeleteSender holds details about calls to the DeleteSender method.
		DeleteSender []struct {
			// Pid is the pid argument value.
			Pid string
		}
		// GetPacket holds details about calls to the GetPacket method.
		GetPacket []struct {
			// Sn is the sn argument value.
			Sn uint16
		}
		// OnCloseHandler holds details about calls to the OnCloseHandler method.
		OnCloseHandler []struct {
			// Fn is the fn argument value.
			Fn func()
		}
		// ReadRTCP holds details about calls to the ReadRTCP method.
		ReadRTCP []struct {
		}
		// SpatialLayer holds details about calls to the SpatialLayer method.
		SpatialLayer []struct {
		}
		// Track holds details about calls to the Track method.
		Track []struct {
		}
		// WriteRTCP holds details about calls to the WriteRTCP method.
		WriteRTCP []struct {
			// In1 is the in1 argument value.
			In1 rtcp.Packet
		}
		// stats holds details about calls to the stats method.
		stats []struct {
		}
	}
	lockAddSender      sync.RWMutex
	lockClose          sync.RWMutex
	lockDeleteSender   sync.RWMutex
	lockGetPacket      sync.RWMutex
	lockOnCloseHandler sync.RWMutex
	lockReadRTCP       sync.RWMutex
	lockSpatialLayer   sync.RWMutex
	lockTrack          sync.RWMutex
	lockWriteRTCP      sync.RWMutex
	lockstats          sync.RWMutex
}

// AddSender calls AddSenderFunc.
func (mock *ReceiverMock) AddSender(sender Sender) {
	if mock.AddSenderFunc == nil {
		panic("ReceiverMock.AddSenderFunc: method is nil but Receiver.AddSender was just called")
	}
	callInfo := struct {
		Sender Sender
	}{
		Sender: sender,
	}
	mock.lockAddSender.Lock()
	mock.calls.AddSender = append(mock.calls.AddSender, callInfo)
	mock.lockAddSender.Unlock()
	mock.AddSenderFunc(sender)
}

// AddSenderCalls gets all the calls that were made to AddSender.
// Check the length with:
//     len(mockedReceiver.AddSenderCalls())
func (mock *ReceiverMock) AddSenderCalls() []struct {
	Sender Sender
} {
	var calls []struct {
		Sender Sender
	}
	mock.lockAddSender.RLock()
	calls = mock.calls.AddSender
	mock.lockAddSender.RUnlock()
	return calls
}

// Close calls CloseFunc.
func (mock *ReceiverMock) Close() {
	if mock.CloseFunc == nil {
		panic("ReceiverMock.CloseFunc: method is nil but Receiver.Close was just called")
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
//     len(mockedReceiver.CloseCalls())
func (mock *ReceiverMock) CloseCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockClose.RLock()
	calls = mock.calls.Close
	mock.lockClose.RUnlock()
	return calls
}

// DeleteSender calls DeleteSenderFunc.
func (mock *ReceiverMock) DeleteSender(pid string) {
	if mock.DeleteSenderFunc == nil {
		panic("ReceiverMock.DeleteSenderFunc: method is nil but Receiver.DeleteSender was just called")
	}
	callInfo := struct {
		Pid string
	}{
		Pid: pid,
	}
	mock.lockDeleteSender.Lock()
	mock.calls.DeleteSender = append(mock.calls.DeleteSender, callInfo)
	mock.lockDeleteSender.Unlock()
	mock.DeleteSenderFunc(pid)
}

// DeleteSenderCalls gets all the calls that were made to DeleteSender.
// Check the length with:
//     len(mockedReceiver.DeleteSenderCalls())
func (mock *ReceiverMock) DeleteSenderCalls() []struct {
	Pid string
} {
	var calls []struct {
		Pid string
	}
	mock.lockDeleteSender.RLock()
	calls = mock.calls.DeleteSender
	mock.lockDeleteSender.RUnlock()
	return calls
}

// GetPacket calls GetPacketFunc.
func (mock *ReceiverMock) GetPacket(sn uint16) *rtp.Packet {
	if mock.GetPacketFunc == nil {
		panic("ReceiverMock.GetPacketFunc: method is nil but Receiver.GetPacket was just called")
	}
	callInfo := struct {
		Sn uint16
	}{
		Sn: sn,
	}
	mock.lockGetPacket.Lock()
	mock.calls.GetPacket = append(mock.calls.GetPacket, callInfo)
	mock.lockGetPacket.Unlock()
	return mock.GetPacketFunc(sn)
}

// GetPacketCalls gets all the calls that were made to GetPacket.
// Check the length with:
//     len(mockedReceiver.GetPacketCalls())
func (mock *ReceiverMock) GetPacketCalls() []struct {
	Sn uint16
} {
	var calls []struct {
		Sn uint16
	}
	mock.lockGetPacket.RLock()
	calls = mock.calls.GetPacket
	mock.lockGetPacket.RUnlock()
	return calls
}

// OnCloseHandler calls OnCloseHandlerFunc.
func (mock *ReceiverMock) OnCloseHandler(fn func()) {
	if mock.OnCloseHandlerFunc == nil {
		panic("ReceiverMock.OnCloseHandlerFunc: method is nil but Receiver.OnCloseHandler was just called")
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
//     len(mockedReceiver.OnCloseHandlerCalls())
func (mock *ReceiverMock) OnCloseHandlerCalls() []struct {
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

// ReadRTCP calls ReadRTCPFunc.
func (mock *ReceiverMock) ReadRTCP() chan rtcp.Packet {
	if mock.ReadRTCPFunc == nil {
		panic("ReceiverMock.ReadRTCPFunc: method is nil but Receiver.ReadRTCP was just called")
	}
	callInfo := struct {
	}{}
	mock.lockReadRTCP.Lock()
	mock.calls.ReadRTCP = append(mock.calls.ReadRTCP, callInfo)
	mock.lockReadRTCP.Unlock()
	return mock.ReadRTCPFunc()
}

// ReadRTCPCalls gets all the calls that were made to ReadRTCP.
// Check the length with:
//     len(mockedReceiver.ReadRTCPCalls())
func (mock *ReceiverMock) ReadRTCPCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockReadRTCP.RLock()
	calls = mock.calls.ReadRTCP
	mock.lockReadRTCP.RUnlock()
	return calls
}

// SpatialLayer calls SpatialLayerFunc.
func (mock *ReceiverMock) SpatialLayer() uint8 {
	if mock.SpatialLayerFunc == nil {
		panic("ReceiverMock.SpatialLayerFunc: method is nil but Receiver.SpatialLayer was just called")
	}
	callInfo := struct {
	}{}
	mock.lockSpatialLayer.Lock()
	mock.calls.SpatialLayer = append(mock.calls.SpatialLayer, callInfo)
	mock.lockSpatialLayer.Unlock()
	return mock.SpatialLayerFunc()
}

// SpatialLayerCalls gets all the calls that were made to SpatialLayer.
// Check the length with:
//     len(mockedReceiver.SpatialLayerCalls())
func (mock *ReceiverMock) SpatialLayerCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockSpatialLayer.RLock()
	calls = mock.calls.SpatialLayer
	mock.lockSpatialLayer.RUnlock()
	return calls
}

// Track calls TrackFunc.
func (mock *ReceiverMock) Track() *webrtc.Track {
	if mock.TrackFunc == nil {
		panic("ReceiverMock.TrackFunc: method is nil but Receiver.Track was just called")
	}
	callInfo := struct {
	}{}
	mock.lockTrack.Lock()
	mock.calls.Track = append(mock.calls.Track, callInfo)
	mock.lockTrack.Unlock()
	return mock.TrackFunc()
}

// TrackCalls gets all the calls that were made to Track.
// Check the length with:
//     len(mockedReceiver.TrackCalls())
func (mock *ReceiverMock) TrackCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockTrack.RLock()
	calls = mock.calls.Track
	mock.lockTrack.RUnlock()
	return calls
}

// WriteRTCP calls WriteRTCPFunc.
func (mock *ReceiverMock) WriteRTCP(in1 rtcp.Packet) error {
	if mock.WriteRTCPFunc == nil {
		panic("ReceiverMock.WriteRTCPFunc: method is nil but Receiver.WriteRTCP was just called")
	}
	callInfo := struct {
		In1 rtcp.Packet
	}{
		In1: in1,
	}
	mock.lockWriteRTCP.Lock()
	mock.calls.WriteRTCP = append(mock.calls.WriteRTCP, callInfo)
	mock.lockWriteRTCP.Unlock()
	return mock.WriteRTCPFunc(in1)
}

// WriteRTCPCalls gets all the calls that were made to WriteRTCP.
// Check the length with:
//     len(mockedReceiver.WriteRTCPCalls())
func (mock *ReceiverMock) WriteRTCPCalls() []struct {
	In1 rtcp.Packet
} {
	var calls []struct {
		In1 rtcp.Packet
	}
	mock.lockWriteRTCP.RLock()
	calls = mock.calls.WriteRTCP
	mock.lockWriteRTCP.RUnlock()
	return calls
}

// stats calls statsFunc.
func (mock *ReceiverMock) stats() string {
	if mock.statsFunc == nil {
		panic("ReceiverMock.statsFunc: method is nil but Receiver.stats was just called")
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
//     len(mockedReceiver.statsCalls())
func (mock *ReceiverMock) statsCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockstats.RLock()
	calls = mock.calls.stats
	mock.lockstats.RUnlock()
	return calls
}

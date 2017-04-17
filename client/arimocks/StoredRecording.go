package arimocks

import ari "github.com/CyCoreSystems/ari"
import mock "github.com/stretchr/testify/mock"

// StoredRecording is an autogenerated mock type for the StoredRecording type
type StoredRecording struct {
	mock.Mock
}

// Copy provides a mock function with given fields: name, dest
func (_m *StoredRecording) Copy(name string, dest string) (ari.StoredRecordingHandle, error) {
	ret := _m.Called(name, dest)

	var r0 ari.StoredRecordingHandle
	if rf, ok := ret.Get(0).(func(string, string) ari.StoredRecordingHandle); ok {
		r0 = rf(name, dest)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ari.StoredRecordingHandle)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(name, dest)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Data provides a mock function with given fields: name
func (_m *StoredRecording) Data(name string) (*ari.StoredRecordingData, error) {
	ret := _m.Called(name)

	var r0 *ari.StoredRecordingData
	if rf, ok := ret.Get(0).(func(string) *ari.StoredRecordingData); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ari.StoredRecordingData)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: name
func (_m *StoredRecording) Delete(name string) error {
	ret := _m.Called(name)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: name
func (_m *StoredRecording) Get(name string) ari.StoredRecordingHandle {
	ret := _m.Called(name)

	var r0 ari.StoredRecordingHandle
	if rf, ok := ret.Get(0).(func(string) ari.StoredRecordingHandle); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ari.StoredRecordingHandle)
		}
	}

	return r0
}

// List provides a mock function with given fields:
func (_m *StoredRecording) List() ([]ari.StoredRecordingHandle, error) {
	ret := _m.Called()

	var r0 []ari.StoredRecordingHandle
	if rf, ok := ret.Get(0).(func() []ari.StoredRecordingHandle); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]ari.StoredRecordingHandle)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StageCopy provides a mock function with given fields: name, dest
func (_m *StoredRecording) StageCopy(name string, dest string) ari.StoredRecordingHandle {
	ret := _m.Called(name, dest)

	var r0 ari.StoredRecordingHandle
	if rf, ok := ret.Get(0).(func(string, string) ari.StoredRecordingHandle); ok {
		r0 = rf(name, dest)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ari.StoredRecordingHandle)
		}
	}

	return r0
}

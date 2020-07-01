// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	v1beta1 "k8s.io/kubelet/pkg/apis/deviceplugin/v1beta1"
)

// ResourcePool is an autogenerated mock type for the ResourcePool type
type ResourcePool struct {
	mock.Mock
}

// GetDeviceSpecs provides a mock function with given fields: deviceIDs
func (_m *ResourcePool) GetDeviceSpecs(deviceIDs []string) []*v1beta1.DeviceSpec {
	ret := _m.Called(deviceIDs)

	var r0 []*v1beta1.DeviceSpec
	if rf, ok := ret.Get(0).(func([]string) []*v1beta1.DeviceSpec); ok {
		r0 = rf(deviceIDs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*v1beta1.DeviceSpec)
		}
	}

	return r0
}

// GetDevices provides a mock function with given fields:
func (_m *ResourcePool) GetDevices() map[string]*v1beta1.Device {
	ret := _m.Called()

	var r0 map[string]*v1beta1.Device
	if rf, ok := ret.Get(0).(func() map[string]*v1beta1.Device); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]*v1beta1.Device)
		}
	}

	return r0
}

// GetEnvs provides a mock function with given fields: deviceIDs
func (_m *ResourcePool) GetEnvs(deviceIDs []string) []string {
	ret := _m.Called(deviceIDs)

	var r0 []string
	if rf, ok := ret.Get(0).(func([]string) []string); ok {
		r0 = rf(deviceIDs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// GetMounts provides a mock function with given fields: deviceIDs
func (_m *ResourcePool) GetMounts(deviceIDs []string) []*v1beta1.Mount {
	ret := _m.Called(deviceIDs)

	var r0 []*v1beta1.Mount
	if rf, ok := ret.Get(0).(func([]string) []*v1beta1.Mount); ok {
		r0 = rf(deviceIDs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*v1beta1.Mount)
		}
	}

	return r0
}

// GetResourceName provides a mock function with given fields:
func (_m *ResourcePool) GetResourceName() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetResourcePrefix provides a mock function with given fields:
func (_m *ResourcePool) GetResourcePrefix() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Probe provides a mock function with given fields:
func (_m *ResourcePool) Probe() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

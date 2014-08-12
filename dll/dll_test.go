package dll

import (
	"testing"
)

func TestFunctions(t *testing.T) {
	GetvJoyVersion()
	VJoyEnabled()
	GetvJoyProductString()
	GetvJoyManufacturerString()
	GetvJoySerialNumberString()
	rid := uint(1)
	GetVJDButtonNumber(rid)
	GetVJDContPovNumber(rid)
	GetVJDAxisExist(rid, AXIS_X)
	var int32val int32
	GetVJDAxisMax(rid, AXIS_X, &int32val)
	GetVJDAxisMin(rid, AXIS_X, &int32val)
	GetVJDStatus(rid)
	ResetAll()
	if AcquireVJD(rid) {
		t.Log("vjd acquired")
		ResetVJD(rid)
		ResetButtons(rid)
		ResetPovs(rid)
		SetAxis(30000, rid, AXIS_X)
		SetBtn(true, rid, 1)
		SetDiscPov(5, rid, 1)
		SetContPov(5, rid, 1)
		RelinquishVJD(rid)
	} else {
		t.Log("no vjd")
	}
}

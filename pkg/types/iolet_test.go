package types

import "testing"

func TestIOletStatus(t *testing.T) {
	ioLetStatus := IOletStatus(0)
	if !ioLetStatus.OK() {
		t.Error("iolet status should be OK")
	}
	ioLetStatus.SetOK(true)
	if !ioLetStatus.OK() {
		t.Error("iolet status should be OK")
	}
	ioLetStatus.SetOK(false)
	if ioLetStatus.OK() {
		t.Error("iolet status should be not OK")
	}
	ioLetStatus.SetOK(false)
	if ioLetStatus.OK() {
		t.Error("iolet status should be not OK")
	}
	ioLetStatus.SetOK(true)
	if !ioLetStatus.OK() {
		t.Error("iolet status should be OK")
	}
}

func TestIOletStatusRunning(t *testing.T) {
	ioLetStatus := IOletStatus(0)
	if ioLetStatus.Running() {
		t.Error("iolet status should be not Running")
	}
	ioLetStatus.SetRunning(true)
	if !ioLetStatus.Running() {
		t.Error("iolet status should be Running")
	}
	ioLetStatus.SetRunning(false)
	if ioLetStatus.Running() {
		t.Error("iolet status should be not Running")
	}
	ioLetStatus.SetRunning(false)
	if ioLetStatus.Running() {
		t.Error("iolet status should be not Running")
	}
	ioLetStatus.SetRunning(true)
	if !ioLetStatus.Running() {
		t.Error("iolet status should be Running")
	}
}

func TestIOletStatusReceiving(t *testing.T) {
	ioLetStatus := IOletStatus(0)
	if ioLetStatus.Receiving() {
		t.Error("iolet status should be not Receiving")
	}
	ioLetStatus.SetReceiving(true)
	if !ioLetStatus.Receiving() {
		t.Error("iolet status should be Receiving")
	}
	ioLetStatus.SetReceiving(false)
	if ioLetStatus.Receiving() {
		t.Error("iolet status should be not Receiving")
	}
	ioLetStatus.SetReceiving(false)
	if ioLetStatus.Receiving() {
		t.Error("iolet status should be not Receiving")
	}
	ioLetStatus.SetReceiving(true)
	if !ioLetStatus.Receiving() {
		t.Error("iolet status should be Receiving")
	}
}

func TestIOletStatusSending(t *testing.T) {
	ioLetStatus := IOletStatus(0)
	if ioLetStatus.Sending() {
		t.Error("iolet status should be not Sending")
	}
	ioLetStatus.SetSending(true)
	if !ioLetStatus.Sending() {
		t.Error("iolet status should be Sending")
	}
	ioLetStatus.SetSending(false)
	if ioLetStatus.Sending() {
		t.Error("iolet status should be not Sending")
	}
	ioLetStatus.SetSending(false)
	if ioLetStatus.Sending() {
		t.Error("iolet status should be not Sending")
	}
	ioLetStatus.SetSending(true)
	if !ioLetStatus.Sending() {
		t.Error("iolet status should be Sending")
	}
}

func TestIOletStatusHigh(t *testing.T) {
	ioLetStatus := IOletStatus(0)
	if ioLetStatus.High() {
		t.Error("iolet status should be not High")
	}
	ioLetStatus.SetHigh(true)
	if !ioLetStatus.High() {
		t.Error("iolet status should be High")
	}
	ioLetStatus.SetHigh(false)
	if ioLetStatus.High() {
		t.Error("iolet status should be not High")
	}
	ioLetStatus.SetHigh(false)
	if ioLetStatus.High() {
		t.Error("iolet status should be not High")
	}
	ioLetStatus.SetHigh(true)
	if !ioLetStatus.High() {
		t.Error("iolet status should be High")
	}
}

func TestIOletStatuses(t *testing.T) {
	ioLetStatus := IOletStatus(0)
	ioLetStatus.SetHigh(true)
	ioLetStatus.SetSending(true)

	if !ioLetStatus.OK() {
		t.Error("should be ok")
	}

	ioLetStatus.SetOK(false)

	if ioLetStatus.OK() {
		t.Error("should be not ok")
	}

	if !ioLetStatus.Sending() {
		t.Error("should be sending")
	}
}

package types

import "testing"

func TestModuleStatus(t *testing.T) {
	var moduleStatus = ModuleStatus(0)
	if !moduleStatus.OK() {
		t.Error("module status should be OK")
	}
	moduleStatus.SetOK(true)
	if !moduleStatus.OK() {
		t.Error("module status should be OK")
	}
	moduleStatus.SetOK(false)
	if moduleStatus.OK() {
		t.Error("module status should be not OK")
	}
	moduleStatus.SetOK(false)
	if moduleStatus.OK() {
		t.Error("module status should be not OK")
	}
	moduleStatus.SetOK(true)
	if !moduleStatus.OK() {
		t.Error("module status should be OK")
	}
}

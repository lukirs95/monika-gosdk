package driver

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lukirs95/monika-gosdk/pkg/types"
)

func logRequestError(logger *log.Logger, r *http.Request, err error) {
	e := fmt.Sprintf("HANDLER ERROR: [%s] %s", r.URL.String(), err.Error())
	logger.Print(e)
}

func (service *Service) handleGetDevices(w http.ResponseWriter, r *http.Request) {
	devices := service.driver.GetDevices()

	w.Header().Add("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(devices); err != nil {
		logRequestError(service.logger, r, err)
	}
}

func (service *Service) handleGetDevice(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	deviceId := types.DeviceId(vars["deviceId"])

	device := service.driver.GetDevice(deviceId)

	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(device); err != nil {
		logRequestError(service.logger, r, err)
	}
}

func (service *Service) handleDeviceControl(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	deviceId := types.DeviceId(vars["deviceId"])
	control := types.DeviceControl(vars["deviceControl"])

	if err := control.Valid(); err != nil {
		logRequestError(service.logger, r, err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := service.driver.RunDeviceControl(deviceId, control); err != nil {
		logRequestError(service.logger, r, err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (service *Service) handleGetModules(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	deviceId := types.DeviceId(vars["deviceId"])

	modules := service.driver.GetModules(deviceId)

	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(modules); err != nil {
		logRequestError(service.logger, r, err)
	}
}

func (service *Service) handleGetModulesByType(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	deviceId := types.DeviceId(vars["deviceId"])
	moduleType := types.ModuleType(vars["moduleType"])

	modules := service.driver.GetModulesByModuleType(deviceId, moduleType)

	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(modules); err != nil {
		logRequestError(service.logger, r, err)
	}
}

func (service *Service) handleGetModule(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	deviceId := types.DeviceId(vars["deviceId"])
	moduleType := types.ModuleType(vars["moduleType"])
	moduleId := types.ModuleId(vars["moduleId"])

	module := service.driver.GetModule(deviceId, moduleType, moduleId)

	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(module); err != nil {
		logRequestError(service.logger, r, err)
	}
}

func (service *Service) handleModuleControl(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	deviceId := types.DeviceId(vars["deviceId"])
	moduleType := types.ModuleType(vars["moduleType"])
	moduleId := types.ModuleId(vars["moduleId"])
	control := types.ModuleControl(vars["moduleControl"])

	if err := control.Valid(); err != nil {
		logRequestError(service.logger, r, err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := service.driver.RunModuleControl(deviceId, moduleType, moduleId, control); err != nil {
		logRequestError(service.logger, r, err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (service *Service) handleGetIOlets(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	deviceId := types.DeviceId(vars["deviceId"])
	moduleType := types.ModuleType(vars["moduleType"])
	moduleId := types.ModuleId(vars["moduleId"])

	iolets := service.driver.GetIOlets(deviceId, moduleType, moduleId)

	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(iolets); err != nil {
		logRequestError(service.logger, r, err)
	}
}

func (service *Service) handleGetIOletsByType(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	deviceId := types.DeviceId(vars["deviceId"])
	moduleType := types.ModuleType(vars["moduleType"])
	moduleId := types.ModuleId(vars["moduleId"])
	ioletType := types.IOletType(vars["ioletType"])

	iolets := service.driver.GetIOletsByIOletType(deviceId, moduleType, moduleId, ioletType)

	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(iolets); err != nil {
		logRequestError(service.logger, r, err)
	}
}

func (service *Service) handleGetIOlet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	deviceId := types.DeviceId(vars["deviceId"])
	moduleType := types.ModuleType(vars["moduleType"])
	moduleId := types.ModuleId(vars["moduleId"])
	ioletType := types.IOletType(vars["ioletType"])
	ioletId := types.IOletId(vars["ioletId"])

	iolet := service.driver.GetIOlet(deviceId, moduleType, moduleId, ioletType, ioletId)

	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(iolet); err != nil {
		logRequestError(service.logger, r, err)
	}
}

func (service *Service) handleIOletControl(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	deviceId := types.DeviceId(vars["deviceId"])
	moduleType := types.ModuleType(vars["moduleType"])
	moduleId := types.ModuleId(vars["moduleId"])
	ioletType := types.IOletType(vars["ioletType"])
	ioletId := types.IOletId(vars["ioletId"])
	control := types.IOletControl(vars["ioletControl"])

	if err := control.Valid(); err != nil {
		logRequestError(service.logger, r, err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := service.driver.RunIOletCommand(deviceId, moduleType, moduleId, ioletType, ioletId, control); err != nil {
		logRequestError(service.logger, r, err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

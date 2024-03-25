package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lukirs95/monika-gosdk/pkg/types"
)

type Service struct {
	router   *mux.Router
	provider Provider
	logger   *log.Logger
	gateway  string
}

func NewService(gateway string, provider Provider, logger *log.Logger) *Service {
	router := mux.NewRouter()

	service := &Service{
		router:   router,
		provider: provider,
		logger:   logger,
		gateway:  gateway,
	}

	// GetAllDevices
	router.HandleFunc("/", service.handleGetDevices).Methods(http.MethodGet)
	router.HandleFunc("/{deviceId}", service.handleGetDevice).Methods(http.MethodGet)
	router.HandleFunc("/{deviceId}/{deviceControl}", service.handleDeviceControl).Methods(http.MethodPost)
	router.HandleFunc("/{deviceId}/modules", service.handleGetModules).Methods(http.MethodGet)
	router.HandleFunc("/{deviceId}/modules/{moduleType}", service.handleGetModulesByType).Methods(http.MethodGet)
	router.HandleFunc("/{deviceId}/modules/{moduleType}/{moduleId}", service.handleGetModule).Methods(http.MethodGet)
	router.HandleFunc("/{deviceId}/modules/{moduleType}/{moduleId}/{moduleControl}", service.handleModuleControl).Methods(http.MethodPost)
	router.HandleFunc("/{deviceId}/modules/{moduleType}/{moduleId}/iolets", service.handleGetIOlets).Methods(http.MethodGet)
	router.HandleFunc("/{deviceId}/modules/{moduleType}/{moduleId}/iolets/{ioletType}", service.handleGetIOletsByType).Methods(http.MethodGet)
	router.HandleFunc("/{deviceId}/modules/{moduleType}/{moduleId}/iolets/{ioletType}/{ioletId}", service.handleGetIOlet).Methods(http.MethodGet)
	router.HandleFunc("/{deviceId}/modules/{moduleType}/{moduleId}/iolets/{ioletType}/{ioletId}/{ioletControl}", service.handleIOletControl).Methods(http.MethodPost)
	return service
}

func logRequestError(logger *log.Logger, r *http.Request, err error) {
	e := fmt.Sprintf("HANDLER ERROR: [%s] %s", r.URL.String(), err.Error())
	logger.Print(e)
}

func (service *Service) handleGetDevices(w http.ResponseWriter, r *http.Request) {
	devices := service.provider.GetDevices()

	w.Header().Add("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(devices); err != nil {
		logRequestError(service.logger, r, err)
	}
}

func (service *Service) handleGetDevice(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	deviceId := types.DeviceId(vars["deviceId"])

	device := service.provider.GetDevice(deviceId)

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

	if err := service.provider.RunDeviceControl(deviceId, control); err != nil {
		logRequestError(service.logger, r, err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (service *Service) handleGetModules(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	deviceId := types.DeviceId(vars["deviceId"])

	modules := service.provider.GetModules(deviceId)

	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(modules); err != nil {
		logRequestError(service.logger, r, err)
	}
}

func (service *Service) handleGetModulesByType(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	deviceId := types.DeviceId(vars["deviceId"])
	moduleType := types.ModuleType(vars["moduleType"])

	modules := service.provider.GetModulesByModuleType(deviceId, moduleType)

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

	module := service.provider.GetModule(deviceId, moduleType, moduleId)

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

	if err := service.provider.RunModuleControl(deviceId, moduleType, moduleId, control); err != nil {
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

	iolets := service.provider.GetIOlets(deviceId, moduleType, moduleId)

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

	iolets := service.provider.GetIOletsByIOletType(deviceId, moduleType, moduleId, ioletType)

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

	iolet := service.provider.GetIOlet(deviceId, moduleType, moduleId, ioletType, ioletId)

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

	if err := service.provider.RunIOletCommand(deviceId, moduleType, moduleId, ioletType, ioletId, control); err != nil {
		logRequestError(service.logger, r, err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (service *Service) sendUpdates() {
	updatedDevices := service.provider.GetUpdates()
	if len(updatedDevices) > 0 {
		body, err := json.Marshal(updatedDevices)
		if err != nil {
			log.Print(err)
			return
		}
		reader := bytes.NewReader(body)
		http.Post(fmt.Sprintf("%s/api/notify/update", service.gateway), "application/json", reader)
	}
}

func (service *Service) connect(port int) error {
	body, err := json.Marshal(&types.Driver{
		DeviceType: service.provider.GetDeviceType(),
		Port:       port,
	})
	if err != nil {
		return err
	}
	reader := bytes.NewReader(body)

	resp, err := http.Post(fmt.Sprintf("%s/driver/connect", service.gateway), "application/json", reader)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("gateway responed with Status %s", resp.Status)
	}
	return nil
}

func (service *Service) disconnect() error {
	body, err := json.Marshal(&types.Driver{
		DeviceType: service.provider.GetDeviceType(),
	})
	if err != nil {
		return err
	}
	reader := bytes.NewReader(body)

	resp, err := http.Post(fmt.Sprintf("%s/driver/disconnect", service.gateway), "application/json", reader)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("gateway responed with Status %s", resp.Status)
	}
	return nil
}

func (service *Service) listenForUpdates(updateChan chan interface{}) {
	for range updateChan {
		service.sendUpdates()
	}
}

func (service *Service) Listen(ctx context.Context, port int, updateChan chan interface{}) error {
	if err := service.connect(port); err != nil {
		return err
	}
	defer service.disconnect()

	go service.listenForUpdates(updateChan)

	return http.ListenAndServe(fmt.Sprintf(":%d", port), service.router)
}

func (service *Service) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	service.router.ServeHTTP(w, r)
}

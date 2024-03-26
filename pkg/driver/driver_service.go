package driver

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
	gateway          string
	driver           Driver
	logger           *log.Logger
	router           *mux.Router
	checkDeviceError types.ErrorCheckerDevice
	checkModuleError types.ErrorCheckerModule
	checkIOletError  types.ErrorCheckerIOlet
	deviceErrors     map[types.DeviceId]*types.Error
	moduleErrors     map[types.ModuleId]*types.Error
	ioletErrors      map[types.IOletId]*types.Error
}

func NewService(gateway string, driver Driver, logger *log.Logger) *Service {
	router := mux.NewRouter()

	service := &Service{
		router:           router,
		driver:           driver,
		logger:           logger,
		gateway:          gateway,
		checkDeviceError: func(device *types.DeviceUpdate) *types.Error { return nil },
		checkModuleError: func(device *types.ModuleUpdate) *types.Error { return nil },
		checkIOletError:  func(device *types.IOletUpdate) *types.Error { return nil },
		deviceErrors:     make(map[types.DeviceId]*types.Error),
		moduleErrors:     make(map[types.ModuleId]*types.Error),
		ioletErrors:      make(map[types.IOletId]*types.Error),
	}

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

func (service *Service) Listen(ctx context.Context, port int, updateChan chan types.Device) error {
	if err := service.connect(port); err != nil {
		return err
	}
	defer service.disconnect()

	go func(updateChan chan types.Device) {
		for device := range updateChan {
			updated := device.Updated()
			if updated != nil {
				service.checkForDeviceErrors(updated)
				service.reportUpdate(updated)
			}
		}
	}(updateChan)

	return http.ListenAndServe(fmt.Sprintf(":%d", port), service.router)
}

func (service *Service) AddErrorCheckDevice(deviceChecker types.ErrorCheckerDevice) {
	service.checkDeviceError = deviceChecker
}

func (service *Service) AddErrorCheckModule(moduleChecker types.ErrorCheckerModule) {
	service.checkModuleError = moduleChecker
}

func (service *Service) AddErrorCheckIOlet(ioletChecker types.ErrorCheckerIOlet) {
	service.checkIOletError = ioletChecker
}

func (service *Service) reportUpdate(device *types.DeviceUpdate) {
	body, err := json.Marshal(device)
	if err != nil {
		service.logger.Print(err)
		return
	}
	reader := bytes.NewReader(body)
	res, err := http.Post(fmt.Sprintf("%s/api/notify/update", service.gateway), "application/json", reader)
	if err != nil {
		service.logger.Print(err)
		return
	}
	if res.StatusCode != http.StatusOK {
		service.logger.Print("could not send update")
		service.logger.Print(res.Status)
	}
}

func (service *Service) reportError(newError *types.PubError) (*types.Error, error) {
	body, err := json.Marshal(newError)
	if err != nil {
		return nil, err
	}

	reader := bytes.NewReader(body)
	res, err := http.Post(fmt.Sprintf("%s/api/notify/error", service.gateway), "application/json", reader)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusCreated {
		service.logger.Print("could not create error: ", res.Status)
		return nil, err
	}

	errorResponse := types.PubErrorResponse{}
	if err := json.NewDecoder(res.Body).Decode(&errorResponse); err != nil {
		service.logger.Print("could not read response from server: ", err)
		return nil, err
	}

	return &types.Error{
		ErrorId:  errorResponse.ErrorId,
		Severity: newError.Severity,
		Message:  newError.Message,
	}, nil
}

func (service *Service) deleteError(oldError *types.Error) error {
	request, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/api/notify/error/%d", service.gateway, oldError.ErrorId), nil)
	if err != nil {
		return err
	}

	client := http.DefaultClient
	res, err := client.Do(request)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("could not create error: %s", res.Status)
	}
	return nil
}

func (service *Service) connect(port int) error {
	body, err := json.Marshal(&types.Driver{
		DeviceType: service.driver.GetDeviceType(),
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
		DeviceType: service.driver.GetDeviceType(),
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

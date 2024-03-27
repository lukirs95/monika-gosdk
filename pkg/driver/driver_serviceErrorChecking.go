package driver

import "github.com/lukirs95/monika-gosdk/pkg/types"

func (service *Service) checkForDeviceErrors(device *types.DeviceUpdate) {
	currentDeviceError, ok := service.deviceErrors[device.Id]
	if deviceError := service.checkDeviceError(device); deviceError != nil { // new error
		if !ok { // no old error
			service.reportDeviceError(device, deviceError)
		} else { // there is an old error reported
			if currentDeviceError.Message != deviceError.Message { // same error
				if deviceError.Severity > currentDeviceError.Severity { // higher severity
					service.deleteDeviceError(device, currentDeviceError)
					service.reportDeviceError(device, deviceError)
				}
			}
		}
	} else { // no new error
		if ok { // there is an old error
			service.deleteDeviceError(device, currentDeviceError)
		}
	}

	for _, module := range device.Modules {
		service.checkForModuleErrors(device, &module)
	}
}

func (service *Service) checkForModuleErrors(device *types.DeviceUpdate, module *types.ModuleUpdate) {
	currentModuleError, ok := service.moduleErrors[module.Id]
	if moduleError := service.checkModuleError(module); moduleError != nil { // new error
		if !ok { // no old error
			service.reportModuleError(device, module, moduleError)
		} else { // there is an old error reported
			if currentModuleError.Message != moduleError.Message { // same error
				if moduleError.Severity > currentModuleError.Severity { // higher severity
					service.deleteModuleError(module, currentModuleError)
					service.reportModuleError(device, module, moduleError)
				}
			}
		}
	} else { // no new error
		if ok { // there is an old error
			service.deleteModuleError(module, currentModuleError)
		}
	}

	for _, iolet := range module.IOlets {
		service.checkForIOletErrors(device, module, &iolet)
	}
}

func (service *Service) checkForIOletErrors(device *types.DeviceUpdate, module *types.ModuleUpdate, iolet *types.IOletUpdate) {
	currentIOletError, ok := service.ioletErrors[iolet.Id]
	if ioletError := service.checkIOletError(iolet); ioletError != nil { // new error
		if !ok { // no old error
			service.reportIOletError(device, module, iolet, ioletError)
		} else { // there is an old error reported
			if currentIOletError.Message != ioletError.Message { // same error
				if ioletError.Severity > currentIOletError.Severity { // higher severity
					service.deleteIOletError(iolet, currentIOletError)
					service.reportIOletError(device, module, iolet, ioletError)
				}
			}
		}
	} else { // no new error
		if ok { // there is an old error
			service.deleteIOletError(iolet, currentIOletError)
		}
	}
}

func (service *Service) reportDeviceError(device *types.DeviceUpdate, deviceError *types.Error) {
	reportedError, err := service.reportError(&types.PubError{
		DeviceId:   device.Id,
		DeviceType: device.Type,
		DeviceName: device.Name,
		Severity:   deviceError.Severity,
		Message:    deviceError.Message,
	})

	if err != nil {
		service.logger.Print(err)
		return
	}
	service.deviceErrors[device.Id] = reportedError
}

func (service *Service) reportModuleError(device *types.DeviceUpdate, module *types.ModuleUpdate, moduleError *types.Error) {
	reportedError, err := service.reportError(&types.PubError{
		DeviceId:   device.Id,
		DeviceType: device.Type,
		DeviceName: device.Name,
		ModuleId:   module.Id,
		ModuleType: module.Type,
		ModuleName: module.Name,
		Severity:   moduleError.Severity,
		Message:    moduleError.Message,
	})

	if err != nil {
		service.logger.Print(err)
		return
	}
	service.moduleErrors[module.Id] = reportedError
}

func (service *Service) reportIOletError(device *types.DeviceUpdate, module *types.ModuleUpdate, iolet *types.IOletUpdate, ioletError *types.Error) {
	reportedError, err := service.reportError(&types.PubError{
		DeviceId:   device.Id,
		DeviceType: device.Type,
		DeviceName: device.Name,
		ModuleId:   module.Id,
		ModuleType: module.Type,
		ModuleName: module.Name,
		IOletId:    iolet.Id,
		IOletType:  iolet.Type,
		IOletName:  iolet.Name,
		Severity:   ioletError.Severity,
		Message:    ioletError.Message,
	})

	if err != nil {
		service.logger.Print(err)
		return
	}
	service.ioletErrors[iolet.Id] = reportedError
}

func (service *Service) deleteDeviceError(device *types.DeviceUpdate, deviceError *types.Error) {
	if err := service.deleteError(deviceError); err != nil {
		service.logger.Print(err)
		return
	}
	delete(service.deviceErrors, device.Id)
}

func (service *Service) deleteModuleError(module *types.ModuleUpdate, moduleError *types.Error) {
	if err := service.deleteError(moduleError); err != nil {
		service.logger.Print(err)
		return
	}
	delete(service.moduleErrors, module.Id)
}

func (service *Service) deleteIOletError(iolet *types.IOletUpdate, ioletError *types.Error) {
	if err := service.deleteError(ioletError); err != nil {
		service.logger.Print(err)
		return
	}
	delete(service.ioletErrors, iolet.Id)
}

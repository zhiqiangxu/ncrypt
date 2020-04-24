package service

import "sync"

// Service for ncrypt
type Service struct {
}

var (
	instanceService *Service
	lockService     sync.Mutex
)

// Instance is singleton for Service
func Instance() *Service {
	if instanceService != nil {
		return instanceService
	}

	lockService.Lock()
	defer lockService.Unlock()
	if instanceService != nil {
		return instanceService
	}

	instanceService = &Service{}

	return instanceService
}

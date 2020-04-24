package instance

import (
	"sync"

	"github.com/zhiqiangxu/ncrypt/pkg/secret"
)

var (
	secretInstance *secret.Secret
	secretMu       sync.Mutex
)

// Secret is singleton for secret.Secret
func Secret() *secret.Secret {
	if secretInstance != nil {
		return secretInstance
	}

	secretMu.Lock()
	defer secretMu.Unlock()

	if secretInstance != nil {
		return secretInstance
	}

	secretInstance = secret.New()
	return secretInstance

}

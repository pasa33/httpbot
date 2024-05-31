package httpbot

import (
	"strings"
)

type DeviceHeader string

const (
	UserAgent            DeviceHeader = "user-agent"
	SecUa                DeviceHeader = "sec-ch-ua"
	SecUaMobile          DeviceHeader = "sec-ch-ua-mobile"
	SecUaArch            DeviceHeader = "sec-ch-ua-arch"
	SecUaFullVersion     DeviceHeader = "sec-ch-ua-full-version"
	SecUaPlatformVersion DeviceHeader = "sec-ch-ua-platform-version"
	SecUaFullVersionList DeviceHeader = "sec-ch-ua-full-version-list"
	SecUaBitness         DeviceHeader = "sec-ch-ua-bitness"
	SecUaModel           DeviceHeader = "sec-ch-ua-model"
	SecUaPlatform        DeviceHeader = "sec-ch-ua-platform"
	AcceptLanguage       DeviceHeader = "accept-language"
)

func (bot *HttpBot) AddDevice(id string, headers map[DeviceHeader]string) {
	bot.deviceMu.Lock()
	defer bot.deviceMu.Unlock()

	id = strings.ToLower(id)

	if bot.devices == nil {
		bot.devices = map[string]map[DeviceHeader]string{}
	}

	bot.devices[id] = headers
	if len(bot.devices) == 1 {
		bot.useDeviceId = id
	}
}

func (bot *HttpBot) EditDevice(id string, headers map[DeviceHeader]string) {
	bot.deviceMu.Lock()
	defer bot.deviceMu.Unlock()

	id = strings.ToLower(id)
	if _, ok := bot.devices[id]; !ok {
		return
	}
	for k, v := range headers {
		bot.devices[id][k] = v
	}
}

func (bot *HttpBot) UseDevice(id string) {
	bot.deviceMu.Lock()
	defer bot.deviceMu.Unlock()

	id = strings.ToLower(id)
	if _, ok := bot.devices[id]; ok {
		bot.useDeviceId = id
	}
}

func (bot *HttpBot) getInUseDeviceValue(key string) string {
	bot.deviceMu.RLock()
	defer bot.deviceMu.RUnlock()

	if val, ok := bot.devices[bot.useDeviceId]; ok {
		key = strings.ToLower(key)
		return val[DeviceHeader(key)]
	}
	return ""
}

func (bot *HttpBot) GetInUseUA() string {
	bot.deviceMu.RLock()
	defer bot.deviceMu.RUnlock()

	if val, ok := bot.devices[bot.useDeviceId]; ok {
		return val[UserAgent]
	}
	return ""
}

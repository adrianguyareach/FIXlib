package header

import (
	"fmt"
	"strings"
	"time"

	"github.com/quickfixgo/field"
	"github.com/quickfixgo/quickfix"
)

func convertGlobalSettingsToMap(inputMap map[quickfix.SessionID]*quickfix.SessionSettings) map[string]string {

	var globalSettings string
	for _, value := range inputMap {

		globalSettings = fmt.Sprintf("%s", *value)
	}
	globalSettings = strings.Trim(globalSettings, "{}")
	keyValuePairs := strings.Split(globalSettings, " ")

	resultMap := make(map[string]string)

	for _, pair := range keyValuePairs {
		// Split each pair into key and value
		parts := strings.Split(pair, ":")
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			resultMap[key] = value
		}
	}
	return resultMap
}

func SetHeader(header quickfix.Header, appSettings *quickfix.Settings) quickfix.Header {
	sessionSettings := appSettings.SessionSettings()

	globalSettings := convertGlobalSettingsToMap(sessionSettings)
	// Access specific session-level settings
	targetCompID := globalSettings["TargetCompID"]
	senderCompID := globalSettings["SenderCompID"]
	beginString := globalSettings["BeginString"]

	// // ... other session-level settings
	// // Set header fields
	header.Set(field.NewBeginString(beginString))
	header.Set(field.NewSenderCompID(senderCompID))
	header.Set(field.NewTargetCompID(targetCompID))
	header.Set(field.NewSendingTime(time.Now().UTC()))

	return header

}

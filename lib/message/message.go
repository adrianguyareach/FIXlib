package message

import (
	"fmt"
	"strings"
	"time"

	"github.com/quickfixgo/field"
	"github.com/quickfixgo/quickfix"
)

type Message struct {
	Content *quickfix.Message
}

type MessageConstructor interface {
	ConstructMessage(appSettings *quickfix.Settings) *quickfix.Message
}

func removeCharacters(input string, charsToRemove string) string {
	result := input
	for _, char := range charsToRemove {
		result = strings.ReplaceAll(result, string(char), "")
	}
	return result
}

func convertGlobalSettingsToMap(inputMap map[quickfix.SessionID]*quickfix.SessionSettings) map[string]string {

	var globalSettings string

	for _, value := range inputMap {

		globalSettings = fmt.Sprintf("%s", *value)
	}

	mapsplit := strings.Split(globalSettings, "[")
	clean := removeCharacters(mapsplit[1], "{[]}")
	keyValuePairs := strings.Split(clean, " ")
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

func (m Message) ConstructMessage(appSettings *quickfix.Settings) *quickfix.Message {
	sessionSettings := appSettings.SessionSettings()

	header := m.Content.Header
	globalSettings := convertGlobalSettingsToMap(sessionSettings)
	// // Set header fields
	header.Set(field.NewBeginString(globalSettings["BeginString"]))
	header.Set(field.NewSenderCompID(globalSettings["SenderCompID"]))
	header.Set(field.NewTargetCompID(globalSettings["TargetCompID"]))
	header.Set(field.NewSendingTime(time.Now().UTC()))

	newMsg := quickfix.NewMessage()
	newMsg.Header = header
	newMsg.Body = m.Content.Body
	newMsg.Trailer = m.Content.Trailer
	return newMsg

}

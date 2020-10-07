package i18n

import (
	"cygo_iris/util/log"
	"fmt"
)

func DefaultMessageFunc(langInput, langMatched, key string, args ...interface{}) string {
	msg := fmt.Sprintf("user language input: %s; matched as: %s; not found key: %s; args: %v", langInput, langMatched, key, args)
	log.Logger.Warn(msg)
	return msg
}

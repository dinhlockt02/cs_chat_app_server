package common

import "github.com/rs/zerolog/log"

func Recovery() {
	if err := recover(); err != nil {
		log.Error().Stack().Err(err.(error)).Msg("")
	}
	return
}

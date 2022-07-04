package api

import (
	"strconv"
	"strings"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

func RequiredUUID(query string, hint string) (uuid.UUID, error) {
	value, err := uuid.Parse(query)
	if err != nil {
		log.Warn("Param '" + hint + "' must a valid UUID")
		return uuid.Nil, err
	}
	return value, nil
}

func OptionalInt(query string, hint string, defaultValue int) int {
	if query == "" {
		return defaultValue
	}

	outcome, err := strconv.Atoi(query)
	if err != nil {
		log.Warn("Param '" + hint + "' is not a valid int")
		return defaultValue
	}

	return outcome
}

func OptionalBool(query string, hint string, defaultValue bool) bool {
	if query == "" {
		return defaultValue
	}

	outcome, err := strconv.ParseBool(query)
	if err != nil {
		log.Warn("Param '" + hint + "' is not a valid bool")
		return defaultValue
	}
	return outcome
}

func OptionalStringList(query string, hint string) []string {
	if len(query) == 0 {
		return []string{}
	}

	outcome := strings.Split(query, ",")

	return outcome
}

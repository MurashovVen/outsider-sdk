package entities

import (
	"strings"
)

type ActionType int64

const (
	ActionWhetherConfigureType ActionType = iota + 1
	ActionWhetherTemperatureConfigureType
)

func ActionTypeParseString(action string) ActionType {
	switch {
	case action == ActionWhetherConfigure:
		return ActionWhetherConfigureType

	case strings.Contains(action, ActionWhetherTemperatureConfigure):
		return ActionWhetherTemperatureConfigureType

	default:
		return -1
	}
}

func (a ActionType) IsWhetherType() bool {
	return a >= 1 && a <= 2
}

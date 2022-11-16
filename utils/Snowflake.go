package utils

import (
	"github.com/bwmarrin/snowflake"
)

func GenerateSnowflake() (string, error) {

	snowflake.Epoch = 1668505810728
	node, err := snowflake.NewNode(10)

	if err != nil {
		return "", err
	}

	return node.Generate().String(), nil
}

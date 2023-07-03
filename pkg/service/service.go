package service

import (
	"encoding/json"
	"fmt"

	"github.com/sabitvrustam/WB_L0/pkg/types"
)

func OrderWrite(order string) {
	var rezult types.Order
	err := json.Unmarshal([]byte(order), &rezult)
	if err != nil {
		fmt.Println(err, rezult)
	}
	fmt.Println(rezult)
}

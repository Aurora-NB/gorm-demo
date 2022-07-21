package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func JsonOut(obj interface{}) {
	o, _ := json.Marshal(obj)
	var buf bytes.Buffer
	_ = json.Indent(&buf, o, "", "\t")
	fmt.Printf("%v\n", buf.String())
}

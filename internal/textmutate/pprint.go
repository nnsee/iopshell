package textmutate

import (
    "encoding/json"
)

func Pprint(input interface{}) string {
    out, err := json.MarshalIndent(input, "", "  ")
    if err == nil {
        return string(out)
    }
    return ""
}

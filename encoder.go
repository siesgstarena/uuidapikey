// Copyright (C) 2019 SIESGSTarena

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//    http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package uuidapikey

import (
	"strconv"
	"strings"

	"github.com/richardlehane/crock32"
)

// Encode will convert your given int64 into base32 crockford encoding format
func Encode(number uint64) string {
	if number < 0 {
		panic("The given input must be greater than or equal to zero")
	}
	encoded := crock32.Encode(number)
	paddingTimes := 7 - len(encoded)
	return strings.ToUpper((strings.Repeat("0", paddingTimes) + encoded))
}

// Decode will convert your given string into original UUID part string
func Decode(decodedString string) string {
	integerDecoded, _ := crock32.Decode(decodedString)
	decoded := strconv.FormatUint(integerDecoded, 16)
	paddingTimes := 8 - len(decoded)
	return (strings.Repeat("0", paddingTimes) + decoded)
}

/*
 * Copyright 2020 Torben Schinke
 *
 * worldiety Enterprise Edition (EE) License
 * See the file LICENSE for allowed usage and redistribution information.
 *
 * Please contact worldiety GmbH or visit www.worldiety.de if you need additional information or have any
 * questions.
 */

package ee

import (
	"fmt"
	"testing"
)

func TestWithPanic(t *testing.T) {
	f := func() error {
		var arr []byte
		_ = arr[0]

		return nil
	}
	err := WithPanic("test call", f)

	if err == nil {
		t.Fatal("should have covered")
	}

	fmt.Println(err)
}

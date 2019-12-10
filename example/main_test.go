// +build e2e

/*
Copyright 2019 The Rigging Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package example

import (
	"github.com/n3wscott/rigging/pkg/installer"
	"testing"

	"knative.dev/pkg/test/logstream"
)

// This test is more for debugging the ko publish process.
func TestKoPublish(t *testing.T) {
	ic, err := installer.ProduceImages()
	if err != nil {
		t.Fatalf("failed to produce images, %s", err)
	}

	for k, v := range ic {
		t.Log(k, "-->", v)
	}
}

// Rest of e2e tests go below:

// TestEcho is an example simple test.
func TestEcho(t *testing.T) {
	cancel := logstream.Start(t)
	defer cancel()

	EchoTestImpl(t)
}

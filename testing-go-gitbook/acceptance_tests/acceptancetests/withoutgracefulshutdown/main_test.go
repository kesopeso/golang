package main_test

import (
	"acceptance_tests/acceptancetests"
	"acceptance_tests/assert"
	"testing"
	"time"
)

const (
	port = "8081"
	url  = "http://localhost:" + port
)

func TestNonGracefulShutdown(t *testing.T) {
	cleanup, sendInterrupt, err := acceptancetests.LaunchTestProgram(port)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(cleanup)

	// check that server works, before shutting down
	assert.CanGet(t, url)

	// fire off a request and before it has a chance to respond send SIGTERM
	time.AfterFunc(50*time.Millisecond, func() {
		assert.NoError(t, sendInterrupt())
	})

	// because we're not using graceful shutdown, this should fail
	assert.CantGet(t, url)

	// after interrupt, the server should be shutdown and no more requests will work
	assert.CantGet(t, url)
}

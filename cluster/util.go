package cluster

import (
	"time"

	"github.com/cenkalti/backoff"
)

// retryUntilSuccess will keep attempting an operation until it succeeds.
//
// Exponential backoff is used to prevent failing attempts from looping madly.
func retryUntilSuccess(name string, operation func() error, strategy backoff.BackOff) {
	errNotifReceiver := func(err error, nextWait time.Duration) {
		log.Error("%s notified of error: %s [next wait=%s]", name, err, nextWait)
	}
	for {
		if err := backoff.RetryNotify(operation, strategy, errNotifReceiver); err != nil {
			log.Error("%s failure: %s [will keep trying]", name, err)
			strategy.Reset()
			continue
		}
		break
	}
}

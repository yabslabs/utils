package context

import (
	"context"
	"os"
	"os/signal"

	"github.com/yabslabs/utils/logging"
)

func NewGracefulShutdownContext() context.Context {
	return GracefulShutdownContext(context.Background())
}

func GracefulShutdownContext(parent context.Context) context.Context {
	return CancelContextOnSignal(parent, os.Interrupt)
}

func CancelContextOnSignal(parent context.Context, signals ...os.Signal) context.Context {
	ctx, cancel := context.WithCancel(parent)

	sigC := make(chan os.Signal, 1)
	signal.Notify(sigC, signals...)

	go func() {
		defer signal.Stop(sigC)

		select {
		case s := <-sigC:
			logging.WithIDFields("CONT-5dfa4c77", "signal", s).Info("received signal, canceling context")
			cancel()
			logging.WithID("CONT-b88cf14a").Info("context canceled.")
		case <-ctx.Done():
		}
	}()

	return ctx
}

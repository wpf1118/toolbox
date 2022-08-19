package tools

import (
	"fmt"
	"github.com/tevino/abool"
	"github.com/wpf1118/toolbox/tools/logging"
	"os"
	"os/signal"
	"syscall"
)

// Service defines the Run() that returns a stop function to call
// Run() needs to run goroutine, and returns in a short period of the time
type Service interface {
	Run() func()
}

// Application gives you basic started, done, stopped to control the service flow
type Application struct {
	started *abool.AtomicBool
	done    chan struct{}
	stopped chan struct{}
	serv    Service
}

// NewApplication creates a new Application struct
func NewApplication(s Service) *Application {
	return &Application{
		started: abool.New(),
		done:    ListenToSignals(),
		stopped: make(chan struct{}),
		serv:    s,
	}
}

// Run runs the service, and is responsible to catch the signals and stops the service.
// This is blocking until Stop() is called or signals are caught.
func (app *Application) Run() {
	if app.started.SetToIf(false, true) {
		stop := app.serv.Run()
		defer func() {
			stop()
			close(app.stopped)
		}()
		<-app.done
	} else {
		panic(fmt.Errorf("the application has been started"))
	}
}

// Stop can be called when there's a manual stop without the signal
func (app *Application) Stop() {
	close(app.done)
	<-app.stopped
}

// ListenToSignals returns a done channel for capturing syscalls to quit the service
func ListenToSignals() chan struct{} {
	sigs := make(chan os.Signal, 1)
	done := make(chan struct{})
	signal.Notify(sigs, syscall.SIGABRT, syscall.SIGILL, syscall.SIGINT, syscall.SIGTERM, syscall.SIGSEGV)
	go func() {
		sig := <-sigs
		logging.InfoF("signal %s", sig.String())
		close(done)
	}()
	return done
}

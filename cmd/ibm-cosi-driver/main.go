package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"k8s.io/klog/v2"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		klog.InfoS("Signal received", "type", sig)
		cancel()

		<-time.After(30 * time.Second)
		os.Exit(1)
	}()

	if err := cmd.ExecuteContext(ctx); err != nil {
		klog.ErrorS(err, "Exiting on error")
	}
}

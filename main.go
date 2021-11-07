package main

import (
	"solverhost/verify"
)

func main() {
	// interrupt := make(chan os.Signal, 1)
	// signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	// ctx, cancel := context.WithCancel(context.Background())
	// defer cancel()

	verify.Do()

	// for {
	// 	select {
	// 	case <-ctx.Done():
	// 		log.Println("all done, shutdown")
	// 		return
	// 	case <-interrupt:
	// 		log.Println("interrupt, shutdown")
	// 		cancel()
	// 		return
	// 	}
	// }
}

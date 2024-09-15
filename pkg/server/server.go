package server

import (
	"errors"
	"fmt"
	"net/http"
)

func ServeDir(destinationDir, serverPort string) error {
	fmt.Printf("Serving %s directory, listening on %s port.\n", destinationDir, serverPort)

	err := http.ListenAndServe(serverPort, http.FileServer(http.Dir(destinationDir)))
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Println("Shutting down server.")
		return nil
	}

	return err
}

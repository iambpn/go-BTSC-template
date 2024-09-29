package route

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/iambpn/go-http-template/cmd/web"
	"github.com/iambpn/go-http-template/internal/handler"

	"github.com/coder/websocket"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func RegisterRoutes() http.Handler {
	// initialize handlers
	indexHandler := handler.NewIndexHandler()
	healthHandler := handler.NewHealthHandler()

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Serving Assets
	fileServer := http.FileServer(http.FS(web.Files))
	r.Handle("/assets/*", fileServer)

	// handler router
	r.Route("/health", healthHandler.RegisterRoutes)
	r.Route("/", indexHandler.RegisterRoutes)

	// Web sockets
	r.Get("/websocket", websocketHandler)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
		w.Write([]byte("route does not exist"))
	})

	r.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(405)
		w.Write([]byte("method is not valid"))
	})

	return r
}

func websocketHandler(w http.ResponseWriter, r *http.Request) {
	socket, err := websocket.Accept(w, r, nil)

	if err != nil {
		log.Printf("could not open websocket: %v", err)
		_, _ = w.Write([]byte("could not open websocket"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer socket.Close(websocket.StatusGoingAway, "server closing websocket")

	ctx := r.Context()
	socketCtx := socket.CloseRead(ctx)

	for {
		payload := fmt.Sprintf("server timestamp: %d", time.Now().UnixNano())
		err := socket.Write(socketCtx, websocket.MessageText, []byte(payload))
		if err != nil {
			break
		}
		time.Sleep(time.Second * 2)
	}
}

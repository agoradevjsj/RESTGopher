package restgopher

import (
	"fmt"
	"log"
	"net/http"

	"github.com/agoradevjsj/RESTGopher/internal/middleware"
	"github.com/agoradevjsj/RESTGopher/internal/router"
)

func main() {
	config := CreateConfigApp(":4545", "", "")
	app := NewServerGopher(config)

	router_gother := router.NewRouterGopher()
	router_gother.HandleFunc("GET", "/hello/:wave", func(w http.ResponseWriter, r *http.Request) {
		req := router_gother.GetRequest(r)
		res := router_gother.GetResponse(&w)
		params := req.GetParams()
		log.Println("Parametros", params)
		response := router.CreateJSON()
		contentWave := router.CreateJSON()
		contentWave["wave"] = params["wave"]
		contentWave["name"] = "Ferico"

		if contentWave["wave"] == "" {
			response = router.CreateJSON()
			response["ok"] = false
			response["message"] = "Falta ingresar el nombre de la empresa."
			res.Status(http.StatusBadRequest).SendJSON(response)
			return
		}

		response["ok"] = true
		response["http_code"] = 12000
		response["message"] = "Este es un mensaje"
		response["person"] = contentWave
		res.Status(http.StatusOK).SendJSON(response)
	})

	router_gother.GET("/test/:id/algo/:animal", func(w http.ResponseWriter, r *http.Request) {
		req := router_gother.GetRequest(r)
		res := router_gother.GetResponse(&w)
		params := req.GetParams()
		res.Status(200).SendJSON(params)
	})
	app.Use(middleware.JSONMiddleware, router_gother.RoutersGopher)

	defer func() {
		r := recover()
		fmt.Println("⏺️ Recovered:", r)
	}()

	errores := app.StartServerGopher()
	if errores != nil {
		fmt.Printf("❌ Error dentro del servidor: %v", errores)
		panic("💻 Error en el servidor.")
	}

}

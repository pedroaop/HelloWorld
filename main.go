package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

// Response define a estrutura da resposta JSON
type Response struct {
	Message string `json:"message"`
}

func main() {
	// Define a porta a partir da variável de ambiente (importante para o Jelastic)
	// Se não estiver definida, usa a 8080 como padrão
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Serve arquivos estáticos (HTML, CSS, JS) do diretório atual
	fs := http.FileServer(http.Dir("."))
	http.Handle("/", fs)

	// Endpoint da API para responder ao clique do botão
	http.HandleFunc("/api/hello", helloHandler)

	fmt.Printf("Servidor iniciando na porta %s...\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

// helloHandler processa a requisição e retorna uma mensagem JSON
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Configura o cabeçalho para JSON
	w.Header().Set("Content-Type", "application/json")
	
	response := Response{
		Message: "Hello World! Sua aplicação Go no Jelastic está funcionando.",
	}

	json.NewEncoder(w).Encode(response)
}

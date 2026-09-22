package main

import (
	"net/http"

	"github.com/DevZank/TaskManagerAPI/config"
	"github.com/DevZank/TaskManagerAPI/handlers"
	"github.com/DevZank/TaskManagerAPI/models"
	"github.com/gorilla/mux"
)

// Declaramos uma função chamada enableCORS.
// Ela recebe como parâmetro "next", que é o handler original (no seu caso, o router do gorilla/mux).
// http.Handler é uma interface do Go: qualquer coisa que sabe "atender" uma requisição HTTP.
// A função retorna também um http.Handler — ou seja, ela pega um handler e devolve outro handler "melhorado".
func enableCORS(next http.Handler) http.Handler {
	// http.HandlerFunc é um "adaptador": ele transforma uma função comum
	// (que recebe ResponseWriter e *Request) em algo que satisfaz a interface http.Handler.
	// Isso é necessário porque a função enableCORS precisa RETORNAR um http.Handler,
	// e não dá pra retornar uma função "crua" sem esse adaptador.
	return http.HandlerFunc(func(writer http.ResponseWriter, response *http.Request) {

		// w é o ResponseWriter: é por ele que escrevemos a RESPOSTA que vai voltar pro navegador.
		// r é o Request: é a requisição que chegou (contém método, URL, headers, corpo, etc).
		// Aqui estamos ADICIONANDO UM HEADER na resposta.
		// Access-Control-Allow-Origin diz ao navegador: "esta origem específica pode ler minha resposta".
		// Trocamos http://localhost:5173 pela URL exata do seu frontend Vite.
		writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")

		// Este header diz quais MÉTODOS HTTP são permitidos vir de outra origem.
		// Sem isso, o navegador bloquearia, por exemplo, um POST vindo do frontend,
		// mesmo que o GET estivesse liberado.
		writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		// Este header diz quais HEADERS customizados o frontend tem permissão de enviar.
		// Como seu frontend manda dados em JSON, ele envia o header "Content-Type: application/json".
		// Se você não liberar "Content-Type" aqui, o navegador bloqueia a requisição.
		writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Aqui verificamos se o MÉTODO da requisição que chegou é "OPTIONS".
		// O navegador manda uma requisição OPTIONS automaticamente ANTES da requisição real,
		// só para "perguntar" ao servidor se ele tem permissão de continuar (isso se chama preflight).
		// O navegador faz isso sozinho, você não precisa programar nada pra ele mandar OPTIONS.
		if response.Method == "OPTIONS" {

			// Se for uma requisição OPTIONS, respondemos com status 200 (OK),
			// avisando ao navegador "pode seguir, está tudo liberado".
			writer.WriteHeader(http.StatusOK)

			// O "return" aqui é importante: ele PARA a execução da função neste ponto.
			// Ou seja, para requisições OPTIONS, NÃO chamamos o handler original (next),
			// porque essa requisição é só uma pergunta, não precisa ser processada de verdade.
			return
		}

		// Se chegou até aqui, significa que NÃO era uma requisição OPTIONS —
		// é a requisição real (GET, POST, etc).
		// Então agora sim chamamos o handler original (o router),
		// passando adiante o ResponseWriter (w) e a Request (r),
		// para que ele processe a requisição normalmente (buscar tasks, criar task, etc).
		next.ServeHTTP(writer, response)
	})
}

func main() {
	dbConn := config.SetupDB()

	_, err := dbConn.Exec(models.CreateTableSQL)
	if err != nil {
		panic(err)
	}

	router := mux.NewRouter()

	taskHandler := handlers.NewTaskHandler(dbConn)

	router.HandleFunc("/tasks", taskHandler.ReadTasks).Methods("GET")
	router.HandleFunc("/tasks", taskHandler.CreateTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", taskHandler.UpdateTask).Methods("PUT")
	router.HandleFunc("/tasks/{id}", taskHandler.DeleteTask).Methods("DELETE")

	// DEFER: Quando a função main() [Função em que ele está em volta] parar de executar: chama a função Close para fechar a conexão com o banco de dados
	defer dbConn.Close()

	panic(http.ListenAndServe(":8080", enableCORS(router)))
}

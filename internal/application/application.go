package application

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"https://github.com/Filatova-Elizaveta/Calculation_0/pkg/calc"
)

type Config struct {
	Addr string
}

type Request struct {
	Expression string `json:"expression"`
}

func ConfigFromEnv() *Config {
	config := new(Config)
	config.Addr = os.Getenv("PORT")
	if config.Addr == "" {
		config.Addr = "8080"
	}
	return config
}

type Application struct {
	config *Config
}

func New() *Application {
	return &Application{
		config: ConfigFromEnv(),
	}
}


func (a *Application) Run() error {
	for {
		log.Println("input expression")
		reader := bufio.NewReader(os.Stdin)
		q, err := reader.ReadString('\n')
		if err != nil {
			log.Println("failed to read expression from console")
		}
		q = strings.TrimSpace(q)
		if q == "exit" {
			log.Println("aplication was successfully closed")
			return nil
		}
		result, err := calculation.Calc(q)
		if err != nil {
			log.Println(q, " calculation failed wit error: ", err)
		} else {
			log.Println(q, "=", result)
		}
	}
}

func HasLetters(s string) bool {
	for _, r := range s {
		if r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' {
			return true
		}
	}
	return false
}


func CalcHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "POST":
		request := new(Request)
		defer r.Body.Close()
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, `{"error": "Internal server error"}`)
			return
		}

		if HasLetters(request.Expression) {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, `{"error": "Expression is not valid"}`)
			return
		}

		result, err := calculation.Calc(request.Expression)
		if err != nil {
			log.Printf("NOT OK: err is %s\n", err.Error())
			w.WriteHeader(http.StatusUnprocessableEntity)
			fmt.Fprintf(w, `{"error": "Expression is not valid"}`)
		} else {
			log.Printf("OK: result is %f\n", result)
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, `{"result": "%f"}`, result)
		}
	default:
		log.Println("not POST request")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"error": "Internal server error"}`))
	}
}

func (a Application) RunServer() {
	log.Println("Server Run...\nhttp://localhost:8080/api/v1/calculate")
	http.HandleFunc("/api/v1/calculate", CalcHandler)
	log.Fatal(http.ListenAndServe(":"+a.config.Addr, nil))
}

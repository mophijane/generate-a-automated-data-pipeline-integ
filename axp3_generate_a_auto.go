package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

type Integration struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type Pipeline struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Source     string `json:"source"`
	Destination string `json:"destination"`
}

type Automation struct {
	ID        string `json:"id"`
	Name       string `json:"name"`
	PipelineID string `json:"pipeline_id"`
}

type Response struct {
	Data  interface{} `json:"data"`
	Error string      `json:"error"`
}

func main() {
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Accept", "Accept-Language", "Content-Type"},
	}))

	r.Group(func(r chi.Router) {
		r.Post("/integrations", createIntegration)
		r.Get("/integrations", get Integrations)
		r.Get("/integrations/{id}", getIntegration)
		r.Put("/integrations/{id}", updateIntegration)
		r.Delete("/integrations/{id}", deleteIntegration)

		r.Post("/pipelines", createPipeline)
		r.Get("/pipelines", getPipelines)
		r.Get("/pipelines/{id}", getPipeline)
		r.Put("/pipelines/{id}", updatePipeline)
		r.Delete("/pipelines/{id}", deletePipeline)

		r.Post("/automations", createAutomation)
		r.Get("/automations", getAutomations)
		r.Get("/automations/{id}", getAutomation)
		r.Put("/automations/{id}", updateAutomation)
		r.Delete("/automations/{id}", deleteAutomation)
	})

	log.Fatal(http.ListenAndServe(":3333", r))
}

func createIntegration(w http.ResponseWriter, r *http.Request) {
	var integration Integration
	err := json.NewDecoder(r.Body).Decode(&integration)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// TO DO: implement creation logic

	response := Response{Data: integration, Error: ""}
	json.NewEncoder(w).Encode(response)
}

func get Integrations(w http.ResponseWriter, r *http.Request) {
	// TO DO: implement retrieval logic

	var integrations []Integration
	response := Response{Data: integrations, Error: ""}
	json.NewEncoder(w).Encode(response)
}

func getIntegration(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	// TO DO: implement retrieval logic

	var integration Integration
	response := Response{Data: integration, Error: ""}
	json.NewEncoder(w).Encode(response)
}

func updateIntegration(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var integration Integration
	err := json.NewDecoder(r.Body).Decode(&integration)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// TO DO: implement update logic

	response := Response{Data: integration, Error: ""}
	json.NewEncoder(w).Encode(response)
}

func deleteIntegration(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	// TO DO: implement deletion logic

	response := Response{Data: id, Error: ""}
	json.NewEncoder(w).Encode(response)
}

func createPipeline(w http.ResponseWriter, r *http.Request) {
	var pipeline Pipeline
	err := json.NewDecoder(r.Body).Decode(&pipeline)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// TO DO: implement creation logic

	response := Response{Data: pipeline, Error: ""}
	json.NewEncoder(w).Encode(response)
}

func getPipelines(w http.ResponseWriter, r *http.Request) {
	// TO DO: implement retrieval logic

	var pipelines []Pipeline
	response := Response{Data: pipelines, Error: ""}
	json.NewEncoder(w).Encode(response)
}

func getPipeline(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	// TO DO: implement retrieval logic

	var pipeline Pipeline
	response := Response{Data: pipeline, Error: ""}
	json.NewEncoder(w).Encode(response)
}

func updatePipeline(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var pipeline Pipeline
	err := json.NewDecoder(r.Body).Decode(&pipeline)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// TO DO: implement update logic

	response := Response{Data: pipeline, Error: ""}
	json.NewEncoder(w).Encode(response)
}

func deletePipeline(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	// TO DO: implement deletion logic

	response := Response{Data: id, Error: ""}
	json.NewEncoder(w).Encode(response)
}

func createAutomation(w http.ResponseWriter, r *http.Request) {
	var automation Automation
	err := json.NewDecoder(r.Body).Decode(&automation)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// TO DO: implement creation logic

	response := Response{Data: automation, Error: ""}
	json.NewEncoder(w).Encode(response)
}

func getAutomations(w http.ResponseWriter, r *http.Request) {
	// TO DO: implement retrieval logic

	var automations []Automation
	response := Response{Data: automations, Error: ""}
	json.NewEncoder(w).Encode(response)
}

func getAutomation(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	// TO DO: implement retrieval logic

	var automation Automation
	response := Response{Data: automation, Error: ""}
	json.NewEncoder(w).Encode(response)
}

func updateAutomation(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var automation Automation
	err := json.NewDecoder(r.Body).Decode(&automation)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// TO DO: implement update logic

	response := Response{Data: automation, Error: ""}
	json.NewEncoder(w).Encode(response)
}

func deleteAutomation(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	// TO DO: implement deletion logic

	response := Response{Data: id, Error: ""}
	json.NewEncoder(w).Encode(response)
}
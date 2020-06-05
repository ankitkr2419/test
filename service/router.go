package service

import (
	"fmt"
	"net/http"

	"mylab/cpagent/config"

	"github.com/gorilla/mux"
)

const (
	versionHeader = "Accept"
)

/* The routing mechanism. Mux helps us define handler functions and the access methods */
func InitRouter(deps Dependencies) (router *mux.Router) {
	router = mux.NewRouter()

	// No version requirement for /ping
	router.HandleFunc("/ping", pingHandler).Methods(http.MethodGet)

	// Version 1 API management
	v1 := fmt.Sprintf("application/vnd.%s.v1", config.AppName())
	fmt.Println("v1", v1)

	router.HandleFunc("/targets", listTargetHandler(deps)).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/template/{id}", updateTemplateHandler(deps)).Methods(http.MethodPut).Headers(versionHeader, v1)
	router.HandleFunc("/template/{id}", showTemplateHandler(deps)).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/template/{id}", deleteTemplateHandler(deps)).Methods(http.MethodDelete).Headers(versionHeader, v1)
	router.HandleFunc("/templates", listTemplateHandler(deps)).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/template", createTemplateHandler(deps)).Methods(http.MethodPost, http.MethodOptions).Headers(versionHeader, v1)
	router.HandleFunc("/stage/{id}", updateStageHandler(deps)).Methods(http.MethodPut).Headers(versionHeader, v1)
	router.HandleFunc("/stage/{id}", showStageHandler(deps)).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/stage/{id}", deleteStageHandler(deps)).Methods(http.MethodDelete).Headers(versionHeader, v1)
	router.HandleFunc("/stages/{template_id}", listStagesHandler(deps)).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/stage", createStageHandler(deps)).Methods(http.MethodPost).Headers(versionHeader, v1)
	router.HandleFunc("/step/{id}", updateStepHandler(deps)).Methods(http.MethodPut).Headers(versionHeader, v1)
	router.HandleFunc("/step/{id}", showStepHandler(deps)).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/step/{id}", deleteStepHandler(deps)).Methods(http.MethodDelete).Headers(versionHeader, v1)
	router.HandleFunc("/steps/{stage_id}", listStepsHandler(deps)).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/step", createStepHandler(deps)).Methods(http.MethodPost).Headers(versionHeader, v1)
	router.HandleFunc("/targets/{template_id}", listTempTargetsHandler(deps)).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/target/{template_id}", updateTempTargetsHandler(deps)).Methods(http.MethodPost).Headers(versionHeader, v1)
	return
}

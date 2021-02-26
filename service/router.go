package service

import (
	"fmt"
	"net/http"

	"mylab/cpagent/config"

	"github.com/gorilla/mux"
	logger "github.com/sirupsen/logrus"
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
	logger.WithField("v1", v1).Info("Accept Header")
	router.HandleFunc("/targets", listTargetHandler(deps)).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/templates/{id}", updateTemplateHandler(deps)).Methods(http.MethodPut).Headers(versionHeader, v1)
	router.HandleFunc("/templates/{id}", showTemplateHandler(deps)).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/templates/{id}", deleteTemplateHandler(deps)).Methods(http.MethodDelete).Headers(versionHeader, v1)
	router.HandleFunc("/templates/{id}/publish", publishTemplateHandler(deps)).Methods(http.MethodPut).Headers(versionHeader, v1)
	router.HandleFunc("/templates", listTemplateHandler(deps)).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/publishtemplates", listPublishedTemplateHandler(deps)).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/templates", createTemplateHandler(deps)).Methods(http.MethodPost, http.MethodOptions).Headers(versionHeader, v1)
	router.HandleFunc("/stages/{id}", updateStageHandler(deps)).Methods(http.MethodPut).Headers(versionHeader, v1)
	router.HandleFunc("/stages/{id}", showStageHandler(deps)).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/stages/{id}", deleteStageHandler(deps)).Methods(http.MethodDelete).Headers(versionHeader, v1)
	router.HandleFunc("/templates/{template_id}/stages", listStagesHandler(deps)).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/steps/{id}", updateStepHandler(deps)).Methods(http.MethodPut).Headers(versionHeader, v1)
	router.HandleFunc("/steps/{id}", showStepHandler(deps)).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/steps/{id}", deleteStepHandler(deps)).Methods(http.MethodDelete).Headers(versionHeader, v1)
	router.HandleFunc("/stages/{stage_id}/steps", listStepsHandler(deps)).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/steps", createStepHandler(deps)).Methods(http.MethodPost).Headers(versionHeader, v1)
	router.HandleFunc("/templates/{template_id}/targets", listTempTargetsHandler(deps)).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/templates/{template_id}/targets", updateTempTargetsHandler(deps)).Methods(http.MethodPost).Headers(versionHeader, v1)
	router.HandleFunc("/experiments/{id}", showExperimentHandler(deps)).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/experiments", listExperimentHandler(deps)).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/experiments", createExperimentHandler(deps)).Methods(http.MethodPost, http.MethodOptions).Headers(versionHeader, v1)
	router.HandleFunc("/experiments/{experiment_id}/targets", listExpTempTargetsHandler(deps)).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/experiments/{experiment_id}/targets", updateExpTempTargetsHandler(deps)).Methods(http.MethodPost).Headers(versionHeader, v1)
	router.HandleFunc("/samples/", findSamplesHandler(deps)).Queries("text", "{text}").Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/experiments/{experiment_id}/wells", listWellsHandler(deps)).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/experiments/{experiment_id}/wells", upsertWellHandler(deps)).Methods(http.MethodPost).Headers(versionHeader, v1)
	router.HandleFunc("/wells/{id}", showWellHandler(deps)).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/wells/{id}", deleteWellHandler(deps)).Methods(http.MethodDelete).Headers(versionHeader, v1)
	router.HandleFunc("/experiments/{experiment_id}/run", runExperimentHandler(deps)).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/monitor", wsHandler(deps))
	router.HandleFunc("/experiments/{experiment_id}/stop", stopExperimentHandler(deps)).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/activewells", listActiveWellsHandler()).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/experiments/{id}/emission", getResultHandler(deps)).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/experiments/{id}/temperature", getTemperatureHandler(deps)).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/users/{username}/validate", validateUserHandler(deps)).Methods(http.MethodPost, http.MethodOptions).Headers(versionHeader, v1)
	router.HandleFunc("/motor", createMotorHandler(deps)).Methods(http.MethodPost, http.MethodOptions).Headers(versionHeader, v1)
	router.HandleFunc("/consumabledistance", createConsumableDistanceHandler(deps)).Methods(http.MethodPost, http.MethodOptions).Headers(versionHeader, v1)
	router.HandleFunc("/tiptube", createTipTubeHandler(deps)).Methods(http.MethodPost, http.MethodOptions).Headers(versionHeader, v1)
	router.HandleFunc("/homing/{deck:[A-B]?}", homingHandler(deps)).Methods(http.MethodGet)
	router.HandleFunc("/manual", manualHandler(deps)).Methods(http.MethodPost)
	router.HandleFunc("/pause/{deck:[A-B]?}", pauseHandler(deps)).Methods(http.MethodGet)
	router.HandleFunc("/resume/{deck:[A-B]?}", resumeHandler(deps)).Methods(http.MethodGet)
	router.HandleFunc("/abort/{deck:[A-B]?}", abortHandler(deps)).Methods(http.MethodGet)
	router.HandleFunc("/piercing", createPiercingHandler(deps)).Methods(http.MethodPost)
	router.HandleFunc("/piercing", listPiercingHandler(deps)).Methods(http.MethodGet)
	router.HandleFunc("/piercing/{id}", showPiercingHandler(deps)).Methods(http.MethodGet)
	router.HandleFunc("/piercing/{id}", deletePiercingHandler(deps)).Methods(http.MethodDelete)
	router.HandleFunc("/piercing/{id}", updatePiercingHandler(deps)).Methods(http.MethodPut)
	router.HandleFunc("/aspireDispense", createAspireDispenseHandler(deps)).Methods(http.MethodPost)
	router.HandleFunc("/aspireDispense", listAspireDispenseHandler(deps)).Methods(http.MethodGet)
	router.HandleFunc("/aspireDispense/{id}", showAspireDispenseHandler(deps)).Methods(http.MethodGet)
	router.HandleFunc("/aspireDispense/{id}", deleteAspireDispenseHandler(deps)).Methods(http.MethodDelete)
	router.HandleFunc("/aspireDispense/{id}", updateAspireDispenseHandler(deps)).Methods(http.MethodPut)
	router.HandleFunc("/recipe", createRecipeHandler(deps)).Methods(http.MethodPost)
	router.HandleFunc("/recipe", listRecipeHandler(deps)).Methods(http.MethodGet)
	router.HandleFunc("/recipe/{id}", showRecipeHandler(deps)).Methods(http.MethodGet)
	router.HandleFunc("/recipe/{id}", deleteRecipeHandler(deps)).Methods(http.MethodDelete)
	router.HandleFunc("/recipe/{id}", updateRecipeHandler(deps)).Methods(http.MethodPut)
	router.HandleFunc("/process", createProcessHandler(deps)).Methods(http.MethodPost)
	router.HandleFunc("/process/{id}", listProcessHandler(deps)).Methods(http.MethodGet)
	router.HandleFunc("/process/{id}", showProcessHandler(deps)).Methods(http.MethodGet)
	router.HandleFunc("/process/{id}", deleteProcessHandler(deps)).Methods(http.MethodDelete)
	router.HandleFunc("/process/{id}", updateProcessHandler(deps)).Methods(http.MethodPut)
	router.HandleFunc("/run/{id}/{deck:[A-B]?}", runRecipeHandler(deps)).Methods(http.MethodGet)
	return
}

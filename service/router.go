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
	router.HandleFunc("/users", createUserHandler(deps)).Methods(http.MethodPost, http.MethodOptions).Headers(versionHeader, v1)
	router.HandleFunc("/motor", createMotorHandler(deps)).Methods(http.MethodPost, http.MethodOptions).Headers(versionHeader, v1)
	router.HandleFunc("/consumabledistance", createConsumableDistanceHandler(deps)).Methods(http.MethodPost, http.MethodOptions).Headers(versionHeader, v1)
	router.HandleFunc("/tiptube", createTipTubeHandler(deps)).Methods(http.MethodPost, http.MethodOptions).Headers(versionHeader, v1)
	router.HandleFunc("/homing/{deck:[A-B]?}", homingHandler(deps)).Methods(http.MethodGet)
	router.HandleFunc("/manual", manualHandler(deps)).Methods(http.MethodPost)
	router.HandleFunc("/pause/{deck:[A-B]}", pauseHandler(deps)).Methods(http.MethodGet)
	router.HandleFunc("/resume/{deck:[A-B]}", resumeHandler(deps)).Methods(http.MethodGet)
	router.HandleFunc("/abort/{deck:[A-B]}", abortHandler(deps)).Methods(http.MethodGet)
	router.HandleFunc("/piercing", createPiercingHandler(deps)).Methods(http.MethodPost)
	router.HandleFunc("/piercing/{id}", showPiercingHandler(deps)).Methods(http.MethodGet)
	router.HandleFunc("/piercing/{id}", updatePiercingHandler(deps)).Methods(http.MethodPut)
	router.HandleFunc("/aspireDispense", createAspireDispenseHandler(deps)).Methods(http.MethodPost)
	router.HandleFunc("/aspireDispense/{id}", showAspireDispenseHandler(deps)).Methods(http.MethodGet)
	router.HandleFunc("/aspireDispense/{id}", updateAspireDispenseHandler(deps)).Methods(http.MethodPut)
	router.HandleFunc("/recipes", createRecipeHandler(deps)).Methods(http.MethodPost)
	router.HandleFunc("/recipes", listRecipesHandler(deps)).Methods(http.MethodGet)
	router.HandleFunc("/recipes/{id}", showRecipeHandler(deps)).Methods(http.MethodGet)
	router.HandleFunc("/recipes/{id}", deleteRecipeHandler(deps)).Methods(http.MethodDelete)
	router.HandleFunc("/recipes/{id}", updateRecipeHandler(deps)).Methods(http.MethodPut)
	router.HandleFunc("/recipes/{id}/publish", publishRecipeHandler(deps)).Methods(http.MethodPost)
	router.HandleFunc("/processes", createProcessHandler(deps)).Methods(http.MethodPost)
	router.HandleFunc("/processes/{id}", listProcessesHandler(deps)).Methods(http.MethodGet)
	router.HandleFunc("/processes/{id}", showProcessHandler(deps)).Methods(http.MethodGet)
	router.HandleFunc("/processes/{id}", deleteProcessHandler(deps)).Methods(http.MethodDelete)
	router.HandleFunc("/processes/{id}", updateProcessHandler(deps)).Methods(http.MethodPut)
	router.HandleFunc("/run/{id}/{deck:[A-B]}", runRecipeHandler(deps)).Methods(http.MethodGet)
	router.HandleFunc("/discard-box/cleanup/{deck:[A-B]}", discardBoxCleanupHandler(deps)).Methods(http.MethodGet)
	router.HandleFunc("/restore-deck/{deck:[A-B]}", restoreDeckHandler(deps)).Methods(http.MethodGet)
	router.HandleFunc("/uv/{time}/{deck:[A-B]}", uvLightHandler(deps)).Methods(http.MethodGet)
	router.HandleFunc("/discard-tip-and-home/{discard}/{deck:[A-B]}", discardAndHomeHandler(deps)).Methods(http.MethodGet)
	router.HandleFunc("/shaking", createShakingHandler(deps)).Methods(http.MethodPost)
	router.HandleFunc("/shaking/{id}", showShakingHandler(deps)).Methods(http.MethodGet)
	router.HandleFunc("/shaking/{id}", updateShakingHandler(deps)).Methods(http.MethodPut)
	router.HandleFunc("/attach-detach", createAttachDetachHandler(deps)).Methods(http.MethodPost)
	router.HandleFunc("/attach-detach/{id}", showAttachDetachHandler(deps)).Methods(http.MethodGet)
	router.HandleFunc("/attach-detach/{id}", updateAttachDetachHandler(deps)).Methods(http.MethodPut)
	router.HandleFunc("/tip-docking", createTipDockHandler(deps)).Methods(http.MethodPost)
	router.HandleFunc("/tip-docking/{id}", showTipDockHandler(deps)).Methods(http.MethodGet)
	router.HandleFunc("/tip-docking/{id}", updateTipDockHandler(deps)).Methods(http.MethodPut)
	router.HandleFunc("/heating", createHeatingHandler(deps)).Methods(http.MethodPost)
	router.HandleFunc("/heating/{id}", showHeatingHandler(deps)).Methods(http.MethodGet)
	router.HandleFunc("/heating/{id}", updateHeatingHandler(deps)).Methods(http.MethodPut)
	router.HandleFunc("/tips-tubes/{tiptube:[a-z]*}", listTipsTubesHandler(deps)).Methods(http.MethodGet)
	router.HandleFunc("/cartridges/list", listCartridgesHandler(deps)).Methods(http.MethodGet)
	router.HandleFunc("/delay", createDelayHandler(deps)).Methods(http.MethodPost)
	router.HandleFunc("/delay/{id}", showDelayHandler(deps)).Methods(http.MethodGet)
	router.HandleFunc("/delay/{id}", updateDelayHandler(deps)).Methods(http.MethodPut)
	router.HandleFunc("/tip-operation", createTipOperationHandler(deps)).Methods(http.MethodPost)
	router.HandleFunc("/tip-operation/{id}", showTipOperationHandler(deps)).Methods(http.MethodGet)
	router.HandleFunc("/tip-operation/{id}", updateTipOperationHandler(deps)).Methods(http.MethodPut)

	return
}

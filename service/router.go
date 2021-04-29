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
	router.HandleFunc("/users/validate/{deck:[A-B]}", validateUserHandler(deps)).Methods(http.MethodPost, http.MethodOptions).Headers(versionHeader, v1)
	router.HandleFunc("/motor", createMotorHandler(deps)).Methods(http.MethodPost, http.MethodOptions).Headers(versionHeader, v1)
	router.HandleFunc("/consumabledistance", authenticateAdmin(createConsumableDistanceHandler(deps), deps)).Methods(http.MethodPost, http.MethodOptions).Headers(versionHeader, v1)
	router.HandleFunc("/tiptube", authenticateAdmin(createTipTubeHandler(deps), deps)).Methods(http.MethodPost, http.MethodOptions).Headers(versionHeader, v1)
	router.HandleFunc("/homing/{deck:[A-B]?}", homingHandler(deps)).Methods(http.MethodGet)
	router.HandleFunc("/manual", manualHandler(deps)).Methods(http.MethodPost)
	router.HandleFunc("/pause/{deck:[A-B]}", pauseHandler(deps)).Methods(http.MethodGet)
	router.HandleFunc("/resume/{deck:[A-B]}", resumeHandler(deps)).Methods(http.MethodGet)
	router.HandleFunc("/abort/{deck:[A-B]}", abortHandler(deps)).Methods(http.MethodGet)
	router.HandleFunc("/piercing", authenticateAdmin(createPiercingHandler(deps), deps)).Methods(http.MethodPost)
	router.HandleFunc("/piercing/{id}", authenticate(showPiercingHandler(deps), deps)).Methods(http.MethodGet)
	router.HandleFunc("/piercing/{id}", authenticateAdmin(updatePiercingHandler(deps), deps)).Methods(http.MethodPut)
	router.HandleFunc("/aspireDispense", authenticateAdmin(createAspireDispenseHandler(deps), deps)).Methods(http.MethodPost)
	router.HandleFunc("/aspireDispense/{id}", authenticate(showAspireDispenseHandler(deps), deps)).Methods(http.MethodGet)
	router.HandleFunc("/aspireDispense/{id}", authenticateAdmin(updateAspireDispenseHandler(deps), deps)).Methods(http.MethodPut)
	router.HandleFunc("/recipes", authenticateAdmin(createRecipeHandler(deps), deps)).Methods(http.MethodPost)
	router.HandleFunc("/recipes", authenticate(listRecipesHandler(deps), deps)).Methods(http.MethodGet)
	router.HandleFunc("/recipes/{id}", authenticate(showRecipeHandler(deps), deps)).Methods(http.MethodGet)
	router.HandleFunc("/recipes/{id}", authenticateAdmin(deleteRecipeHandler(deps), deps)).Methods(http.MethodDelete)
	router.HandleFunc("/recipes/{id}", authenticateAdmin(updateRecipeHandler(deps), deps)).Methods(http.MethodPut)
	router.HandleFunc("/recipes/{id}/publish", publishRecipeHandler(deps)).Methods(http.MethodPost)
	router.HandleFunc("/processes", createProcessHandler(deps)).Methods(http.MethodPost)
	router.HandleFunc("/recipe/{id}/processes", listProcessesHandler(deps)).Methods(http.MethodGet)
	router.HandleFunc("/processes/{id}", showProcessHandler(deps)).Methods(http.MethodGet)
	router.HandleFunc("/processes/{id}", deleteProcessHandler(deps)).Methods(http.MethodDelete)
	router.HandleFunc("/processes/{id}", updateProcessHandler(deps)).Methods(http.MethodPut)
	router.HandleFunc("/run/{id}/{deck:[A-B]}", runRecipeHandler(deps, false)).Methods(http.MethodGet)
	router.HandleFunc("/step-run/{id}/{deck:[A-B]}", runRecipeHandler(deps, true)).Methods(http.MethodGet)
	router.HandleFunc("/run-next-step/{deck:[A-B]}", runNextStepHandler(deps)).Methods(http.MethodGet)
	router.HandleFunc("/discard-box/cleanup/{deck:[A-B]}", discardBoxCleanupHandler(deps)).Methods(http.MethodGet)
	router.HandleFunc("/restore-deck/{deck:[A-B]}", restoreDeckHandler(deps)).Methods(http.MethodGet)
	router.HandleFunc("/uv/{time}/{deck:[A-B]}", uvLightHandler(deps)).Methods(http.MethodGet)
	router.HandleFunc("/discard-tip-and-home/{discard}/{deck:[A-B]}", discardAndHomeHandler(deps)).Methods(http.MethodGet)
	router.HandleFunc("/shaking", authenticateAdmin(createShakingHandler(deps), deps)).Methods(http.MethodPost)
	router.HandleFunc("/shaking/{id}", showShakingHandler(deps)).Methods(http.MethodGet)
	router.HandleFunc("/shaking/{id}", authenticateAdmin(updateShakingHandler(deps), deps)).Methods(http.MethodPut)
	router.HandleFunc("/attach-detach", authenticateAdmin(createAttachDetachHandler(deps), deps)).Methods(http.MethodPost)
	router.HandleFunc("/attach-detach/{id}", authenticate(showAttachDetachHandler(deps), deps)).Methods(http.MethodGet)
	router.HandleFunc("/attach-detach/{id}", authenticateAdmin(updateAttachDetachHandler(deps), deps)).Methods(http.MethodPut)
	router.HandleFunc("/tip-docking", authenticateAdmin(createTipDockHandler(deps), deps)).Methods(http.MethodPost)
	router.HandleFunc("/tip-docking/{id}", authenticate(showTipDockHandler(deps), deps)).Methods(http.MethodGet)
	router.HandleFunc("/tip-docking/{id}", authenticateAdmin(updateTipDockHandler(deps), deps)).Methods(http.MethodPut)
	router.HandleFunc("/heating", authenticateAdmin(createHeatingHandler(deps), deps)).Methods(http.MethodPost)
	router.HandleFunc("/heating/{id}", authenticate(showHeatingHandler(deps), deps)).Methods(http.MethodGet)
	router.HandleFunc("/heating/{id}", authenticateAdmin(updateHeatingHandler(deps), deps)).Methods(http.MethodPut)
	router.HandleFunc("/delay", authenticateAdmin(createDelayHandler(deps), deps)).Methods(http.MethodPost)
	router.HandleFunc("/delay/{id}", authenticate(showDelayHandler(deps), deps)).Methods(http.MethodGet)
	router.HandleFunc("/delay/{id}", authenticateAdmin(updateDelayHandler(deps), deps)).Methods(http.MethodPut)
	router.HandleFunc("/tip-operation", authenticateAdmin(createTipOperationHandler(deps), deps)).Methods(http.MethodPost)
	router.HandleFunc("/tip-operation/{id}", authenticate(showTipOperationHandler(deps), deps)).Methods(http.MethodGet)
	router.HandleFunc("/tip-operation/{id}", authenticateAdmin(updateTipOperationHandler(deps), deps)).Methods(http.MethodPut)
	router.HandleFunc("/tips-tubes/{tiptube:[a-z]*}", listTipsTubesHandler(deps)).Methods(http.MethodGet)
	router.HandleFunc("/cartridges", listCartridgesHandler(deps)).Methods(http.MethodGet)

	return
}

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

	//Websocket router
	router.HandleFunc("/monitor", wsHandler(deps))

	router.HandleFunc("/experiments/{experiment_id}/stop", stopExperimentHandler(deps)).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/activewells", listActiveWellsHandler()).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/experiments/{id}/emission", getResultHandler(deps)).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/experiments/{id}/temperature", getTemperatureHandler(deps)).Methods(http.MethodGet).Headers(versionHeader, v1)
	//usercreate
	router.HandleFunc("/users", authenticate(createUserHandler(deps), deps, supervisor, admin)).Methods(http.MethodPost, http.MethodOptions).Headers(versionHeader, v1)
	//userlogin
	router.HandleFunc("/login/{deck:[A-B]?}", validateUserHandler(deps)).Methods(http.MethodPost, http.MethodOptions).Headers(versionHeader, v1)
	//userlogout
	router.HandleFunc("/logout/{deck:[A-B]?}", authenticate(logoutUserHandler(deps), deps)).Methods(http.MethodDelete, http.MethodOptions).Headers(versionHeader, v1)

	router.HandleFunc("/motor", authenticate(createMotorHandler(deps), deps, admin, engineer)).Methods(http.MethodPost, http.MethodOptions).Headers(versionHeader, v1)
	router.HandleFunc("/consumable-distance", authenticate(createConsumableDistanceHandler(deps), deps, engineer)).Methods(http.MethodPost, http.MethodOptions).Headers(versionHeader, v1)
	router.HandleFunc("/tiptube", authenticate(createTipTubeHandler(deps), deps, engineer)).Methods(http.MethodPost, http.MethodOptions).Headers(versionHeader, v1)

	//homing
	router.HandleFunc("/homing/{deck:[A-B]?}", homingHandler(deps)).Methods(http.MethodGet).Headers(versionHeader, v1)
	//manual
	router.HandleFunc("/manual", authenticate(manualHandler(deps), deps, engineer)).Methods(http.MethodPost).Headers(versionHeader, v1)
	router.HandleFunc("/pause/{deck:[A-B]}", authenticate(pauseHandler(deps), deps)).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/resume/{deck:[A-B]}", authenticate(resumeHandler(deps), deps)).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/abort/{deck:[A-B]}", authenticate(abortHandler(deps), deps)).Methods(http.MethodGet).Headers(versionHeader, v1)

	//processes CRUD
	router.HandleFunc("/piercing/{recipe_id}", authenticate(createPiercingHandler(deps), deps, admin, engineer)).Methods(http.MethodPost).Headers(versionHeader, v1)
	router.HandleFunc("/piercing/{id}", authenticate(showPiercingHandler(deps), deps)).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/piercing/{id}", authenticate(updatePiercingHandler(deps), deps, admin, engineer)).Methods(http.MethodPut).Headers(versionHeader, v1)
	router.HandleFunc("/aspire-dispense/{recipe_id}", authenticate(createAspireDispenseHandler(deps), deps, admin, engineer)).Methods(http.MethodPost).Headers(versionHeader, v1)
	router.HandleFunc("/aspire-dispense/{id}", authenticate(showAspireDispenseHandler(deps), deps)).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/aspire-dispense/{id}", authenticate(updateAspireDispenseHandler(deps), deps, admin, engineer)).Methods(http.MethodPut).Headers(versionHeader, v1)
	router.HandleFunc("/recipes", authenticate(createRecipeHandler(deps), deps, admin, engineer)).Methods(http.MethodPost).Headers(versionHeader, v1)
	router.HandleFunc("/recipes", authenticate(listRecipesHandler(deps), deps)).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/recipes/{id}", authenticate(showRecipeHandler(deps), deps)).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/recipes/{id}", authenticate(deleteRecipeHandler(deps), deps, admin, engineer)).Methods(http.MethodDelete).Headers(versionHeader, v1)
	router.HandleFunc("/recipes/{id}", authenticate(updateRecipeHandler(deps), deps, admin, engineer)).Methods(http.MethodPut).Headers(versionHeader, v1)
	router.HandleFunc("/recipes/{id}/{publish:[a-z]*}", authenticate(publishRecipeHandler(deps), deps, admin)).Methods(http.MethodPost).Headers(versionHeader, v1)
	router.HandleFunc("/processes", authenticate(createProcessHandler(deps), deps, admin, engineer)).Methods(http.MethodPost).Headers(versionHeader, v1)
	router.HandleFunc("/recipe/{id}/processes", authenticate(listProcessesHandler(deps), deps)).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/duplicate-process/{process_id}", authenticate(duplicateProcessHandler(deps), deps, admin, engineer)).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/rearrange-processes/{recipe_id}", authenticate(rearrangeProcessesHandler(deps), deps, admin, engineer)).Methods(http.MethodPost).Headers(versionHeader, v1)
	router.HandleFunc("/processes/{id}", authenticate(showProcessHandler(deps), deps)).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/processes/{id}", authenticate(deleteProcessHandler(deps), deps, admin, engineer)).Methods(http.MethodDelete).Headers(versionHeader, v1)
	router.HandleFunc("/processes/{id}", authenticate(updateProcessHandler(deps), deps, admin)).Methods(http.MethodPut).Headers(versionHeader, v1)
	router.HandleFunc("/run/{id}/{deck:[A-B]}", authenticate(runRecipeHandler(deps, false), deps)).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/step-run/{id}/{deck:[A-B]}", authenticate(runRecipeHandler(deps, true), deps, admin, engineer)).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/run-next-step/{deck:[A-B]}", authenticate(runNextStepHandler(deps), deps, admin, engineer)).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/discard-box/cleanup/{deck:[A-B]}", authenticate(discardBoxCleanupHandler(deps), deps)).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/restore-deck/{deck:[A-B]}", authenticate(restoreDeckHandler(deps), deps)).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/uv/{time}/{deck:[A-B]}", authenticate(uvLightHandler(deps), deps)).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/discard-tip-and-home/{discard}/{deck:[A-B]}", authenticate(discardAndHomeHandler(deps), deps)).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/shaking/{recipe_id}", authenticate(createShakingHandler(deps), deps, admin, engineer)).Methods(http.MethodPost).Headers(versionHeader, v1)
	router.HandleFunc("/shaking/{id}", authenticate(showShakingHandler(deps), deps)).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/shaking/{id}", authenticate(updateShakingHandler(deps), deps, admin, engineer)).Methods(http.MethodPut).Headers(versionHeader, v1)
	router.HandleFunc("/attach-detach/{recipe_id}", authenticate(createAttachDetachHandler(deps), deps, admin, engineer)).Methods(http.MethodPost).Headers(versionHeader, v1)
	router.HandleFunc("/attach-detach/{id}", authenticate(showAttachDetachHandler(deps), deps)).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/attach-detach/{id}", authenticate(updateAttachDetachHandler(deps), deps, admin, engineer)).Methods(http.MethodPut).Headers(versionHeader, v1)
	router.HandleFunc("/tip-docking/{recipe_id}", authenticate(createTipDockHandler(deps), deps, admin, engineer)).Methods(http.MethodPost).Headers(versionHeader, v1)
	router.HandleFunc("/tip-docking/{id}", authenticate(showTipDockHandler(deps), deps)).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/tip-docking/{id}", authenticate(updateTipDockHandler(deps), deps, admin, engineer)).Methods(http.MethodPut).Headers(versionHeader, v1)
	router.HandleFunc("/heating/{recipe_id}", authenticate(createHeatingHandler(deps), deps, admin, engineer)).Methods(http.MethodPost).Headers(versionHeader, v1)
	router.HandleFunc("/heating/{id}", authenticate(showHeatingHandler(deps), deps)).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/heating/{id}", authenticate(updateHeatingHandler(deps), deps, admin, engineer)).Methods(http.MethodPut).Headers(versionHeader, v1)
	router.HandleFunc("/delay/{recipe_id}", authenticate(createDelayHandler(deps), deps, admin, engineer)).Methods(http.MethodPost).Headers(versionHeader, v1)
	router.HandleFunc("/delay/{id}", authenticate(showDelayHandler(deps), deps)).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/delay/{id}", authenticate(updateDelayHandler(deps), deps, admin, engineer)).Methods(http.MethodPut).Headers(versionHeader, v1)
	router.HandleFunc("/tip-operation/{recipe_id}", authenticate(createTipOperationHandler(deps), deps, admin, engineer)).Methods(http.MethodPost).Headers(versionHeader, v1)
	router.HandleFunc("/tip-operation/{id}", authenticate(showTipOperationHandler(deps), deps)).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/tip-operation/{id}", authenticate(updateTipOperationHandler(deps), deps, admin, engineer)).Methods(http.MethodPut).Headers(versionHeader, v1)
	router.HandleFunc("/tips-tubes/{tiptube:[a-z]*}", authenticate(listTipsTubesHandler(deps), deps)).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/cartridges", authenticate(listCartridgesHandler(deps), deps)).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/logout/{deck:[A-B]?}", authenticate(logoutUserHandler(deps), deps)).Methods(http.MethodDelete, http.MethodOptions).Headers(versionHeader, v1)
	router.HandleFunc("/safe-to-upgrade", safeToUpgradeHandler(deps)).Methods(http.MethodGet).Headers(versionHeader, v1)

	return
}

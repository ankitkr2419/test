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
	router.HandleFunc("/targets", authenticate(listTargetHandler(deps), deps, RTPCR)).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/templates/{id}", authenticate(updateTemplateHandler(deps), deps, RTPCR)).Methods(http.MethodPut).Headers(versionHeader, v1)
	router.HandleFunc("/templates/{id}", authenticate(showTemplateHandler(deps), deps, RTPCR)).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/templates/{id}", authenticate(deleteTemplateHandler(deps), deps, RTPCR)).Methods(http.MethodDelete).Headers(versionHeader, v1)
	router.HandleFunc("/templates/{id}/publish", authenticate(publishTemplateHandler(deps), deps, RTPCR)).Methods(http.MethodPut).Headers(versionHeader, v1)
	router.HandleFunc("/templates", authenticate(listTemplateHandler(deps), deps, RTPCR)).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/publishtemplates", authenticate(listPublishedTemplateHandler(deps), deps, RTPCR)).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/templates", authenticate(createTemplateHandler(deps), deps, RTPCR)).Methods(http.MethodPost, http.MethodOptions).Headers(versionHeader, v1)
	router.HandleFunc("/stages/{id}", authenticate(updateStageHandler(deps), deps, RTPCR)).Methods(http.MethodPut).Headers(versionHeader, v1)
	router.HandleFunc("/stages/{id}", authenticate(showStageHandler(deps), deps, RTPCR)).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/stages/{id}", authenticate(deleteStageHandler(deps), deps, RTPCR)).Methods(http.MethodDelete).Headers(versionHeader, v1)
	router.HandleFunc("/templates/{template_id}/stages", authenticate(listStagesHandler(deps), deps, RTPCR)).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/steps/{id}", authenticate(updateStepHandler(deps), deps, RTPCR)).Methods(http.MethodPut).Headers(versionHeader, v1)
	router.HandleFunc("/steps/{id}", authenticate(showStepHandler(deps), deps, RTPCR)).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/steps/{id}", authenticate(deleteStepHandler(deps), deps, RTPCR)).Methods(http.MethodDelete).Headers(versionHeader, v1)
	router.HandleFunc("/stages/{stage_id}/steps", authenticate(listStepsHandler(deps), deps, RTPCR)).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/steps", authenticate(createStepHandler(deps), deps, RTPCR)).Methods(http.MethodPost).Headers(versionHeader, v1)
	router.HandleFunc("/templates/{template_id}/targets", authenticate(listTempTargetsHandler(deps), deps, RTPCR)).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/templates/{template_id}/targets", authenticate(updateTempTargetsHandler(deps), deps, RTPCR)).Methods(http.MethodPost).Headers(versionHeader, v1)
	router.HandleFunc("/experiments/{id}", authenticate(showExperimentHandler(deps), deps, RTPCR)).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/experiments", authenticate(listExperimentHandler(deps), deps, RTPCR)).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/experiments", authenticate(createExperimentHandler(deps), deps, RTPCR)).Methods(http.MethodPost, http.MethodOptions).Headers(versionHeader, v1)
	router.HandleFunc("/experiments/{experiment_id}/targets", authenticate(listExpTempTargetsHandler(deps), deps, RTPCR)).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/experiments/{experiment_id}/targets", authenticate(updateExpTempTargetsHandler(deps), deps, RTPCR)).Methods(http.MethodPost).Headers(versionHeader, v1)
	router.HandleFunc("/experiments/{experiment_id}/wells", authenticate(listWellsHandler(deps), deps, RTPCR)).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/experiments/{experiment_id}/wells", authenticate(upsertWellHandler(deps), deps, RTPCR)).Methods(http.MethodPost).Headers(versionHeader, v1)
	router.HandleFunc("/wells/{id}", authenticate(showWellHandler(deps), deps, RTPCR)).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/wells/{id}", authenticate(deleteWellHandler(deps), deps, RTPCR)).Methods(http.MethodDelete).Headers(versionHeader, v1)
	router.HandleFunc("/experiments/{experiment_id}/run", authenticate(runExperimentHandler(deps), deps, RTPCR)).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/finish/template/{id}", authenticate(finishTemplateHandler(deps), deps, RTPCR, admin)).Methods(http.MethodPut).Headers(versionHeader, v1)
	router.HandleFunc("/finished/templates", authenticate(listFinishedTemplateHandler(deps), deps, RTPCR)).Methods(http.MethodGet).Headers(versionHeader, v1)

	//Websocket router
	router.HandleFunc("/monitor", wsHandler(deps)).Methods(http.MethodGet)

	router.HandleFunc("/experiments/{experiment_id}/stop", authenticate(stopExperimentHandler(deps), deps, RTPCR)).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/activewells", authenticate(listActiveWellsHandler(), deps, RTPCR)).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/experiments/{id}/emission", authenticate(getResultHandler(deps), deps, RTPCR)).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/experiments/{id}/temperature", authenticate(getTemperatureHandler(deps), deps, RTPCR)).Methods(http.MethodGet).Headers(versionHeader, v1)
	//usercreate
	router.HandleFunc("/users", authenticate(createUserHandler(deps), deps, Combined, supervisor, admin)).Methods(http.MethodPost, http.MethodOptions).Headers(versionHeader, v1)
	// update user
	router.HandleFunc("/users/{old_username}", authenticate(updateUserHandler(deps), deps, Combined, supervisor, admin)).Methods(http.MethodPut, http.MethodOptions).Headers(versionHeader, v1)
	//userlogin
	router.HandleFunc("/login/{deck:[A-B]?}", validateUserHandler(deps)).Methods(http.MethodPost, http.MethodOptions).Headers(versionHeader, v1)
	//userlogout
	router.HandleFunc("/logout/{deck:[A-B]?}", authenticate(logoutUserHandler(deps), deps, Combined)).Methods(http.MethodDelete, http.MethodOptions).Headers(versionHeader, v1)

	// configs
	router.HandleFunc("/configs", authenticate(getConfigHandler(deps), deps, Combined, engineer, admin)).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/configs", authenticate(updateConfigHandler(deps), deps, Combined, engineer, admin)).Methods(http.MethodPut, http.MethodOptions).Headers(versionHeader, v1)

	router.HandleFunc("/motor", authenticate(createMotorHandler(deps), deps, Extraction, admin, engineer)).Methods(http.MethodPost, http.MethodOptions).Headers(versionHeader, v1)
	router.HandleFunc("/motor/{id}", authenticate(updateMotorHandler(deps), deps, Extraction, admin, engineer)).Methods(http.MethodPut, http.MethodOptions).Headers(versionHeader, v1)
	router.HandleFunc("/motors", authenticate(listMotorsHandler(deps), deps, Extraction, admin, engineer)).Methods(http.MethodGet, http.MethodOptions).Headers(versionHeader, v1)
	router.HandleFunc("/motor/{id}", authenticate(deleteMotorHandler(deps), deps, Extraction, admin, engineer)).Methods(http.MethodDelete, http.MethodOptions).Headers(versionHeader, v1)

	router.HandleFunc("/consumable-distance", authenticate(createConsumableDistanceHandler(deps), deps, Extraction, engineer, admin)).Methods(http.MethodPost, http.MethodOptions).Headers(versionHeader, v1)
	router.HandleFunc("/tiptube", authenticate(createTipTubeHandler(deps), deps, Extraction, engineer, admin)).Methods(http.MethodPost, http.MethodOptions).Headers(versionHeader, v1)
	router.HandleFunc("/tips-tubes/{id}", authenticate(deleteTipTubeHandler(deps), deps, Extraction, admin, engineer)).Methods(http.MethodDelete, http.MethodOptions).Headers(versionHeader, v1)

	//homing
	router.HandleFunc("/homing/{deck:[A-B]?}", authenticate(homingHandler(deps), deps, Extraction)).Methods(http.MethodGet).Headers(versionHeader, v1)
	//manual
	router.HandleFunc("/manual", authenticate(manualHandler(deps), deps, Extraction, engineer, admin)).Methods(http.MethodPost).Headers(versionHeader, v1)
	router.HandleFunc("/pause/{deck:[A-B]}", authenticate(pauseHandler(deps), deps, Extraction)).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/resume/{deck:[A-B]}", authenticate(resumeHandler(deps), deps, Extraction)).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/abort/{deck:[A-B]}", authenticate(abortHandler(deps), deps, Extraction)).Methods(http.MethodGet).Headers(versionHeader, v1)

	//processes CRUD
	router.HandleFunc("/piercing/{recipe_id}", authenticate(createPiercingHandler(deps), deps, Extraction, admin, engineer)).Methods(http.MethodPost).Headers(versionHeader, v1)
	router.HandleFunc("/piercing/{id}", authenticate(showPiercingHandler(deps), deps, Extraction)).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/piercing/{id}", authenticate(updatePiercingHandler(deps), deps, Extraction, admin, engineer)).Methods(http.MethodPut).Headers(versionHeader, v1)
	router.HandleFunc("/aspire-dispense/{recipe_id}", authenticate(createAspireDispenseHandler(deps), deps, Extraction, admin, engineer)).Methods(http.MethodPost).Headers(versionHeader, v1)
	router.HandleFunc("/aspire-dispense/{id}", authenticate(showAspireDispenseHandler(deps), deps, Extraction)).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/aspire-dispense/{id}", authenticate(updateAspireDispenseHandler(deps), deps, Extraction, admin, engineer)).Methods(http.MethodPut).Headers(versionHeader, v1)
	router.HandleFunc("/recipes", authenticate(createRecipeHandler(deps), deps, Extraction, admin, engineer)).Methods(http.MethodPost).Headers(versionHeader, v1)
	router.HandleFunc("/recipes", authenticate(listRecipesHandler(deps), deps, Extraction)).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/recipes/{id}", authenticate(showRecipeHandler(deps), deps, Extraction)).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/recipes/{id}", authenticate(deleteRecipeHandler(deps), deps, Extraction, admin, engineer)).Methods(http.MethodDelete).Headers(versionHeader, v1)
	router.HandleFunc("/recipes/{id}", authenticate(updateRecipeHandler(deps), deps, Extraction, admin, engineer)).Methods(http.MethodPut).Headers(versionHeader, v1)
	router.HandleFunc("/recipes/{id}/{publish:[a-z]*}", authenticate(publishRecipeHandler(deps), deps, Extraction, admin)).Methods(http.MethodPost).Headers(versionHeader, v1)
	router.HandleFunc("/processes", authenticate(createProcessHandler(deps), deps, Extraction, admin, engineer)).Methods(http.MethodPost).Headers(versionHeader, v1)
	router.HandleFunc("/recipe/{id}/processes", authenticate(listProcessesHandler(deps), deps, Extraction)).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/duplicate-process/{process_id}", authenticate(duplicateProcessHandler(deps), deps, Extraction, admin, engineer)).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/rearrange-processes/{recipe_id}", authenticate(rearrangeProcessesHandler(deps), deps, Extraction, admin, engineer)).Methods(http.MethodPost).Headers(versionHeader, v1)
	router.HandleFunc("/processes/{id}", authenticate(showProcessHandler(deps), deps, Extraction)).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/processes/{id}", authenticate(deleteProcessHandler(deps), deps, Extraction, admin, engineer)).Methods(http.MethodDelete).Headers(versionHeader, v1)
	router.HandleFunc("/processes/{id}", authenticate(updateProcessHandler(deps), deps, Extraction, admin)).Methods(http.MethodPut).Headers(versionHeader, v1)
	router.HandleFunc("/run/{id}/{deck:[A-B]}", authenticate(runRecipeHandler(deps, false), deps, Extraction)).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/step-run/{id}/{deck:[A-B]}", authenticate(runRecipeHandler(deps, true), deps, Extraction, admin, engineer)).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/run-next-step/{deck:[A-B]}", authenticate(runNextStepHandler(deps), deps, Extraction, admin, engineer)).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/discard-box/cleanup/{deck:[A-B]}", authenticate(discardBoxCleanupHandler(deps), deps, Extraction)).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/restore-deck/{deck:[A-B]}", authenticate(restoreDeckHandler(deps), deps, Extraction)).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/uv/{time}/{deck:[A-B]}", authenticate(uvLightHandler(deps), deps, Extraction)).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/discard-tip-and-home/{discard}/{deck:[A-B]}", authenticate(discardAndHomeHandler(deps), deps, Extraction)).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/shaking/{recipe_id}", authenticate(createShakingHandler(deps), deps, Extraction, admin, engineer)).Methods(http.MethodPost).Headers(versionHeader, v1)
	router.HandleFunc("/shaking/{id}", authenticate(showShakingHandler(deps), deps, Extraction)).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/shaking/{id}", authenticate(updateShakingHandler(deps), deps, Extraction, admin, engineer)).Methods(http.MethodPut).Headers(versionHeader, v1)
	router.HandleFunc("/attach-detach/{recipe_id}", authenticate(createAttachDetachHandler(deps), deps, Extraction, admin, engineer)).Methods(http.MethodPost).Headers(versionHeader, v1)
	router.HandleFunc("/attach-detach/{id}", authenticate(showAttachDetachHandler(deps), deps, Extraction)).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/attach-detach/{id}", authenticate(updateAttachDetachHandler(deps), deps, Extraction, admin, engineer)).Methods(http.MethodPut).Headers(versionHeader, v1)
	router.HandleFunc("/tip-docking/{recipe_id}", authenticate(createTipDockHandler(deps), deps, Extraction, admin, engineer)).Methods(http.MethodPost).Headers(versionHeader, v1)
	router.HandleFunc("/tip-docking/{id}", authenticate(showTipDockHandler(deps), deps, Extraction)).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/tip-docking/{id}", authenticate(updateTipDockHandler(deps), deps, Extraction, admin, engineer)).Methods(http.MethodPut).Headers(versionHeader, v1)
	router.HandleFunc("/heating/{recipe_id}", authenticate(createHeatingHandler(deps), deps, Extraction, admin, engineer)).Methods(http.MethodPost).Headers(versionHeader, v1)
	router.HandleFunc("/heating/{id}", authenticate(showHeatingHandler(deps), deps, Extraction)).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/heating/{id}", authenticate(updateHeatingHandler(deps), deps, Extraction, admin, engineer)).Methods(http.MethodPut).Headers(versionHeader, v1)
	router.HandleFunc("/delay/{recipe_id}", authenticate(createDelayHandler(deps), deps, Extraction, admin, engineer)).Methods(http.MethodPost).Headers(versionHeader, v1)
	router.HandleFunc("/delay/{id}", authenticate(showDelayHandler(deps), deps, Extraction)).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/delay/{id}", authenticate(updateDelayHandler(deps), deps, Extraction, admin, engineer)).Methods(http.MethodPut).Headers(versionHeader, v1)
	router.HandleFunc("/tip-operation/{recipe_id}", authenticate(createTipOperationHandler(deps), deps, Extraction, admin, engineer)).Methods(http.MethodPost).Headers(versionHeader, v1)
	router.HandleFunc("/tip-operation/{id}", authenticate(showTipOperationHandler(deps), deps, Extraction)).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/tip-operation/{id}", authenticate(updateTipOperationHandler(deps), deps, Extraction, admin, engineer)).Methods(http.MethodPut).Headers(versionHeader, v1)
	router.HandleFunc("/tips-tubes/{tiptube:[a-z]*}", authenticate(listTipsTubesHandler(deps), deps, Extraction)).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/tips-tubes/{tiptube:[a-z]*}/{position:[1-9]+}", authenticate(listTipsTubesPositionHandler(deps), deps, Extraction)).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/cartridges", authenticate(listCartridgesHandler(deps), deps, Extraction)).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/cartridge", authenticate(createCartridgeHandler(deps), deps, Extraction, admin, engineer)).Methods(http.MethodPost, http.MethodOptions).Headers(versionHeader, v1)
	router.HandleFunc("/cartridge/{id}", authenticate(deleteCartridgeHandler(deps), deps, Extraction, admin, engineer)).Methods(http.MethodDelete, http.MethodOptions).Headers(versionHeader, v1)

	router.HandleFunc("/safe-to-upgrade", safeToUpgradeHandler(deps)).Methods(http.MethodGet).Headers(versionHeader, v1)
	//TODO: allow only for engineer
	router.HandleFunc("/pid-calibration/{deck:[A-B]}", pidCalibrationHandler(deps)).Methods(http.MethodGet).Headers(versionHeader, v1)
	// shaker and heater for engineer
	router.HandleFunc("/shaker/{deck:[A-B]}", shakerHandler(deps)).Methods(http.MethodPost, http.MethodOptions).Headers(versionHeader, v1)
	router.HandleFunc("/heater/{deck:[A-B]}", heaterHandler(deps)).Methods(http.MethodPost, http.MethodOptions).Headers(versionHeader, v1)

	router.HandleFunc("/app-info", appInfoHandler(deps)).Methods(http.MethodGet).Headers(versionHeader, v1)

	//rt-pcr funcs
	router.HandleFunc("/rt-pcr/homing", authenticate(rtpcrHomingHandler(deps), deps, RTPCR)).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/rt-pcr/reset", authenticate(rtpcrResetHandler(deps), deps, RTPCR)).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/rt-pcr/cycle", authenticate(rtpcrStartCycleHandler(deps), deps, RTPCR)).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/rt-pcr/monitor", authenticate(rtpcrMonitorHandler(deps), deps, RTPCR)).Methods(http.MethodGet).Headers(versionHeader, v1)

	// tec funcs
	router.HandleFunc("/tec/set-temp-and-ramp", authenticate(setTempAndRampHandler(deps), deps, RTPCR)).Methods(http.MethodPost).Headers(versionHeader, v1)
	router.HandleFunc("/tec/run", authenticate(runProfileHandler(deps), deps, RTPCR)).Methods(http.MethodPost).Headers(versionHeader, v1)
	router.HandleFunc("/tec/auto-tune", authenticate(autoTuneHandler(deps), deps, RTPCR)).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/tec/reset-device", authenticate(resetDeviceHandler(deps), deps, RTPCR)).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/tec/run", authenticate(runTECHandler(deps), deps, RTPCR)).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/tec/get-all", authenticate(getAllTECHandler(deps), deps, RTPCR)).Methods(http.MethodGet).Headers(versionHeader, v1)

	return
}

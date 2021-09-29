package responses

import "fmt"

var (
	Pong = "pong"

	AppInfoFetch     = "application information is being fetched"
	AppInfoRequested = "application information was requested"

	PiercingCreateSuccess        = "piercing record created successfully"
	PiercingFetchSuccess         = "piercing record fetched successfully"
	PiercingUpdateSuccess        = "piercing record updated successfully"
	PiercingInitialisedState     = "piercing initialised"
	PiercingCompletedState       = "piercing completed"
	PiercingListInitialisedState = "piercing list initialised"
	PiercingListCompletedState   = "piercing list completed"

	AspireDispenseCreateSuccess        = "aspire dispense record created successfully"
	AspireDispenseFetchSuccess         = "aspire dispense record fetched successfully"
	AspireDispenseUpdateSuccess        = "aspire dispense record updated successfully"
	AspireDispenseInitialisedState     = "aspire dispense initialised"
	AspireDispenseCompletedState       = "aspire dispense completed"
	AspireDispenseListInitialisedState = "aspire dispense list initialised"
	AspireDispenseListCompletedState   = "aspire dispense list completed"

	DelayCreateSuccess        = "delay record created successfully"
	DelayFetchSuccess         = "delay record fetched successfully"
	DelayUpdateSuccess        = "delay record updated successfully"
	DelayInitialisedState     = "delay initialised"
	DelayCompletedState       = "delay completed"
	DelayListInitialisedState = "delay list initialised"
	DelayListCompletedState   = "delay list completed"

	AttachDetachCreateSuccess    = "attach detach record created successfully"
	AttachDetachFetchSuccess     = "attach detach record fetched successfully"
	AttachDetachUpdateSuccess    = "attach detach record updated successfully"
	AttachDetachInitialisedState = "attach detach initialised"
	AttachDetachCompletedState   = "attach detach completed"

	ShakingCreateSuccess    = "shaking record created successfully"
	ShakingFetchSuccess     = "shaking record fetched successfully"
	ShakingUpdateSuccess    = "shaking record updated successfully"
	ShakingInitialisedState = "shaking initialised"
	ShakingCompletedState   = "shaking completed"

	HeatingCreateSuccess    = "heating record created successfully"
	HeatingFetchSuccess     = "heating record fetched successfully"
	HeatingUpdateSuccess    = "heating record updated successfully"
	HeatingInitialisedState = "heating initialised"
	HeatingCompletedState   = "heating completed"

	TipDockingCreateSuccess    = "tip docking record created successfully"
	TipDockingFetchSuccess     = "tip docking record fetched successfully"
	TipDockingUpdateSuccess    = "tip docking record updated successfully"
	TipDockingInitialisedState = "tip docking initialised"
	TipDockingCompletedState   = "tip docking completed"

	TipOperationCreateSuccess    = "tip operation record created successfully"
	TipOperationFetchSuccess     = "tip operation record fetched successfully"
	TipOperationUpdateSuccess    = "tip operation record updated successfully"
	TipOperationInitialisedState = "tip operation initialised"
	TipOperationCompletedState   = "tip operation completed"

	ProcessDuplicationSuccess = "duplicate process record created successfully"

	UserLoginSuccess  = "user successfully logged in."
	UserCreateSuccess = "user successfully created"
	UserUpdateSuccess = "user successfully updated"
	UserLogoutSuccess = "user successfully logged out"

	ProcessesRearrangeSuccess = "rearranging processes success"

	CartridgeFetchSuccess          = "cartridge record fetched successfully"
	ConsumableDistanceFetchSuccess = "consumable distances records fetched successfully"

	CartridgeInitialisedState     = "cartridge initialised"
	CartridgeCompletedState       = "cartridge completed"
	CartridgeListInitialisedState = "cartridge list initialised"
	CartridgeListCompletedState   = "cartridge list completed"
	CartridgeCreatedSuccess       = "cartridge created successfully"
	CartridgeDeletedSuccess       = "cartridge deleted successfully"

	ConsumableDistanceInitialisedState = "consumable initialised"
	ConsumableDistanceCompletedState   = "ConsumableDistance completed"

	CalibrationsFetchSuccess    = "calibration were fetched successfully"
	CalibrationUpdateSuccess    = "calibration were updated successfully"
	CalibrationInitialisedState = "calibrations initialised"
	CalibrationCompletedState   = "calibration completed"

	TipTubeFetchSuccess  = "tip tube record fetched successfully"
	TipTubeCreateSuccess = "tip tube record created successfully"

	TipTubeInitialisedState = "tip operation initialised"
	TipTubeCompletedState   = "tip operation initialised"

	RunRecipeInitialisedState = "run recipe initialised"
	RunRecipeCheckStepRun     = "check if the step-run is in progress"
	RunRecipeStepRunSuccess   = "next step run is in progress"

	RecipeInitialisedState     = "recipe initialised"
	RecipeCompletedState       = "recipe completed"
	RecipePublishedState       = "recipe publish completed"
	RecipeListInitialisedState = "recipe list initialised"
	RecipeListCompletedState   = "recipe list Completed"

	RunRecipeProgress     = "next step run is in progress"
	RecipeRunInProgress   = "recipe run is in progress"
	NextStepRunInProgress = "next step run is in progress"
	NextStepWillRun       = "next step will be run"
	StepRunWillAbort      = "step run will be aborted"
	WaitingRunNextProcess = "waiting to run next process"
	NextProcessInProgress = "next process is in progress"

	SafeToUpgrade          = "safe to upgrade the cpagent version"
	TempSettingBothDeckRun = "temporary setting both deck run in progress"
	ResettingBothDeckRun   = "resetting both deck run in progress"

	RecipeCreateSuccess    = "Recipe record created successfully"
	RecipeFetchSuccess     = "Recipe record fetched successfully"
	RecipeListFetchSuccess = "Recipe list fetched successfully"
	RecipeUpdateSuccess    = "Recipe record updated successfully"
	RecipeDeleteSuccess    = "Recipe record deleted successfully"
	RecipePublishSuccess   = "Recipe published successfully"
	RecipeUnPublishSuccess = "Recipe unpublished successfully"

	DiscardBoxMovedSuccess = "Discard box moved to cleanup position successful"
	RestoreDeckSuccess     = "Restore Deck successful"

	UVCleanupProgress = "uv light clean up in progress"

	DiscardBoxInitialisedState     = "discard box operation initialised"
	DiscardBoxCompletedState       = "discard box operation completed"
	RestoreDeckInitialisedState    = "restore deck operation initialised"
	RestoreDeckCompletedState      = "restore deck operation completed"
	UvLightInitialisedState        = "uv light operation initialised"
	UvLightCompletedState          = "uv light operation completed"
	DiscardTipHomeInitialisedState = "discard tip and home operation initialised"
	DiscardTipHomeCompletedState   = "discard tip and home operation completed"
	UvCleanUpSuccess               = "uv light clean up in progress"
	DiscardTipHomeSuccess          = "discard-tip-and-home in progress"

	ProcessInitialisedState          = "process initialised"
	ProcessCompletedState            = "process completed"
	ProcessesFetchSuccess            = "processes fetch success"
	ProcessFetchSuccess              = "process fetch success"
	ProcessDeleteSuccess             = "process delete success"
	ProcessCreateSuccess             = "process create success"
	ProcessUpdateSuccess             = "process updated successfully"
	ProcessListInitialisedState      = "process list initialised"
	ProcessListCompletedState        = "process list completed"
	DuplicateProcessInitialisedState = "duplicate process initialised"
	DuplicateProcessCompletedState   = "duplicate process completed"
	RearrangeProcessInitialisedState = "rearrange process initialised"
	RearrangeProcessCompletedState   = "rearrange process completed"

	//RTPCR
	RTPCRHomingSuccess  = "Homing Success"
	RTPCRResetSuccess   = "Reset Success"
	LidPIDTuningStopped = "PID Tuning stopped successfully"

	// Configs
	UpdateConfigSuccess         = "Config was updated successfully"
	DyeToleranceProgressSuccess = "dye tolerance calculation in progress"
	UserDeleteSuccess           = "user deleted successfully"

	ConsumableDistanceUpdateSuccess = "consumable distance record updated successfully"
	DyeCreateSuccess                = "dyes record created successfully"
	DyeListSuccess                  = "dyes record listed successfully"
)

func GetMachineOperationMessage(operation string, state string) (message string) {

	return fmt.Sprintf("operation %s %s", operation, state)
}

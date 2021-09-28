package responses

import (
	"fmt"
)

var (
	PiercingDecodeError     = fmt.Errorf("error decoding piercing record")
	PiercingValidationError = fmt.Errorf("error validating piercing record")
	PiercingCreateError     = fmt.Errorf("error creating piercing record")
	PiercingFetchError      = fmt.Errorf("error fetching piercing record")
	PiercingUpdateError     = fmt.Errorf("error updating piercing record")

	AspireDispenseDecodeError     = fmt.Errorf("error decoding aspire dispense record")
	AspireDispenseValidationError = fmt.Errorf("error validating aspire dispense record")
	AspireDispenseCreateError     = fmt.Errorf("error creating aspire dispense record")
	AspireDispenseFetchError      = fmt.Errorf("error fetching aspire dispense record")
	AspireDispenseUpdateError     = fmt.Errorf("error updating aspire dispense record")

	DelayDecodeError     = fmt.Errorf("error decoding delay record")
	DelayValidationError = fmt.Errorf("error validating delay record")
	DelayCreateError     = fmt.Errorf("error creating delay record")
	DelayFetchError      = fmt.Errorf("error fetching delay record")
	DelayUpdateError     = fmt.Errorf("error updating delay record")

	AttachDetachDecodeError     = fmt.Errorf("error decoding attach detach record")
	AttachDetachValidationError = fmt.Errorf("error validating attach detach record")
	AttachDetachCreateError     = fmt.Errorf("error creating attach detach record")
	AttachDetachFetchError      = fmt.Errorf("error fetching attach detach record")
	AttachDetachUpdateError     = fmt.Errorf("error updating attach detach record")

	ShakingDecodeError     = fmt.Errorf("error decoding shaking record")
	ShakingValidationError = fmt.Errorf("error validating shaking record")
	ShakingCreateError     = fmt.Errorf("error creating shaking record")
	ShakingFetchError      = fmt.Errorf("error fetching shaking record")
	ShakingUpdateError     = fmt.Errorf("error updating shaking record")
	InvalidShakerTemp      = fmt.Errorf("error please check shaker temperature range")

	HeatingDecodeError     = fmt.Errorf("error decoding heating record")
	HeatingValidationError = fmt.Errorf("error validating heating record")
	HeatingCreateError     = fmt.Errorf("error creating heating record")
	HeatingFetchError      = fmt.Errorf("error fetching heating record")
	HeatingUpdateError     = fmt.Errorf("error updating heating record")

	TipDockingDecodeError     = fmt.Errorf("error decoding tip docking record")
	TipDockingValidationError = fmt.Errorf("error validating tip docking record")
	TipDockingCreateError     = fmt.Errorf("error creating tip docking record")
	TipDockingFetchError      = fmt.Errorf("error fetching tip docking record")
	TipDockingUpdateError     = fmt.Errorf("error updating tip docking record")

	TipOperationDecodeError     = fmt.Errorf("error decoding tip operation record")
	TipOperationValidationError = fmt.Errorf("error validating tip operation record")
	TipOperationCreateError     = fmt.Errorf("error creating tip operation record")
	TipOperationFetchError      = fmt.Errorf("error fetching tip operation record")
	TipOperationUpdateError     = fmt.Errorf("error updating tip operation record")
	TipOperationConvertError    = fmt.Errorf("error converting tip operation to its specific type")

	TipTubeDecodeError       = fmt.Errorf("error decoding tip tube record")
	TipTubeCreateError       = fmt.Errorf("error creating tip tube record")
	TipTubeFetchError        = fmt.Errorf("error fetching tip tube record")
	TipTubeArgumentsError    = fmt.Errorf("error invalid tip tube arguments")
	TipTubeCreateConfigError = fmt.Errorf("error creating tip tube config")

	UUIDParseError = fmt.Errorf("error parsing uuid")

	ProcessValidationError  = fmt.Errorf("error validating process record")
	ProcessDecodeError      = fmt.Errorf("error decoding process record")
	ProcessFetchError       = fmt.Errorf("error fetching process record")
	ProcessCreateError      = fmt.Errorf("error creating process record")
	ProcessIDInvalidError   = fmt.Errorf("error process id is invalid")
	ProcessDuplicationError = fmt.Errorf("error creating duplicate process record")
	ProcessTypeInvalid      = fmt.Errorf("error process type is wrong")
	ProcessesRearrangeError = fmt.Errorf("error rearranging processes")
	ProcessesDecodeSeqError = fmt.Errorf("error while decoding process sequence data")
	ProcessDeleteError      = fmt.Errorf("error deleting process record")
	ProcessUpdateError      = fmt.Errorf("error updating process record")

	RecipeIDInvalidError = fmt.Errorf("error recipe id is invalid")
	RecipeRunPaniced     = fmt.Errorf("error recipe panicked")

	RecipeDecodeError     = fmt.Errorf("error decoding Recipe record")
	RecipeValidationError = fmt.Errorf("error validating Recipe record")
	RecipeCreateError     = fmt.Errorf("error creating Recipe record")
	RecipeFetchError      = fmt.Errorf("error fetching Recipe record")
	RecipeListFetchError  = fmt.Errorf("error fetching Recipe list")
	RecipeUpdateError     = fmt.Errorf("error updating Recipe record")
	RecipeDeleteError     = fmt.Errorf("error deleting Recipe record")
	RecipePublishError    = fmt.Errorf("error recipe already published/unpublished")
	RecipeRunError        = fmt.Errorf("error occured while recipe was running")
	RecipeWasPausedError  = fmt.Errorf("error running recipe was paused atleast once")

	InvalidInterfaceConversionError = fmt.Errorf("error interface conversion failed")
	DelayRangeInvalid               = fmt.Errorf("error invalid delay range allowed range is (0, 100]")

	SimulatorReservedDelayError = fmt.Errorf("error delay is allowed only for simulator")

	StepRunNotInProgressError = fmt.Errorf("error step run is not in progress")
	StepRunAborted            = fmt.Errorf("error step run aborted")
	DeckNameInvalid           = fmt.Errorf("error deck name is invalid")
	PleaseHomeMachineError    = fmt.Errorf("error please home the machine first")
	PickupPositionInvalid     = fmt.Errorf("position is invalid to pickup the tip")
	WebsocketMarshallingError = fmt.Errorf("error in marshalling web socket data")
	UrlArgumentInvalid        = fmt.Errorf("error invalid url argument")

	//user
	UserDecodeError           = fmt.Errorf("error decoding user record")
	UserDeckLoginError        = fmt.Errorf("error login in deck")
	UserInvalidError          = fmt.Errorf("error invalid user")
	UserAuthError             = fmt.Errorf("error in storing user authentication data")
	UserTokenEncodeError      = fmt.Errorf("error in fetching token")
	UserMarshallingError      = fmt.Errorf("error in marshalling token")
	UserInsertError           = fmt.Errorf("error in inserting user")
	UserUpdateError           = fmt.Errorf("error in updating user")
	UsernameBlankError        = fmt.Errorf("error username is blank")
	UserAuthDataFetchError    = fmt.Errorf("error in authenticating user")
	UserAuthDataDeleteError   = fmt.Errorf("error in deleting authenticated user data")
	UserTokenApplicationError = fmt.Errorf("error in token application type")
	UserInvalidDeckError      = fmt.Errorf("error invalid deck")
	UserDeleteError           = fmt.Errorf("error delete user")
	SameUserDeleteError       = fmt.Errorf("error trying to delete same user")

	//user authenticate
	UserUnauthorised              = fmt.Errorf("error user unauthorised")
	UserTokenEmptyError           = fmt.Errorf("error empty token")
	UserTokenDecodeError          = fmt.Errorf("error in decoding token")
	UserTokenRoleEmptyError       = fmt.Errorf("error failed to fetch role")
	UserTokenDeckError            = fmt.Errorf("error failed to fetch deck")
	UserTokenInvalidRoleError     = fmt.Errorf("error invalid role")
	UserTokenCrossDeckError       = fmt.Errorf("error wrong token for deck")
	UserTokenLoggedOutDeckError   = fmt.Errorf("error already logged out ")
	UserTokenUsernameError        = fmt.Errorf("error username not in token ")
	UserTokenApplicationTypeError = fmt.Errorf("error application type not in token ")
	UserTokenAppMismatchError     = fmt.Errorf("error application type mismatch ")
	UserTokenAppNotExistError     = fmt.Errorf("error application type does not exist or is incorrect ")
	UserTokenAuthIdError          = fmt.Errorf("error auth_id not in token ")
	UserTokenAuthIdParseError     = fmt.Errorf("error auth_id parse error ")
	UserAuthNotFoundError         = fmt.Errorf("error user already logged out")

	AuditLogFetchError  = fmt.Errorf("error failed fetching log")
	AuditLogCreateError = fmt.Errorf("error failed saving log")

	CartridgeFetchError     = fmt.Errorf("error fetching cartridge record")
	CartridgeDecodeError    = fmt.Errorf("Error while decoding Cartridge data")
	CartridegInsertionError = fmt.Errorf("Error while inserting Cartridge")
	CartridegDeletionError  = fmt.Errorf("Error while deleting Cartridge")
	CartridgeIDParseError   = fmt.Errorf("Error while parsing Cartridge id")

	InvalidTipTubePositionError = fmt.Errorf("Error Invalid Position for Tip Tube")
	InvalidCartridgeIDError     = fmt.Errorf("Error Invalid Cartridge ID Error")

	InvalidSourcePosition                 = fmt.Errorf("error source position is invalid")
	InvalidDestinationPosition            = fmt.Errorf("error destination position is invalid")
	InvalidDeckPosition                   = fmt.Errorf("error deck position is invalid")
	RecipeCartridge1Missing               = fmt.Errorf("error cartridge 1 is missing")
	RecipeCartridge2Missing               = fmt.Errorf("error cartridge 2 is missing")
	MissingCartridgeWellsError            = fmt.Errorf("error cartridge wells are missing")
	InvalidCartridgeType                  = fmt.Errorf("error cartridge type is invalid")
	CartridgeWellsMismatchWithHeightError = fmt.Errorf("error cartridge wells mismatch with heights")
	InvalidAspireWell                     = fmt.Errorf("error aspire well settings are invalid")
	InvalidTipDockWell                    = fmt.Errorf("error tip docking on well settings are invalid")
	InvalidDispenseWell                   = fmt.Errorf("error dispense well position is invalid")
	InvalidCategoryAspireDispense         = fmt.Errorf("error category for aspire dispense is invalid")
	InvalidPiercingWell                   = fmt.Errorf("error piercing well settings are invalid")

	WrongDeckError       = fmt.Errorf("error invalid deck name")
	TipDoesNotExistError = fmt.Errorf("error specified tip does not exist")
	TipMissingError      = fmt.Errorf("error specified tip does not exist in recipe")

	RunInProgressForSomeDeckError = fmt.Errorf("error run is in progress for either of the decks")

	DiscardBoxMoveError    = fmt.Errorf("error discard box moving was unsuccessful")
	RestoreDeckError       = fmt.Errorf("error restore deck was unsuccessful")
	DiscardBoolOptionError = fmt.Errorf("Invalid boolean value for tip discard option")

	CUDNotAllowedError = "this %v is in progress, so not allowed to %v"

	//RTPCR
	RTPCRHomingError = fmt.Errorf("error in homing rt-pcr")
	RTPCRResetError  = fmt.Errorf("error in reseting rt-pcr")

	// Configs
	ConfigDataDecodeError = fmt.Errorf("error while decoding config data")
	ConfigDataFetchError  = fmt.Errorf("error fetching Config data")
	ConfigDataUpdateError = fmt.Errorf("error Updating Config data")

	PLCDataUpdateError = fmt.Errorf("error updatinf PLC data")

	InvalidExperimentID  = fmt.Errorf("Invalid experiment id")
	ScaleDecodeError     = fmt.Errorf("error while decoding scale data")
	InvalidScaleRange    = fmt.Errorf("error invalid scale range")
	ExperimentFetchError = fmt.Errorf("error fetching experiment data")
	ConfTargetFetchError = fmt.Errorf("error fetching target data")
	ResultFetchError     = fmt.Errorf("error fetching result data")

	InvalidEmailIDError = fmt.Errorf("error fetching result data")

	PreviousExperimentProgressError = fmt.Errorf("error previous experiment already in progress")

	ReportAbsent            = fmt.Errorf("report is absent in form data!")
	DyeToleranceDecodeError = fmt.Errorf("error decoding dye tolerance data!")
	InvalidKitIDError       = fmt.Errorf("error invalid kit id!")

	ConsumableDistanceDecodeError       = fmt.Errorf("error decoding consumable distance record")
	ConsumableDistanceCreateError       = fmt.Errorf("error creating consumable distance record")
	ConsumableDistanceFetchError        = fmt.Errorf("error fetching consumable distance record")
	ConsumableDistanceArgumentsError    = fmt.Errorf("error invalid consumable distance arguments")
	ConsumableDistanceUpdateError       = fmt.Errorf("error updating Consumable distance record")
	ConsumableDistanceUpdateConfigError = fmt.Errorf("error updating Consumable distance config")
	CartridgeCreateConfigError          = fmt.Errorf("error creating Cartridge config")

	CalibrationsFetchError             = fmt.Errorf("error fetching calibration records")
	CalibrationDecodeError             = fmt.Errorf("error decoding calibration record")
	CalibrationUpdateConfigError       = fmt.Errorf("error updating calibrations config")
	CalibrationUpdateError             = fmt.Errorf("error updating calibrations")
	CalibrationVariableMissingError    = fmt.Errorf("error calibration variable is missing")
	CalibrationsPositionCalculateError = fmt.Errorf("error position calculating calibrations")
	CalibrationsCalculateError         = fmt.Errorf("error calculating calibrations")
	CalibrationMethodUnsetError        = fmt.Errorf("error calibration method is unset")

	DyeDecodeError = fmt.Errorf("error decoding dyes record")
	DyeFetchError  = fmt.Errorf("error fetching dyes record")

	DyeInsertError  = fmt.Errorf("error Inserting dyes record")
	DyeMarshalError = fmt.Errorf("error marshalling dyes record")

	UVTimeFormatDecodeError = fmt.Errorf("error decoding uv time record")
	UVMinimumTimeError      = fmt.Errorf("error user given uv time is less than minimum allowed time record")
)

// Special errors which are in []byte format
var (
	DataMarshallingError = []byte(`error marshalling data`)
)

func DefineCUDNotAllowedError(processOrRecipe string, operation string) (err error) {
	msg := fmt.Sprintf(CUDNotAllowedError, processOrRecipe, operation)
	return fmt.Errorf(msg)
}

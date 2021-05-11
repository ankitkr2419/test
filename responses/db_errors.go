package responses

import (
	"fmt"
)

var (
	AspireDispenseUpdateNameError      = fmt.Errorf("error in updating aspire dispense process name in db")
	AspireDispenseInitiateDBTxError    = fmt.Errorf("error while initiating database transaction for aspire dispense")
	AspireDispenseDBFetchError         = fmt.Errorf("error in fetching aspire dispense from db")
	AspireDispenseDBCreateError        = fmt.Errorf("Error creating aspire dispense in db")
	AspireDispenseDuplicateCreateError = fmt.Errorf("Error creating duplicate aspire dispense in db")

	AttachDetachUpdateNameError      = fmt.Errorf("error in updating attach detach process name in db")
	AttachDetachInitiateDBTxError    = fmt.Errorf("error while initiating database transaction for attach detach")
	AttachDetachDBFetchError         = fmt.Errorf("error in fetching attach detach from db")
	AttachDetachDBCreateError        = fmt.Errorf("Error creating attach detach in db")
	AttachDetachDuplicateCreateError = fmt.Errorf("Error creating duplicate attach detach in db")

	TipOperationUpdateNameError      = fmt.Errorf("error in updating tip operation process name in db")
	TipOperationInitiateDBTxError    = fmt.Errorf("error while initiating database transaction for tip operation")
	TipOperationDBFetchError         = fmt.Errorf("error in fetching tip operation from db")
	TipOperationDBCreateError        = fmt.Errorf("Error creating tip operation in db")
	TipOperationDuplicateCreateError = fmt.Errorf("Error creating duplicate tip operation in db")

	TipDockingUpdateNameError      = fmt.Errorf("error in updating tip docking process name in db")
	TipDockingInitiateDBTxError    = fmt.Errorf("error while initiating database transaction for tip docking")
	TipDockingDBFetchError         = fmt.Errorf("error in fetching tip docking from db")
	TipDockingDBCreateError        = fmt.Errorf("Error creating tip docking in db")
	TipDockingDuplicateCreateError = fmt.Errorf("Error creating duplicate tip docking in db")

	HeatingUpdateNameError      = fmt.Errorf("error in updating heating process name in db")
	HeatingInitiateDBTxError    = fmt.Errorf("error while initiating database transaction for heating")
	HeatingDBFetchError         = fmt.Errorf("error in fetching heating from db")
	HeatingDBCreateError        = fmt.Errorf("Error creating heating in db")
	HeatingDuplicateCreateError = fmt.Errorf("Error creating duplicate heating in db")

	ShakingUpdateNameError      = fmt.Errorf("error in updating shaking process name in db")
	ShakingInitiateDBTxError    = fmt.Errorf("error while initiating database transaction for shaking")
	ShakingDBFetchError         = fmt.Errorf("error in fetching shaking from db")
	ShakingDBCreateError        = fmt.Errorf("Error creating shaking in db")
	ShakingDuplicateCreateError = fmt.Errorf("Error creating duplicate shaking in db")

	PiercingUpdateNameError      = fmt.Errorf("error in updating piercing process name in db")
	PiercingInitiateDBTxError    = fmt.Errorf("error while initiating database transaction for piercing")
	PiercingDBFetchError         = fmt.Errorf("error in fetching piercing from db")
	PiercingDBCreateError        = fmt.Errorf("Error creating piercing in db")
	PiercingDuplicateCreateError = fmt.Errorf("Error creating duplicate piercing in db")

	DelayUpdateNameError      = fmt.Errorf("error in updating delay process name in db")
	DelayInitiateDBTxError    = fmt.Errorf("error while initiating database transaction for delay")
	DelayDBFetchError         = fmt.Errorf("error in fetching delay from db")
	DelayDBCreateError        = fmt.Errorf("Error creating delay in db")
	DelayDuplicateCreateError = fmt.Errorf("Error creating duplicate delay in db")

	ProcessUpdateNameError         = fmt.Errorf("error in updating process name in db")
	ProcessInitiateDBTxError       = fmt.Errorf("error while initiating database transaction for process")
	ProcessDBFetchError            = fmt.Errorf("error in fetching process from db")
	ProcessDBCreateError           = fmt.Errorf("Error creating process in db")
	ProcessDuplicateCreationError  = fmt.Errorf("error creating duplicate process in db")
	ProcessHighestSeqNumFetchError = fmt.Errorf("error getting highest sequence number of process in db")

	AuditLogDBCreateError = fmt.Errorf("error in creating audit log in db")
	AuditLogDBShowError   = fmt.Errorf("error in showing audit log ")
)

package gocapng

import "errors"

// Initialized errors (based on documentation)
var (
	ErrNotInitialized                               = errors.New("not initialized")
	ErrSelectBoundsFailureDropBoundingSetCapability = errors.New("SelectBounds and failure to drop a bounding set capability")
	ErrSelectBoundsAndFailureToReReadBoundingSet    = errors.New("SelectBounds and failure to re-read bounding set")
	ErrSelectBoundsCAPSetPCap                       = errors.New("SelectBounds and process does not have CAPSetPCap")
	ErrSelectCapsCapsetSyscall                      = errors.New("SelectCaps and failure in capset syscall")
	ErrSelectAmbientAndProcessClearing              = errors.New("SelectAmbient and process has no capabilities and failed clearing ambient capabilities")
	ErrSelectAmbientProcessCapabilitiesClearing     = errors.New("SelectAmbient and process has capabilities and failed clearing ambient capabilities")
	ErrSelectAmbientProcessCapabilitiesSetting      = errors.New("SelectAmbient and process has capabilities and failed setting an ambient capability")
	ErrCAPNGNotInittedProperly                      = errors.New("means capng has not been initted properly")
	ErrFailureRequestingCapabilitiesUidChange       = errors.New("means a failure requesting to keep capabilities across the uid change")
	ErrApplyingIntermediateCapabilitiesFailed       = errors.New("means that applying the intermediate capabilities failed")
	ErrChangingGIDFailed                            = errors.New("means changing gid failed")
	ErrDroppingSupplementalGroupsFailed             = errors.New("means dropping supplemental groups failed")
	ErrChangingUIDFailed                            = errors.New("means changing the uid failed")
	ErrDroppingAbilityRetainUIDChangeFailed         = errors.New("means dropping the ability to retain caps across a uid change failed")
	ErrClearingBoundingSet                          = errors.New("means clearing the bounding set failed")
	ErrDroppingCAPSETPCAP                           = errors.New("means dropping CAP_SETPCAP or ambient capabilities failed")
	ErrInitializedSupplementalGroups                = errors.New("means initializing supplemental groups failed")
	ErrFDIsNotRegularFile                           = errors.New("fd is not a regular file")
	ErrNonRootNamespaceIDUsedForRootID              = errors.New("non-root namespace id is being used for rootid")
	ErrCapabilityNotFound                           = errors.New("Capability not found")
)

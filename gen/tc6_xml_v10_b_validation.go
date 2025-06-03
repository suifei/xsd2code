package generated

// Generated validation functions

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

// Validator interface for all generated types
type Validator interface {
	Validate() error
}

// Validate validates the DataTypeString struct
func (v *DataTypeString) Validate() error {
	return nil
}

// Validate validates the DataTypeWstring struct
func (v *DataTypeWstring) Validate() error {
	return nil
}

// Validate validates the DataTypeArray struct
func (v *DataTypeArray) Validate() error {
	if len(v.Dimension) < 1 {
		return fmt.Errorf("DataTypeArray.Dimension must have at least 1 elements")
	}
	return nil
}

// Validate validates the DataTypeDerived struct
func (v *DataTypeDerived) Validate() error {
	if v.Name == "" {
		return fmt.Errorf("DataTypeDerived.Name is required")
	}
	return nil
}

// Validate validates the DataTypeEnumValuesValue struct
func (v *DataTypeEnumValuesValue) Validate() error {
	if v.Name == "" {
		return fmt.Errorf("DataTypeEnumValuesValue.Name is required")
	}
	return nil
}

// Validate validates the DataTypeEnumValues struct
func (v *DataTypeEnumValues) Validate() error {
	return nil
}

// Validate validates the DataTypeEnum struct
func (v *DataTypeEnum) Validate() error {
	return nil
}

// Validate validates the DataTypeSubrangeSigned struct
func (v *DataTypeSubrangeSigned) Validate() error {
	return nil
}

// Validate validates the DataTypeSubrangeUnsigned struct
func (v *DataTypeSubrangeUnsigned) Validate() error {
	return nil
}

// Validate validates the DataTypePointer struct
func (v *DataTypePointer) Validate() error {
	return nil
}

// Validate validates the DataType struct
func (v *DataType) Validate() error {
	if v.DATE != nil {
		if err := validateDateTime(*v.DATE); err != nil {
			return fmt.Errorf("DataType.DATE: %v", err)
		}
	}
	if v.DT != nil {
		if err := validateDateTime(*v.DT); err != nil {
			return fmt.Errorf("DataType.DT: %v", err)
		}
	}
	if v.TOD != nil {
		if err := validateDateTime(*v.TOD); err != nil {
			return fmt.Errorf("DataType.TOD: %v", err)
		}
	}
	return nil
}

// Validate validates the RangeSigned struct
func (v *RangeSigned) Validate() error {
	return nil
}

// Validate validates the RangeUnsigned struct
func (v *RangeUnsigned) Validate() error {
	return nil
}

// Validate validates the ValueSimpleValue struct
func (v *ValueSimpleValue) Validate() error {
	return nil
}

// Validate validates the ValueArrayValueValue struct
func (v *ValueArrayValueValue) Validate() error {
	return nil
}

// Validate validates the ValueArrayValue struct
func (v *ValueArrayValue) Validate() error {
	return nil
}

// Validate validates the ValueStructValueValue struct
func (v *ValueStructValueValue) Validate() error {
	return nil
}

// Validate validates the ValueStructValue struct
func (v *ValueStructValue) Validate() error {
	return nil
}

// Validate validates the Value struct
func (v *Value) Validate() error {
	return nil
}

// Validate validates the BodyFBDComment struct
func (v *BodyFBDComment) Validate() error {
	return nil
}

// Validate validates the BodyFBDError struct
func (v *BodyFBDError) Validate() error {
	return nil
}

// Validate validates the BodyFBDConnector struct
func (v *BodyFBDConnector) Validate() error {
	if v.Name == "" {
		return fmt.Errorf("BodyFBDConnector.Name is required")
	}
	return nil
}

// Validate validates the BodyFBDContinuation struct
func (v *BodyFBDContinuation) Validate() error {
	if v.Name == "" {
		return fmt.Errorf("BodyFBDContinuation.Name is required")
	}
	return nil
}

// Validate validates the BodyFBDActionBlockActionReference struct
func (v *BodyFBDActionBlockActionReference) Validate() error {
	if v.Name == "" {
		return fmt.Errorf("BodyFBDActionBlockActionReference.Name is required")
	}
	return nil
}

// Validate validates the BodyFBDActionBlockAction struct
func (v *BodyFBDActionBlockAction) Validate() error {
	return nil
}

// Validate validates the BodyFBDActionBlock struct
func (v *BodyFBDActionBlock) Validate() error {
	return nil
}

// Validate validates the BodyFBDBlockInputVariablesVariable struct
func (v *BodyFBDBlockInputVariablesVariable) Validate() error {
	if v.FormalParameter == "" {
		return fmt.Errorf("BodyFBDBlockInputVariablesVariable.FormalParameter is required")
	}
	return nil
}

// Validate validates the BodyFBDBlockInputVariables struct
func (v *BodyFBDBlockInputVariables) Validate() error {
	return nil
}

// Validate validates the BodyFBDBlockInOutVariablesVariable struct
func (v *BodyFBDBlockInOutVariablesVariable) Validate() error {
	if v.FormalParameter == "" {
		return fmt.Errorf("BodyFBDBlockInOutVariablesVariable.FormalParameter is required")
	}
	return nil
}

// Validate validates the BodyFBDBlockInOutVariables struct
func (v *BodyFBDBlockInOutVariables) Validate() error {
	return nil
}

// Validate validates the BodyFBDBlockOutputVariablesVariable struct
func (v *BodyFBDBlockOutputVariablesVariable) Validate() error {
	if v.FormalParameter == "" {
		return fmt.Errorf("BodyFBDBlockOutputVariablesVariable.FormalParameter is required")
	}
	return nil
}

// Validate validates the BodyFBDBlockOutputVariables struct
func (v *BodyFBDBlockOutputVariables) Validate() error {
	return nil
}

// Validate validates the BodyFBDBlock struct
func (v *BodyFBDBlock) Validate() error {
	if v.TypeName == "" {
		return fmt.Errorf("BodyFBDBlock.TypeName is required")
	}
	return nil
}

// Validate validates the BodyFBDInVariable struct
func (v *BodyFBDInVariable) Validate() error {
	if v.Expression == "" {
		return fmt.Errorf("BodyFBDInVariable.Expression is required")
	}
	return nil
}

// Validate validates the BodyFBDOutVariable struct
func (v *BodyFBDOutVariable) Validate() error {
	if v.Expression == "" {
		return fmt.Errorf("BodyFBDOutVariable.Expression is required")
	}
	return nil
}

// Validate validates the BodyFBDInOutVariable struct
func (v *BodyFBDInOutVariable) Validate() error {
	if v.Expression == "" {
		return fmt.Errorf("BodyFBDInOutVariable.Expression is required")
	}
	return nil
}

// Validate validates the BodyFBDLabel struct
func (v *BodyFBDLabel) Validate() error {
	if v.Label == "" {
		return fmt.Errorf("BodyFBDLabel.Label is required")
	}
	return nil
}

// Validate validates the BodyFBDJump struct
func (v *BodyFBDJump) Validate() error {
	if v.Label == "" {
		return fmt.Errorf("BodyFBDJump.Label is required")
	}
	return nil
}

// Validate validates the BodyFBDReturn struct
func (v *BodyFBDReturn) Validate() error {
	return nil
}

// Validate validates the BodyFBD struct
func (v *BodyFBD) Validate() error {
	return nil
}

// Validate validates the BodyLDComment struct
func (v *BodyLDComment) Validate() error {
	return nil
}

// Validate validates the BodyLDError struct
func (v *BodyLDError) Validate() error {
	return nil
}

// Validate validates the BodyLDConnector struct
func (v *BodyLDConnector) Validate() error {
	if v.Name == "" {
		return fmt.Errorf("BodyLDConnector.Name is required")
	}
	return nil
}

// Validate validates the BodyLDContinuation struct
func (v *BodyLDContinuation) Validate() error {
	if v.Name == "" {
		return fmt.Errorf("BodyLDContinuation.Name is required")
	}
	return nil
}

// Validate validates the BodyLDActionBlockActionReference struct
func (v *BodyLDActionBlockActionReference) Validate() error {
	if v.Name == "" {
		return fmt.Errorf("BodyLDActionBlockActionReference.Name is required")
	}
	return nil
}

// Validate validates the BodyLDActionBlockAction struct
func (v *BodyLDActionBlockAction) Validate() error {
	return nil
}

// Validate validates the BodyLDActionBlock struct
func (v *BodyLDActionBlock) Validate() error {
	return nil
}

// Validate validates the BodyLDBlockInputVariablesVariable struct
func (v *BodyLDBlockInputVariablesVariable) Validate() error {
	if v.FormalParameter == "" {
		return fmt.Errorf("BodyLDBlockInputVariablesVariable.FormalParameter is required")
	}
	return nil
}

// Validate validates the BodyLDBlockInputVariables struct
func (v *BodyLDBlockInputVariables) Validate() error {
	return nil
}

// Validate validates the BodyLDBlockInOutVariablesVariable struct
func (v *BodyLDBlockInOutVariablesVariable) Validate() error {
	if v.FormalParameter == "" {
		return fmt.Errorf("BodyLDBlockInOutVariablesVariable.FormalParameter is required")
	}
	return nil
}

// Validate validates the BodyLDBlockInOutVariables struct
func (v *BodyLDBlockInOutVariables) Validate() error {
	return nil
}

// Validate validates the BodyLDBlockOutputVariablesVariable struct
func (v *BodyLDBlockOutputVariablesVariable) Validate() error {
	if v.FormalParameter == "" {
		return fmt.Errorf("BodyLDBlockOutputVariablesVariable.FormalParameter is required")
	}
	return nil
}

// Validate validates the BodyLDBlockOutputVariables struct
func (v *BodyLDBlockOutputVariables) Validate() error {
	return nil
}

// Validate validates the BodyLDBlock struct
func (v *BodyLDBlock) Validate() error {
	if v.TypeName == "" {
		return fmt.Errorf("BodyLDBlock.TypeName is required")
	}
	return nil
}

// Validate validates the BodyLDInVariable struct
func (v *BodyLDInVariable) Validate() error {
	if v.Expression == "" {
		return fmt.Errorf("BodyLDInVariable.Expression is required")
	}
	return nil
}

// Validate validates the BodyLDOutVariable struct
func (v *BodyLDOutVariable) Validate() error {
	if v.Expression == "" {
		return fmt.Errorf("BodyLDOutVariable.Expression is required")
	}
	return nil
}

// Validate validates the BodyLDInOutVariable struct
func (v *BodyLDInOutVariable) Validate() error {
	if v.Expression == "" {
		return fmt.Errorf("BodyLDInOutVariable.Expression is required")
	}
	return nil
}

// Validate validates the BodyLDLabel struct
func (v *BodyLDLabel) Validate() error {
	if v.Label == "" {
		return fmt.Errorf("BodyLDLabel.Label is required")
	}
	return nil
}

// Validate validates the BodyLDJump struct
func (v *BodyLDJump) Validate() error {
	if v.Label == "" {
		return fmt.Errorf("BodyLDJump.Label is required")
	}
	return nil
}

// Validate validates the BodyLDReturn struct
func (v *BodyLDReturn) Validate() error {
	return nil
}

// Validate validates the BodyLDLeftPowerRailConnectionPointOut struct
func (v *BodyLDLeftPowerRailConnectionPointOut) Validate() error {
	return nil
}

// Validate validates the BodyLDLeftPowerRail struct
func (v *BodyLDLeftPowerRail) Validate() error {
	return nil
}

// Validate validates the BodyLDRightPowerRail struct
func (v *BodyLDRightPowerRail) Validate() error {
	return nil
}

// Validate validates the BodyLDCoil struct
func (v *BodyLDCoil) Validate() error {
	if v.Variable == "" {
		return fmt.Errorf("BodyLDCoil.Variable is required")
	}
	return nil
}

// Validate validates the BodyLDContact struct
func (v *BodyLDContact) Validate() error {
	if v.Variable == "" {
		return fmt.Errorf("BodyLDContact.Variable is required")
	}
	return nil
}

// Validate validates the BodyLD struct
func (v *BodyLD) Validate() error {
	return nil
}

// Validate validates the BodySFCComment struct
func (v *BodySFCComment) Validate() error {
	return nil
}

// Validate validates the BodySFCError struct
func (v *BodySFCError) Validate() error {
	return nil
}

// Validate validates the BodySFCConnector struct
func (v *BodySFCConnector) Validate() error {
	if v.Name == "" {
		return fmt.Errorf("BodySFCConnector.Name is required")
	}
	return nil
}

// Validate validates the BodySFCContinuation struct
func (v *BodySFCContinuation) Validate() error {
	if v.Name == "" {
		return fmt.Errorf("BodySFCContinuation.Name is required")
	}
	return nil
}

// Validate validates the BodySFCActionBlockActionReference struct
func (v *BodySFCActionBlockActionReference) Validate() error {
	if v.Name == "" {
		return fmt.Errorf("BodySFCActionBlockActionReference.Name is required")
	}
	return nil
}

// Validate validates the BodySFCActionBlockAction struct
func (v *BodySFCActionBlockAction) Validate() error {
	return nil
}

// Validate validates the BodySFCActionBlock struct
func (v *BodySFCActionBlock) Validate() error {
	return nil
}

// Validate validates the BodySFCBlockInputVariablesVariable struct
func (v *BodySFCBlockInputVariablesVariable) Validate() error {
	if v.FormalParameter == "" {
		return fmt.Errorf("BodySFCBlockInputVariablesVariable.FormalParameter is required")
	}
	return nil
}

// Validate validates the BodySFCBlockInputVariables struct
func (v *BodySFCBlockInputVariables) Validate() error {
	return nil
}

// Validate validates the BodySFCBlockInOutVariablesVariable struct
func (v *BodySFCBlockInOutVariablesVariable) Validate() error {
	if v.FormalParameter == "" {
		return fmt.Errorf("BodySFCBlockInOutVariablesVariable.FormalParameter is required")
	}
	return nil
}

// Validate validates the BodySFCBlockInOutVariables struct
func (v *BodySFCBlockInOutVariables) Validate() error {
	return nil
}

// Validate validates the BodySFCBlockOutputVariablesVariable struct
func (v *BodySFCBlockOutputVariablesVariable) Validate() error {
	if v.FormalParameter == "" {
		return fmt.Errorf("BodySFCBlockOutputVariablesVariable.FormalParameter is required")
	}
	return nil
}

// Validate validates the BodySFCBlockOutputVariables struct
func (v *BodySFCBlockOutputVariables) Validate() error {
	return nil
}

// Validate validates the BodySFCBlock struct
func (v *BodySFCBlock) Validate() error {
	if v.TypeName == "" {
		return fmt.Errorf("BodySFCBlock.TypeName is required")
	}
	return nil
}

// Validate validates the BodySFCInVariable struct
func (v *BodySFCInVariable) Validate() error {
	if v.Expression == "" {
		return fmt.Errorf("BodySFCInVariable.Expression is required")
	}
	return nil
}

// Validate validates the BodySFCOutVariable struct
func (v *BodySFCOutVariable) Validate() error {
	if v.Expression == "" {
		return fmt.Errorf("BodySFCOutVariable.Expression is required")
	}
	return nil
}

// Validate validates the BodySFCInOutVariable struct
func (v *BodySFCInOutVariable) Validate() error {
	if v.Expression == "" {
		return fmt.Errorf("BodySFCInOutVariable.Expression is required")
	}
	return nil
}

// Validate validates the BodySFCLabel struct
func (v *BodySFCLabel) Validate() error {
	if v.Label == "" {
		return fmt.Errorf("BodySFCLabel.Label is required")
	}
	return nil
}

// Validate validates the BodySFCJump struct
func (v *BodySFCJump) Validate() error {
	if v.Label == "" {
		return fmt.Errorf("BodySFCJump.Label is required")
	}
	return nil
}

// Validate validates the BodySFCReturn struct
func (v *BodySFCReturn) Validate() error {
	return nil
}

// Validate validates the BodySFCLeftPowerRailConnectionPointOut struct
func (v *BodySFCLeftPowerRailConnectionPointOut) Validate() error {
	return nil
}

// Validate validates the BodySFCLeftPowerRail struct
func (v *BodySFCLeftPowerRail) Validate() error {
	return nil
}

// Validate validates the BodySFCRightPowerRail struct
func (v *BodySFCRightPowerRail) Validate() error {
	return nil
}

// Validate validates the BodySFCCoil struct
func (v *BodySFCCoil) Validate() error {
	if v.Variable == "" {
		return fmt.Errorf("BodySFCCoil.Variable is required")
	}
	return nil
}

// Validate validates the BodySFCContact struct
func (v *BodySFCContact) Validate() error {
	if v.Variable == "" {
		return fmt.Errorf("BodySFCContact.Variable is required")
	}
	return nil
}

// Validate validates the BodySFCStepConnectionPointOut struct
func (v *BodySFCStepConnectionPointOut) Validate() error {
	return nil
}

// Validate validates the BodySFCStepConnectionPointOutAction struct
func (v *BodySFCStepConnectionPointOutAction) Validate() error {
	return nil
}

// Validate validates the BodySFCStep struct
func (v *BodySFCStep) Validate() error {
	if v.Name == "" {
		return fmt.Errorf("BodySFCStep.Name is required")
	}
	return nil
}

// Validate validates the BodySFCMacroStep struct
func (v *BodySFCMacroStep) Validate() error {
	return nil
}

// Validate validates the BodySFCJumpStep struct
func (v *BodySFCJumpStep) Validate() error {
	if v.TargetName == "" {
		return fmt.Errorf("BodySFCJumpStep.TargetName is required")
	}
	return nil
}

// Validate validates the BodySFCTransitionConditionReference struct
func (v *BodySFCTransitionConditionReference) Validate() error {
	if v.Name == "" {
		return fmt.Errorf("BodySFCTransitionConditionReference.Name is required")
	}
	return nil
}

// Validate validates the BodySFCTransitionConditionInline struct
func (v *BodySFCTransitionConditionInline) Validate() error {
	return nil
}

// Validate validates the BodySFCTransitionCondition struct
func (v *BodySFCTransitionCondition) Validate() error {
	if len(v.Connection) < 1 {
		return fmt.Errorf("BodySFCTransitionCondition.Connection must have at least 1 elements")
	}
	return nil
}

// Validate validates the BodySFCTransition struct
func (v *BodySFCTransition) Validate() error {
	return nil
}

// Validate validates the BodySFCSelectionDivergenceConnectionPointOut struct
func (v *BodySFCSelectionDivergenceConnectionPointOut) Validate() error {
	return nil
}

// Validate validates the BodySFCSelectionDivergence struct
func (v *BodySFCSelectionDivergence) Validate() error {
	return nil
}

// Validate validates the BodySFCSelectionConvergenceConnectionPointIn struct
func (v *BodySFCSelectionConvergenceConnectionPointIn) Validate() error {
	return nil
}

// Validate validates the BodySFCSelectionConvergence struct
func (v *BodySFCSelectionConvergence) Validate() error {
	return nil
}

// Validate validates the BodySFCSimultaneousDivergenceConnectionPointOut struct
func (v *BodySFCSimultaneousDivergenceConnectionPointOut) Validate() error {
	return nil
}

// Validate validates the BodySFCSimultaneousDivergence struct
func (v *BodySFCSimultaneousDivergence) Validate() error {
	return nil
}

// Validate validates the BodySFCSimultaneousConvergence struct
func (v *BodySFCSimultaneousConvergence) Validate() error {
	return nil
}

// Validate validates the BodySFC struct
func (v *BodySFC) Validate() error {
	return nil
}

// Validate validates the Body struct
func (v *Body) Validate() error {
	return nil
}

// Validate validates the VarList struct
func (v *VarList) Validate() error {
	return nil
}

// Validate validates the VarListPlainVariable struct
func (v *VarListPlainVariable) Validate() error {
	if v.Name == "" {
		return fmt.Errorf("VarListPlainVariable.Name is required")
	}
	return nil
}

// Validate validates the VarListPlain struct
func (v *VarListPlain) Validate() error {
	return nil
}

// Validate validates the Position struct
func (v *Position) Validate() error {
	return nil
}

// Validate validates the Connection struct
func (v *Connection) Validate() error {
	return nil
}

// Validate validates the ConnectionPointIn struct
func (v *ConnectionPointIn) Validate() error {
	if len(v.Connection) < 1 {
		return fmt.Errorf("ConnectionPointIn.Connection must have at least 1 elements")
	}
	return nil
}

// Validate validates the ConnectionPointOut struct
func (v *ConnectionPointOut) Validate() error {
	return nil
}

// Validate validates the PouInstance struct
func (v *PouInstance) Validate() error {
	if v.Name == "" {
		return fmt.Errorf("PouInstance.Name is required")
	}
	if v.Type == "" {
		return fmt.Errorf("PouInstance.Type is required")
	}
	return nil
}

// Validate validates the FormattedText struct
func (v *FormattedText) Validate() error {
	return nil
}

// Validate validates the EdgeModifierType struct
func (v *EdgeModifierType) Validate() error {
	return nil
}

// Validate validates the StorageModifierType struct
func (v *StorageModifierType) Validate() error {
	return nil
}

// Validate validates the PouType struct
func (v *PouType) Validate() error {
	return nil
}

// Validate validates the ProjectFileHeader struct
func (v *ProjectFileHeader) Validate() error {
	if v.CompanyName == "" {
		return fmt.Errorf("ProjectFileHeader.CompanyName is required")
	}
	if v.ProductName == "" {
		return fmt.Errorf("ProjectFileHeader.ProductName is required")
	}
	if v.ProductVersion == "" {
		return fmt.Errorf("ProjectFileHeader.ProductVersion is required")
	}
	if err := validateDateTime(v.CreationDateTime); err != nil {
		return fmt.Errorf("ProjectFileHeader.CreationDateTime: %v", err)
	}
	return nil
}

// Validate validates the ProjectContentHeaderCoordinateInfoPageSize struct
func (v *ProjectContentHeaderCoordinateInfoPageSize) Validate() error {
	return nil
}

// Validate validates the ProjectContentHeaderCoordinateInfoFbdScaling struct
func (v *ProjectContentHeaderCoordinateInfoFbdScaling) Validate() error {
	return nil
}

// Validate validates the ProjectContentHeaderCoordinateInfoFbd struct
func (v *ProjectContentHeaderCoordinateInfoFbd) Validate() error {
	return nil
}

// Validate validates the ProjectContentHeaderCoordinateInfoLdScaling struct
func (v *ProjectContentHeaderCoordinateInfoLdScaling) Validate() error {
	return nil
}

// Validate validates the ProjectContentHeaderCoordinateInfoLd struct
func (v *ProjectContentHeaderCoordinateInfoLd) Validate() error {
	return nil
}

// Validate validates the ProjectContentHeaderCoordinateInfoSfcScaling struct
func (v *ProjectContentHeaderCoordinateInfoSfcScaling) Validate() error {
	return nil
}

// Validate validates the ProjectContentHeaderCoordinateInfoSfc struct
func (v *ProjectContentHeaderCoordinateInfoSfc) Validate() error {
	return nil
}

// Validate validates the ProjectContentHeaderCoordinateInfo struct
func (v *ProjectContentHeaderCoordinateInfo) Validate() error {
	return nil
}

// Validate validates the ProjectContentHeader struct
func (v *ProjectContentHeader) Validate() error {
	if v.Name == "" {
		return fmt.Errorf("ProjectContentHeader.Name is required")
	}
	if v.ModificationDateTime != nil {
		if err := validateDateTime(*v.ModificationDateTime); err != nil {
			return fmt.Errorf("ProjectContentHeader.ModificationDateTime: %v", err)
		}
	}
	return nil
}

// Validate validates the ProjectTypesDataTypesDataType struct
func (v *ProjectTypesDataTypesDataType) Validate() error {
	if v.Name == "" {
		return fmt.Errorf("ProjectTypesDataTypesDataType.Name is required")
	}
	return nil
}

// Validate validates the ProjectTypesDataTypes struct
func (v *ProjectTypesDataTypes) Validate() error {
	return nil
}

// Validate validates the ProjectTypesPousPouInterfaceLocalVars struct
func (v *ProjectTypesPousPouInterfaceLocalVars) Validate() error {
	return nil
}

// Validate validates the ProjectTypesPousPouInterfaceTempVars struct
func (v *ProjectTypesPousPouInterfaceTempVars) Validate() error {
	return nil
}

// Validate validates the ProjectTypesPousPouInterfaceInputVars struct
func (v *ProjectTypesPousPouInterfaceInputVars) Validate() error {
	return nil
}

// Validate validates the ProjectTypesPousPouInterfaceOutputVars struct
func (v *ProjectTypesPousPouInterfaceOutputVars) Validate() error {
	return nil
}

// Validate validates the ProjectTypesPousPouInterfaceInOutVars struct
func (v *ProjectTypesPousPouInterfaceInOutVars) Validate() error {
	return nil
}

// Validate validates the ProjectTypesPousPouInterfaceExternalVars struct
func (v *ProjectTypesPousPouInterfaceExternalVars) Validate() error {
	return nil
}

// Validate validates the ProjectTypesPousPouInterfaceGlobalVars struct
func (v *ProjectTypesPousPouInterfaceGlobalVars) Validate() error {
	return nil
}

// Validate validates the ProjectTypesPousPouInterface struct
func (v *ProjectTypesPousPouInterface) Validate() error {
	return nil
}

// Validate validates the ProjectTypesPousPouActionsAction struct
func (v *ProjectTypesPousPouActionsAction) Validate() error {
	if v.Name == "" {
		return fmt.Errorf("ProjectTypesPousPouActionsAction.Name is required")
	}
	return nil
}

// Validate validates the ProjectTypesPousPouActions struct
func (v *ProjectTypesPousPouActions) Validate() error {
	return nil
}

// Validate validates the ProjectTypesPousPouTransitionsTransition struct
func (v *ProjectTypesPousPouTransitionsTransition) Validate() error {
	if v.Name == "" {
		return fmt.Errorf("ProjectTypesPousPouTransitionsTransition.Name is required")
	}
	return nil
}

// Validate validates the ProjectTypesPousPouTransitions struct
func (v *ProjectTypesPousPouTransitions) Validate() error {
	return nil
}

// Validate validates the ProjectTypesPousPou struct
func (v *ProjectTypesPousPou) Validate() error {
	if v.Name == "" {
		return fmt.Errorf("ProjectTypesPousPou.Name is required")
	}
	return nil
}

// Validate validates the ProjectTypesPous struct
func (v *ProjectTypesPous) Validate() error {
	return nil
}

// Validate validates the ProjectTypes struct
func (v *ProjectTypes) Validate() error {
	return nil
}

// Validate validates the ProjectInstancesConfigurationsConfigurationResourceTask struct
func (v *ProjectInstancesConfigurationsConfigurationResourceTask) Validate() error {
	if v.Name == "" {
		return fmt.Errorf("ProjectInstancesConfigurationsConfigurationResourceTask.Name is required")
	}
	if v.Priority == "" {
		return fmt.Errorf("ProjectInstancesConfigurationsConfigurationResourceTask.Priority is required")
	}
	return nil
}

// Validate validates the ProjectInstancesConfigurationsConfigurationResource struct
func (v *ProjectInstancesConfigurationsConfigurationResource) Validate() error {
	if v.Name == "" {
		return fmt.Errorf("ProjectInstancesConfigurationsConfigurationResource.Name is required")
	}
	return nil
}

// Validate validates the ProjectInstancesConfigurationsConfiguration struct
func (v *ProjectInstancesConfigurationsConfiguration) Validate() error {
	if v.Name == "" {
		return fmt.Errorf("ProjectInstancesConfigurationsConfiguration.Name is required")
	}
	return nil
}

// Validate validates the ProjectInstancesConfigurations struct
func (v *ProjectInstancesConfigurations) Validate() error {
	return nil
}

// Validate validates the ProjectInstances struct
func (v *ProjectInstances) Validate() error {
	return nil
}

// Validate validates the Project struct
func (v *Project) Validate() error {
	return nil
}

// Helper validation functions

func validateDateTime(dt time.Time) error {
	if dt.IsZero() {
		return fmt.Errorf("invalid datetime")
	}
	return nil
}

func validatePattern(value, pattern string) error {
	matched, err := regexp.MatchString(pattern, value)
	if err != nil {
		return fmt.Errorf("invalid pattern: %v", err)
	}
	if !matched {
		return fmt.Errorf("value does not match pattern %s", pattern)
	}
	return nil
}

func validateIntRange(value, min, max int) error {
	if value < min || value > max {
		return fmt.Errorf("value %d is out of range [%d, %d]", value, min, max)
	}
	return nil
}

// applyWhiteSpaceProcessing applies XSD whiteSpace facet processing
func applyWhiteSpaceProcessing(value, whiteSpaceAction string) string {
	switch whiteSpaceAction {
	case "replace":
		// Replace tab, newline, and carriage return with space
		value = strings.ReplaceAll(value, "\t", " ")
		value = strings.ReplaceAll(value, "\n", " ")
		value = strings.ReplaceAll(value, "\r", " ")
		return value
	case "collapse":
		// First apply replace processing
		value = strings.ReplaceAll(value, "\t", " ")
		value = strings.ReplaceAll(value, "\n", " ")
		value = strings.ReplaceAll(value, "\r", " ")
		// Then collapse sequences of spaces and trim
		value = regexp.MustCompile(`\\s+`).ReplaceAllString(value, " ")
		value = strings.TrimSpace(value)
		return value
	case "preserve":
		fallthrough
	default:
		// Preserve all whitespace as-is
		return value
	}
}

func validateFixedValue(value, expectedValue string) error {
	if value != expectedValue {
		return fmt.Errorf("value '%s' does not match fixed value '%s'", value, expectedValue)
	}
	return nil
}

import React, { useMemo, useCallback } from "react";
import PropTypes from "prop-types";
import { Button, CheckBox, Select, Input } from "core-components";
import {
  TargetList,
  TargetListHeader,
  TargetListItem,
  Text,
} from "shared-components";

import { covertToSelectOption } from "utils/helpers";
import TargetHeader from "./TargetHeader";
import { checkIfIdPresentInList, validateThreshold } from "./targetHelper";
import { MIN_THRESHOLD, MAX_THRESHOLD } from "./targetConstants";
import { TargetActions } from "./Style";

const TargetComponent = (props) => {
  const {
    // master target list
    listTargetReducer,
    // local state target list used for manipulation
    selectedTargetState,
    onCheckedHandler,
    onTargetSelect,
    onThresholdChange,
    onVolumeChange,
    onSaveClick,
    onNextClick,
    isLoginTypeAdmin,
    isLoginTypeOperator,
    isTargetListUpdated,
    selectedTemplateDetails,
    isViewStagesEnabled,
    navigateToStepWizard,
    editTemplate,
    isNoTargetSelected,
    setThresholdError,
    isThresholdInvalid,
    onRemoveTarget,
  } = props;

  const isTargetDisabled = (ele) => {
    if (ele.selectedTarget === undefined || ele.threshold === undefined) {
      return true;
    }
    return false;
  };

  // set threshold error flag true if threshold is invalid
  const onThresholdBlurHandler = useCallback(
    (threshold, index) => {
      if (validateThreshold(threshold) === false) {
        setThresholdError(true, index);
      }
    },
    [setThresholdError]
  );

  // reset threshold error flag to false on focus on input field
  const onThresholdFocusHandler = useCallback(
    (index) => {
      setThresholdError(false, index);
    },
    [setThresholdError]
  );

  const getFilteredOptionsList = useMemo(() => {
    if (isLoginTypeAdmin === true) {
      return listTargetReducer.filter(
        (ele) => !checkIfIdPresentInList(ele.get("id"), selectedTargetState)
      );
    }
  }, [listTargetReducer, selectedTargetState, isLoginTypeAdmin]);

  const incrementThreshold = (ele, index) => {
    let val = ele.threshold ? ele.threshold + 1 : 1;
    onThresholdChange(val, index);
  };

  const decrementThreshold = (ele, index) => {
    let val = ele.threshold ? ele.threshold - 1 : 0;
    onThresholdChange(val, index);
  };

  const getTargetRows = useMemo(
    () =>
      selectedTargetState.map((ele, index) => (
        <TargetListItem key={index}>
          {isLoginTypeOperator === true && (
            <CheckBox
              onChange={(event) => {
                onCheckedHandler(event, index);
              }}
              className="mr-2"
              id={`target${index}`}
              checked={ele.isChecked}
              disabled={isTargetDisabled(ele)}
            />
          )}
          {isLoginTypeAdmin === true && (
            // if it's a admin he can select targets from master targets
            <Select
              className="flex-100 px-2"
              options={covertToSelectOption(
                getFilteredOptionsList,
                "name",
                "id"
              )}
              placeholder="Please select target."
              onChange={(selectedTarget) => {
                onTargetSelect(selectedTarget, index);
              }}
              value={ele.selectedTarget ? ele.selectedTarget : null}
            />
          )}
          {isLoginTypeOperator === true && (
            // for operator will show in disabled state
            <Input
              className="flex-100 mr-2"
              type="text"
              placeholder="Type here..."
              defaultValue={ele.selectedTarget && ele.selectedTarget.label}
              disabled
            />
          )}
          <Input
            className="flex-40 pl-2"
            type="number"
            name={`threshold${index}`}
            index={`threshold${index}`}
            placeholder={`${MIN_THRESHOLD} - ${MAX_THRESHOLD}`}
            value={ele.threshold === undefined ? "" : ele.threshold}
            onChange={(event) => {
              onThresholdChange(event.target.value, index);
            }}
            onBlur={() => onThresholdBlurHandler(ele.threshold, index)}
            onFocus={() => onThresholdFocusHandler(index)}
            invalid={ele.thresholdError}
            //disabled for operator always OR for admin it should be disabled till target is selected
            disabled={
              isLoginTypeOperator ||
              (isLoginTypeAdmin === true && ele?.selectedTarget?.label == null)
            }
          />

          {isLoginTypeAdmin === true &&
            Object.keys(ele).length !== 0 &&
            ele?.selectedTarget?.label !== null && (
              <>
                {/** -1 */}
                <Button
                  color="secondary"
                  className="ml-2"
                  style={{ width: "auto" }}
                  onClick={() => decrementThreshold(ele, index)}
                >
                  -
                </Button>
                {/** +1 */}
                <Button
                  color="secondary"
                  className="ml-2"
                  style={{ width: "auto" }}
                  onClick={() => incrementThreshold(ele, index)}
                >
                  +
                </Button>

                {/**show remove target if target is available */}
                <Button
                  color="secondary"
                  className="ml-2"
                  style={{ width: "auto" }}
                  onClick={() => onRemoveTarget(ele)}
                >
                  X
                </Button>
              </>
            )}
        </TargetListItem>
      )),
    [
      onCheckedHandler,
      onTargetSelect,
      onThresholdChange,
      selectedTargetState,
      isLoginTypeOperator,
      isLoginTypeAdmin,
      getFilteredOptionsList,
      onThresholdBlurHandler,
      onThresholdFocusHandler,
    ]
  );
  return (
    <div className="d-flex flex-column overflow-hidden flex-100 pt-5 pb-4">
      <TargetHeader
        isLoginTypeAdmin={isLoginTypeAdmin}
        isLoginTypeOperator={isLoginTypeOperator}
        selectedTemplateDetails={selectedTemplateDetails}
        editTemplate={editTemplate}
        onVolumeChange={onVolumeChange}
      />
      <div className="d-flex overflow-hidden">
        <div className="flex-70 scroll-y p-3">
          <TargetList className="list-target">
            <TargetListHeader>
              <Text className="mb-2 mr-2" />
              <Text className="flex-100 mb-2 px-4">Target</Text>
              <Text className="flex-40 mb-2 px-4">Threshold</Text>
            </TargetListHeader>
            {getTargetRows}
          </TargetList>
        </div>

        <TargetActions className="d-flex flex-column flex-30">
          {isLoginTypeAdmin === true && (
            <Button
              color="primary"
              onClick={navigateToStepWizard}
              className="mx-auto mb-3"
              disabled={
                isTargetListUpdated === true || isViewStagesEnabled === false
              }
            >
              View Steps
            </Button>
          )}
          {isLoginTypeOperator === true && (
            <Button
              color="primary"
              onClick={onNextClick}
              className="mx-auto mb-3"
              disabled={isTargetListUpdated}
            >
              Next
            </Button>
          )}
          <Button
            color="primary"
            onClick={onSaveClick}
            className="mx-auto"
            disabled={
              isTargetListUpdated === false ||
              isNoTargetSelected === true ||
              isThresholdInvalid === true
            }
          >
            Save
          </Button>
        </TargetActions>
      </div>
      {isThresholdInvalid && (
        <div className="d-flex">
          <div className="flex-70 px-5">
            <Text
              Tag="p"
              size={14}
              className="text-right text-danger px-5 mb-0"
            >
              Threshold value should be between {MIN_THRESHOLD} -{" "}
              {MAX_THRESHOLD}
            </Text>
          </div>
        </div>
      )}
    </div>
  );
};

TargetComponent.propTypes = {
  selectedTargetState: PropTypes.object.isRequired,
  selectedTemplateDetails: PropTypes.object.isRequired,
  onCheckedHandler: PropTypes.func.isRequired,
  onThresholdChange: PropTypes.func.isRequired,
  onSaveClick: PropTypes.func.isRequired,
  isLoginTypeAdmin: PropTypes.bool.isRequired,
  isLoginTypeOperator: PropTypes.bool.isRequired,
  isTargetListUpdated: PropTypes.bool.isRequired,
  listTargetReducer: PropTypes.object,
  onTargetSelect: PropTypes.func,
  onNextClick: PropTypes.func,
  isViewStagesEnabled: PropTypes.bool,
  navigateToStageWizard: PropTypes.func,
  editTemplate: PropTypes.func,
};

export default TargetComponent;

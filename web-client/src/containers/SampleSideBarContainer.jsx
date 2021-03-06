import React, { useEffect, useReducer, useCallback } from "react";
import { useDispatch, useSelector } from "react-redux";
import PropTypes from "prop-types";
import SidebarSample from "components/Plate/Sidebar/Sample/SidebarSample";
import { getSamples } from "selectors/sampleSelectors";
import { getWellsSavedStatus } from "selectors/wellSelectors";
import {
  fetchSamples as fetchSamplesAction,
  addSampleLocallyCreated,
} from "action-creators/sampleActionCreators";
import createSampleStateReducer, {
  createSampleInitialState,
  createSampleActions,
  validate,
  getSampleRequestData,
} from "components/Plate/Sidebar/Sample/createSampleState";
import { addWell, addWellReset } from "action-creators/wellActionCreators";
import { taskOptions } from "components/Plate/plateConstant";
import {
  getSampleTargetList,
  getInitSampleTargetList,
} from "components/Plate/Sidebar/Sample/sampleHelper";
import { EXPERIMENT_STATUS } from "appConstants";
import { toast } from "react-toastify";

const SampleSideBarContainer = (props) => {
  // constant
  const {
    experimentTargetsList,
    positions,
    experimentStatus,
    experimentId,
    updateWell,
    isExpanded,
    onWellUpdateClickHandler,
  } = props;
  const dispatch = useDispatch();
  // useSelector
  const samplesListReducer = useSelector(getSamples);
  const { list: sampleList, isLoading: isSampleListLoading } =
    samplesListReducer.toJS();
  const areWellsCreated = useSelector(getWellsSavedStatus);

  //get login reducer details
  const loginReducer = useSelector((state) => state.loginReducer);
  const loginReducerData = loginReducer.toJS();
  let activeDeckObj = loginReducerData?.decks.find((deck) => deck.isActive);
  const { token } = activeDeckObj;

  // useReducer
  const [sampleState, sampleStateDispatch] = useReducer(
    createSampleStateReducer,
    createSampleInitialState
  );
  const isSampleStateValid = validate(sampleState);

  // helper function to update local state
  const updateCreateSampleWrapper = (key, value) => {
    sampleStateDispatch({
      type: createSampleActions.SET_VALUES,
      key,
      value,
    });
  };

  // update targets to local state, so every time list will contain
  // original target list with each target selected
  const addTargetsToLocalState = useCallback(() => {
    if (experimentTargetsList !== null && experimentTargetsList.size !== 0) {
      updateCreateSampleWrapper(
        "targets",
        getInitSampleTargetList(experimentTargetsList)
      );
    }
  }, [experimentTargetsList]);

  // reset local state
  const resetLocalState = useCallback(() => {
    sampleStateDispatch({ type: createSampleActions.RESET_VALUES });
    // after local state reset update targets to local state so that for a newly selected well
    // original target list is shown
    addTargetsToLocalState();
  }, [addTargetsToLocalState]);

  useEffect(() => {
    // if well is selected and data is prefilled, then we empty it
    if (updateWell === null) {
      resetLocalState();
    }
  }, [updateWell]);

  useEffect(() => {
    // on page load, Load target list to local
    addTargetsToLocalState();
  }, [addTargetsToLocalState]);

  useEffect(() => {
    // once well is created reset localState, addWellReducer and restore original target list
    if (areWellsCreated === true) {
      resetLocalState();
      dispatch(addWellReset());
      addTargetsToLocalState();
    }
  }, [areWellsCreated, addTargetsToLocalState, dispatch, resetLocalState]);

  useEffect(() => {
    // this effect will run when operator is trying to update well data
    if (updateWell !== null) {
      const { sample_name, sample_id, task, position, targets } = updateWell;
      // set data to local state for update
      sampleStateDispatch({
        type: createSampleActions.UPDATE_STATE,
        value: {
          isEdit: true,
          position,
          isSideBarOpen: true,
          sample: {
            label: sample_name,
            value: sample_id,
          },
          targets: getSampleTargetList(targets, experimentTargetsList),
          task: {
            label: task,
            value: task,
          },
        },
      });
    }
  }, [updateWell, experimentTargetsList]);

  const fetchSamples = (text) => {
    dispatch(fetchSamplesAction(text, token));
  };

  const addNewLocalSample = (sample) => {
    dispatch(addSampleLocallyCreated(sample));
  };
  // helper function to select or unselect a target stored in targets list in local state
  const onTargetClickHandler = (index) => {
    /** We can select only one target for a dye,
     * So while selection only checking whether any target is already selected having same dye as dye of selected target*/
    const targets = sampleState.get("targets").toJS();
    if (targets[index].isSelected === false) {
      const selectedTargets = targets?.filter((ele) => ele.isSelected);
      const dyeOfSelectedTarget = targets[index]?.dye;
      let isDyeFoundInSelectedTargets = false;
      selectedTargets.map((ele) => {
        if (ele.dye === dyeOfSelectedTarget) {
          isDyeFoundInSelectedTargets = true;
        }
      });
      if (isDyeFoundInSelectedTargets === true) {
        toast.error("Two targets for same dye are not allowed!", {
          autoClose: false,
        });
        return;
      }
    }

    sampleStateDispatch({
      type: createSampleActions.TOGGLE_TARGET,
      value: index,
    });
  };

  const addButtonClickHandler = () => {
    const requestObject = getSampleRequestData(sampleState, positions.toJS());
    dispatch(addWell(experimentId, requestObject, token));
    // reset well update
    onWellUpdateClickHandler(null);
  };

  return (
    <SidebarSample
      sampleState={sampleState}
      updateCreateSampleWrapper={updateCreateSampleWrapper}
      experimentTargetsList={experimentTargetsList}
      fetchSamples={fetchSamples}
      addNewLocalSample={addNewLocalSample}
      sampleOptions={sampleList}
      isSampleListLoading={isSampleListLoading}
      taskOptions={taskOptions}
      onTargetClickHandler={onTargetClickHandler}
      addButtonClickHandler={addButtonClickHandler}
      isSampleStateValid={isSampleStateValid}
      resetLocalState={resetLocalState}
      isDisabled={
        (experimentStatus === EXPERIMENT_STATUS.running ||
          experimentStatus === EXPERIMENT_STATUS.success ||
          experimentStatus === EXPERIMENT_STATUS.stopped ||
          positions.size === 0 ||
          isExpanded === true) &&
        updateWell === null
      }
    />
  );
};

SampleSideBarContainer.propTypes = {
  experimentTargetsList: PropTypes.object.isRequired,
  positions: PropTypes.object.isRequired,
  experimentId: PropTypes.string.isRequired,
  // updated well will contain data of well which is to be updated
  updateWell: PropTypes.object,
};

export default SampleSideBarContainer;

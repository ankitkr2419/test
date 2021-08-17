import React, { useEffect } from "react";
import Plate from "components/Plate";
import { useSelector, useDispatch } from "react-redux";
import { getWells, getWellsPosition } from "selectors/wellSelectors";
import {
  setSelectedWell as setSelectedWellAction,
  setMultiSelectedWell as setMultiSelectedWellAction,
  toggleMultiSelectOption as toggleMultiSelectOptionAction,
  resetSelectedWells as resetSelectedWellAction,
  selectAllWellsOption as selectAllWellsAction,
  fetchWells,
} from "action-creators/wellActionCreators";
import { getExperimentTargets } from "selectors/experimentTargetSelector";
import { fetchExperimentTargets } from "action-creators/experimentTargetActionCreators";
import {
  getExperimentId,
  getExperimentTemplate,
} from "selectors/experimentSelector";
import { setIsPlateRoute } from "action-creators/loginActionCreators";
import { getActiveLoadedWells } from "selectors/activeWellSelector";
import { EXPERIMENT_STATUS, MAX_NO_OF_WELLS } from "appConstants";
import { mailReportInitiated } from "action-creators/activityLogActionCreators";
import { toast } from "react-toastify";
import { getRunExperimentReducer } from "selectors/runExperimentSelector";
import { resetGraphInitiated } from "action-creators/wellGraphActionCreators";
import {
  temperatureApiGraphInitiated,
  temperatureDataSucceeded,
  temperatureGraphInitiated,
} from "action-creators/temperatureGraphActionCreators";

const PlateContainer = () => {
  const dispatch = useDispatch();

  const runExperimentReducer = useSelector(getRunExperimentReducer);
  const runExpProgressReducer = useSelector(
    (state) => state.runExpProgressReducer
  );

  // isExpanded: boolean -> Determines whether this page is redirect via normal flow
  // or by expanding
  const createExperimentReducer = useSelector(
    (state) => state.createExperimentReducer
  );
  const isExpanded = createExperimentReducer.get("isExpanded");

  // current status of experiment
  const experimentStatus = runExperimentReducer.get("experimentStatus");

  //get login reducer details
  const loginReducer = useSelector((state) => state.loginReducer);
  const loginReducerData = loginReducer.toJS();
  let activeDeckObj = loginReducerData?.decks.find((deck) => deck.isActive);
  const { token } = activeDeckObj;

  // get temperature reducer data
  const tempDataReducer = useSelector((state) => state.temperatureGraphReducer);
  const { temperatureData } = tempDataReducer.toJS();

  // experiment targets
  const experimentTargets = useSelector(getExperimentTargets);
  const experimentTargetsList = experimentTargets.get("list");
  // get wells data from server
  const wellListReducer = useSelector(getWells);
  // running experiment id
  const experimentId = useSelector(getExperimentId);
  // running template details
  const experimentTemplate = useSelector(getExperimentTemplate);
  // selected wells positions i.e wells with isSelected/isMultiSelected flag
  const positions = getWellsPosition(wellListReducer);
  // activeWells means the well which are allowed to configure
  const activeWells = useSelector(getActiveLoadedWells);

  // set isPlateRoute true on mount and false on unmount
  useEffect(() => {
    // isPlateRoute use in appHeader to manage visibility of header buttons
    dispatch(setIsPlateRoute(true));
    return () => {
      dispatch(setIsPlateRoute(false));
    };
  }, [dispatch]);

  useEffect(() => {
    if (experimentId !== null) {
      // fetching configured wells data
      dispatch(fetchWells(experimentId, token));
      // fetching experiment targets to show while configuring sample and graph filter
      dispatch(fetchExperimentTargets(experimentId, token));

      if (isExpanded === true) {
        // API to render amplification plot after clicking on expand
        dispatch(resetGraphInitiated({ experimentId: experimentId, token: token }));
        // API to render temperature plot after clicking on expand
        dispatch(temperatureApiGraphInitiated({ experimentId, token }));
      }
    }
    return () => {
      // isPlateRoute use in appHeader to manage visibility of header buttons
      dispatch(setIsPlateRoute(false));
    };
  }, [experimentId, experimentStatus, isExpanded, dispatch]);

  const setSelectedWell = (index, isWellSelected) => {
    dispatch(setSelectedWellAction(index, isWellSelected));
  };

  const resetSelectedWells = () => {
    dispatch(resetSelectedWellAction());
  };

  const setMultiSelectedWell = (index, isWellSelected) => {
    dispatch(setMultiSelectedWellAction(index, isWellSelected));
  };

  const toggleMultiSelectOption = () => {
    // multi selection option for well selection to view it on graph
    dispatch(toggleMultiSelectOptionAction());
  };

  const toggleAllWellSelectedOption = (isAllWellsSelected) => {
    dispatch(selectAllWellsAction(isAllWellsSelected));
  };

  const mailBtnHandler = () => {
    //API call to send an email
    dispatch(
      mailReportInitiated({
        token: token,
        experimentId: experimentId,
        showToast: true,
      })
    );
    toast.info("Sending mail...");
  };

  return (
    <Plate
      wells={wellListReducer.get("defaultList")}
      setSelectedWell={setSelectedWell}
      resetSelectedWells={resetSelectedWells}
      experimentTargetsList={experimentTargetsList}
      positions={positions}
      experimentId={experimentId}
      setMultiSelectedWell={setMultiSelectedWell}
      isMultiSelectionOptionOn={wellListReducer.get("isMultiSelectionOptionOn")}
      isAllWellsSelected={positions.toJS().length === MAX_NO_OF_WELLS}
      toggleMultiSelectOption={toggleMultiSelectOption}
      toggleAllWellSelectedOption={toggleAllWellSelectedOption}
      activeWells={activeWells}
      experimentTemplate={experimentTemplate}
      headerData={runExpProgressReducer.toJS()}
      temperatureData={temperatureData}
      mailBtnHandler={mailBtnHandler}
      token={token}
      isExpanded={isExpanded}
    />
  );
};

export default PlateContainer;

import React, { useEffect, useState } from "react";
import { useHistory } from "react-router";
import { useDispatch, useSelector } from "react-redux";
import CalibrationExtractionComponent from "components/CalibrationExtraction";
import {
  deckBlockInitiated,
  logoutInitiated,
} from "action-creators/loginActionCreators";
import {
  abort,
  commonDetailsInitiated,
  fetchPidInitiated,
  motorInitiated,
  runPid,
  runPidReset,
  updateCommonDetailsInitiated,
  updateMotorDetailsInitiated,
  updatePidInitiated,
  createTipsOrTubesInitiated,
  resetCreatingTipsOrTubes,
  createCartridgesInitiated,
  deleteCartridgesInitiated,
  fetchConsumableInitiated,
  updateConsumableInitiated,
  addConsumableInitiated,
  senseAndHitInitiated,
  fetchCalibrationsDeckAInitiated,
  fetchCalibrationsDeckBInitiated,
} from "action-creators/calibrationActionCreators";
import { DECKNAME, PID_STATUS } from "appConstants";
import { useFormik } from "formik";
import {
  formikInitialState,
  formikToArray,
  formikInitialStateForTipsTubes,
} from "components/CalibrationExtraction/helpers";
import {
  deckHomingActionInitiated,
  showHomingModal,
} from "action-creators/homingActionCreators";

const CalibrationExtractionContainer = () => {
  const dispatch = useDispatch();

  const [showConfirmationModal, setConfirmModal] = useState(false);
  const [showSenseHitHomingModel, setSenseHitHomingModel] = useState(false);
  const [motorSelected, setMotorSelected] = useState([]);

  const formik = useFormik({
    initialValues: formikInitialState,
    enableReinitialize: true,
  });

  //get login reducer details
  const loginReducer = useSelector((state) => state.loginReducer);
  const loginReducerData = loginReducer.toJS();
  let activeDeckObj = loginReducerData?.decks.find((deck) => deck.isActive);
  const { isAdmin, name, token } = activeDeckObj;

  const heaterReducer = useSelector((state) => state.heaterProgressReducer);
  const heaterProgressReducerData = heaterReducer.toJS();
  const { data } = heaterProgressReducerData;

  //create tips or tubes reducer details to reset formik after api success
  const createTipTubeReducer = useSelector(
    (state) => state.createTipTubeReducer
  );
  const createTipTubeReducerData = createTipTubeReducer.toJS();
  const isLoadingCreateTipTube = createTipTubeReducerData?.isLoading;
  const isErrorCreateTipTube = createTipTubeReducerData?.error;

  const pidProgessReducer = useSelector((state) => state.pidProgessReducer);
  const pidProgessReducerData = pidProgessReducer.toJS();
  const progressData = pidProgessReducerData.decks.find(
    (deckObj) => deckObj.deckName === name
  );

  const pidReducer = useSelector((state) => state.pidReducer);
  const pidReducerData = pidReducer.toJS();
  const { pidStatus, pidData, isPidUpdateApi } = pidReducerData;

  const commonDetailsReducer = useSelector(
    (state) => state.commonDetailsReducer
  );
  const commonDetailsReducerData = commonDetailsReducer.toJS();
  const { isUpdateApi, details } = commonDetailsReducerData;

  //Tolerance Variables
  const consumableReducer = useSelector((state) => state.consumableReducer);
  const consumableReducerData = consumableReducer.toJS();

  //Calibrations valriables for Deck A
  const calibrationDeckAReducer = useSelector(
    (state) => state.calibrationDeckAReducer
  );
  const calibrationDeckAReducerData = calibrationDeckAReducer.toJS();

  //Calibrations valriables for Deck B
  const calibrationDeckBReducer = useSelector(
    (state) => state.calibrationDeckBReducer
  );
  const calibrationDeckBReducerData = calibrationDeckBReducer.toJS();

  // fetch pidDetails API (pidTemp, pidMinutes) called initially
  useEffect(() => {
    dispatch(fetchPidInitiated(token));
  }, []);

  // fetch commonDetails (name, email, roomTemp) API called initially
  useEffect(() => {
    dispatch(commonDetailsInitiated(token));
  }, []);

  // fetch consumable distance
  useEffect(() => {
    dispatch(fetchConsumableInitiated(token));
  }, []);

  // fetch calibrations for Deck A and B
  useEffect(() => {
    dispatch(fetchCalibrationsDeckAInitiated(token, DECKNAME.DeckAShort));
    dispatch(fetchCalibrationsDeckBInitiated(token, DECKNAME.DeckBShort));
  }, []);

  useEffect(() => {
    if (
      commonDetailsReducerData.error === false &&
      commonDetailsReducerData.isLoading === false
    ) {
      // populate formik data with fetched values
      if (isUpdateApi === false) {
        const {
          receiver_name,
          receiver_email,
          room_temperature,
          serial_number,
          software_version,
          machine_version,
          manufacturing_year,
          contact_number,
        } = details;
        formik.setFieldValue("name.value", receiver_name);
        formik.setFieldValue("email.value", receiver_email);
        formik.setFieldValue("roomTemperature.value", room_temperature);
        formik.setFieldValue("serialNo.value", serial_number);
        formik.setFieldValue("softwareVersion.value", software_version);
        formik.setFieldValue("machineVersion.value", machine_version);
        formik.setFieldValue("manufacturingYear.value", manufacturing_year);
        formik.setFieldValue("contactNumber.value", contact_number);
      } else {
        // fetch updated data after updation
        dispatch(commonDetailsInitiated(token));
      }
    }
  }, [
    commonDetailsReducerData.error,
    commonDetailsReducerData.isLoading,
    isUpdateApi,
  ]);

  useEffect(() => {
    if (
      consumableReducerData.error === false &&
      consumableReducerData.isLoading === false &&
      consumableReducerData.isUpdateApi === true
    )
      // fetch updated data after updation for consumables and calibrations
      dispatch(fetchConsumableInitiated(token));
    dispatch(fetchCalibrationsDeckAInitiated(token, DECKNAME.DeckAShort));
    dispatch(fetchCalibrationsDeckBInitiated(token, DECKNAME.DeckBShort));
  }, [
    consumableReducerData.error,
    consumableReducerData.isLoading,
    consumableReducerData.isUpdateApi,
  ]);

  useEffect(() => {
    if (pidReducerData.error === false && pidReducerData.isLoading === false) {
      if (isPidUpdateApi === false) {
        // populate formik data with fetched values
        formik.setFieldValue("pidTemperature.value", pidData.pid_temperature);
      } else {
        // fetch updated data after updation
        dispatch(fetchPidInitiated(token));
      }
    }
  }, [pidReducerData.error, pidReducerData.isLoading, isPidUpdateApi]);

  /**another deck must be blocked**/
  useEffect(() => {
    dispatch(deckBlockInitiated({ deckName: name }));
  }, []);

  //if api call is success then reset formik for tipsTubesFields and reducer
  useEffect(() => {
    if (isLoadingCreateTipTube === false && isErrorCreateTipTube === false) {
      handleOnChange("tipTubeId", formikInitialStateForTipsTubes.tipTubeId);
      handleOnChange("tipTubeName", formikInitialStateForTipsTubes.tipTubeName);
      handleOnChange("tipTubeType", formikInitialStateForTipsTubes.tipTubeType);
      handleOnChange(
        "allowedPositions",
        formikInitialStateForTipsTubes.allowedPositions
      );
      handleOnChange("volume", formikInitialStateForTipsTubes.volume);
      handleOnChange("height", formikInitialStateForTipsTubes.height);
      handleOnChange("ttBase", formikInitialStateForTipsTubes.ttBase);

      //clear reducer
      dispatch(resetCreatingTipsOrTubes());
    }
  }, [isLoadingCreateTipTube, isErrorCreateTipTube]);

  const handleLogout = () => {
    dispatch(
      logoutInitiated({ deckName: name, token: token, showToast: true })
    );
  };

  const handlePidBtn = () => {
    const deckName =
      name === DECKNAME.DeckA ? DECKNAME.DeckAShort : DECKNAME.DeckBShort;
    if (pidStatus === PID_STATUS.running) {
      // dispatch abort API if progressing
      dispatch(abort(token, deckName));
    } else {
      // dispatch run PID progressing
      dispatch(runPid(token, deckName));
    }
  };

  const handleMotorBtn = (e) => {
    e.preventDefault();

    const { motorNumber, direction, distance } = formik.values;

    const body = {
      deck: name === DECKNAME.DeckA ? DECKNAME.DeckAShort : DECKNAME.DeckBShort,
      motor_number: motorNumber.value,
      direction: direction.value,
      distance: distance.value,
    };
    if (!motorSelected.includes(motorNumber.value)) {
      setMotorSelected((arr) => [...arr, motorNumber.value]);
    }

    dispatch(motorInitiated(token, body));
  };

  const handleSenseAndHitBtn = () => {
    setSenseHitHomingModel(true);
    dispatch(showHomingModal());
  };

  const handleSaveDetailsBtn = (data) => {
    const {
      name,
      email,
      roomTemperature,
      serialNo,
      manufacturingYear,
      machineVersion,
      softwareVersion,
      contactNumber,
    } = data;
    const requestBody = {
      receiver_name: name.value,
      receiver_email: email.value,
      room_temperature: roomTemperature.value,
      serial_number: serialNo.value,
      manufacturing_year: manufacturingYear.value,
      machine_version: machineVersion.value,
      software_version: softwareVersion.value,
      contact_number: contactNumber.value,
    };
    dispatch(updateCommonDetailsInitiated({ token, data: requestBody }));
  };

  const handlePidUpdateBtn = (pidData) => {
    const { pidTemperature } = pidData;
    const requestBody = {
      pid_temperature: pidTemperature.value,
      pid_minutes: 30, // will be removed in future
      micro_lit_pulses: 25, // will be removed in future
      shaker_steps_per_revolution: 800, // will be removed in future
    };
    dispatch(updatePidInitiated(token, requestBody));
  };

  const handleUpdateMotorDetailsBtn = ({
    id,
    deck,
    number,
    name,
    ramp,
    steps,
    slow,
    fast,
  }) => {
    const requestBody = {
      id: parseInt(id.value),
      deck: deck.value,
      number: parseInt(number.value),
      name: name.value,
      ramp: parseInt(ramp.value),
      steps: parseInt(steps.value),
      slow: parseInt(slow.value),
      fast: parseInt(fast.value),
    };
    dispatch(updateMotorDetailsInitiated({ requestBody, token }));
  };

  const toggleConfirmModal = () => {
    dispatch(showHomingModal());
    setConfirmModal(!showConfirmationModal);
  };

  const handleTipesTubesButton = (e) => {
    e.preventDefault();

    let {
      tipTubeId,
      tipTubeName,
      tipTubeType,
      allowedPositions,
      volume,
      height,
      ttBase,
    } = formik.values;
    let arrayOfAllowedPositions = formikToArray(allowedPositions);

    let body = {
      id: tipTubeId.value,
      name: tipTubeName.value,
      type: tipTubeType.value.value,
      allowed_positions: arrayOfAllowedPositions,
      volume: volume.value,
      height: height.value,
      tt_base: ttBase.value,
    };

    dispatch(createTipsOrTubesInitiated(token, body));
  };

  const handleCreateCartridgeBtn = (body) => {
    dispatch(createCartridgesInitiated(token, body));
  };

  const handleDeleteCartridgeBtn = (id) => {
    dispatch(deleteCartridgesInitiated(token, id));
  };

  const handleOnChange = (key, value) => {
    formik.setFieldValue(key, value);
  };

  const addNewConsumableDistance = (requestBody, isUpdate) => {
    if (isUpdate) {
      dispatch(updateConsumableInitiated({ token, requestBody }));
    } else {
      dispatch(addConsumableInitiated({ token, requestBody }));
    }
  };

  const handleHomingDeck = () => {
    let deckName =
      name === DECKNAME.DeckA
        ? DECKNAME.DeckAShort
        : name === DECKNAME.DeckB
        ? DECKNAME.DeckBShort
        : null;
    if (deckName === null) {
      console.error("deckName not found");
      return;
    }
    dispatch(deckHomingActionInitiated({ deckName, token }));
  };

  return (
    <CalibrationExtractionComponent
      toggleConfirmModal={toggleConfirmModal}
      handleLogout={handleLogout}
      handleBtnClick={handlePidBtn}
      handleMotorBtn={handleMotorBtn}
      handleSenseAndHitBtn={handleSenseAndHitBtn}
      handleSaveDetailsBtn={handleSaveDetailsBtn}
      handlePidUpdateBtn={handlePidUpdateBtn}
      handleUpdateMotorDetailsBtn={handleUpdateMotorDetailsBtn}
      handleCreateCartridgeBtn={handleCreateCartridgeBtn}
      handleDeleteCartridgeBtn={handleDeleteCartridgeBtn}
      showConfirmationModal={showConfirmationModal}
      setConfirmModal={setConfirmModal}
      showSenseHitHomingModel={showSenseHitHomingModel}
      setSenseHitHomingModel={setSenseHitHomingModel}
      motorSelected={motorSelected}
      setMotorSelected={setMotorSelected}
      heaterData={data}
      progressData={progressData}
      pidStatus={pidStatus}
      deckName={name}
      formik={formik}
      isAdmin={isAdmin}
      handleTipesTubesButton={handleTipesTubesButton}
      addNewConsumableDistance={addNewConsumableDistance}
      consumableDistanceData={consumableReducerData.data || null}
      calibrationsDataForDeckA={calibrationDeckAReducerData.data || null}
      calibrationsDataForDeckB={calibrationDeckBReducerData.data || null}
      handleHomingDeck={handleHomingDeck}
    />
  );
};

export default React.memo(CalibrationExtractionContainer);

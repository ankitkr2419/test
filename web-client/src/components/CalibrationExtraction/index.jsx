import React from "react";
import { useHistory } from "react-router";

import { Button, Card, CardBody } from "core-components";
import {
  ButtonIcon,
  ColoredCircle,
  Icon,
  MlModal,
  Text,
} from "shared-components";
import { MODAL_BTN, MODAL_MESSAGE, ROOT_URL_PATH, ROUTES } from "appConstants";
import MotorComponent from "./MotorComponent";
import CommonFieldsComponent from "./CommonFieldsComponent";

import { HeadingTitle } from "./HeadingTitle";
import PidProgressComponent from "./PidProgressComponent";
import PidComponent from "./PidComponent";
import UpdateMotorComponent from "./UpdateMotorComponent";
import TipsTubesComponent from "./TipsTubesComponent";
import CartridgeComponent from "./CartridgeComponent";
import DeleteCartridgeComponent from "./DeleteCartridgeComponent";
import ConsumableDistancesComponent from "./ConsumableDistancesComponent";
import ShakerAndHeaterComponent from "./ShakerAndHeaterComponent";
import HomingCurrentDeckComponent from "./HomingCurrentDeckComponent";
import HomingWithSenseAndHitComponent from "./HomingWithSenseAndHitComponent";
import LogoutComponent from "./LogoutComponent";

const CalibrationExtractionComponent = (props) => {
  const {
    deckName,
    heaterData,
    progressData,
    pidStatus,
    handleBtnClick,
    handleLogout,
    handlePidUpdateBtn,
    handleMotorBtn,
    handleSenseAndHitBtn,
    handleSaveDetailsBtn,
    handleUpdateMotorDetailsBtn,
    handleCreateCartridgeBtn,
    handleDeleteCartridgeBtn,
    showConfirmationModal,
    setConfirmModal,
    showSenseHitHomingModel,
    setSenseHitHomingModel,
    motorSelected,
    setMotorSelected,
    toggleConfirmModal,
    formik,
    isAdmin,
    handleTipesTubesButton,
    addNewConsumableDistance,
    addNewCalibration,
    consumableDistanceData,
    calibrationsDataForDeckA,
    calibrationsDataForDeckB,
    handleHomingDeck,
  } = props;

  const { shaker_1_temp, shaker_2_temp, heater_on } = heaterData;

  const history = useHistory();

  const handleBack = () => {
    history.push(ROUTES.recipeListing);
  };

  return (
    <div className="calibration-content px-5 pt-3">
      <div className="d-flex align-items-center">
        {isAdmin && (
          <div style={{ cursor: "pointer" }} onClick={handleBack}>
            <Icon name="angle-left" size={32} className="text-white" />
            <HeadingTitle
              Tag="h5"
              className="text-white font-weight-bold ml-3 mb-0"
            >
              Go back to recipe listing
            </HeadingTitle>
          </div>
        )}

        <Card default className="ml-auto rounded-lg">
          <CardBody className="d-flex p-2">
            <Text className="font-weight-bold mr-3 text-muted">
              Heater Status: <ColoredCircle isOnline={heater_on} />
              {"  "}
            </Text>
            <Text className="font-weight-bold m-0 text-muted">
              Heater Temperature 1: {shaker_1_temp ? shaker_1_temp : 0}° C
              <br />
              Heater Temperature 2: {shaker_2_temp ? shaker_2_temp : 0}° C
            </Text>
          </CardBody>
        </Card>

        <ButtonIcon
          name="logout"
          size={28}
          className="ml-3 bg-white border-primary"
          onClick={toggleConfirmModal}
        />
      </div>

      <Card
        default
        className="mt-3"
        style={{ maxHeight: 419, overflow: "scroll" }}
      >
        <CardBody className="px-5 py-4 text-center">
          {/* Shaker and Heater card */}
          <ShakerAndHeaterComponent />

          <div className="d-flex">
            {/* {PID Start/Abort Progress Component} */}
            <PidProgressComponent
              pidStatus={pidStatus}
              progressData={progressData}
              handleBtnClick={handleBtnClick}
            />

            {/* PID Details update component */}
            <PidComponent formik={formik} handlePidBtn={handlePidUpdateBtn} />
          </div>

          {/* Common fields - name, email, room temperature */}
          <CommonFieldsComponent
            formik={formik}
            handleSaveDetailsBtn={handleSaveDetailsBtn}
          />

          {/* Motor Component -   */}
          <MotorComponent
            formik={formik}
            handleMotorBtn={handleMotorBtn}
            handleSenseAndHitBtn={handleSenseAndHitBtn}
            setMotorSelected={setMotorSelected}
            motorSelected={motorSelected}
          />

          {/* Update Motor Component -   */}
          <UpdateMotorComponent
            handleUpdateMotorDetailsBtn={handleUpdateMotorDetailsBtn}
          />

          {/* Create Cartridges Component -   */}
          <CartridgeComponent
            handleCreateCartridgeBtn={handleCreateCartridgeBtn}
          />

          {/* Delate cartridge component */}
          <DeleteCartridgeComponent
            handleDeleteBtn={handleDeleteCartridgeBtn}
          />

          {/**Tips & tubes Component */}
          <TipsTubesComponent
            formik={formik}
            handleTipesTubesButton={handleTipesTubesButton}
          />

          {/* Consumable Distances Component */}
          <ConsumableDistancesComponent
            addNewData={addNewConsumableDistance}
            data={consumableDistanceData}
            heading={"Consumable Distances"}
          />

          {/* Calibrations Component for Deck A */}
          <ConsumableDistancesComponent
            addNewData={addNewConsumableDistance}
            data={calibrationsDataForDeckA}
            heading={"Calibrations for Deck A"}
            isReadOnly
          />

          {/* Calibrations Component for Deck B */}
          <ConsumableDistancesComponent
            addNewData={addNewConsumableDistance}
            data={calibrationsDataForDeckB}
            heading={"Calibrations for Deck B"}
            isReadOnly
          />

          {/** Homing for current deck */}
          <HomingCurrentDeckComponent
            heading={`Homing for ${deckName}`}
            handleHomingDeck={handleHomingDeck}
          />
        </CardBody>
      </Card>

      {showConfirmationModal && (
        <LogoutComponent
          handleLogout={handleLogout}
          setConfirmModal={setConfirmModal}
        />
      )}

      {showSenseHitHomingModel && (
        <HomingWithSenseAndHitComponent
          formik={formik}
          setSenseHitHomingModel={setSenseHitHomingModel}
          motorSelected={motorSelected}
          setMotorSelected={setMotorSelected}
        />
      )}
    </div>
  );
};

export default React.memo(CalibrationExtractionComponent);

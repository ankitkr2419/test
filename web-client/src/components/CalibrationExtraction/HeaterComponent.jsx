import React, { useEffect } from "react";
import { Redirect } from "react-router";
import { useDispatch, useSelector } from "react-redux";
import { useFormik } from "formik";
import { toast } from "react-toastify";

import { Card, CardBody } from "core-components";
import {
  ButtonIcon,
  ButtonBar,
  Text,
  ColoredCircle,
  HomingModal,
} from "shared-components";
import {
  abort,
  heaterInitiated,
} from "action-creators/calibrationActionCreators";
import { showHomingModal as showHomingModalAction } from "action-creators/homingActionCreators";
import { getHeaterRequestBody, heaterInitialFormikState } from "./helpers";
import { PageBody, HeatingBox, TopContent } from "./Style";
import HeatingProcess from "./HeatingProcess";
import TopHeading from "shared-components/TopHeading";
import { DECKNAME, HEATER_RUN_STATUS, ROUTES } from "appConstants";
import { Spinner } from "reactstrap";

const HeaterComponent = (props) => {
  const dispatch = useDispatch();

  const loginReducer = useSelector((state) => state.loginReducer);
  const loginReducerData = loginReducer.toJS();
  let activeDeckObj =
    loginReducerData && loginReducerData.decks.find((deck) => deck.isActive);

  const heaterRunProgessReducer = useSelector(
    (state) => state.heaterRunProgessReducer
  );
  const { heaterRunStatus } = heaterRunProgessReducer.toJS();
  const { progressing, progressAborted } = HEATER_RUN_STATUS;

  const heaterReducer = useSelector((state) => state.heaterProgressReducer);
  const heaterProgressReducerData = heaterReducer.toJS();
  const { heater_on, shaker_1_temp, shaker_2_temp } =
    heaterProgressReducerData.data;

  const formik = useFormik({
    initialValues: heaterInitialFormikState,
    enableReinitialize: true,
  });

  const { name, token } = activeDeckObj;

  // if progress is aborted then open homing modal
  useEffect(() => {
    if (heaterRunStatus === progressAborted) {
      dispatch(showHomingModalAction());
    }
  }, [heaterRunStatus]);

  const handleSaveBtn = () => {
    const body = getHeaterRequestBody(formik);

    if (body) {
      const requestBody = {
        body: body,
        token: token,
        deckName:
          name === DECKNAME.DeckA ? DECKNAME.DeckAShort : DECKNAME.DeckBShort,
      };
      dispatch(heaterInitiated(requestBody));
    } else {
      //error
      toast.error("Invalid Request", { autoClose: false });
    }
  };

  const handleAbortBtn = () => {
    const deckName =
      name === DECKNAME.DeckA ? DECKNAME.DeckAShort : DECKNAME.DeckBShort;
    dispatch(abort(token, deckName));
  };

  const getRightBtnLabel = () => {
    if (heaterRunStatus === progressing) {
      return (
        <div className="d-flex">
          <Spinner size="sm" />
          <Text className="ml-3">Heating</Text>
        </div>
      );
    } else {
      return "Start";
    }
  };

  if (!activeDeckObj.isLoggedIn) {
    return <Redirect to={`/${ROUTES.landing}`} />;
  }

  return (
    <>
      <HomingModal />
      <PageBody>
        <HeatingBox>
          <div className="process-content process-heating px-2">
            <TopContent className="d-flex justify-content-between align-items-center mx-5">
              <div className="d-flex flex-column">
                <div className="d-flex align-items-center frame-icon">
                  <ButtonIcon
                    size={60}
                    name="heating"
                    className="text-primary bg-white border-gray"
                  />
                  <TopHeading titleHeading="Heating" />
                </div>
              </div>
              <Card default className="ml-auto mr-5 rounded-lg bg-transparent">
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
            </TopContent>
            <Card className="fix-height-card">
              <CardBody className="p-0 overflow-hidden">
                <HeatingProcess formik={formik} />
              </CardBody>
            </Card>
            <ButtonBar
              rightBtnLabel={getRightBtnLabel()}
              leftBtnLabel="Abort"
              handleRightBtn={handleSaveBtn}
              handleLeftBtn={handleAbortBtn}
              btnBarClassname={"btn-bar-adjust-heating"}
              isRightBtnDisabled={heaterRunStatus === progressing}
              isLeftBtnDisabled={!(heaterRunStatus === progressing)}
            />
          </div>
        </HeatingBox>
      </PageBody>
    </>
  );
};

HeaterComponent.propTypes = {};

export default React.memo(HeaterComponent);

import React from "react";
import { Redirect } from "react-router";
import { useDispatch, useSelector } from "react-redux";
import { useFormik } from "formik";
import { toast } from "react-toastify";

import { Card, CardBody } from "core-components";
import { ButtonIcon, ButtonBar, Text } from "shared-components";
import {
  abort,
  heaterInitiated,
} from "action-creators/calibrationActionCreators";
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
  const { progressing } = HEATER_RUN_STATUS;

  const formik = useFormik({
    initialValues: heaterInitialFormikState,
    enableReinitialize: true,
  });

  const { name, token } = activeDeckObj;

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
      toast.error("Invalid Request");
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

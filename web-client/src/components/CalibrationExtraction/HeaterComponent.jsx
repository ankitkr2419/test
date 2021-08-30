import React from "react";

import { Card, CardBody } from "core-components";
import { ButtonIcon, ButtonBar } from "shared-components";

import HeatingProcess from "./HeatingProcess";
import TopHeading from "shared-components/TopHeading";
import { useDispatch, useSelector } from "react-redux";
import { useFormik } from "formik";
import { getHeaterRequestBody, heaterInitialFormikState } from "./helpers";
import { PageBody, HeatingBox, TopContent } from "./Style";
import { toast } from "react-toastify";
import { Redirect } from "react-router";
import { DECKNAME, ROUTES } from "appConstants";
import {
  abort,
  heaterInitiated,
} from "action-creators/calibrationActionCreators";

const HeaterComponent = (props) => {
  const dispatch = useDispatch();

  const loginReducer = useSelector((state) => state.loginReducer);
  const loginReducerData = loginReducer.toJS();
  let activeDeckObj =
    loginReducerData && loginReducerData.decks.find((deck) => deck.isActive);

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
            <Card>
              <CardBody className="p-0 overflow-hidden">
                <HeatingProcess formik={formik} />
              </CardBody>
            </Card>
            <ButtonBar
              rightBtnLabel="Start"
              leftBtnLabel="Abort"
              handleRightBtn={handleSaveBtn}
              handleLeftBtn={handleAbortBtn}
              btnBarClassname={"btn-bar-adjust-heating"}
            />
          </div>
        </HeatingBox>
      </PageBody>
    </>
  );
};

HeaterComponent.propTypes = {};

export default React.memo(HeaterComponent);

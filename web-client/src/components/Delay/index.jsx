import React from "react";

import {
  Card,
  CardBody,
  FormGroup,
  Label,
  Input,
  FormError,
} from "core-components";
import { ButtonIcon, ButtonBar, ImageIcon } from "shared-components";

import delayProcessGraphics from "assets/images/delay-process-graphics.svg";
import TopHeading from "shared-components/TopHeading";
import { DelayBox, PageBody, TopContent } from "./Style";
import { Redirect } from "react-router";
import { useDispatch, useSelector } from "react-redux";
import { useFormik } from "formik";
import { ROUTES } from "appConstants";
import { saveDelayInitiated } from "action-creators/processesActionCreators";

const DelayComponent = (props) => {
  const loginReducer = useSelector((state) => state.loginReducer);
  const loginReducerData = loginReducer.toJS();

  const dispatch = useDispatch();

  const formik = useFormik({
    initialValues: { hours: 0, mins: 0 },
  });

  const recipeDetailsReducer = useSelector(
    (state) => state.updateRecipeDetailsReducer
  );
  const recipeID = recipeDetailsReducer.recipeDetails.id;
  const token = recipeDetailsReducer.token;

  const saveBtnHandler = () => {
    const hours = parseInt(formik.values.hours);
    const mins = parseInt(formik.values.mins);
    const time = hours * 60 * 60 + mins * 60;

    const requestBody = {
      body: { delay_time: time },
      recipeID: recipeID,
      token: token,
    };
    dispatch(saveDelayInitiated(requestBody));
  };

  let activeDeckObj =
    loginReducerData && loginReducerData.decks.find((deck) => deck.isActive);
  if (!activeDeckObj.isLoggedIn) {
    return <Redirect to={`/${ROUTES.landing}`} />;
  }

  return (
    <>
      <PageBody>
        <DelayBox>
          <div className="process-content -delay px-2">
            <TopContent className="d-flex justify-content-between align-items-center mx-5">
              <div className="d-flex flex-column">
                <div className="d-flex align-items-center frame-icon">
                  <ButtonIcon
                    size={60}
                    name="delay"
                    className="text-primary bg-white border-gray"
                  />
                  <TopHeading titleHeading="Delay" />
                </div>
              </div>
            </TopContent>
            <Card>
              <CardBody className="delay-inner-box">
                <div className="delay-box d-flex justify-content-center align-items-center">
                  <FormGroup className="d-flex align-items-center px-4">
                    <Label for="deck-position" className="label-name mb-0">
                      Hold for
                    </Label>
                    <div className="d-flex flex-column input-field">
                      <Input
                        type="text"
                        name="hours"
                        id="hours"
                        placeholder="Type here"
                        onChange={(e) =>
                          formik.setFieldValue("hours", e.target.value)
                        }
                      />
                      <Label
                        for="delay"
                        className="font-weight-bold delay-note mt-2"
                      >
                        Hours
                      </Label>
                      <FormError>Incorrect Hours</FormError>
                    </div>

                    <div className="d-flex flex-column input-field ml-4">
                      <Input
                        type="text"
                        name="minutes"
                        id="minutes"
                        placeholder="Type here"
                        onChange={(e) =>
                          formik.setFieldValue("mins", e.target.value)
                        }
                      />
                      <Label
                        for="delay"
                        className="font-weight-bold delay-note mt-2"
                      >
                        Minutes
                      </Label>
                      <FormError>Incorrect Minutes</FormError>
                    </div>
                  </FormGroup>
                </div>
                <ImageIcon
                  src={delayProcessGraphics}
                  alt="Tip Pickup Process"
                  className="process-image"
                />
              </CardBody>
            </Card>
            <ButtonBar rightBtnLabel="Save" handleRightBtn={saveBtnHandler} />
          </div>
        </DelayBox>
      </PageBody>
    </>
  );
};

DelayComponent.propTypes = {};

export default DelayComponent;

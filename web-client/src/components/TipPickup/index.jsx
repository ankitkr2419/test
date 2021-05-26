import React from "react";
import { useFormik } from "formik";
import { useDispatch, useSelector } from "react-redux";
import { useHistory } from "react-router";

import {
  Card,
  CardBody,
  FormGroup,
  Label,
  Select,
  FormError,
} from "core-components";
import { ButtonIcon, ButtonBar, ImageIcon } from "shared-components";

import AppFooter from "components/AppFooter";
import tipPickupProcessGraphics from "assets/images/tip-pickup-process-graphics.svg";
import TopHeading from "shared-components/TopHeading";
import { PageBody, TipPickupBox, TopContent } from "./Style";
import { saveTipPickupInitiated } from "action-creators/processesActionCreators";
import { TEST_RECIPE_ID, TEST_TOKEN } from "appConstants";

const TipPickupComponent = (props) => {
  const dispatch = useDispatch();
  const history = useHistory();

  const recipeDetailsReducer = useSelector(
    (state) => state.updateRecipeDetailsReducer
  );
  const recipeID = recipeDetailsReducer.recipeDetails.id;
  const token = recipeDetailsReducer.token;

  const taskOptions = [
    { value: "1", label: "1" },
    { value: "2", label: "2" },
    { value: "3", label: "3" },
    { value: "4", label: "4" },
    { value: "5", label: "5" },
  ];

  const formik = useFormik({
    initialValues: { tipPosition: null },
  });

  const handleRightBtn = () => {
    const tipPosition = formik.values.tipPosition;
    dispatch(
      saveTipPickupInitiated({
        recipeID: recipeID, //TEST_RECIPE_ID, //For testing, will be removed later.
        position: tipPosition,
        token: token, //TEST_TOKEN, //For testing, will be removed later.
      })
    );
  };

  return (
    <>
      <PageBody>
        <TipPickupBox>
          <div className="process-content -tip-pickup px-2">
            <TopContent className="d-flex justify-content-between align-items-center mx-5">
              <div className="d-flex flex-column">
                <div className="d-flex align-items-center frame-icon">
                  <ButtonIcon
                    size={60}
                    name="tip-pickup"
                    className="text-primary bg-white border-gray"
                  />
                  <TopHeading titleHeading="Tip Pickup" />
                </div>
              </div>
            </TopContent>
            <Card>
              <CardBody className="tip-pickup-inner-box">
                <div className="tip-pickup-box d-flex justify-content-center align-items-center">
                  <FormGroup className="d-flex align-items-center px-4">
                    <Label for="deck-position" className="label-name mb-0">
                      Tip Position
                    </Label>
                    <div className="d-flex flex-column input-field">
                      <Select
                        placeholder="Select Option"
                        className=""
                        size="sm"
                        options={taskOptions}
                        onChange={(e) =>
                          formik.setFieldValue("tipPosition", e.value)
                        }
                      />
                      <Label
                        for="tip-pickup"
                        className="font-weight-bold tip-pickup-note mt-2"
                      >
                        200 ul
                      </Label>
                      <FormError>Incorrect Option</FormError>
                    </div>
                  </FormGroup>
                </div>
                <ImageIcon
                  src={tipPickupProcessGraphics}
                  alt="Tip Pickup Process"
                  className="process-image"
                />
              </CardBody>
            </Card>
            <ButtonBar rightBtnLabel="Save" handleRightBtn={handleRightBtn} />
          </div>
        </TipPickupBox>
        <AppFooter />
      </PageBody>
    </>
  );
};

TipPickupComponent.propTypes = {};

export default TipPickupComponent;

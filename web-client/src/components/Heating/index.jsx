import React from "react";

import { Card, CardBody } from "core-components";
import { ButtonIcon, ButtonBar } from "shared-components";

import HeatingProcess from "./HeatingProcess";
import TopHeading from "shared-components/TopHeading";
import { useDispatch, useSelector } from "react-redux";
import { useFormik } from "formik";
import { getFormikInitialState, getRequestBody } from "./functions";
import { PageBody, HeatingBox, TopContent } from "./Style";
import { saveHeatingInitiated } from "action-creators/processesActionCreators";
import { toast } from "react-toastify";

const HeatingComponent = (props) => {
  const dispatch = useDispatch();

  const formik = useFormik({
    initialValues: getFormikInitialState(),
  });

  const recipeDetailsReducer = useSelector(
    (state) => state.updateRecipeDetailsReducer
  );
  const recipeID = recipeDetailsReducer.recipeDetails.id;
  const token = recipeDetailsReducer.token;

  const handleSaveBtn = () => {
    const body = getRequestBody(formik);
    if (body) {
      const requestBody = {
        body: body,
        recipeID: recipeID,
        token: token,
      };
      dispatch(saveHeatingInitiated(requestBody));
    } else {
      //error
      toast.error("Invalid time");
    }
  };

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
            <ButtonBar rightBtnLabel="Save" handleRightBtn={handleSaveBtn} />
          </div>
        </HeatingBox>
      </PageBody>
    </>
  );
};

HeatingComponent.propTypes = {};

export default React.memo(HeatingComponent);

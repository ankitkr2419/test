import React, { useEffect } from "react";

import { Card, CardBody } from "core-components";
import { ButtonIcon, ButtonBar } from "shared-components";

import HeatingProcess from "./HeatingProcess";
import TopHeading from "shared-components/TopHeading";
import { useDispatch, useSelector } from "react-redux";
import { useFormik } from "formik";
import { getFormikInitialState, getRequestBody } from "./functions";
import { PageBody, HeatingBox, TopContent } from "./Style";
import { saveProcessInitiated } from "action-creators/processesActionCreators";
import { toast } from "react-toastify";
import { Redirect, useHistory } from "react-router";
import { API_ENDPOINTS, HTTP_METHODS, ROUTES } from "appConstants";

const HeatingComponent = (props) => {
  const dispatch = useDispatch();
  const history = useHistory();

  const editReducer = useSelector((state) => state.editProcessReducer);
  const editReducerData = editReducer.toJS();
  const processesReducer = useSelector((state) => state.processesReducer);
  const loginReducer = useSelector((state) => state.loginReducer);
  const loginReducerData = loginReducer.toJS();
  let activeDeckObj =
    loginReducerData && loginReducerData.decks.find((deck) => deck.isActive);
  const recipeDetailsReducer = useSelector(
    (state) => state.updateRecipeDetailsReducer
  );

  const formik = useFormik({
    initialValues: editReducerData.process_id
      ? getFormikInitialState(editReducerData)
      : getFormikInitialState(),
    enableReinitialize: true,
  });

  const recipeID = recipeDetailsReducer.recipeDetails.id;
  const token = activeDeckObj.token;

  const errorInAPICall = processesReducer.error;
  useEffect(() => {
    if (errorInAPICall === false) {
      history.push(ROUTES.processListing);
    }
  }, [errorInAPICall]);

  const handleSaveBtn = () => {
    const body = getRequestBody(formik);
    if (body) {
      const requestBody = {
        body: body,
        id: editReducerData?.process_id ? editReducerData.process_id : recipeID,
        token: token,
        api: API_ENDPOINTS.heating,
        method: editReducerData?.process_id
          ? HTTP_METHODS.PUT
          : HTTP_METHODS.POST,
      };
      dispatch(saveProcessInitiated(requestBody));
    } else {
      //error
      toast.error("Invalid Request");
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
            <Card>
              <CardBody className="p-0 overflow-hidden">
                <HeatingProcess formik={formik} />
              </CardBody>
            </Card>
            <ButtonBar
              rightBtnLabel="Save"
              handleRightBtn={handleSaveBtn}
              btnBarClassname={"btn-bar-adjust-heating"}
            />
          </div>
        </HeatingBox>
      </PageBody>
    </>
  );
};

HeatingComponent.propTypes = {};

export default React.memo(HeatingComponent);

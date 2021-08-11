import React, { useEffect } from "react";
import { useFormik } from "formik";
import { useDispatch, useSelector } from "react-redux";

import {
  Card,
  CardBody,
  FormGroup,
  Label,
  Select,
  FormError,
} from "core-components";
import { ButtonIcon, ButtonBar, ImageIcon } from "shared-components";

import tipPickupProcessGraphics from "assets/images/tip-pickup-process-graphics.svg";
import TopHeading from "shared-components/TopHeading";
import { PageBody, TipPickupBox, TopContent } from "./Style";
import { saveProcessInitiated } from "action-creators/processesActionCreators";
import { Redirect, useHistory } from "react-router";
import { API_ENDPOINTS, HTTP_METHODS, ROUTES } from "appConstants";
import { TIP_PICKUP_PROCESS_OPTIONS } from "appConstants";

const TipPickupComponent = () => {
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
  const recipeID = recipeDetailsReducer.recipeDetails.id;
  const token = activeDeckObj.token;

  const formik = useFormik({
    initialValues: {
      tipPosition: editReducerData?.position ? editReducerData.position : null,
    },
    enableReinitialize: true,
  });

  const errorInAPICall = processesReducer.error;
  useEffect(() => {
    if (errorInAPICall === false) {
      history.push(ROUTES.processListing);
    }
  }, [errorInAPICall]);

  const handleRightBtn = () => {
    const tipPosition = formik.values.tipPosition;
    const requestBody = {
      body: { position: parseInt(tipPosition), type: "pickup" },
      id: editReducerData?.process_id ? editReducerData.process_id : recipeID,
      token: token,
      api: API_ENDPOINTS.tipOperation,
      method: editReducerData?.process_id
        ? HTTP_METHODS.PUT
        : HTTP_METHODS.POST,
    };
    dispatch(saveProcessInitiated(requestBody));
  };

  if (!activeDeckObj.isLoggedIn) {
    return <Redirect to={`/${ROUTES.landing}`} />;
  }

  return (
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
                      value={
                        TIP_PICKUP_PROCESS_OPTIONS[
                          formik.values.tipPosition - 1
                        ]
                      }
                      options={TIP_PICKUP_PROCESS_OPTIONS}
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
          <ButtonBar
            rightBtnLabel="Save"
            handleRightBtn={handleRightBtn}
            btnBarClassname={"btn-bar-adjust-tipPickup"}
          />
        </div>
      </TipPickupBox>
    </PageBody>
  );
};

export default React.memo(TipPickupComponent);

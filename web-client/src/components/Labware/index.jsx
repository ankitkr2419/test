import React, { useEffect, useState } from "react";
import { useDispatch, useSelector } from "react-redux";
import { useFormik } from "formik";

import { Card, CardBody } from "core-components";
import { ButtonBar, MlModal } from "shared-components";

import {
  LABWARE_INITIAL_STATE,
  DECKNAME,
  MODAL_BTN,
  ROUTES,
} from "appConstants";
import { updateAllTicks } from "./updateAllTicks";
import { LabwareBox, PageBody } from "./Styles";
import {
  updateRecipeActionInitiated,
  updateRecipeActionReset,
} from "action-creators/saveNewRecipeActionCreators";
import { Redirect, useHistory } from "react-router";
import { getRequestBody } from "./functions";
import Preview from "./Preview";
import SelectProcesses from "./SelectProcesses";

const LabWareComponent = (props) => {
  const [activeTab, setActiveTab] = useState("1");
  const [preview, setPreview] = useState(false);
  const [showConfirmModal, setShowConfirmModal] = useState(false);

  const loginReducer = useSelector((state) => state.loginReducer);
  const recipeDetailsReducer = useSelector(
    (state) => state.updateRecipeDetailsReducer
  );

  const isSuccess = recipeDetailsReducer.isSuccess;

  const dispatch = useDispatch();
  const history = useHistory();

  const formik = useFormik({
    initialValues: LABWARE_INITIAL_STATE,
    enableReinitialize: true,
  });

  useEffect(() => {
    if (isSuccess) {
      setShowConfirmModal(!showConfirmModal);
      dispatch(updateRecipeActionReset());
    }
  }, [isSuccess]);

  const loginReducerData = loginReducer.toJS();
  let activeDeckObj =
    loginReducerData && loginReducerData.decks.find((deck) => deck.isActive);
  if (!activeDeckObj.isLoggedIn) {
    return <Redirect to={`/${ROUTES.landing}`} />;
  }

  const tipsOptions = recipeDetailsReducer.tipsOptions;
  const tubesOptions = recipeDetailsReducer.tubesOptions;
  const cartridgeOptions = recipeDetailsReducer.cartridgeOptions;
  const newRecipeName = recipeDetailsReducer.recipeDetails.name;

  const handleSaveBtn = () => {
    const requestBody = getRequestBody(newRecipeName, formik.values);
    dispatch(
      updateRecipeActionInitiated({
        requestBody: requestBody,
        deckName: DECKNAME.DeckA,
        token: activeDeckObj.token,
      })
    );
  };

  const toggle = (tab) => activeTab !== tab && setActiveTab(tab);

  const handleSuccessBtn = () => {
    setShowConfirmModal(!showConfirmModal);
    history.push(ROUTES.selectProcess); //redirect to next page
  };

  return (
    <>
      {showConfirmModal && (
        <MlModal
          textHead={activeDeckObj.name}
          textBody={"Labware is set!"}
          isOpen={showConfirmModal}
          successBtn={MODAL_BTN.okay}
          handleSuccessBtn={handleSuccessBtn}
          handleCrossBtn={() => {
            setShowConfirmModal(!showConfirmModal);
          }}
        />
      )}
      <PageBody>
        <LabwareBox>
          <div className="process-content process-labware px-2">
            <Card className="labware-card-box">
              <CardBody className="p-0 overflow-hidden">
                {/* Preview Body */}
                {preview ? (
                  <Preview formik={formik} />
                ) : (
                  // { /* Select Process */ }
                  <SelectProcesses
                    activeTab={activeTab}
                    toggle={toggle}
                    formik={formik}
                    tubesOptions={tubesOptions}
                    tipsOptions={tipsOptions}
                    cartridgeOptions={cartridgeOptions}
                  />
                )}
              </CardBody>

              {/* Bottom Button Bar */}
              <div className="bottom-btn-bar">
                {preview ? (
                  <ButtonBar
                    handleLeftBtn={() => setPreview(!preview)}
                    handleRightBtn={handleSaveBtn}
                    leftBtnLabel={"Modify"}
                    rightBtnLabel={"Save"}
                    btnBarClassname={"btn-bar-adjust-labware"}
                  />
                ) : (
                  <ButtonBar
                    handleRightBtn={() => {
                      setPreview(!preview);
                      updateAllTicks(formik);
                    }}
                    rightBtnLabel={"Preview"}
                    btnBarClassname={"btn-bar-adjust-labware"}
                  />
                )}
              </div>
            </Card>
          </div>
        </LabwareBox>
      </PageBody>
    </>
  );
};

export default React.memo(LabWareComponent);

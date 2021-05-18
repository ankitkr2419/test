import React, { useEffect, useState } from "react";
import { useDispatch, useSelector } from "react-redux";
import { useFormik } from "formik";

import { Card, CardBody } from "core-components";
import { ButtonBar, MlModal } from "shared-components";
import AppFooter from "components/AppFooter";

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
  });

  const loginReducer = useSelector((state) => state.loginReducer);
  const recipeDetailsReducer = useSelector(
    (state) => state.updateRecipeDetailsReducer
  );

  const loginReducerData = loginReducer.toJS();
  let activeDeckObj =
    loginReducerData && loginReducerData.decks.find((deck) => deck.isActive);
  if (!activeDeckObj.isLoggedIn) {
    return <Redirect to={`/${ROUTES.landing}`} />;
  }

  const deckIndex = activeDeckObj.name === DECKNAME.DeckA ? 0 : 1;
  const tubesOptions = recipeDetailsReducer.decks[deckIndex].tubesOptions;
  const tipsOptions = recipeDetailsReducer.decks[deckIndex].tipsOptions;
  const tipsAndTubesOptions =
    recipeDetailsReducer.decks[deckIndex].recipeOptions;
  const cartridgeOptions =
    recipeDetailsReducer.decks[deckIndex].cartridgeOptions;
  const newRecipeName =
    recipeDetailsReducer.decks[deckIndex].recipeDetails.name;

  const isSuccess = recipeDetailsReducer.isSuccess;

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
                    tipsAndTubesOptions={tipsAndTubesOptions}
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
                  />
                ) : (
                  <ButtonBar
                    handleRightBtn={() => {
                      setPreview(!preview);
                      updateAllTicks(formik);
                    }}
                    rightBtnLabel={"Preview"}
                  />
                )}
              </div>
            </Card>
          </div>
          <AppFooter />
        </LabwareBox>
      </PageBody>
    </>
  );
};

export default React.memo(LabWareComponent);

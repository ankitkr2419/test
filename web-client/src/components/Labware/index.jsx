import React, { useEffect, useState } from "react";
import { useDispatch, useSelector } from "react-redux";
import { useFormik } from "formik";

import { Row, Col, Card, CardBody } from "core-components";
import { ButtonBar, ImageIcon, Text, Icon, MlModal } from "shared-components";
import { TabContent, TabPane, Nav } from "reactstrap";

import AppFooter from "components/AppFooter";
import labwarePlate from "assets/images/labware-plate.png";
import {
  LABWARE_INITIAL_STATE,
  DECKNAME,
  MODAL_BTN,
  ROUTES,
} from "appConstants";
import {
  getSideBarNavItems,
  getTipsAtPosition,
  getTipPiercingAtPosition,
  getPreviewInfo,
  updateAllTicks,
  getFieldAtPosition,
} from "./HelperFunctions";
import { LabwareBox, PageBody, ProcessSetting } from "./Styles";
import {
  updateRecipeActionInitiated,
  updateRecipeActionReset,
} from "action-creators/saveNewRecipeActionCreators";
import { Redirect, useHistory } from "react-router";
import { getRequestBody, getOptions } from "./functions";
// import { Preview } from "./Preview";

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

  const loginReducer = useSelector((state) => state.loginReducer);
  const recipeDetailsReducer = useSelector(
    (state) => state.updateRecipeDetailsReducer
  );

  const isSuccess = recipeDetailsReducer.isSuccess;

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
  const tipsAndTubesOptions = recipeDetailsReducer.tipsAndTubesOptions;
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

  const toggle = (tab) => {
    if (activeTab !== tab) setActiveTab(tab);
  };

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
                  <div className="w-100 h-100 preview-box">
                    <Row>
                      <Col
                        md={12}
                        className="d-flex align-items-center font-weight-bold text-center top-heading"
                      >
                        Preview
                      </Col>
                    </Row>
                    <div className="d-flex justify-content-between">
                      <div className="labware-selection-info w-100">
                        <Text className="setting-info font-weight-bold selected-positions">
                          Selected Positions
                        </Text>
                        <ul className="list-unstyled">
                          {getPreviewInfo(formik)}
                          {/* {formik.values && <Preview recipeData={formik.values} />} */}
                        </ul>
                      </div>

                      <div className="img-box">
                        <ProcessSetting>
                          <div className="tips-info">
                            <ul className="list-unstyled tip-position active">
                              {formik.values.tips.processDetails.tipPosition1
                                .id && (
                                <li className="highlighted tip-position-1"></li>
                              )}
                              {formik.values.tips.processDetails.tipPosition2
                                .id && (
                                <li className="highlighted tip-position-2"></li>
                              )}
                              {formik.values.tips.processDetails.tipPosition3
                                .id && (
                                <li className="highlighted tip-position-3"></li>
                              )}
                            </ul>
                          </div>

                          <div className="piercing-info">
                            <ul className="list-unstyled piercing-position active">
                              {formik.values.tipPiercing.processDetails
                                .position1.id && (
                                <li className="highlighted piercing-position-1"></li>
                              )}
                              {formik.values.tipPiercing.processDetails
                                .position2.id && (
                                <li className="highlighted piercing-position-2"></li>
                              )}
                            </ul>
                          </div>

                          <div className="deck-position-info">
                            <ul className="list-unstyled deck-position active">
                              {formik.values.deckPosition1.processDetails
                                .tubeType.id && (
                                <li className="highlighted deck-position-1 active" />
                              )}
                              {formik.values.deckPosition2.processDetails
                                .tubeType.id && (
                                <li className="highlighted deck-position-2 active" />
                              )}
                              {formik.values.deckPosition3.processDetails
                                .tubeType.id && (
                                <li className="highlighted deck-position-3 active" />
                              )}
                              {formik.values.deckPosition4.processDetails
                                .tubeType.id && (
                                <li className="highlighted deck-position-4 active" />
                              )}
                            </ul>
                          </div>

                          <div className="cartridge-position-info">
                            <ul className="list-unstyled cartridge-position active">
                              {formik.values.cartridge1.processDetails
                                .cartridgeType.id && (
                                <li className="highlighted cartridge-position-1 active" />
                              )}
                              {formik.values.cartridge2.processDetails
                                .cartridgeType.id && (
                                <li className="highlighted cartridge-position-2 active" />
                              )}
                            </ul>
                          </div>

                          <ImageIcon
                            src={labwarePlate}
                            alt="Labware Plate"
                            className=""
                          />
                        </ProcessSetting>
                      </div>
                    </div>
                  </div>
                ) : (
                  // Select Process
                  <div className="d-flex">
                    <Nav tabs className="d-flex flex-column border-0 side-bar">
                      <Text className="d-flex justify-content-center align-items-center px-3 pt-3 pb-3 mb-0 font-weight-bold text-white">
                        <Icon name="setting" size={18} />
                        <Text Tag="span" className="ml-2">
                          Settings{" "}
                        </Text>
                      </Text>
                      {getSideBarNavItems(formik, activeTab, toggle)}
                    </Nav>

                    <TabContent activeTab={activeTab} className="flex-grow-1">
                      <TabPane tabId="1">
                        {getOptions(1, 3, tipsAndTubesOptions) &&
                          getTipsAtPosition(
                            1,
                            formik,
                            getOptions(1, 3, tipsAndTubesOptions)
                          )}
                      </TabPane>
                      <TabPane tabId="2">
                        {getTipPiercingAtPosition(1, formik)}
                      </TabPane>
                      <TabPane tabId="3">
                        {getFieldAtPosition(
                          1,
                          formik,
                          tubesOptions,
                          "deckPosition"
                        )}
                      </TabPane>
                      <TabPane tabId="4">
                        {getFieldAtPosition(
                          2,
                          formik,
                          tubesOptions,
                          "deckPosition"
                        )}
                      </TabPane>
                      <TabPane tabId="5">
                        {getFieldAtPosition(
                          1,
                          formik,
                          cartridgeOptions,
                          "cartridge"
                        )}
                      </TabPane>
                      <TabPane tabId="6">
                        {getFieldAtPosition(
                          3,
                          formik,
                          tubesOptions,
                          "deckPosition"
                        )}
                      </TabPane>
                      <TabPane tabId="7">
                        {getFieldAtPosition(
                          2,
                          formik,
                          cartridgeOptions,
                          "cartridge"
                        )}
                      </TabPane>
                      <TabPane tabId="8">
                        {getFieldAtPosition(
                          4,
                          formik,
                          tubesOptions,
                          "deckPosition"
                        )}
                      </TabPane>
                    </TabContent>
                  </div>
                )}
              </CardBody>
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

LabWareComponent.propTypes = {};

export default React.memo(LabWareComponent);

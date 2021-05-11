import React, { useState } from "react";
import { useSelector } from "react-redux";
import { useFormik } from "formik";

import { Row, Col, Card, CardBody } from "core-components";
import { ButtonBar, ImageIcon, Text, Icon } from "shared-components";
import { TabContent, TabPane, Nav } from "reactstrap";

import AppFooter from "components/AppFooter";
import labwarePlate from "assets/images/labware-plate.png";
import { LABWARE_INITIAL_STATE, DECKNAME } from "appConstants";
import {
  getSideBarNavItems,
  getDeckAtPosition,
  getCartidgeAtPosition,
  getTipsAtPosition,
  getTipPiercingAtPosition,
  getPreviewInfo,
  updateAllTicks,
} from "./HelperFunctions";
import { LabwareBox, PageBody, ProcessSetting } from "./Styles";

const LabWareComponent = (props) => {
  const [activeTab, setActiveTab] = useState("1");
  const [preview, setPreview] = useState(false);

  const formik = useFormik({
    initialValues: LABWARE_INITIAL_STATE,
    enableReinitialize: true,
  });

  const loginReducer = useSelector((state) => state.loginReducer);
  const loginReducerData = loginReducer.toJS();
  let activeDeckObj =
    loginReducerData && loginReducerData.decks.find((deck) => deck.isActive);
  // if (!activeDeckObj.isLoggedIn) {
  //   return <Redirect to={`/${ROUTES.landing}`} />;
  // }

  const recipeDetailsReducer = useSelector(
    (state) => state.updateRecipeDetailsReducer
  );
  const deckIndex = activeDeckObj.name === DECKNAME.DeckA ? 0 : 1;
  const tipsAndTubesOptions =
    recipeDetailsReducer.decks[deckIndex].recipeOptions;
  const cartridgeOptions =
    recipeDetailsReducer.decks[deckIndex].cartridgeOptions;

  const getOptions = (lowerLimit, higherLimit, options) => {
    const optionsObj = [];
    if (options) {
      options.forEach((optionObj) => {
        if (optionObj.id >= lowerLimit && optionObj.id <= higherLimit) {
          optionsObj.push({
            value: optionObj.id,
            label: optionObj.name ? optionObj.name : optionObj.description,
          });
        }
      });
    }
    return optionsObj;
  };

  const handleSaveBtn = () => {
    const selectedOptions = formik.values;
    const requestBody = {
      pos_1: selectedOptions.tips.processDetails.tipPosition1.id,
      pos_2: selectedOptions.tips.processDetails.tipPosition2.id,
      pos_3: selectedOptions.tips.processDetails.tipPosition3.id,
      pos_4: selectedOptions.tipPiercing.processDetails.position1.id,
      pos_5: selectedOptions.tipPiercing.processDetails.position2.id,
      pos_6: selectedOptions.deckPosition1.processDetails.tubeType.id,
      pos_7: selectedOptions.deckPosition2.processDetails.tubeType.id,
      pos_cartridge_1: selectedOptions.cartridge1.processDetails.cartridgeType.id,
      pos_9: selectedOptions.deckPosition3.processDetails.tubeType.id,
      pos_cartridge_2: selectedOptions.cartridge2.processDetails.cartridgeType.id,
      pos_11: selectedOptions.deckPosition4.processDetails.tubeType.id,
    };
    console.log(requestBody);
  };

  const toggle = (tab) => {
    if (activeTab !== tab) setActiveTab(tab);
  };

  return (
    <>
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
                        </ul>
                      </div>

                      <div className="img-box">
                        <ProcessSetting>
                          <div className="tips-info">
                            <ul class="list-unstyled tip-position active">
                              {formik.values.tips.processDetails
                                .tipPosition1.id && (
                                <li class="highlighted tip-position-1"></li>
                              )}
                              {formik.values.tips.processDetails
                                .tipPosition2.id && (
                                <li class="highlighted tip-position-2"></li>
                              )}
                              {formik.values.tips.processDetails
                                .tipPosition3.id && (
                                <li class="highlighted tip-position-3"></li>
                              )}
                            </ul>
                          </div>

                          <div className="piercing-info">
                            <ul class="list-unstyled piercing-position active">
                              {formik.values.tipPiercing.processDetails
                                .position1.id && (
                                <li class="highlighted piercing-position-1"></li>
                              )}
                              {formik.values.tipPiercing.processDetails
                                .position2.id && (
                                <li class="highlighted piercing-position-2"></li>
                              )}
                            </ul>
                          </div>

                          <div className="deck-position-info">
                            <ul class="list-unstyled deck-position active">
                              {formik.values.deckPosition1.processDetails
                                .tubeType.id && (
                                <li class="highlighted deck-position-1 active" />
                              )}
                              {formik.values.deckPosition2.processDetails
                                .tubeType.id && (
                                <li class="highlighted deck-position-2 active" />
                              )}
                              {formik.values.deckPosition3.processDetails
                                .tubeType.id && (
                                <li class="highlighted deck-position-3 active" />
                              )}
                              {formik.values.deckPosition4.processDetails
                                .tubeType.id && (
                                <li class="highlighted deck-position-4 active" />
                              )}
                            </ul>
                          </div>

                          <div className="cartridge-position-info">
                            <ul class="list-unstyled cartridge-position active">
                              {formik.values.cartridge1.processDetails
                                .cartridgeType.id && (
                                <li class="highlighted cartridge-position-1 active" />
                              )}
                              {formik.values.cartridge2.processDetails
                                .cartridgeType.id && (
                                <li class="highlighted cartridge-position-2 active" />
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
                        {getOptions(4, 4, tipsAndTubesOptions) &&
                          getDeckAtPosition(
                            1,
                            formik,
                            getOptions(4, 4, tipsAndTubesOptions)
                          )}
                      </TabPane>
                      <TabPane tabId="4">
                        {getOptions(5, 5, tipsAndTubesOptions) &&
                          getDeckAtPosition(
                            2,
                            formik,
                            getOptions(5, 5, tipsAndTubesOptions)
                          )}
                      </TabPane>
                      <TabPane tabId="5">
                        {getOptions(1, 1, cartridgeOptions) &&
                          getCartidgeAtPosition(
                            1,
                            formik,
                            getOptions(1, 1, cartridgeOptions)
                          )}
                      </TabPane>
                      <TabPane tabId="6">
                        {getOptions(6, 6, tipsAndTubesOptions) &&
                          getDeckAtPosition(
                            3,
                            formik,
                            getOptions(6, 6, tipsAndTubesOptions)
                          )}
                      </TabPane>
                      <TabPane tabId="7">
                        {getOptions(2, 2, cartridgeOptions) &&
                          getCartidgeAtPosition(
                            2,
                            formik,
                            getOptions(2, 2, cartridgeOptions)
                          )}
                      </TabPane>
                      <TabPane tabId="8">
                        {getOptions(7, 7, tipsAndTubesOptions) &&
                          getDeckAtPosition(
                            4,
                            formik,
                            getOptions(7, 7, tipsAndTubesOptions)
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

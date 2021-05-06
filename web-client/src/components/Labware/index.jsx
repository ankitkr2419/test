import React, { useState } from "react";
import { useSelector } from "react-redux";
import { useFormik } from "formik";

import { Row, Col, Card, CardBody } from "core-components";
import { ButtonBar, ImageIcon, Text, Icon } from "shared-components";
import { TabContent, TabPane, Nav } from "reactstrap";

import AppFooter from "components/AppFooter";
import labwarePlate from "assets/images/labware-plate.png";
import { LABWARE_INITIAL_STATE, ROUTES } from "appConstants";
import {
  getSideBarNavItems,
  getDeckAtPosition,
  getCartidgeAtPosition,
  getTipsAtPosition,
  getTipPiercingAtPosition,
  getPreviewInfo,
} from "./HelperFunctions";
import { LabwareBox, PageBody, ProcessSetting } from "./Styles";
import { Redirect } from "react-router";

const LabWareComponent = (props) => {
  const [activeTab, setActiveTab] = useState("1");
  const [preview, setPreview] = useState(true);

  const formik = useFormik({
    initialValues: LABWARE_INITIAL_STATE,
  });

  // const loginReducer = useSelector((state) => state.loginReducer);
  // const loginReducerData = loginReducer.toJS();
  // let activeDeckObj =
  //   loginReducerData && loginReducerData.decks.find((deck) => deck.isActive);
  // if (!activeDeckObj.isLoggedIn) {
  //   return <Redirect to={`/${ROUTES.landing}`} />;
  // }

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
                          <div className="deck-position-info">
                            <ul class="list-unstyled deck-position active">
                              <li class="highlighted deck-position-4 active" />
                            </ul>
                            <ImageIcon
                              src={labwarePlate}
                              alt="Labware Plate"
                              className=""
                            />
                          </div>
                        </ProcessSetting>
                      </div>
                    </div>
                  </div>
                ) : (
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
                        {getTipsAtPosition(1, formik)}
                      </TabPane>
                      <TabPane tabId="2">
                        {getTipPiercingAtPosition(1, formik)}
                      </TabPane>
                      <TabPane tabId="3">
                        {getDeckAtPosition(1, formik)}
                      </TabPane>
                      <TabPane tabId="4">
                        {getDeckAtPosition(2, formik)}
                      </TabPane>
                      <TabPane tabId="5">
                        {getCartidgeAtPosition(1, formik)}
                      </TabPane>
                      <TabPane tabId="6">
                        {getDeckAtPosition(3, formik)}
                      </TabPane>
                      <TabPane tabId="7">
                        {getCartidgeAtPosition(2, formik)}
                      </TabPane>
                      <TabPane tabId="8">
                        {getDeckAtPosition(4, formik)}
                      </TabPane>
                    </TabContent>
                  </div>
                )}
              </CardBody>
              <div className="bottom-btn-bar">
                {preview ? (
                  <ButtonBar
                    handleLeftBtn={() => setPreview(!preview)}
                    handleRightBtn={getPreviewInfo}
                    leftBtnLabel={"Modify"}
                    rightBtnLabel={"Save"}
                  />
                ) : (
                  <ButtonBar
                    handleRightBtn={() => setPreview(!preview)}
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

export default LabWareComponent;

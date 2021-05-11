import React, { useState } from "react";
import { useSelector } from "react-redux";
import { useFormik } from "formik";

import { Row, Col, Card, CardBody } from "core-components";
import { ButtonBar, ImageIcon, Text, Icon } from "shared-components";
import { TabContent, TabPane, Nav } from "reactstrap";

import AppFooter from "components/AppFooter";
import labwarePlate from "assets/images/labware-plate.png";
import {
  LABWARE_INITIAL_STATE, ROUTES,
} from "appConstants";
import {
  getSideBarNavItems,
  getDeckAtPosition,
  getCartidgeAtPosition,
  getTipsAtPosition,
  getTipPiercingAtPosition,
} from "./HelperFunctions";
import { LabwareBox, PageBody, ProcessSetting } from "./Styles";
import { Redirect } from "react-router";

const LabWareComponent = (props) => {
  const [activeTab, setActiveTab] = useState("1");
  
  const formik = useFormik({
    initialValues: LABWARE_INITIAL_STATE,
  });

  const loginReducer = useSelector((state) => state.loginReducer);

  const loginReducerData = loginReducer.toJS();
  let activeDeckObj =
    loginReducerData && loginReducerData.decks.find((deck) => deck.isActive);
  if (!activeDeckObj.isLoggedIn) {
    return <Redirect to={`/${ROUTES.landing}`} />;
  }

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
                    <TabPane tabId="1">{getTipsAtPosition(1, formik)}</TabPane>
                    <TabPane tabId="2">{getTipPiercingAtPosition(1, formik)}</TabPane>
                    <TabPane tabId="3">{getDeckAtPosition(1, formik)}</TabPane>
                    <TabPane tabId="4">{getDeckAtPosition(2, formik)}</TabPane>
                    <TabPane tabId="5">{getCartidgeAtPosition(1, formik)}</TabPane>
                    <TabPane tabId="6">{getDeckAtPosition(3, formik)}</TabPane>
                    <TabPane tabId="7">{getCartidgeAtPosition(2, formik)}</TabPane>
                    <TabPane tabId="8">{getDeckAtPosition(4, formik)}</TabPane>
                  </TabContent>
                </div>

                <div className="w-100 h-100 preview-box d-none">
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
                        <li className="d-flex justify-content-between">
                          <div className="w-25 font-weight-bold">
                            <Text>Tips : </Text>
                          </div>
                          <div className="w-75">
                            <div className="ml-2 setting-value">
                              <Text>
                                <Text Tag="span" className="font-weight-bold">
                                  Tip Position 1 :
                                </Text>{" "}
                                <Text Tag="span" className="">
                                  Extraction 200ul{" "}
                                </Text>
                              </Text>
                              <Text>
                                <Text Tag="span" className="font-weight-bold">
                                  Tip Position 3 :{" "}
                                </Text>
                                <Text Tag="span" className="">
                                  PCR 40ul{" "}
                                </Text>
                              </Text>
                            </div>
                          </div>
                        </li>
                        <li className="">
                          <div className="d-flex">
                            <Text className="w-25 font-weight-bold">
                              Tip Piercing :
                            </Text>
                            <Text className="w-75">
                              <div className="ml-2 setting-value">
                                Position 2
                              </div>
                            </Text>
                          </div>
                        </li>
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
              </CardBody>
              <div className="bottom-btn-bar">
                <ButtonBar handleTemp={() => { console.log(formik.values); }}/>
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

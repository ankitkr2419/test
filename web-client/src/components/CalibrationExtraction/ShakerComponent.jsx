import React, { useState } from "react";
import { useFormik } from "formik";
import { useDispatch, useSelector } from "react-redux";
import { Redirect, useHistory } from "react-router";
import classnames from "classnames";
import { toast } from "react-toastify";

import { Card, CardBody } from "core-components";
import { ButtonIcon, ButtonBar } from "shared-components";

import { TabContent, TabPane, Nav, NavItem, NavLink } from "reactstrap";
import ShakingProcess from "./ShakingProcess";
import TopHeading from "shared-components/TopHeading";
import { PageBody, TopContent, ShakingBox } from "./Style";
import { shakerInitialFormikState, getShakerRequestBody } from "./helpers";
import { DECKNAME, ROUTES } from "appConstants";
import {
  abort,
  shakerInitiated,
} from "action-creators/calibrationActionCreators";

const ShakerComponent = (props) => {
  const [activeTab, setActiveTab] = useState("2");
  const dispatch = useDispatch();
  const history = useHistory();

  const loginReducer = useSelector((state) => state.loginReducer);
  const loginReducerData = loginReducer.toJS();
  let activeDeckObj =
    loginReducerData && loginReducerData.decks.find((deck) => deck.isActive);

  const { isLoggedIn, token, name } = activeDeckObj;

  const formik = useFormik({
    initialValues: shakerInitialFormikState,
  });

  const toggle = (tab) => {
    if (activeTab !== tab) setActiveTab(tab);
  };

  const handleStartBtn = () => {
    const requestBody = {
      body: getShakerRequestBody(formik, activeTab),
      token: token,
      deckName:
        name === DECKNAME.DeckA ? DECKNAME.DeckAShort : DECKNAME.DeckBShort,
    };
    if (requestBody) {
      dispatch(shakerInitiated(requestBody));
    } else {
      toast.warning("Please check inputted values");
    }
  };

  const handleAbortBtn = () => {
    const deckName =
      name === DECKNAME.DeckA ? DECKNAME.DeckAShort : DECKNAME.DeckBShort;
    dispatch(abort(token, deckName));
  };

  if (!isLoggedIn) {
    return <Redirect to={`/${ROUTES.landing}`} />;
  }

  return (
    <>
      <PageBody>
        <ShakingBox>
          <div className="process-content process-shaking px-2">
            <TopContent className="d-flex justify-content-between align-items-center mx-5">
              <div className="d-flex flex-column">
                <div className="d-flex align-items-center frame-icon">
                  <ButtonIcon
                    size={60}
                    name="shaking"
                    className="text-primary bg-white border-gray"
                    // onClick={toggleExportDataModal}
                  />
                  <TopHeading titleHeading="Shaking" />
                </div>
              </div>
            </TopContent>
            <Card>
              <CardBody className="p-0 overflow-hidden">
                <Nav
                  tabs
                  className="bg-white px-3 pb-0 d-flex justify-content-center align-items-center border-0"
                >
                  <NavItem className="text-center flex-fill px-2 pt-2">
                    <NavLink
                      className={classnames({ active: activeTab === "1" })}
                      onClick={() => {
                        toggle("1");
                      }}
                    >
                      Without heating
                    </NavLink>
                  </NavItem>
                  <NavItem className="text-center flex-fill px-2 pt-2">
                    <NavLink
                      className={classnames({ active: activeTab === "2" })}
                      onClick={() => {
                        toggle("2");
                      }}
                    >
                      With heating
                    </NavLink>
                  </NavItem>
                </Nav>
                <TabContent activeTab={activeTab} className="p-5">
                  <TabPane tabId="1">
                    <ShakingProcess formik={formik} activeTab={activeTab} />
                  </TabPane>
                  <TabPane tabId="2">
                    <ShakingProcess
                      formik={formik}
                      activeTab={activeTab}
                      temperature
                    />
                  </TabPane>
                </TabContent>
              </CardBody>
            </Card>
            <ButtonBar
              rightBtnLabel="Start"
              leftBtnLabel="Abort"
              handleRightBtn={handleStartBtn}
              handleLeftBtn={handleAbortBtn}
              btnBarClassname={"btn-bar-adjust-shaking"}
            />
          </div>
        </ShakingBox>
      </PageBody>
    </>
  );
};

ShakerComponent.propTypes = {};

export default ShakerComponent;

import React, { useEffect, useState } from "react";
import { useFormik } from "formik";
import { useDispatch, useSelector } from "react-redux";
import { Redirect, useHistory } from "react-router";
import classnames from "classnames";
import { toast } from "react-toastify";

import { Card, CardBody } from "core-components";
import {
  ButtonIcon,
  ButtonBar,
  Text,
  ColoredCircle,
  HomingModal,
} from "shared-components";

import {
  TabContent,
  TabPane,
  Nav,
  NavItem,
  NavLink,
  Spinner,
} from "reactstrap";
import { showHomingModal as showHomingModalAction } from "action-creators/homingActionCreators";
import ShakingProcess from "./ShakingProcess";
import TopHeading from "shared-components/TopHeading";
import { PageBody, TopContent, ShakingBox } from "./Style";
import { shakerInitialFormikState, getShakerRequestBody } from "./helpers";
import { DECKNAME, ROUTES, SHAKER_RUN_STATUS } from "appConstants";
import {
  abort,
  shakerInitiated,
} from "action-creators/calibrationActionCreators";

const ShakerComponent = (props) => {
  const [activeTab, setActiveTab] = useState("2");
  const dispatch = useDispatch();

  const loginReducer = useSelector((state) => state.loginReducer);
  const loginReducerData = loginReducer.toJS();
  let activeDeckObj =
    loginReducerData && loginReducerData.decks.find((deck) => deck.isActive);

  const heaterReducer = useSelector((state) => state.heaterProgressReducer);
  const heaterProgressReducerData = heaterReducer.toJS();
  const { heater_on, shaker_1_temp, shaker_2_temp } =
    heaterProgressReducerData.data;

  const shakerRunProgessReducer = useSelector(
    (state) => state.shakerRunProgessReducer
  );
  const { shakerRunStatus } = shakerRunProgessReducer.toJS();
  const { progressing, progressAborted } = SHAKER_RUN_STATUS;

  const { isLoggedIn, token, name } = activeDeckObj;

  // if progress is aborted then open homing modal
  useEffect(() => {
    if (shakerRunStatus === progressAborted) {
      dispatch(showHomingModalAction());
    }
  }, [shakerRunStatus]);

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
      toast.warning("Please check inputted values", { autoClose: false });
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

  const getRightBtnLabel = () => {
    if (shakerRunStatus === progressing) {
      return (
        <div className="d-flex">
          <Spinner size="sm" />
          <Text className="ml-3">Shaking</Text>
        </div>
      );
    } else {
      return "Start";
    }
  };

  return (
    <>
      <HomingModal />
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
              <Card default className="ml-auto mr-5 rounded-lg bg-transparent">
                <CardBody className="d-flex p-2">
                  <Text className="font-weight-bold mr-3 text-muted">
                    Heater Status: <ColoredCircle isOnline={heater_on} />
                    {"  "}
                  </Text>
                  <Text className="font-weight-bold m-0 text-muted">
                    Heater Temperature 1: {shaker_1_temp ? shaker_1_temp : 0}?? C
                    <br />
                    Heater Temperature 2: {shaker_2_temp ? shaker_2_temp : 0}?? C
                  </Text>
                </CardBody>
              </Card>
            </TopContent>
            <Card className="fix-height-card">
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
              rightBtnLabel={getRightBtnLabel()}
              leftBtnLabel="Abort"
              handleRightBtn={handleStartBtn}
              handleLeftBtn={handleAbortBtn}
              btnBarClassname={"btn-bar-adjust-shaking"}
              isRightBtnDisabled={shakerRunStatus === progressing}
              isLeftBtnDisabled={!(shakerRunStatus === progressing)}
            />
          </div>
        </ShakingBox>
      </PageBody>
    </>
  );
};

ShakerComponent.propTypes = {};

export default ShakerComponent;

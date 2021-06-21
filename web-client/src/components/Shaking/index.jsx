import React, { useEffect, useState } from "react";

import { Card, CardBody } from "core-components";
import { ButtonIcon, ButtonBar } from "shared-components";

import { TabContent, TabPane, Nav, NavItem, NavLink } from "reactstrap";
import classnames from "classnames";
import ShakingProcess from "./ShakingProcess";
import TopHeading from "shared-components/TopHeading";
import { PageBody, TopContent, ShakingBox } from "./Style";
import { useFormik } from "formik";
import { isDisabled, getFormikInitialState, getRequestBody } from "./helpers";
import { useDispatch, useSelector } from "react-redux";
import { saveProcessInitiated } from "action-creators/processesActionCreators";
import { toast } from "react-toastify";
import { Redirect, useHistory } from "react-router";
import { API_ENDPOINTS, HTTP_METHODS, ROUTES } from "appConstants";

const ShakingComponent = (props) => {
  const [activeTab, setActiveTab] = useState("2");
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

  useEffect(() => {
    if (editReducerData.process_id) {
      const selectedTab = editReducerData.with_temp ? "2" : "1";
      setActiveTab(selectedTab);
    }
  }, [editReducerData.process_id]);

  const errorInAPICall = processesReducer.error;
  useEffect(() => {
    if (errorInAPICall === false) {
      history.push(ROUTES.processListing);
    }
  }, [errorInAPICall]);

  const toggle = (tab) => {
    if (activeTab !== tab) setActiveTab(tab);
  };

  const handleSaveBtn = () => {
    const body = getRequestBody(formik, activeTab);

    console.log(body);

    if (body) {
      const requestBody = {
        body: body,
        id: editReducerData?.process_id ? editReducerData.process_id : recipeID,
        token: token,
        api: API_ENDPOINTS.shaking,
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
                      // disabled={isDisabled.withoutHeating} //This feature may get activated later
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
                      // disabled={isDisabled.withHeating}  //This feature may get activated later
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
                      temperature={true}
                    />
                  </TabPane>
                </TabContent>
              </CardBody>
            </Card>
            <ButtonBar
              rightBtnLabel="Save"
              handleRightBtn={handleSaveBtn}
              btnBarClassname={"btn-bar-adjust-shaking"}
            />
          </div>
        </ShakingBox>
      </PageBody>
    </>
  );
};

ShakingComponent.propTypes = {};

export default ShakingComponent;

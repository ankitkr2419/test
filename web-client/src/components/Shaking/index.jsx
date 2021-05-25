import React, { useState } from "react";

import { Card, CardBody } from "core-components";
import { ButtonIcon, ButtonBar } from "shared-components";

import AppFooter from "components/AppFooter";
import { TabContent, TabPane, Nav, NavItem, NavLink } from "reactstrap";
import classnames from "classnames";
import ShakingProcess from "./ShakingProcess";
import TopHeading from "shared-components/TopHeading";
import { PageBody, TopContent, ShakingBox } from "./Style";

const ShakingComponent = (props) => {
  const [activeTab, setActiveTab] = useState("1");
  const withHeating = useState(true);

  const toggle = (tab) => {
    if (activeTab !== tab) setActiveTab(tab);
  };

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
                    <ShakingProcess />
                  </TabPane>
                  <TabPane tabId="2">
                    <ShakingProcess temperature={true} />
                  </TabPane>
                </TabContent>
              </CardBody>
            </Card>
            <ButtonBar rightBtnLabel="Save" />
          </div>
        </ShakingBox>
      </PageBody>
    </>
  );
};

ShakingComponent.propTypes = {};

export default ShakingComponent;

import React, { useState } from "react";

import { Card, CardBody } from "core-components";
import { ButtonIcon, ButtonBar, Text, Icon } from "shared-components";

import { TabContent, TabPane, Nav, NavItem, NavLink } from "reactstrap";
import classnames from "classnames";
import AspireCommonField from "./AspireCommonField";
import DispenseCommonField from "./DispenseCommonField";
import TopHeading from "shared-components/TopHeading";
import { AspireDispenseBox, PageBody, TopContent } from "./Style";
import { ASPIRE_DISPENSE_SIDEBAR_LABELS } from "appConstants";
import { WellComponent } from "./WellComponent";
import { getArray } from "./functions";
import CommonDeckPosition from "./CommonDeckPosition";
import { useSelector } from "react-redux";

const aspireCart1Wells = getArray(8, 0);
const aspireCart2Wells = getArray(8, 0);
const dispenseCart1Wells = getArray(8, 1);
const dispenseCart2Wells = getArray(8, 1);

const AspireDispenseComponent = (props) => {
  const [activeTab, setActiveTab] = useState("1");
  const [isAspire, setIsAspire] = useState(true);
  const [selectedTab, setSelectedTab] = useState(false);

  const recipeDetailsReducer = useSelector(
    (state) => state.updateRecipeDetailsReducer
  );
  const recipeID = recipeDetailsReducer.recipeDetails.id;
  const token = recipeDetailsReducer.token;

  const toggle = (tab) => {
    if (activeTab !== tab) setActiveTab(tab);
  };

  const handleWellClick = () => {};

  const handleTabElementChange = (e) => {
    !selectedTab && setSelectedTab(activeTab);
    console.log("Selected Tab: ", selectedTab);
  };

  return (
    <>
      <PageBody>
        <AspireDispenseBox>
          <div className="process-content process-aspire-dispense px-2">
            <TopContent className="d-flex justify-content-between align-items-center mx-5">
              <div className="d-flex flex-column">
                <div className="d-flex align-items-center frame-icon">
                  <ButtonIcon
                    size={60}
                    name="aspire-dispense"
                    className="text-primary bg-white border-gray"
                    // onClick={toggleExportDataModal}
                  />
                  <TopHeading titleHeading="Aspire & Dispense" />
                </div>
              </div>
            </TopContent>

            <Card>
              <CardBody className="p-0 overflow-hidden">
                <div className="d-flex">
                  <Nav tabs className="d-flex flex-column border-0 side-bar">
                    <Text className="d-flex justify-content-center align-items-center  px-4 py-4 mb-0 font-weight-bold text-white">
                      <Icon
                        name={isAspire ? "upward-magnet" : "downward-magnet"}
                        size={29}
                        className="text-primary"
                      />
                      <Text Tag="span" className="ml-2">
                        {isAspire ? "Aspire Target" : "Dispense Target"}
                      </Text>
                    </Text>

                    {ASPIRE_DISPENSE_SIDEBAR_LABELS.map((label, index) => (
                      <NavItem>
                        <NavLink
                          className={classnames({
                            active: activeTab === `${index + 1}`,
                          })}
                          onClick={() => {
                            toggle(`${index + 1}`);
                          }}
                          disabled={
                            selectedTab && !(`${index + 1}` === selectedTab)
                          }
                        >
                          {label}
                        </NavLink>
                      </NavItem>
                    ))}
                  </Nav>

                  <TabContent activeTab={activeTab} className="flex-grow-1">
                    <Text className="d-flex justify-content-end align-items-center bg-white flex-fill mb-0 tab-content-top-heading">
                      <Text Tag="span" className="">
                        <Icon className="" name={"upward-magnet"} size={19} />
                        {"Aspire Target: Cartridge 1: Well no. 3 "}
                        {!isAspire && (
                          <>
                            <Icon
                              className=""
                              name={"downward-magnet"}
                              size={19}
                            />
                            Dispense Target: Cartridge 1: Well no. 3
                          </>
                        )}
                      </Text>
                    </Text>

                    <TabPane
                      tabId={"1"}
                      onChange={(e) => handleTabElementChange(e)}
                    >
                      <>
                        <WellComponent
                          wellsObjArray={
                            isAspire ? aspireCart1Wells : dispenseCart1Wells
                          }
                          wellClickHandler={handleWellClick}
                        />
                        {isAspire ? (
                          <AspireCommonField />
                        ) : (
                          <DispenseCommonField />
                        )}
                      </>
                    </TabPane>

                    <TabPane
                      tabId="2"
                      onChange={(e) => handleTabElementChange(e)}
                    >
                      <>
                        <WellComponent
                          wellsObjArray={
                            isAspire ? aspireCart2Wells : dispenseCart2Wells
                          }
                          wellClickHandler={handleWellClick}
                        />
                        {isAspire ? (
                          <AspireCommonField />
                        ) : (
                          <DispenseCommonField />
                        )}
                      </>
                    </TabPane>

                    <TabPane
                      tabId="3"
                      onChange={(e) => handleTabElementChange(e)}
                    >
                      {isAspire ? (
                        <div className="py-3">
                          <AspireCommonField />
                        </div>
                      ) : (
                        <DispenseCommonField />
                      )}
                    </TabPane>

                    <TabPane
                      tabId="4"
                      onChange={(e) => handleTabElementChange(e)}
                    >
                      <div className="mb-4 border-bottom-line">
                        <CommonDeckPosition />
                        {isAspire ? (
                          <AspireCommonField />
                        ) : (
                          <DispenseCommonField />
                        )}
                      </div>
                    </TabPane>
                  </TabContent>
                </div>
              </CardBody>
            </Card>
            <ButtonBar
              rightBtnLabel={isAspire ? "Next" : "Save"}
              handleRightBtn={() => {
                setIsAspire(!isAspire);
              }}
            />
          </div>
        </AspireDispenseBox>
      </PageBody>
    </>
  );
};

export default React.memo(AspireDispenseComponent);

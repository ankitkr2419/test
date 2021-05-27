import React, { useState } from "react";

import { Card, CardBody } from "core-components";
import { ButtonIcon, ButtonBar, TopHeading } from "shared-components";

import { TabContent, TabPane, Nav, NavItem, NavLink } from "reactstrap";
import classnames from "classnames";
import { PageBody, PiercingBox, TopContent } from "./Style";
import { WellComponent } from "./WellComponent";
import HeightModal from "components/modals/HeightModal";
import { useDispatch, useSelector } from "react-redux";
import { savePiercingInitiated } from "action-creators/processesActionCreators";
import { TEST_RECIPE_ID, TEST_TOKEN } from "appConstants";

const extractionWells = [
  { id: 1, type: 0, label: "1", footerText: "", isDisabled: false, isSelected: false },
  { id: 2, type: 0, label: "2", footerText: "", isDisabled: false, isSelected: false },
  { id: 3, type: 0, label: "3", footerText: "", isDisabled: false, isSelected: false },
  { id: 4, type: 0, label: "4", footerText: "", isDisabled: false, isSelected: false },
  { id: 5, type: 0, label: "5", footerText: "", isDisabled: false, isSelected: false },
  { id: 6, type: 0, label: "6", footerText: "", isDisabled: false, isSelected: false },
  { id: 7, type: 0, label: "7", footerText: "", isDisabled: false, isSelected: false },
  { id: 8, type: 0, label: "8", footerText: "", isDisabled: false, isSelected: false },
];

const pcrWells = [
  { id: 1, type: 1, label: "1", footerText: "", isDisabled: false, isSelected: false },
  { id: 2, type: 1, label: "2", footerText: "", isDisabled: false, isSelected: false },
  { id: 3, type: 1, label: "3", footerText: "", isDisabled: false, isSelected: false },
  { id: 4, type: 1, label: "4", footerText: "", isDisabled: false, isSelected: false },
];

const PiercingComponent = (props) => {
  const [activeTab, setActiveTab] = useState("0");
  const [showHeightModal, setShowHieghtModal] = useState(false);
  const [currentWellObj, setCurrentWellObj] = useState({});

  const recipeDetailsReducer = useSelector(
    (state) => state.updateRecipeDetailsReducer
  );
  const recipeID = recipeDetailsReducer.recipeDetails.id;
  const token = recipeDetailsReducer.token;

  const dispatch = useDispatch();

  const toggle = (tab) => {
    if (activeTab !== tab) setActiveTab(tab);
  };

  const handleSuccessBtn = (height, type) => {
    setShowHieghtModal(!showHeightModal);
    // here type : extractionWellsArray and 1 for pcrObjArray
    const wellsObjArray = type === 0 ? extractionWells : pcrWells;
    wellsObjArray.map((obj) => {
      // obj.isDisabled = !(obj.id === currentWellObj.id);
      if (!obj.isSelected) {
        obj.isSelected = obj.id === currentWellObj.id;
        obj.footerText =
          obj.id === currentWellObj.id ? `Height: ${height}mm` : "";
      }
      return obj;
    });
  };

  const wellClickHandler = (id, type) => {
    setShowHieghtModal(!showHeightModal);
    // here type : extractionWellsArray and 1 for pcrObjArray
    const wellsObjArray = type === 0 ? extractionWells : pcrWells;
    const currentWellObj = wellsObjArray.find((wellObj) => {
      if (wellObj.id === id) {
        return wellObj;
      }
    });
    setCurrentWellObj(currentWellObj);
  };

  const handleSaveBtn = () => {
    const type = parseInt(activeTab);
    const wellsObjArray = type === 0 ? extractionWells : pcrWells;
    const cartridgeWells = wellsObjArray
      .filter((obj) => obj.isSelected)
      .map((obj) => obj.id);

    const requestBody = {
      type: type,
      cartridgeWells: cartridgeWells,
      recipeID: recipeID,
      token: token,//TEST_TOKEN,
    };

    dispatch(savePiercingInitiated(requestBody));
  };

  return (
    <>
      {showHeightModal && (
        <HeightModal
          isOpen={showHeightModal}
          handleCrossBtn={() => setShowHieghtModal(!showHeightModal)}
          handleSuccessBtn={handleSuccessBtn}
          wellObj={currentWellObj}
        />
      )}
      <PageBody>
        <PiercingBox>
          <div className="process-content process-piercing px-2">
            <TopContent className="d-flex justify-content-between align-items-center mx-5">
              <div className="d-flex flex-column">
                <div className="d-flex align-items-center frame-icon">
                  <ButtonIcon
                    size={60}
                    name="piercing"
                    className="text-primary bg-white border-gray"
                  />
                  <TopHeading titleHeading="Piercing" />
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
                      className={classnames({ active: activeTab === "0" })}
                      onClick={() => toggle("0")}
                    >
                      Extraction
                    </NavLink>
                  </NavItem>
                  <NavItem className="text-center flex-fill px-2 pt-2">
                    <NavLink
                      className={classnames({ active: activeTab === "1" })}
                      onClick={() => toggle("1")}
                    >
                      PCR
                    </NavLink>
                  </NavItem>
                </Nav>
                <TabContent activeTab={activeTab} className="p-5">
                  <TabPane tabId="0">
                    <WellComponent
                      wellsObjArray={extractionWells}
                      wellClickHandler={wellClickHandler}
                    />
                  </TabPane>
                  <TabPane tabId="1">
                    <WellComponent
                      wellsObjArray={pcrWells}
                      wellClickHandler={wellClickHandler}
                    />
                  </TabPane>
                </TabContent>
              </CardBody>
            </Card>
            <ButtonBar rightBtnLabel="Save" handleRightBtn={handleSaveBtn} />
          </div>
        </PiercingBox>
      </PageBody>
    </>
  );
};

PiercingComponent.propTypes = {};

export default PiercingComponent;

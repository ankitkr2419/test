import React, { useEffect, useState } from "react";

import { Card, CardBody } from "core-components";
import { ButtonIcon, ButtonBar, TopHeading } from "shared-components";

import { TabContent, TabPane, Nav, NavItem, NavLink } from "reactstrap";
import classnames from "classnames";
import { PageBody, PiercingBox, TopContent } from "./Style";
import { WellComponent } from "./WellComponent";
import HeightModal from "components/modals/HeightModal";
import { useDispatch, useSelector } from "react-redux";
import { saveProcessInitiated } from "action-creators/processesActionCreators";
import { API_ENDPOINTS, HTTP_METHODS, ROUTES } from "appConstants";
import { Redirect, useHistory } from "react-router";
import {
  getWellsArrayForEdit,
  getWellsInitialArray,
  updateWellsArray,
} from "./functions";

const PiercingComponent = (props) => {
  const { editReducerData } = props;

  const dispatch = useDispatch();
  const history = useHistory();
  const [activeTab, setActiveTab] = useState("0");
  const [showHeightModal, setShowHieghtModal] = useState(false);
  const [currentWellObj, setCurrentWellObj] = useState({});
  const [extractionWells, setExtractionWell] = useState(
    getWellsInitialArray(8, 0)
  );
  const [pcrWells, setPcrWell] = useState(getWellsInitialArray(4, 1));

  const loginReducer = useSelector((state) => state.loginReducer);
  const loginReducerData = loginReducer.toJS();
  let activeDeckObj =
    loginReducerData && loginReducerData.decks.find((deck) => deck.isActive);
  const recipeDetailsReducer = useSelector(
    (state) => state.updateRecipeDetailsReducer
  );
  const processesReducer = useSelector((state) => state.processesReducer);

  // if data from editReducer is NOT NULL than updated wellsArrays
  useEffect(() => {
    if (editReducerData?.cartridge_wells) {
      setActiveTab(editReducerData.type === "cartridge_1" ? "0" : "1"); //change tab accordingly

      const upadatedExtractionWells = getWellsArrayForEdit(
        extractionWells,
        editReducerData
      );
      setExtractionWell((extractionWells) =>
        extractionWells.map((wellObj, index) => {
          return { ...wellObj, ...upadatedExtractionWells[index] };
        })
      );
      const upadatedPcrWells = getWellsArrayForEdit(pcrWells, editReducerData);
      setPcrWell((pcrWells) =>
        pcrWells.map((wellObj, index) => {
          return { ...wellObj, ...upadatedPcrWells[index] };
        })
      );
    }
  }, [editReducerData]);

  const recipeID = recipeDetailsReducer.recipeDetails.id;
  const token = activeDeckObj.token;

  //if no error occurs while saving process, redirect to next page
  const errorInAPICall = processesReducer.error;
  useEffect(() => {
    if (errorInAPICall === false) {
      history.push(ROUTES.processListing);
    }
  }, [errorInAPICall]);

  const toggle = (tab) => {
    if (activeTab !== tab) setActiveTab(tab);
  };

  //modal 'okay' btn
  const handleSuccessBtn = (height, type) => {
    setShowHieghtModal(!showHeightModal);
    const wellsObjArray = type === 0 ? extractionWells : pcrWells;
    const updatedWellObjArray = updateWellsArray(wellsObjArray, currentWellObj);
    type === 0
      ? setExtractionWell(updatedWellObjArray)
      : setPcrWell(updatedWellObjArray);
  };

  //when any well is clicked
  const wellClickHandler = (id, type) => {
    const wellsObjArray = type === 0 ? extractionWells : pcrWells;
    const currentWellObj = wellsObjArray.find((wellObj) => {
      if (wellObj.id === id) {
        return wellObj;
      }
    });
    setCurrentWellObj(currentWellObj);
    // if already selected then de-select
    if (currentWellObj.isSelected) {
      const updatedWellObjArray = updateWellsArray(
        wellsObjArray,
        currentWellObj
      );
      type === 0
        ? setExtractionWell(updatedWellObjArray)
        : setPcrWell(updatedWellObjArray);
    }
    // else open height modal and select
    else {
      setShowHieghtModal(!showHeightModal);
    }
  };

  const handleSaveBtn = () => {
    const type = parseInt(activeTab);
    const wellsObjArray = type === 0 ? extractionWells : pcrWells;
    const cartridgeWells = wellsObjArray
      .filter((obj) => obj.isSelected)
      .map((obj) => obj.id);

    const body = {
      type: `cartridge_${type + 1}`,
      cartridge_wells: cartridgeWells,
    };
    const requestBody = {
      body: body,
      id: editReducerData?.process_id ? editReducerData.process_id : recipeID,
      api: API_ENDPOINTS.piercing,
      token: token,
      method: editReducerData?.process_id
        ? HTTP_METHODS.PUT
        : HTTP_METHODS.POST,
    };
    dispatch(saveProcessInitiated(requestBody));
  };

  if (!activeDeckObj.isLoggedIn) return <Redirect to={`/${ROUTES.landing}`} />;

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

export default React.memo(PiercingComponent);

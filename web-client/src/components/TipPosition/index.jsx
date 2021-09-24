import React, { useEffect, useState } from "react";
import { useFormik } from "formik";
import { useDispatch, useSelector } from "react-redux";

import { ButtonIcon, ButtonBar } from "shared-components";

import TopHeading from "shared-components/TopHeading";
import { PageBody, TipPositionBox, TopContent } from "./Style";
import {
  getFormikInitialState,
  typeName,
  typeNameAPI,
  updateWellsArray,
} from "./functions";
import HeightModal from "components/modals/HeightModal";
import { TabsContent } from "./TabsContent";
import {
  API_ENDPOINTS,
  CARTRIDGE_1_WELLS,
  HTTP_METHODS,
  ROUTES,
  TIP_HEIGHT_MAX_ALLOWED_VALUE,
  TIP_HEIGHT_MIN_ALLOWED_VALUE,
  TIP_POSTION_ERROR_MSG,
} from "appConstants";
import { getPosition, tabApiNames } from "./functions";
import { Redirect, useHistory } from "react-router";
import { saveProcessInitiated } from "action-creators/processesActionCreators";
import { toast } from "react-toastify";

const TipPositionComponent = (props) => {
  const { editReducerData } = props;

  const [activeTab, setActiveTab] = useState("1");
  const [currentWellObj, setCurrentWellObj] = useState({});
  const [showHeightModal, setShowHieghtModal] = useState(false);

  const dispatch = useDispatch();
  const history = useHistory();

  const loginReducer = useSelector((state) => state.loginReducer);
  const loginReducerData = loginReducer.toJS();
  let activeDeckObj =
    loginReducerData && loginReducerData.decks.find((deck) => deck.isActive);
  const recipeDetailsReducer = useSelector(
    (state) => state.updateRecipeDetailsReducer
  );
  const processesReducer = useSelector((state) => state.processesReducer);

  const recipeID = recipeDetailsReducer.recipeDetails.id;
  const token = activeDeckObj.token;

  const formik = useFormik({
    initialValues: getFormikInitialState(editReducerData),
    enableReinitialize: true,
  });

  useEffect(() => {
    if (editReducerData?.process_id) {
      const value = editReducerData.type;
      const type = tabApiNames[value];

      setActiveTab(`${tabApiNames[value]}`);

      // disable other tabs
      for (const key in formik.values) {
        formik.setFieldValue(`${key}.isDisabled`, key !== typeName[type]);
      }
    }
  }, [editReducerData]);

  const errorInAPICall = processesReducer.error;
  useEffect(() => {
    if (errorInAPICall === false) {
      history.push(ROUTES.processListing);
    }
  }, [errorInAPICall]);

  //when any well is clicked
  const wellClickHandler = (id, type) => {
    // Please note,
    // type 0 is for cartridge-1, and,
    // type 1 is for cartridge-2
    const cartridge = formik.values[`cartridge${type + 1}`];
    const wellsObjArray = cartridge.wellsArray;

    // deck and the other cartridge tab must be disabled
    const otherTabToDisable =
      type === CARTRIDGE_1_WELLS ? "cartridge2" : "cartridge1";

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
      formik.setFieldValue(
        `cartridge${type + 1}.wellsArray`,
        updatedWellObjArray
      );
      //check for tipHeight and enable other tabs accordingly
      if (!cartridge.tipHeight) {
        formik.setFieldValue(`${otherTabToDisable}.isDisabled`, false);
        formik.setFieldValue(`deck.isDisabled`, false);
      }
    }
    // else select well
    else {
      //selecting one well and also sets height
      const wellsObjArray = formik.values[`cartridge${type + 1}`].wellsArray;
      wellsObjArray.forEach((wellObj, index) => {
        formik.setFieldValue(
          `cartridge${type + 1}.wellsArray.${index}.isSelected`,
          wellObj.id === currentWellObj.id
        );
      });
      //disable other tabs
      formik.setFieldValue(`${otherTabToDisable}.isDisabled`, true);
      formik.setFieldValue(`deck.isDisabled`, true);
    }
  };

  const handleSaveBtn = () => {
    const type = typeNameAPI[activeTab];
    const position =
      type === "deck"
        ? formik.values.deck.deckPosition
        : getPosition(formik.values[typeName[activeTab]].wellsArray);
    const tipHeight = formik.values[typeName[activeTab]].tipHeight;

    // check for valid data; if invalid data is sent, throw toast;
    // TipHeight: max:25 min:0;
    if (
      !position ||
      !tipHeight ||
      position === parseInt(0) ||
      tipHeight < TIP_HEIGHT_MIN_ALLOWED_VALUE ||
      tipHeight > TIP_HEIGHT_MAX_ALLOWED_VALUE
    ) {
      toast.error(TIP_POSTION_ERROR_MSG, { autoClose: false });
      return;
    }

    const body = {
      type: type,
      position: position,
      height: tipHeight,
    };

    const requestBody = {
      body: body,
      id: editReducerData?.process_id ? editReducerData.process_id : recipeID,
      api: API_ENDPOINTS.tipDocking,
      token: token,
      method: editReducerData?.process_id
        ? HTTP_METHODS.PUT
        : HTTP_METHODS.POST,
    };
    dispatch(saveProcessInitiated(requestBody));
  };

  const toggle = (tab) => activeTab !== tab && setActiveTab(tab);

  if (!activeDeckObj.isLoggedIn) return <Redirect to={`/${ROUTES.landing}`} />;

  return (
    <>
      <PageBody>
        <TipPositionBox>
          <div className="process-content process-tip-position px-2">
            {/* Header */}
            <TopContent className="d-flex justify-content-between align-items-center mx-5">
              <div className="d-flex flex-column">
                <div className="d-flex align-items-center frame-icon">
                  <ButtonIcon
                    size={60}
                    name="tip-position"
                    className="text-primary bg-white border-gray"
                  />
                  <TopHeading titleHeading="Tip Position" />
                </div>
              </div>
            </TopContent>

            {/* Body */}
            <TabsContent
              formik={formik}
              activeTab={activeTab}
              toggle={toggle}
              wellClickHandler={wellClickHandler}
            />
            <ButtonBar rightBtnLabel="Save" handleRightBtn={handleSaveBtn} />
          </div>
        </TipPositionBox>
      </PageBody>
    </>
  );
};

TipPositionComponent.propTypes = {};

export default React.memo(TipPositionComponent);

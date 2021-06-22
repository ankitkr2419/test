import React, { useEffect, useState } from "react";
import { useDispatch, useSelector } from "react-redux";
import { useFormik } from "formik";
import { API_ENDPOINTS, HTTP_METHODS, ROUTES } from "appConstants";
import { Redirect, useHistory } from "react-router";

import { Card, CardBody } from "core-components";
import { ButtonBar, ButtonIcon } from "shared-components";

import TopHeading from "shared-components/TopHeading";
import { AspireDispenseBox, PageBody, TopContent } from "./Style";
import {
  setFormikField,
  getFormikInitialState,
  getRequestBody,
  disabledTabInitTab,
  toggler,
} from "./helpers";
import AspireDispenseTabsContent from "./AspireDispenseTabsContent";
import { saveProcessInitiated } from "action-creators/processesActionCreators";

const AspireDispenseComponent = () => {
  const [activeTab, setActiveTab] = useState("1");
  const [isAspire, setIsAspire] = useState(true);
  const [disabledTab, setDisabledTab] = useState(disabledTabInitTab);

  const dispatch = useDispatch();
  const history = useHistory();

  const editReducer = useSelector((state) => state.editProcessReducer);
  const editReducerData = editReducer.toJS();
  const loginReducer = useSelector((state) => state.loginReducer);
  const loginReducerData = loginReducer.toJS();
  const recipeDetailsReducer = useSelector(
    (state) => state.updateRecipeDetailsReducer
  );
  const processesReducer = useSelector((state) => state.processesReducer);

  let activeDeckObj =
    loginReducerData && loginReducerData.decks.find((deck) => deck.isActive);

  const formik = useFormik({
    initialValues: getFormikInitialState(editReducerData),
    enableReinitialize: true,
  });

  useEffect(() => {
    setActiveTab(
      formik.values[isAspire ? "aspire" : "dispense"].selectedCategory
    );
  });

  const errorInAPICall = processesReducer.error;
  useEffect(() => {
    if (errorInAPICall === false) {
      history.push(ROUTES.processListing);
    }
  }, [errorInAPICall]);

  const recipeID = recipeDetailsReducer.recipeDetails.id;
  const token = activeDeckObj.token;

  const toggle = (tab) => {
    formik.setFieldValue(
      `${isAspire ? "aspire" : "dispense"}.selectedCategory`,
      `${tab}`
    );
  };

  const wellClickHandler = (id, type) => {
    // here type : aspire and 1 for dispense
    const aspire = formik.values.aspire;
    const dispense = formik.values.dispense;

    let wellsObjArray;
    if (type === 0) {
      wellsObjArray =
        activeTab === 1 ? aspire.cartridge1Wells : aspire.cartridge2Wells;
    } else {
      wellsObjArray =
        activeTab === 1 ? dispense.cartridge1Wells : dispense.cartridge2Wells;
    }

    wellsObjArray.forEach((wellObj, index) => {
      //get current selected value : true or false?
      let isSelected =
        formik.values[type === 0 ? "aspire" : "dispense"][
          `cartridge${activeTab}Wells`
        ][index].isSelected;

      setFormikField(
        formik,
        isAspire,
        activeTab,
        `cartridge${activeTab}Wells.${index}.isSelected`,
        wellObj.id === id ? !isSelected : false
      );
    });
  };

  useEffect(() => {
    let updatedDisabledTabState = toggler(formik, isAspire);
    setDisabledTab({ ...disabledTab, ...updatedDisabledTabState });
  }, [formik.values]);

  const handleModifyBtn = () => {
    setIsAspire(!isAspire);
    setActiveTab(formik.values.aspire.selectedCategory);
    formik.setFieldValue("dispense.selectedCategory", activeTab);
  };

  const handleNextBtn = () => {
    setIsAspire(!isAspire);
    setActiveTab(formik.values.dispense.selectedCategory);
    formik.setFieldValue("aspire.selectedCategory", activeTab);
  };

  const handleSaveBtn = () => {
    const aspire = formik.values.aspire;
    const dispense = formik.values.dispense;

    const requestBody = {
      body: getRequestBody(activeTab, aspire, dispense),
      id: editReducerData?.process_id ? editReducerData.process_id : recipeID,
      token: token,
      api: API_ENDPOINTS.aspireDispense,
      method: editReducerData?.process_id
        ? HTTP_METHODS.PUT
        : HTTP_METHODS.POST,
    };
    dispatch(saveProcessInitiated(requestBody));
  };

  if (!activeDeckObj.isLoggedIn) return <Redirect to={`/${ROUTES.landing}`} />;

  return (
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
                  onClick={() => setIsAspire(!isAspire)}
                />
                <TopHeading titleHeading="Aspire & Dispense" />
              </div>
            </div>
          </TopContent>

          <Card>
            <CardBody className="p-0 overflow-hidden">
              <AspireDispenseTabsContent
                formik={formik}
                isAspire={isAspire}
                toggle={toggle}
                activeTab={activeTab}
                wellClickHandler={wellClickHandler}
                disabledTab={disabledTab}
              />
            </CardBody>
          </Card>
          <ButtonBar
            leftBtnLabel={isAspire ? null : "Modify"}
            rightBtnLabel={isAspire ? "Next" : "Save"}
            handleLeftBtn={() => (isAspire ? null : handleModifyBtn())}
            handleRightBtn={() =>
              isAspire ? handleNextBtn() : handleSaveBtn()
            }
            btnBarClassname={"btn-bar-adjust-aspireDispense"}
          />
        </div>
      </AspireDispenseBox>
    </PageBody>
  );
};

export default React.memo(AspireDispenseComponent);

import React, { useState } from "react";

import { Card, CardBody } from "core-components";
import { ButtonBar, ButtonIcon } from "shared-components";

import TopHeading from "shared-components/TopHeading";
import { AspireDispenseBox, PageBody, TopContent } from "./Style";
import { getFormikInitialState } from "./functions";

import { useDispatch, useSelector } from "react-redux";
import { useFormik } from "formik";
import AspireDispenseTabsContent from "./AspireDispenseTabsContent";
import { setFormikField } from "./functions";
import { saveAspireDispenseInitiated } from "action-creators/processesActionCreators";

const AspireDispenseComponent = (props) => {
  const [activeTab, setActiveTab] = useState("1");
  const [isAspire, setIsAspire] = useState(true);

  const dispatch = useDispatch();

  const formik = useFormik({
    initialValues: getFormikInitialState(),
  });

  const recipeDetailsReducer = useSelector(
    (state) => state.updateRecipeDetailsReducer
  );
  const recipeID = recipeDetailsReducer.recipeDetails.id;
  const token = recipeDetailsReducer.token;

  const toggle = (tab) => {
    if (activeTab !== tab) setActiveTab(tab);
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
      setFormikField(
        formik,
        isAspire,
        activeTab,
        `cartridge${activeTab}Wells.${index}.isSelected`,
        wellObj.id === id
      );
    });
  };

  const handleSaveBtn = () => {
    const aspire = formik.values.aspire;
    const dispense = formik.values.dispense;

    //will be completed after clarification of APIs
    const body = {
      // category: //find out category
      // cartridge_type:
      // source_position:
      aspire_height: aspire.aspireHeight,
      aspire_mixing_volumne: aspire.mixingVolume,
      aspire_no_of_cycles: aspire.nCycles,
      aspire_volume: aspire.aspireVolume,
      aspire_air_volume: aspire.airVolume,
      dispense_height: dispense.dispenseHeight,
      dispense_mixing_volume: dispense.mixingVolume,
      dispense_no_of_cycles: dispense.nCycles,
      // destination_position: find cartridge
    };

    const requestBody = {
      body: body,
      recipeID: recipeID,
      token: token,
    };

    dispatch(saveAspireDispenseInitiated(requestBody));
  };

  const handleTabElementChange = () => {};
  // console.log(formik.values.dispense);

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
                />
              </CardBody>
            </Card>
            <ButtonBar
              leftBtnLabel={isAspire ? null : "Modify"}
              rightBtnLabel={isAspire ? "Next" : "Save"}
              handleLeftBtn={() =>
                isAspire ? null : setIsAspire(!isAspire)
              }
              handleRightBtn={() =>
                isAspire ? setIsAspire(!isAspire) : handleSaveBtn()
              }
            />
          </div>
        </AspireDispenseBox>
      </PageBody>
    </>
  );
};

export default React.memo(AspireDispenseComponent);

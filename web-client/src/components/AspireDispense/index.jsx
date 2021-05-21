import React, { useState } from "react";

import { Card, CardBody } from "core-components";
import { ButtonIcon } from "shared-components";

import TopHeading from "shared-components/TopHeading";
import { AspireDispenseBox, PageBody, TopContent } from "./Style";
import { getFormikInitialState } from "./functions";

import { useSelector } from "react-redux";
import { useFormik } from "formik";
import AspireDispenseTabsContent from "./AspireDispenseTabsContent";

const AspireDispenseComponent = (props) => {
  const [activeTab, setActiveTab] = useState("1");
  const [isAspire, setIsAspire] = useState(true);

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
      formik.setFieldValue(
        `${
          isAspire ? "aspire" : "dispense"
        }.cartridge${activeTab}Wells.${index}.isSelected`,
        wellObj.id === id
      );
      formik.setFieldValue(
        `${
          isAspire ? "aspire" : "dispense"
        }.cartridge${activeTab}Wells.${index}.isDisabled`,
        !(wellObj.id === id)
      );
    });
  };

  const handleTabElementChange = () => {};

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
                  handleTabElementChange={handleTabElementChange}
                  wellClickHandler={wellClickHandler}
                />
              </CardBody>
            </Card>
            {/* <ButtonBar
              rightBtnLabel={isAspire ? "Next" : "Save"}
              handleRightBtn={() => {
                setIsAspire(!isAspire);
              }}
            /> */}
          </div>
        </AspireDispenseBox>
      </PageBody>
    </>
  );
};

export default React.memo(AspireDispenseComponent);

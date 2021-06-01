import React from "react";

import { Text, Icon } from "shared-components";
import AspireCommonField from "./AspireCommonField";
import DispenseCommonField from "./DispenseCommonField";
import { WellComponent } from "./WellComponent";
import CommonDeckPosition from "./CommonDeckPosition";
import { ASPIRE_DISPENSE_SIDEBAR_LABELS, CATEGORY_LABEL } from "appConstants";
import { TabContent, TabPane, Nav, NavItem, NavLink } from "reactstrap";
import classnames from "classnames";
import { disabledTab } from "./functions";

const AspireDispenseTabsContent = (props) => {
  const { formik, isAspire, toggle, activeTab, wellClickHandler } = props;

  const disabledTabObj = isAspire ? disabledTab.aspire : disabledTab.dispense;
  const aspireCategoryLabel = formik.values.aspire.selectedCategory
    ? CATEGORY_LABEL[formik.values.aspire.selectedCategory]
    : CATEGORY_LABEL[activeTab];
  const dispenseCategoryLabel = CATEGORY_LABEL[activeTab];

  return (
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

        {ASPIRE_DISPENSE_SIDEBAR_LABELS.map((label, index) => {
          let key = Object.keys(disabledTabObj)[index];
          return (
            <NavItem key={index}>
              <NavLink
                className={classnames({
                  active: activeTab === `${index + 1}`,
                })}
                onClick={() => {
                  toggle(`${index + 1}`);
                }}
                disabled={disabledTabObj[key]}
              >
                {label}
              </NavLink>
            </NavItem>
          );
        })}
      </Nav>

      <TabContent activeTab={activeTab} className="flex-grow-1">
        <Text className="d-flex justify-content-end align-items-center bg-white flex-fill mb-0 tab-content-top-heading">
          <Text Tag="span" className="">
            <Icon className="" name={"upward-magnet"} size={19} />
            {`Aspire Target: ${aspireCategoryLabel}: Well no. 3 `}
            {!isAspire && (
              <>
                <Icon className="" name={"downward-magnet"} size={19} />
                {`Dispense Target: ${dispenseCategoryLabel}: Well no. 3`}
              </>
            )}
          </Text>
        </Text>

        <TabPane tabId={"1"}>
          <WellComponent
            wellsObjArray={
              isAspire
                ? formik.values.aspire.cartridge1Wells
                : formik.values.dispense.cartridge1Wells
            }
            wellClickHandler={wellClickHandler}
          />
          {isAspire ? (
            <AspireCommonField formik={formik} currentTab={activeTab} />
          ) : (
            <DispenseCommonField formik={formik} currentTab={activeTab} />
          )}
        </TabPane>

        <TabPane tabId="2">
          <WellComponent
            wellsObjArray={
              isAspire
                ? formik.values.aspire.cartridge2Wells
                : formik.values.dispense.cartridge2Wells
            }
            wellClickHandler={wellClickHandler}
          />
          {isAspire ? (
            <AspireCommonField formik={formik} currentTab={activeTab} />
          ) : (
            <DispenseCommonField formik={formik} currentTab={activeTab} />
          )}
        </TabPane>

        <TabPane tabId="3">
          {isAspire ? (
            <div className="py-3">
              <AspireCommonField formik={formik} currentTab={activeTab} />
            </div>
          ) : (
            <DispenseCommonField formik={formik} currentTab={activeTab} />
          )}
        </TabPane>

        <TabPane tabId="4">
          <div className="mb-4 border-bottom-line">
            <CommonDeckPosition
              formik={formik}
              isAspire={isAspire}
              currentTab={activeTab}
            />
            {isAspire ? (
              <AspireCommonField formik={formik} currentTab={activeTab} />
            ) : (
              <DispenseCommonField formik={formik} currentTab={activeTab} />
            )}
          </div>
        </TabPane>
      </TabContent>
    </div>
  );
};

export default React.memo(AspireDispenseTabsContent);

import React from "react";

import { Text, Icon } from "shared-components";
import AspireCommonField from "./AspireCommonField";
import DispenseCommonField from "./DispenseCommonField";
import { WellComponent } from "./WellComponent";
import CommonDeckPosition from "./CommonDeckPosition";
import { ASPIRE_DISPENSE_SIDEBAR_LABELS, CATEGORY_LABEL } from "appConstants";
import { TabContent, TabPane, Nav, NavItem, NavLink } from "reactstrap";
import classnames from "classnames";
import { getPosition, getWellLabel } from "./helpers";

const AspireDispenseTabsContent = (props) => {
  const { formik, isAspire, toggle, activeTab, wellClickHandler, disabledTab } =
    props;

  const { aspire, dispense } = formik.values;
  const disabledTabObj = isAspire ? disabledTab.aspire : disabledTab.dispense;
  const aspireCategoryLabel = CATEGORY_LABEL[aspire.selectedCategory];
  const dispenseCategoryLabel = CATEGORY_LABEL[dispense.selectedCategory];

  // get selected well ID to display
  const aspireWellId = getPosition(
    aspire[`cartridge${aspire.selectedCategory}Wells`]
  );
  const dispenseWellId = getPosition(
    dispense[`cartridge${dispense.selectedCategory}Wells`]
  );

  // get label
  const aspireWellLabel = getWellLabel(aspireCategoryLabel, aspireWellId);
  const dispenseWellLabel = getWellLabel(dispenseCategoryLabel, dispenseWellId);

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
          <Text Tag="span">
            <Icon name="upward-magnet" size={19} className="mx-3" />
            {`Aspire Target: ${aspireCategoryLabel}${aspireWellLabel}`}
          </Text>

          <Text Tag="span" className="border-left ml-3">
            <Icon name="downward-magnet" size={19} className="mx-3" />
            {`Dispense Target: ${dispenseCategoryLabel}${dispenseWellLabel}`}
          </Text>
        </Text>

        <TabPane tabId={"1"}>
          <WellComponent
            wellsObjArray={
              isAspire ? aspire.cartridge1Wells : dispense.cartridge1Wells
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
              isAspire ? aspire.cartridge2Wells : dispense.cartridge2Wells
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

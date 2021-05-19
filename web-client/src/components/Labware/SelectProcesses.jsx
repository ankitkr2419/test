import React from "react";

import { Text, Icon } from "shared-components";

import { TabContent, TabPane, Nav } from "reactstrap";
import TubesAndCartridgeSelection from "./TubesAndCartridgeSelection";
import SideBarNavItems from "./SideBarNavItems";
import TipPiercingCheckbox from "./TipPiercingCheckbox";
import TipsDropDown from "./TipsDropDown";

const SelectProcesses = (props) => {
  const {
    formik,
    toggle,
    activeTab,
    tubesOptions,
    tipsOptions,
    tipsAndTubesOptions,
    cartridgeOptions,
  } = props;
  return (
    <div className="d-flex">
      <Nav tabs className="d-flex flex-column border-0 side-bar">
        <Text className="d-flex justify-content-center align-items-center px-3 pt-3 pb-3 mb-0 font-weight-bold text-white">
          <Icon name="setting" size={18} />
          <Text Tag="span" className="ml-2">
            Settings{" "}
          </Text>
        </Text>
        <SideBarNavItems
          formik={formik}
          activeTab={activeTab}
          toggle={toggle}
        />
      </Nav>

      <TabContent activeTab={activeTab} className="flex-grow-1">
        <TabPane tabId="1">
          <TipsDropDown formik={formik} tipsOptions={tipsAndTubesOptions} />
        </TabPane>
        <TabPane tabId="2">
          <TipPiercingCheckbox formik={formik} />
        </TabPane>
        <TabPane tabId="3">
          <TubesAndCartridgeSelection
            formik={formik}
            isDeck={true}
            position={1}
            allOptions={tubesOptions}
          />
        </TabPane>
        <TabPane tabId="4">
          <TubesAndCartridgeSelection
            formik={formik}
            isDeck={true}
            position={2}
            allOptions={tubesOptions}
          />
        </TabPane>
        <TabPane tabId="5">
          <TubesAndCartridgeSelection
            formik={formik}
            isDeck={false}
            position={1}
            allOptions={cartridgeOptions}
          />
        </TabPane>
        <TabPane tabId="6">
          <TubesAndCartridgeSelection
            formik={formik}
            isDeck={true}
            position={3}
            allOptions={tubesOptions}
          />
        </TabPane>
        <TabPane tabId="7">
          <TubesAndCartridgeSelection
            formik={formik}
            isDeck={false}
            position={2}
            allOptions={cartridgeOptions}
          />
        </TabPane>
        <TabPane tabId="8">
          <TubesAndCartridgeSelection
            formik={formik}
            isDeck={true}
            position={4}
            allOptions={tubesOptions}
          />
        </TabPane>
      </TabContent>
    </div>
  );
};

export default React.memo(SelectProcesses);

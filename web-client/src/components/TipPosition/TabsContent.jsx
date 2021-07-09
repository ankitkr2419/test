import React from "react";

import { Card, CardBody } from "core-components";

import { TabContent, TabPane, Nav, NavItem, NavLink } from "reactstrap";
import classnames from "classnames";
import TipPositionInfo from "./TipPositionInfo";
import DeckPositionInfo from "./DeckPositionInfo";
import {
  TAB_TYPE_CARTRIDGE_1,
  TAB_TYPE_CARTRIDGE_2,
  TAB_TYPE_DECK,
} from "appConstants";

export const TabsContent = (props) => {
  const { formik, activeTab, toggle, wellClickHandler } = props;

  return (
    <Card className="card-box">
      <CardBody className="p-0 overflow-hidden">
        <Nav
          tabs
          className="bg-white px-3 pb-0 d-flex justify-content-center align-items-center border-0"
        >
          <NavItem className="text-center flex-fill px-2 pt-2">
            <NavLink
              className={classnames({
                active: activeTab === TAB_TYPE_CARTRIDGE_1,
              })}
              onClick={() => {
                toggle(TAB_TYPE_CARTRIDGE_1);
              }}
              disabled={formik.values.cartridge1.isDisabled}
            >
              Cartridge 1
            </NavLink>
          </NavItem>
          <NavItem className="text-center flex-fill px-2 pt-2">
            <NavLink
              className={classnames({ active: activeTab === TAB_TYPE_DECK })}
              onClick={() => {
                toggle(TAB_TYPE_DECK);
              }}
              disabled={formik.values.deck.isDisabled}
            >
              Deck
            </NavLink>
          </NavItem>
          <NavItem className="text-center flex-fill px-2 pt-2">
            <NavLink
              className={classnames({
                active: activeTab === TAB_TYPE_CARTRIDGE_2,
              })}
              onClick={() => {
                toggle(TAB_TYPE_CARTRIDGE_2);
              }}
              disabled={formik.values.cartridge2.isDisabled}
            >
              Cartridge 2
            </NavLink>
          </NavItem>
        </Nav>
        <TabContent activeTab={activeTab} className="p-5">
          <TabPane tabId="1">
            <TipPositionInfo
              formik={formik}
              activeTab={activeTab}
              wellClickHandler={wellClickHandler}
            />
          </TabPane>
          <TabPane tabId="2">
            <DeckPositionInfo formik={formik} activeTab={activeTab} />
          </TabPane>
          <TabPane tabId="3">
            <TipPositionInfo
              formik={formik}
              activeTab={activeTab}
              wellClickHandler={wellClickHandler}
            />
          </TabPane>
        </TabContent>
      </CardBody>
    </Card>
  );
};

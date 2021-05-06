import React from "react";
import {
  LABWARE_CARTRIDGE_1_OPTIONS,
  LABWARE_CARTRIDGE_2_OPTIONS,
  LABWARE_DECK_POS_1_OPTIONS,
  LABWARE_DECK_POS_2_OPTIONS,
  LABWARE_DECK_POS_3_OPTIONS,
  LABWARE_DECK_POS_4_OPTIONS,
  LABWARE_ITEMS_NAME,
  LABWARE_TIPS_OPTIONS,
} from "appConstants";

import labwareTips from "assets/images/labware-plate-tips.png";
import labwarePiercing from "assets/images/labware-plate-piercing.png";
import labwareDeckPosition1 from "assets/images/labware-plate-deck-position-1.png";
import labwareDeckPosition2 from "assets/images/labware-plate-deck-position-2.png";
import labwareDeckPosition3 from "assets/images/labware-plate-deck-position-3.png";
import labwareDeckPosition4 from "assets/images/labware-plate-deck-position-4.png";
import labwareCartridePosition1 from "assets/images/labware-plate-cartridge-1.png";
import labwareCartridePosition2 from "assets/images/labware-plate-cartridge-2.png";

import { Icon, ImageIcon } from "shared-components";
import { NavItem, NavLink } from "reactstrap";
import classnames from "classnames";
import { FormGroup, Label, FormError, Select, CheckBox } from "core-components";

import TubeSelection from "./TubeSelection";
import CartridgeSelection from "./CartrideSelection";
import { ProcessSetting } from "./Styles";

export const getTick = (index, formik) => {
  const currentState = formik.values;
  const name = currentState[Object.keys(currentState)[index]];

  let tick = false;
  Object.values(name).forEach((value) => {
    if (value) tick = true;
  });
  return tick;
};

export const getSideBarNavItems = (formik, activeTab, toggle) => {
  const navItems = [];
  LABWARE_ITEMS_NAME.forEach((name, index) => {
    navItems.push(
      <NavItem>
        <NavLink
          className={classnames({ active: activeTab === `${index + 1}` })}
          onClick={() => {
            toggle(`${index + 1}`);
          }}
        >
          {name}
          {getTick(index, formik) ? (
            <Icon name="tick" size={12} className="ml-auto" />
          ) : null}
        </NavLink>
      </NavItem>
    );
  });
  return navItems;
};

export const getTipPiercingCheckbox = (formik, nCheckboxes = 2) => {
  const tipsPiercingCheckbox = [];
  for (let index = 0; index < nCheckboxes; index++) {
    tipsPiercingCheckbox.push(
      <CheckBox
        id={`position${index + 1}`}
        name={`position${index + 1}`}
        label={`Position ${index + 1}`}
        className={index > 0 ? "ml-4" : ""}
        onChange={(e) => {
          formik.setFieldValue(
            `tipPiercing.position${index + 1}`,
            e.target.checked
          );
        }}
      />
    );
  }
  return tipsPiercingCheckbox;
};

export const getTipsDropdown = (formik, nDropdown = 3) => {
  const tipsOptions = [];
  for (let index = 0; index < nDropdown; index++) {
    tipsOptions.push(
      <FormGroup className="d-flex align-items-center mb-4">
        <Label for={`tip-position-${index + 1}`} className="px-0 label-name">
          Tip Position {index + 1}
        </Label>
        <div className="d-flex flex-column input-field position-relative">
          <Select
            placeholder="Select Option"
            className=""
            size="sm"
            options={LABWARE_TIPS_OPTIONS}
            onChange={(e) =>
              formik.setFieldValue(`tips.tipPosition${index + 1}`, e.value)
            }
          />
          <FormError>Incorrect Tip Position {index + 1}</FormError>
        </div>
      </FormGroup>
    );
  }
  return tipsOptions;
};

export const getTipsAtPosition = (position, formik) => {
  return (
    <>
      <div className="">
        <div className="mb-3">
          <FormGroup row>
            <Label
              for="select-tip-position"
              md={12}
              className="mb-3 font-weight-bold"
            >
              Select Tip Position
            </Label>
          </FormGroup>
        </div>
        <div className="">
          {getTipsDropdown(formik)}
          {/* <CommonField /> */}
        </div>
      </div>
      <ProcessSetting>
        <div className="tips-info">
          <ul class="list-unstyled tip-position active">
            <li class="highlighted tip-position-1"></li>
            <li class="highlighted tip-position-2 active"></li>
            <li class="highlighted tip-position-3 active"></li>
          </ul>
          <ImageIcon src={labwareTips} alt="Tip Pickup Process" className="" />
        </div>
      </ProcessSetting>
    </>
  );
};

export const getTipPiercingAtPosition = (position, formik) => {
  return (
    <>
      <div className="mb-3">
        <FormGroup row>
          <Label
            for="select-tip-piercing"
            md={12}
            className="mb-3 font-weight-bold"
          >
            Select Tip Piercing
          </Label>
        </FormGroup>
      </div>
      <div className="d-flex align-items-center">
        {getTipPiercingCheckbox(formik)}
        <ProcessSetting>
          <div className="piercing-info">
            <ul class="list-unstyled piercing-position active">
              <li class="highlighted piercing-position-1"></li>
              <li class="highlighted piercing-position-2 active"></li>
            </ul>
            <ImageIcon
              src={labwarePiercing}
              alt="Piercing Process"
              className=""
            />
          </div>
        </ProcessSetting>
      </div>
    </>
  );
};

export const getDeckAtPosition = (position, formik) => {
  const deckPositionOptions = [
    LABWARE_DECK_POS_1_OPTIONS,
    LABWARE_DECK_POS_2_OPTIONS,
    LABWARE_DECK_POS_3_OPTIONS,
    LABWARE_DECK_POS_4_OPTIONS,
  ];
  const deckImages = [
    labwareDeckPosition1,
    labwareDeckPosition2,
    labwareDeckPosition3,
    labwareDeckPosition4,
  ];
  return (
    <>
      <TubeSelection
        handleOptionChange={(e) => {
          formik.setFieldValue(`deckPosition${position}.tubeType`, e.value);
        }}
        options={deckPositionOptions[position - 1]}
      />
      <ProcessSetting>
        <div className="deck-position-info">
          <ul class="list-unstyled deck-position active">
            <li class={`highlighted deck-position-${position} active`} />
          </ul>
          <ImageIcon
            src={deckImages[position - 1]}
            alt={`Deck Position ${position} Process`}
            className=""
          />
        </div>
      </ProcessSetting>
    </>
  );
};

export const getCartidgeAtPosition = (position, formik) => {
  const cartidgePositionOptions = [
    LABWARE_CARTRIDGE_1_OPTIONS,
    LABWARE_CARTRIDGE_2_OPTIONS,
  ];
  const cartridgeImages = [labwareCartridePosition1, labwareCartridePosition2];
  return (
    <>
      <CartridgeSelection
        handleOptionChange={(e) => {
          formik.setFieldValue(`cartridge${position}.cartridgeType`, e.value);
        }}
        options={cartidgePositionOptions[position - 1]}
      />
      <ProcessSetting>
        <div className="cartridge-position-info">
          <ul class="list-unstyled cartridge-position active">
            <li class={`highlighted cartridge-position-${position} active`} />
          </ul>
          <ImageIcon
            src={cartridgeImages[position - 1]}
            alt={`Cartridge Position ${position} Process`}
            className=""
          />
        </div>
      </ProcessSetting>
    </>
  );
};

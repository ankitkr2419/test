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
  LABWARE_NAME,
} from "appConstants";

import labwareTips from "assets/images/labware-plate-tips.png";
import labwarePiercing from "assets/images/labware-plate-piercing.png";
import labwareDeckPosition1 from "assets/images/labware-plate-deck-position-1.png";
import labwareDeckPosition2 from "assets/images/labware-plate-deck-position-2.png";
import labwareDeckPosition3 from "assets/images/labware-plate-deck-position-3.png";
import labwareDeckPosition4 from "assets/images/labware-plate-deck-position-4.png";
import labwareCartridePosition1 from "assets/images/labware-plate-cartridge-1.png";
import labwareCartridePosition2 from "assets/images/labware-plate-cartridge-2.png";

import { Icon, ImageIcon, Text } from "shared-components";
import { NavItem, NavLink } from "reactstrap";
import classnames from "classnames";
import { FormGroup, Label, FormError, Select, CheckBox } from "core-components";

import TubeSelection from "./TubeSelection";
import CartridgeSelection from "./CartrideSelection";
import { ProcessSetting } from "./Styles";

export const updateAllTicks = (formik) => {
  const currentState = formik.values;

  Object.keys(currentState).forEach((key) => {
    const processDetails = currentState[key].processDetails;  
    // const tick = currentState[key].isTicked;
    let tick = false;

    Object.values(processDetails).forEach((value) => {
      if (value) {
        tick = true;
        return;
      }
    });
    formik.setFieldValue(`${key}.isTicked`, tick);
  });

  // const processDetails = currentState[key].processDetails;
  // const tick = currentState[key].isTicked;

  // Object.values(processDetails).forEach((value) => {
  //   if (value && !tick) {
  //     formik.setFieldValue(`${key}.isTicked`, true);
  //     return;
  //   }
  // });
};

export const getSideBarNavItems = (formik, activeTab, toggle) => {
  const navItems = [];
  LABWARE_ITEMS_NAME.forEach((name, index) => {
    const currentState = formik.values;
    const key = Object.keys(currentState)[index];
    // getTicked(key, formik);
    navItems.push(
      <NavItem>
        <NavLink
          className={classnames({ active: activeTab === `${index + 1}` })}
          onClick={() => {
            toggle(`${index + 1}`);
            updateAllTicks(formik);
          }}
        >
          {name}
          {currentState[key].isTicked ? (
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
    let checked =
      formik.values.tipPiercing.processDetails[`position${index + 1}`];
    tipsPiercingCheckbox.push(
      <CheckBox
        id={`position${index + 1}`}
        name={`position${index + 1}`}
        label={`Position ${index + 1}`}
        className={index > 0 ? "ml-4" : ""}
        checked={checked}
        onChange={(e) => {
          formik.setFieldValue(
            `tipPiercing.processDetails.position${index + 1}`,
            e.target.checked
          );
        }}
      />
    );
  }
  return tipsPiercingCheckbox;
};

export const getTipsDropdown = (formik, options) => {
  const tips = formik.values.tips;
  const nDropdown = 3;
  const tipsOptions = [];
  for (let i = 0; i < nDropdown; i++) {
    let tipPosition = tips.processDetails[`tipPosition${i + 1}`];
    let index = options.map((item) => item.value).indexOf(
      tipPosition
    );
    tipsOptions.push(
      <FormGroup className="d-flex align-items-center mb-4">
        <Label for={`tip-position-${i + 1}`} className="px-0 label-name">
          Tip Position {i + 1}
        </Label>
        <div className="d-flex flex-column input-field position-relative">
          <Select
            placeholder="Select Option"
            className=""
            size="sm"
            value={options[index]}
            options={options}
            onChange={(e) =>
              formik.setFieldValue(
                `tips.processDetails.tipPosition${i + 1}`,
                e.value
              )
            }
          />
          <FormError>Incorrect Tip Position {index + 1}</FormError>
        </div>
      </FormGroup>
    );
  }
  return tipsOptions;
};

export const getTipsAtPosition = (position, formik, options) => {
  const tipPosition1Value = formik.values.tips.processDetails.tipPosition1;
  const tipPosition2Value = formik.values.tips.processDetails.tipPosition2;
  const tipPosition3Value = formik.values.tips.processDetails.tipPosition3;

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
          {getTipsDropdown(formik, options)}
          {/* <CommonField /> */}
        </div>
      </div>
      <ProcessSetting>
        <div className="tips-info">
          <ul class="list-unstyled tip-position active">
            {tipPosition1Value && <li class="highlighted tip-position-1"></li>}
            {tipPosition2Value && (
              <li class="highlighted tip-position-2 active"></li>
            )}
            {tipPosition3Value && (
              <li class="highlighted tip-position-3 active"></li>
            )}
          </ul>
          <ImageIcon src={labwareTips} alt="Tip Pickup Process" className="" />
        </div>
      </ProcessSetting>
    </>
  );
};

export const getTipPiercingAtPosition = (position, formik) => {
  const position1 = formik.values.tipPiercing.processDetails.position1;
  const position2 = formik.values.tipPiercing.processDetails.position2;

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
              {position1 && <li class="highlighted piercing-position-1"></li>}
              {position2 && (
                <li class="highlighted piercing-position-2 active"></li>
              )}
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

export const getDeckAtPosition = (position, formik, options) => {
  const deckImages = [
    labwareDeckPosition1,
    labwareDeckPosition2,
    labwareDeckPosition3,
    labwareDeckPosition4,
  ];
  const deckPositionValue = formik.values[`deckPosition${position}`].processDetails;
  const index = options
    .map((item) => item.value)
    .indexOf(deckPositionValue);
  
  return (
    <>
      <TubeSelection
        handleOptionChange={(e) => {
          formik.setFieldValue(
            `deckPosition${position}.processDetails.tubeType`,
            e.value
          );
        }}
        value={options[index]}
        options={options}
      />
      <ProcessSetting>
        <div className="deck-position-info">
          <ul class="list-unstyled deck-position active">
            {deckPositionValue && (
              <li class={`highlighted deck-position-${position} active`} />
            )}
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

export const getCartidgeAtPosition = (position, formik, options) => {
  const cartidgePositionOptions = [
    LABWARE_CARTRIDGE_1_OPTIONS,
    LABWARE_CARTRIDGE_2_OPTIONS,
  ];
  const cartridgeImages = [labwareCartridePosition1, labwareCartridePosition2];

  const cartridgeValue =
    formik.values[`cartridge${position}`].processDetails.cartridgeType;
  const index = cartidgePositionOptions[position - 1]
    .map((item) => item.value)
    .indexOf(cartridgeValue);

  const cartridgeType =
    formik.values[`cartridge${position}`].processDetails.cartridgeType;
  return (
    <>
      <CartridgeSelection
        handleOptionChange={(e) => {
          formik.setFieldValue(
            `cartridge${position}.processDetails.cartridgeType`,
            e.value
          );
        }}
        options={cartidgePositionOptions[position - 1]}
        value={cartidgePositionOptions[position - 1][index]}
      />
      <ProcessSetting>
        <div className="cartridge-position-info">
          <ul class="list-unstyled cartridge-position active">
            {cartridgeType && (
              <li class={`highlighted cartridge-position-${position} active`} />
            )}
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

export const getSubHead = (key, formik) => {
  const recipeData = formik.values;
  const nestedKeys = Object.keys(recipeData[key].processDetails);
  const LEN = nestedKeys.length;
  const previewInfoSubHead = [];

  nestedKeys.forEach((nestedKey) => {
    recipeData[key].processDetails[nestedKey] &&
      previewInfoSubHead.push(
        <Text>
          {LEN > 1 && (
            <Text Tag="span" className="font-weight-bold">
              {LABWARE_NAME[nestedKey]} :{" "}
            </Text>
          )}
          <Text Tag="span" className={LEN === 1 ? "font-weight-bold" : ""}>
            {recipeData[key].processDetails[nestedKey]}{" "}
          </Text>
        </Text>
      );
  });
  return previewInfoSubHead;
};

export const getPreviewInfo = (formik) => {
  const previewInfoHead = [];
  const recipeData = formik.values;
  Object.keys(recipeData).forEach((key) => {
    recipeData[key].isTicked &&
      previewInfoHead.push(
        <li className="d-flex justify-content-between">
          <Text className="d-flex w-25 font-weight-bold">
            {LABWARE_NAME[key]} :{" "}
          </Text>
          <div className="w-75">
            <div className="ml-2 setting-value">
              <Text>{getSubHead(key, formik)}</Text>
            </div>
          </div>
        </li>
      );
  });
  return previewInfoHead;
};

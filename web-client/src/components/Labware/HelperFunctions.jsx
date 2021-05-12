import React from "react";
import { LABWARE_ITEMS_NAME, LABWARE_NAME } from "appConstants";

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
import { ProcessSetting } from "./Styles";
import HeaderAndLabel from "./HeaderAndLabel";
import { getOptionsForTubesAndCartridges } from "./functions";

export const updateAllTicks = (formik) => {
  const currentState = formik.values;

  Object.keys(currentState).forEach((key, index) => {
    const processDetails = currentState[key].processDetails;
    let tick = false;

    for (const key in processDetails) {
      if (processDetails[key].id) {
        tick = true;
        break;
      }
    }
    formik.setFieldValue(`${key}.isTicked`, tick);
  });
};

export const getSideBarNavItems = (formik, activeTab, toggle) => {
  const navItems = [];
  LABWARE_ITEMS_NAME.forEach((name, index) => {
    const currentState = formik.values;
    const key = Object.keys(currentState)[index];
    navItems.push(
      <NavItem key={key}>
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
    let isChecked =
      formik.values.tipPiercing.processDetails[`position${index + 1}`].id;
    tipsPiercingCheckbox.push(
      <CheckBox
        id={`position${index + 1}`}
        name={`position${index + 1}`}
        label={`Position ${index + 1}`}
        className={index > 0 ? "ml-4" : ""}
        checked={isChecked}
        onChange={(e) => {
          formik.setFieldValue(
            `tipPiercing.processDetails.position${index + 1}.id`,
            e.target.checked ? 3 : 0
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
    let tipPosition = tips.processDetails[`tipPosition${i + 1}`].id;
    let index = options.map((item) => item.value).indexOf(tipPosition);
    tipsOptions.push(
      <FormGroup key={i} className="d-flex align-items-center mb-4">
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
            onChange={(e) => {
              formik.setFieldValue(
                `tips.processDetails.tipPosition${i + 1}.id`,
                e.value
              );
              formik.setFieldValue(
                `tips.processDetails.tipPosition${i + 1}.label`,
                e.label
              );
            }}
          />
          <FormError>Incorrect Tip Position {index + 1}</FormError>
        </div>
      </FormGroup>
    );
  }
  return tipsOptions;
};

export const getTipsAtPosition = (position, formik, options) => {
  const tips = formik.values.tips;
  const tipPosition1Value = tips.processDetails.tipPosition1.id;
  const tipPosition2Value = tips.processDetails.tipPosition2.id;
  const tipPosition3Value = tips.processDetails.tipPosition3.id;

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
          <ul className="list-unstyled tip-position active">
            {tipPosition1Value && (
              <li className="highlighted tip-position-1"></li>
            )}
            {tipPosition2Value && (
              <li className="highlighted tip-position-2 active"></li>
            )}
            {tipPosition3Value && (
              <li className="highlighted tip-position-3 active"></li>
            )}
          </ul>
          <ImageIcon src={labwareTips} alt="Tip Pickup Process" className="" />
        </div>
      </ProcessSetting>
    </>
  );
};

export const getTipPiercingAtPosition = (position, formik) => {
  const position1 = formik.values.tipPiercing.processDetails.position1.id;
  const position2 = formik.values.tipPiercing.processDetails.position2.id;

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
            <ul className="list-unstyled piercing-position active">
              {position1 && (
                <li className="highlighted piercing-position-1"></li>
              )}
              {position2 && (
                <li className="highlighted piercing-position-2 active"></li>
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

const handleOptionChange = (formik, position, e, key, type) => {
  formik.setFieldValue(`${key}${position}.processDetails.${type}.id`, e.value);
  formik.setFieldValue(
    `${key}${position}.processDetails.${type}.label`,
    e.label
  );
};

const deckImages = [
  labwareDeckPosition1,
  labwareDeckPosition2,
  labwareDeckPosition3,
  labwareDeckPosition4,
];
const cartridgeImages = [labwareCartridePosition1, labwareCartridePosition2];

export const getFieldAtPosition = (position, formik, allOptions, key) => {
  const images = key === "deckPosition" ? deckImages : cartridgeImages;

  if (allOptions) {
    //id in response from backend starts from 4 for tubes and
    // starts with 1 for catridge.
    const n = key === "deckPosition" ? 3 : 0;
    const options = getOptionsForTubesAndCartridges(allOptions, position + n);
    const type = key === "deckPosition" ? "tubeType" : "cartridgeType";

    const selectedOptionID =
      formik.values[`${key}${position}`].processDetails[type].id;
    const index = options.map((item) => item.value).indexOf(selectedOptionID);

    return (
      <>
        <HeaderAndLabel
          key={key}
          headerText={
            key === "deckPosition" ? "Select Deck" : "Select Cartridge"
          }
          label={key === "deckPosition" ? "Tube Type" : "Cartridge Type"}
          handleOptionChange={(e) =>
            handleOptionChange(formik, position, e, key, type)
          }
          value={options[index]}
          options={options}
          position={position}
          images={images}
          typeValue={selectedOptionID}
        />
      </>
    );
  }
};

export const getSubHead = (key, formik) => {
  const recipeData = formik.values;
  const nestedKeys = Object.keys(recipeData[key].processDetails);
  const LEN = nestedKeys.length;
  const previewInfoSubHead = [];

  nestedKeys.forEach((nestedKey) => {
    recipeData[key].processDetails[nestedKey].id &&
      previewInfoSubHead.push(
        <Text>
          {LEN > 1 && (
            <Text Tag="span" className="font-weight-bold">
              {LABWARE_NAME[nestedKey]} :{" "}
            </Text>
          )}
          <Text Tag="span" className={LEN === 1 ? "font-weight-bold" : ""}>
            {recipeData[key].processDetails[nestedKey].label}{" "}
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

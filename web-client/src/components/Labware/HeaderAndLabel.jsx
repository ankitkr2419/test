import React from "react";
import { FormGroup, Label, FormError, Select } from "core-components";
import { CommmonTubeFields } from "./CommmonTubeFields";
import { ImageIcon } from "shared-components";
import { ProcessSetting } from "./Styles";

const HeaderAndLabel = (props) => {
  const {
    key,
    handleOptionChange,
    options,
    value,
    headerText,
    label,
    images,
    position,
    typeValue,
  } = props;

  const type = (key==="deckPosition") ? "deck" : "cartridge";
  return (
    <>
      <CommmonTubeFields>
        <FormGroup>
          <Label className="mb-3 font-weight-bold px-0">{headerText}</Label>
        </FormGroup>
        <FormGroup className="d-flex align-items-center mb-2">
          <Label className="px-0 label-name">{label}</Label>
          <div className="d-flex flex-column input-field position-relative">
            <Select
              placeholder="Select Option"
              className=""
              size="sm"
              value={value}
              options={options}
              onChange={handleOptionChange}
            />
            <FormError>Incorrect {label}</FormError>
          </div>
        </FormGroup>
      </CommmonTubeFields>

      <ProcessSetting>
        <div className={`${type}-position-info`}>
          <ul className={`list-unstyled ${type}-position active`}>
            {typeValue && (
              <li
                className={`highlighted ${type}-position-${position} active`}
              />
            )}
          </ul>
          {images && (
            <ImageIcon
              src={images[position - 1]}
              alt={`${type} Position ${position} Process`}
              className=""
            />
          )}
        </div>
      </ProcessSetting>
    </>
  );
};

HeaderAndLabel.propTypes = {};

export default React.memo(HeaderAndLabel);

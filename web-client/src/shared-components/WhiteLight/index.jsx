import React from "react";
import PropTypes from "prop-types";
import { Switch } from "core-components";
import { ImageIcon } from "shared-components";
import bulb from "assets/images/bulb.svg";
import { StyledWhiteLightComponent } from "./Style";
import { DECKNAME } from "appConstants";
import { whiteLightDeckInitiated } from "action-creators/whiteLightActionCreators";
import { useDispatch, useSelector } from "react-redux";

const WhiteLight = (props) => {
  const { className } = props;

  const dispatch = useDispatch();
  const lightReducer = useSelector((state) => state.whiteLightReducer);
  let lightReducerForDeckA = lightReducer.decks.find(
    (deckObj) => deckObj.name === DECKNAME.DeckA
  );
  let lightReducerForDeckB = lightReducer.decks.find(
    (deckObj) => deckObj.name === DECKNAME.DeckB
  );

  const handleWhiteLightClick = () => {
    let deckA = DECKNAME.DeckAShort;
    let deckB = DECKNAME.DeckBShort;
    if (lightReducerForDeckA.isLightOn || lightReducerForDeckB.isLightOn) {
      dispatch(whiteLightDeckInitiated({ deck: deckA, lightStatus: 0 }));
      dispatch(whiteLightDeckInitiated({ deck: deckB, lightStatus: 0 }));
    } else {
      dispatch(whiteLightDeckInitiated({ deck: deckA, lightStatus: 1 }));
      dispatch(whiteLightDeckInitiated({ deck: deckB, lightStatus: 1 }));
    }
  };

  return (
    <StyledWhiteLightComponent
      className={`d-flex align-items-center ${className}`}
    >
      <ImageIcon className="bulb" src={bulb} />
      <label className="switch" style={{ marginLeft: "10px" }}>
        <input
          type="checkbox"
          checked={
            lightReducerForDeckA.isLightOn || lightReducerForDeckB.isLightOn
          }
          onClick={handleWhiteLightClick}
        />
        <span className="slider round"> </span>
      </label>
    </StyledWhiteLightComponent>
  );
};

WhiteLight.propTypes = {
  isLightOn: PropTypes.bool,
  className: PropTypes.string,
};

WhiteLight.defaultProps = {
  className: "",
};

export default WhiteLight;

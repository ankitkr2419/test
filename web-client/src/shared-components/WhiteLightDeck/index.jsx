import React from "react";
import PropTypes from "prop-types";
import { ImageIcon } from "shared-components";
import bulb from "assets/images/bulb.svg";
import { StyledWhiteLightComponent } from "./Style";
import { DECKNAME } from "appConstants";
import { useSelector } from "react-redux";

const WhiteLightDeck = (props) => {
  const { handleWhiteLightClick, className, deckName } = props;

  const lightReducer = useSelector((state) => state.whiteLightReducer);
  let lightReducerForDeckA = lightReducer.decks.find(
    (deckObj) => deckObj.name === DECKNAME.DeckA
  );
  let lightReducerForDeckB = lightReducer.decks.find(
    (deckObj) => deckObj.name === DECKNAME.DeckB
  );

  let lightDeck =
    deckName == DECKNAME.DeckA ? lightReducerForDeckA : lightReducerForDeckB;

  return (
    <StyledWhiteLightComponent
      className={`d-flex align-items-center ${className}`}
    >
      <ImageIcon className="bulb" src={bulb} />
      <label className="switch">
        <input
          type="checkbox"
          onClick={handleWhiteLightClick}
          checked={lightDeck.isLightOn}
        />
        <span className="slider round"> </span>
      </label>
    </StyledWhiteLightComponent>
  );
};

WhiteLightDeck.propTypes = {
  isLightOn: PropTypes.bool,
  className: PropTypes.string,
};

WhiteLightDeck.defaultProps = {
  className: "",
};

export default WhiteLightDeck;

import React from "react";

import { Col, FormGroup, Button } from "core-components";
import { Text } from "shared-components";
import { deckPositionNames } from "./functions";
import { DeckPositionInfoBox } from "./Style";
import { CommonTipHeightComponent } from "./CommonTipHeightComponent";

const DeckPositionInfo = (props) => {
  const { formik, activeTab } = props;

  const handleDeckClick = (id) => {
    let otherTabsAreDisabled = true;
    let deckPosition = deckPositionNames[parseInt(id)];
    const deck = formik.values.deck;
    const formikDeckPosition = deck.deckPosition;
    const tipHeightValue = deck.tipHeight;

    // if already selected than we de-select it
    // also here we check and change isFilled parameter
    if (formikDeckPosition && formikDeckPosition === deckPosition) {
      deckPosition = null;

      // check if tipheight for deck is null or empty, if true we enable other tabs.
      if (!tipHeightValue) {
        otherTabsAreDisabled = false;
      }
    }
    // set deckPosition value
    formik.setFieldValue(`deck.deckPosition`, deckPosition);

    //enable/disable other tabs
    formik.setFieldValue(`cartridge1.isDisabled`, otherTabsAreDisabled);
    formik.setFieldValue(`cartridge2.isDisabled`, otherTabsAreDisabled);
  };

  return (
    <>
      <DeckPositionInfoBox>
        <div className="process-box deck-position-box mx-auto">
          <div className="mb-3 border-bottom-line">
            <FormGroup row>
              <Text Tag="h5" md={12} className="title-heading">
                Select Deck Position
              </Text>
              <Col md={12} className="deck-position-options">
                {deckPositionNames.map((deckPositionName, index) => {
                  return (
                    <Button
                      key={index}
                      id={index}
                      outline
                      className={
                        formik.values.deck.deckPosition ===
                        deckPositionNames[index]
                          ? "selected-opt"
                          : ""
                      }
                      onClick={(e) => handleDeckClick(e.target.id)}
                    >
                      {deckPositionName}
                    </Button>
                  );
                })}
              </Col>
            </FormGroup>
          </div>
          <CommonTipHeightComponent formik={formik} activeTab={activeTab} />
        </div>
      </DeckPositionInfoBox>
    </>
  );
};

DeckPositionInfo.propTypes = {};

export default DeckPositionInfo;

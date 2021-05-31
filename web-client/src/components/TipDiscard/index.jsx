import React from "react";

import { Card, CardBody, Radio } from "core-components";
import { ButtonIcon, ButtonBar, ImageIcon, Text } from "shared-components";

import tipDiscardProcessGraphics from "assets/images/tip-discard-process-graphics.svg";
import graphicsAtPickup from "assets/images/graphics-at-pickup.svg";
import graphicsAtDiscard from "assets/images/graphics-at-discard.svg";
import longDownArrow from "assets/images/long-down-arrow.svg";
import TopHeading from "shared-components/TopHeading";
import { PageBody, TipDiscardBox, TopContent } from "./Style";
import { useDispatch, useSelector } from "react-redux";
import { Redirect } from "react-router";
import { ROUTES } from "appConstants";
import { saveTipDiscardInitiated } from "action-creators/processesActionCreators";

const TipDiscardComponent = (props) => {
  const dispatch = useDispatch();

  const loginReducer = useSelector((state) => state.loginReducer);

  const recipeDetailsReducer = useSelector(
    (state) => state.updateRecipeDetailsReducer
  );
  const recipeID = recipeDetailsReducer.recipeDetails.id;
  const token = recipeDetailsReducer.token;

  //this API call is not completed from backend.
  //It will be added in future.
  //Will update it then.
  const saveBtnHandler = () => {
    const body = {
      type: "discard",
      // "position": null
    };
    const requestBody = {
      body: body,
      recipeID: recipeID,
      token: token,
    };

    dispatch(saveTipDiscardInitiated(requestBody));
  };

  const loginReducerData = loginReducer.toJS();
  let activeDeckObj =
    loginReducerData && loginReducerData.decks.find((deck) => deck.isActive);
  if (!activeDeckObj.isLoggedIn) {
    return <Redirect to={`/${ROUTES.landing}`} />;
  }

  return (
    <>
      <PageBody>
        <TipDiscardBox>
          <div className="process-content -tip-discard px-2">
            <TopContent className="d-flex justify-content-between align-items-center mx-5">
              <div className="d-flex flex-column">
                <div className="d-flex align-items-center frame-icon">
                  <ButtonIcon
                    size={60}
                    name="tip-discard"
                    className="text-primary bg-white border-gray"
                  />
                  <TopHeading titleHeading="Tip Discard" />
                </div>
              </div>
            </TopContent>
            <Card>
              <CardBody className="p-5 overflow-hidden">
                <div className="process-box mx-auto py-5 d-flex">
                  <div className="magnet-large-btn d-flex justify-content-around align-items-center flex-column">
                    <Radio
                      id="attach"
                      name="magnet-action"
                      label="At Pickup"
                      className=""
                    />
                    <div className="position-relative">
                      <Text Tag="span" className="pickup-point" />
                      <ImageIcon
                        src={graphicsAtPickup}
                        alt="Tip Pickup Process"
                      />
                    </div>
                  </div>
                  <div className="magnet-large-btn d-flex justify-content-around align-items-center flex-column bg-white ml-4 position-relative selected">
                    <Radio
                      id="detach"
                      name="magnet-action"
                      label="At Discard"
                      checked
                    />
                    <div className="position-relative">
                      <Text Tag="span" className="discard-point" />
                      <ImageIcon
                        src={graphicsAtDiscard}
                        alt="Tip Discard Process"
                      />
                      <ImageIcon
                        src={longDownArrow}
                        alt="Down Arrow"
                        className="long-down-arrow"
                      />
                    </div>
                  </div>
                  <ImageIcon
                    src={tipDiscardProcessGraphics}
                    alt="Tip Discard"
                    className="process-image"
                  />
                </div>
              </CardBody>
            </Card>
            <ButtonBar rightBtnLabel="Save" handleRightBtn={saveBtnHandler} />
          </div>
        </TipDiscardBox>
      </PageBody>
    </>
  );
};

TipDiscardComponent.propTypes = {};

export default React.memo(TipDiscardComponent);

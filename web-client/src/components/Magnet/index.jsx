import React from "react";

import { useDispatch, useSelector } from "react-redux";
import { useState } from "react";

import { Card, CardBody, Radio } from "core-components";
import { ButtonIcon, ButtonBar, ImageIcon, Icon } from "shared-components";

import magnetProcessGraphics from "assets/images/magnet-process-graphics.svg";
import TopHeading from "shared-components/TopHeading";
import { PageBody, MagnetBox, TopContent } from "./Style";
import { saveMagnetInitiated } from "action-creators/processesActionCreators";
import { Redirect } from "react-router";
import { ROUTES } from "appConstants";

const MagnetComponent = (props) => {
  const [isAttach, setIsAttach] = useState(null);
  const dispatch = useDispatch();

  const loginReducer = useSelector((state) => state.loginReducer);
  const recipeDetailsReducer = useSelector(
    (state) => state.updateRecipeDetailsReducer
  );
  const recipeID = recipeDetailsReducer.recipeDetails.id;
  const token = recipeDetailsReducer.token;

  const saveBtnHandler = () => {
    const body = {
      operation: isAttach ? "attach" : "detach",
      operation_type: "wash", // will change in future
    };
    const requestBody = {
      body: body,
      recipeID: recipeID,
      token: token,
    };
    dispatch(saveMagnetInitiated(requestBody));
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
        <MagnetBox>
          <div className="process-content process-magnet px-2">
            <TopContent className="d-flex justify-content-between align-items-center mx-5">
              <div className="d-flex flex-column">
                <div className="d-flex align-items-center frame-icon">
                  <ButtonIcon
                    size={60}
                    name="magnet"
                    className="text-primary bg-white border-gray"
                  />
                  <TopHeading titleHeading="Magnet" />
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
                      label="Attach"
                      className=""
                      onClick={() => setIsAttach(true)}
                    />
                    <div className="d-flex justify-content-center align-items-center flex-column animated-attach-icon">
                      <Icon name="upward-magnet" size={35} />
                      <Icon name="downward-magnet" size={35} />
                    </div>
                  </div>
                  <div className="magnet-large-btn d-flex justify-content-around align-items-center flex-column bg-white ml-4">
                    <Radio
                      id="detach"
                      name="magnet-action"
                      label="Detach"
                      className=""
                      onClick={() => setIsAttach(false)}
                    />
                    <div className="d-flex justify-content-center align-items-center flex-column animated-detach-icon">
                      <Icon name="upward-magnet" size={35} />
                      <Icon name="downward-magnet" size={35} />
                    </div>
                  </div>
                  <ImageIcon
                    src={magnetProcessGraphics}
                    alt="Magnet Process"
                    className="process-image"
                  />
                </div>
              </CardBody>
            </Card>
            <ButtonBar rightBtnLabel="Save" handleRightBtn={saveBtnHandler} />
          </div>
        </MagnetBox>
      </PageBody>
    </>
  );
};

MagnetComponent.propTypes = {};

export default React.memo(MagnetComponent);

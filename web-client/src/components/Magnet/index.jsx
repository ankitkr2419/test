import React, { useEffect } from "react";

import { useDispatch, useSelector } from "react-redux";
import { useState } from "react";

import { Card, CardBody, Radio, Input } from "core-components";
import {
  ButtonIcon,
  ButtonBar,
  ImageIcon,
  Icon,
  Text,
} from "shared-components";

import magnetProcessGraphics from "assets/images/magnet-process-graphics.svg";
import TopHeading from "shared-components/TopHeading";
import { PageBody, MagnetBox, TopContent } from "./Style";
import { saveProcessInitiated } from "action-creators/processesActionCreators";
import { Redirect, useHistory } from "react-router";
import { API_ENDPOINTS, HTTP_METHODS, ROUTES } from "appConstants";
import { toast } from "react-toastify";

const MagnetComponent = (props) => {
  const dispatch = useDispatch();
  const history = useHistory();
  const [isAttach, setIsAttach] = useState(false);
  const [height, setHeight] = useState(0);

  const editReducer = useSelector((state) => state.editProcessReducer);
  const editReducerData = editReducer.toJS();
  const processesReducer = useSelector((state) => state.processesReducer);
  const loginReducer = useSelector((state) => state.loginReducer);
  const loginReducerData = loginReducer.toJS();
  let activeDeckObj =
    loginReducerData && loginReducerData.decks.find((deck) => deck.isActive);
  const recipeDetailsReducer = useSelector(
    (state) => state.updateRecipeDetailsReducer
  );

  const recipeID = recipeDetailsReducer.recipeDetails.id;
  const token = activeDeckObj.token;

  const isAttachFromAPI = editReducerData?.operation;
  useEffect(() => {
    if (editReducerData?.operation) {
      setIsAttach(isAttachFromAPI === "attach");
      setHeight(editReducerData.height);
    }
  }, [isAttachFromAPI]);

  const errorInAPICall = processesReducer.error;
  useEffect(() => {
    if (errorInAPICall === false) {
      history.push(ROUTES.processListing);
    }
  }, [errorInAPICall, isAttachFromAPI]);

  const handleHeightBlur = (event) => {
    if (event.target.value === "") {
      if (editReducerData?.operation) {
        setHeight(editReducerData.height);
      } else {
        setHeight(0);
      }
    }
  };

  const saveBtnHandler = () => {
    const heightIsInt = `${height}`.match(/^[0-9]\d*$/);
    if (isAttach && !heightIsInt) {
      toast.error("Please enter valid height", { autoClose: false });
      return;
    }
    const body = {
      operation: isAttach ? "attach" : "detach",
      operation_type: "wash", // will change in future
      height: parseInt(height),
    };
    const requestBody = {
      body: body,
      id: editReducerData?.process_id ? editReducerData.process_id : recipeID,
      token: token,
      api: API_ENDPOINTS.magnet,
      method: editReducerData?.process_id
        ? HTTP_METHODS.PUT
        : HTTP_METHODS.POST,
    };
    dispatch(saveProcessInitiated(requestBody));
  };

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
            <Card className="fix-height-card">
              <CardBody className="p-5 overflow-hidden">
                <div className="process-box mx-auto py-4 d-flex">
                  <div className="magnet-large-btn d-flex justify-content-around align-items-center flex-column">
                    <div style={{ width: "8rem" }} className="d-flex">
                      <Text className={isAttach ? "" : "text-muted"}>
                        Height
                      </Text>
                      <Input
                        type="number"
                        name="height"
                        id="height"
                        placeholder=""
                        value={height}
                        disabled={!isAttach}
                        onChange={(e) => setHeight(e.target.value)}
                        onBlur={(e) => handleHeightBlur(e)}
                      />
                    </div>

                    <Radio
                      id="attach"
                      name="magnet-action"
                      label="Attach"
                      className=""
                      checked={isAttach}
                      onClick={() => setIsAttach(!isAttach)}
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
                      checked={!isAttach}
                      onClick={() => setIsAttach(!isAttach)}
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
            <ButtonBar
              rightBtnLabel="Save"
              handleRightBtn={saveBtnHandler}
              btnBarClassname={"btn-bar-adjust-magnet"}
            />
          </div>
        </MagnetBox>
      </PageBody>
    </>
  );
};

MagnetComponent.propTypes = {};

export default React.memo(MagnetComponent);

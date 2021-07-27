import React, { useEffect } from "react";
import { useHistory } from "react-router";
import { useDispatch, useSelector } from "react-redux";
import CalibrationExtractionComponent from "components/CalibrationExtraction";

const CalibrationExtractionContainer = () => {
  const dispatch = useDispatch();
  const history = useHistory();

  //get login reducer details
  const loginReducer = useSelector((state) => state.loginReducer);
  const loginReducerData = loginReducer.toJS();
  let activeDeckObj = loginReducerData?.decks.find((deck) => deck.isActive);
  const { token } = activeDeckObj;

  //api call to get configurations
  // useEffect(() => {
  //   if (token) {
  //     //TODO initial api's if required
  //   }
  // }, [dispatch, token]);

  return <CalibrationExtractionComponent />;
};

export default React.memo(CalibrationExtractionContainer);

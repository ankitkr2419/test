import React, { useEffect, useState } from "react";
import { useHistory } from "react-router";
import { useDispatch, useSelector } from "react-redux";
import CalibrationExtractionComponent from "components/CalibrationExtraction";
import { logoutInitiated } from "action-creators/loginActionCreators";

const CalibrationExtractionContainer = () => {
  const dispatch = useDispatch();
  const history = useHistory();

  const [showConfirmationModal, setConfirmModal] = useState(false);

  //get login reducer details
  const loginReducer = useSelector((state) => state.loginReducer);
  const loginReducerData = loginReducer.toJS();
  let activeDeckObj = loginReducerData?.decks.find((deck) => deck.isActive);
  const { name, token } = activeDeckObj;

  const handleLogout = () => {
    dispatch(
      logoutInitiated({ deckName: name, token: token, showToast: true })
    );
  };

  const toggleConfirmModal = () => setConfirmModal(!showConfirmationModal);

  //api call to get configurations
  // useEffect(() => {
  //   if (token) {
  //     //TODO initial api's if required
  //   }
  // }, [dispatch, token]);

  return (
    <CalibrationExtractionComponent
      toggleConfirmModal={toggleConfirmModal}
      handleLogout={handleLogout}
      showConfirmationModal={showConfirmationModal}
      deckName={name}
    />
  );
};

export default React.memo(CalibrationExtractionContainer);

import React, { useEffect, useState } from "react";
import { StyledProcessListing } from "./Styles";
import { useDispatch, useSelector } from "react-redux";
import { Redirect } from "react-router";
import {
  ROUTES,
  MODAL_MESSAGE,
  MODAL_BTN,
  CARTRIDGE_1,
  CARTRIDGE_2,
} from "appConstants";
import TopContentComponent from "components/RecipeListing/TopContentComponent";
import ProcessListingCards from "./ProcessListingCards";
import { ButtonBar } from "shared-components";
import MlModal from "shared-components/MlModal";
import { useHistory } from "react-router";
import {
  getCartridge1ActionInitiated,
  getCartridge2ActionInitiated,
  getCartridgeActionInitiated,
} from "action-creators/saveNewRecipeActionCreators";
import { ResetProcessPageState } from "action-creators/PageActionCreators";

const ProcessListComponent = (props) => {
  let {
    recipeDetails,
    processList,
    toggleIsOpen,
    draggedProcessId,
    setDraggedProcessId,
    handleChangeSequenceTo,
    handleProcessMove,
    createDuplicateProcess,
    handleEditProcess,
    onFinishConfirmation,
    handleAddProcessClick,
    handleDeleteProcess,
  } = props;

  const [finishModal, setFinishModal] = useState(false);
  const [deleteModal, setDeleteModal] = useState(false);
  const [deleteProcessId, setDeleteProcessId] = useState(null);

  const history = useHistory();
  const dispatch = useDispatch();

  /** get isApiCalled field */
  const cartridge1DetailsReducer = useSelector(
    (state) => state.cartridge1DetailsReducer
  );
  const cartridge2DetailsReducer = useSelector(
    (state) => state.cartridge2DetailsReducer
  );

  /**get active login deck data*/
  const loginReducer = useSelector((state) => state.loginReducer);
  const loginReducerData = loginReducer.toJS();
  let activeDeckObj = loginReducerData?.decks.find((deck) => deck.isActive);
  let deckName = activeDeckObj.name;

  /** call and fill cartridge1 & cartridge2 reducers */
  useEffect(() => {
    const { token } = activeDeckObj;
    const { name, pos_cartridge_1, pos_cartridge_2 } = recipeDetails;

    /** Checks if API is already called */
    const { isApiCalled: cartridge1ApiCalled } = cartridge1DetailsReducer;
    const { isApiCalled: cartridge2ApiCalled } = cartridge2DetailsReducer;

    let params = { deckName: name, token: token };
    if (pos_cartridge_1 && cartridge1ApiCalled === false) {
      params = { ...params, id: pos_cartridge_1, type: CARTRIDGE_1 };
      dispatch(getCartridge1ActionInitiated(params));
    }
    if (pos_cartridge_2 && cartridge2ApiCalled === false) {
      params = { ...params, id: pos_cartridge_2, type: CARTRIDGE_2 };
      dispatch(getCartridge2ActionInitiated(params));
    }
  }, [recipeDetails]);

  /**
   * if user is not logged in, go to landing page
   */
  if (!activeDeckObj.isLoggedIn) {
    return <Redirect to={`/${ROUTES.landing}`} />;
  }

  const toggleFinishModal = () => {
    setFinishModal(!finishModal);
  };
  const onFinishConfirmationClick = () => {
    dispatch(ResetProcessPageState(deckName));
    toggleFinishModal();
    onFinishConfirmation();
  };

  const handleDeleteProcessClick = (processId) => {
    setDeleteProcessId(processId);
    toggleDeleteModal();
  };

  const toggleDeleteModal = () => {
    setDeleteModal(!deleteModal);
  };

  const onConfirmedDeleteProcess = () => {
    toggleDeleteModal();
    handleDeleteProcess(deleteProcessId);
  };

  const handleBackToRecipeList = () => {
    dispatch(ResetProcessPageState(deckName));
    history.push(ROUTES.recipeListing);
  };

  return (
    <StyledProcessListing>
      <div className="landing-content px-2">
        {/**TopContentComponent: to show recipe details at top */}
        <TopContentComponent
          isProcessListingPage={true}
          recipeName={recipeDetails.name}
          createdAt={recipeDetails.created_at}
          updatedAt={recipeDetails.updated_at}
          processListBackButtonHandler={handleBackToRecipeList}
        />

        {/** ProcessListingCards: pagination/ process list */}
        <ProcessListingCards
          processList={processList}
          toggleIsOpen={toggleIsOpen}
          draggedProcessId={draggedProcessId}
          setDraggedProcessId={setDraggedProcessId}
          handleChangeSequenceTo={handleChangeSequenceTo}
          handleProcessMove={handleProcessMove}
          createDuplicateProcess={createDuplicateProcess}
          handleEditProcess={handleEditProcess}
          handleDeleteProcess={handleDeleteProcessClick}
        />

        {/* Action buttons (add process, finish)*/}
        <ButtonBar
          leftBtnLabel="Add Process"
          rightBtnLabel="Finish"
          handleLeftBtn={handleAddProcessClick}
          handleRightBtn={toggleFinishModal}
          handleBackToRecipeList={handleBackToRecipeList}
          pageReset={true}
        />

        {/**finish confirmation modal */}
        {finishModal && (
          <MlModal
            isOpen={finishModal}
            textHead={deckName}
            textBody={`${MODAL_MESSAGE.finishProcessListConfirmation}${recipeDetails.name}`}
            handleSuccessBtn={onFinishConfirmationClick}
            handleCrossBtn={toggleFinishModal}
            successBtn={MODAL_BTN.yes}
            failureBtn={MODAL_BTN.no}
          />
        )}

        {/**delete confirmation modal */}
        {deleteModal && (
          <MlModal
            isOpen={deleteModal}
            textHead={deckName}
            textBody={MODAL_MESSAGE.deleteProcessConfirmation}
            handleSuccessBtn={onConfirmedDeleteProcess}
            handleCrossBtn={toggleDeleteModal}
            successBtn={MODAL_BTN.yes}
            failureBtn={MODAL_BTN.no}
          />
        )}
      </div>
    </StyledProcessListing>
  );
};

export default React.memo(ProcessListComponent);

import React, { useState } from "react";
import { StyledProcessListing } from "./Styles";
import { useSelector } from "react-redux";
import { Redirect, useHistory } from "react-router";
import { ROUTES, MODAL_MESSAGE, MODAL_BTN } from "appConstants";
import TopContentComponent from "components/RecipeListing/TopContentComponent";
import ProcessListingCards from "./ProcessListingCards";
import { ButtonBar } from "shared-components";
import MlModal from "shared-components/MlModal";

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
    } = props;

    const [finishModal, setFinishModal] = useState(false);
    const history = useHistory();

    /**get active login deck data*/
    const loginReducer = useSelector((state) => state.loginReducer);
    const loginReducerData = loginReducer.toJS();
    let activeDeckObj = loginReducerData?.decks.find((deck) => deck.isActive);
    let deckName = activeDeckObj.name;
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
        toggleFinishModal();
        onFinishConfirmation();
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
                />

                {/* Action buttons (add process, finish)*/}
                <ButtonBar
                    leftBtnLabel="Add Process"
                    rightBtnLabel="Finish"
                    handleLeftBtn={handleAddProcessClick}
                    handleRightBtn={toggleFinishModal}
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
            </div>
        </StyledProcessListing>
    );
};

export default React.memo(ProcessListComponent);

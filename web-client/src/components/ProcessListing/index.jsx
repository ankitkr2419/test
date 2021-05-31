import React from "react";
import { StyledProcessListing } from "./Styles";
import { useSelector } from "react-redux";
import { Redirect, useHistory } from "react-router";
import { ROUTES } from "appConstants";
import TopContentComponent from "components/RecipeListing/TopContentComponent";
import ProcessListingCards from "./ProcessListingCards";
import { ButtonBar } from "shared-components";

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
    } = props;

    const history = useHistory();

    /**get active login deck data*/
    const loginReducer = useSelector((state) => state.loginReducer);
    const loginReducerData = loginReducer.toJS();
    let activeDeckObj = loginReducerData?.decks.find((deck) => deck.isActive);

    /**
     * if user is not logged in, go to landing page
     */
    if (!activeDeckObj.isLoggedIn) {
        return <Redirect to={`/${ROUTES.landing}`} />;
    }

    const handleAddProcessClick = () => {
        history.push(ROUTES.selectProcess);
    };

    const handleFinishClick = () => {
        //TODO: required api calls
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
                    handleRightBtn={handleFinishClick}
                />
            </div>
        </StyledProcessListing>
    );
};

export default React.memo(ProcessListComponent);

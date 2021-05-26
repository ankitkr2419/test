import React, { useEffect, useState } from "react";
import { useDispatch, useSelector } from "react-redux";
import ProcessListComponent from "components/ProcessListing";
import {
    changeProcessSequences,
    sortProcessListBySequence,
} from "components/ProcessListing/helper";
import {
    processListInitiated,
    duplicateProcessInitiated,
    setProcessList,
} from "action-creators/processActionCreators";
import { Loader } from "shared-components";
import { toast } from "react-toastify";

const ProcessListingContainer = (props) => {
    //if we have draggedProcessId, means user selected this process to change its sequence (move)
    const [draggedProcessId, setDraggedProcessId] = useState(null);

    const dispatch = useDispatch();
    const processListReducer = useSelector(
        (state) => state.processListReducer
    ).toJS();
    const { isLoading, processList } = processListReducer;

    /**get active login deck data*/
    const loginReducer = useSelector((state) => state.loginReducer);
    const loginReducerData = loginReducer.toJS();
    let activeDeckObj = loginReducerData?.decks.find((deck) => deck.isActive);

    /**get recipeDetails */
    const recipeDetailsReducer = useSelector(
        (state) => state.updateRecipeDetailsReducer
    );
    const recipeDetails = recipeDetailsReducer.recipeDetails;

    /**Get process list of a recipe */
    useEffect(() => {
        const recipeId = recipeDetails.id;
        const token = activeDeckObj.token;
        if (recipeId) {
            dispatch(processListInitiated({ recipeId, token }));
        }
    }, [recipeDetails.id]);

    //toggle isOpen field of process object to toggle process menu
    const toggleIsOpen = (processId) => {
        const newProcessList = processList.map((processObj) => {
            return processObj.id === processId
                ? { ...processObj, isOpen: !processObj.isOpen }
                : processObj;
        });
        dispatch(setProcessList({ processList: newProcessList }));
    };

    /** purpose=> setting draggedProcessId will let us know that this process is dragged(move), to toggle process menu for drop operations */
    const handleDraggedProcessId = (processId) => {
        //if draggedProcessId not found then store it, else clear old one
        draggedProcessId
            ? setDraggedProcessId(null)
            : setDraggedProcessId(processId);

        //if processId not found, means 'move' operation already done, no need to toggle
        processId && toggleIsOpen(processId);
    };

    //move
    const handleChangeSequenceTo = (droppedProcessId) => {
        moveProcessAndSave(draggedProcessId, droppedProcessId);

        //reset drag-drop (move)
        handleDraggedProcessId(null);
    };

    //up or down
    const handleProcessMove = (processId, sequenceNumber, direction) => {
        const dropProcess = processList.find((obj) => {
            return direction === "up"
                ? obj.sequence_num === sequenceNumber - 1
                : obj.sequence_num === sequenceNumber + 1;
        });

        const droppedProcessId = dropProcess?.id;

        if (!droppedProcessId) {
            toast.error(
                direction === "up"
                    ? "We can not move first process up"
                    : "We can not move last process down"
            );
            //hide menu
            toggleIsOpen(processId);
            return;
        }

        moveProcessAndSave(processId, droppedProcessId);
    };

    //common method for up/down/move operations
    const moveProcessAndSave = (draggedProcessId, droppedProcessId) => {
        let arr = changeProcessSequences(
            processList,
            draggedProcessId,
            droppedProcessId
        );
        let sortedArr = sortProcessListBySequence(arr);
        dispatch(setProcessList({ processList: sortedArr }));
    };

    const createDuplicateProcess = (processId) => {
        const token = activeDeckObj.token;
        dispatch(duplicateProcessInitiated({ processId, token }));
    };

    return (
        <>
            {isLoading && <Loader />}
            <ProcessListComponent
                recipeDetails={recipeDetails}
                processList={processList}
                toggleIsOpen={toggleIsOpen}
                draggedProcessId={draggedProcessId}
                setDraggedProcessId={handleDraggedProcessId} //move dragged
                handleChangeSequenceTo={handleChangeSequenceTo} //move dropped
                handleProcessMove={handleProcessMove} //up and down
                createDuplicateProcess={createDuplicateProcess}
            />
        </>
    );
};

export default React.memo(ProcessListingContainer);

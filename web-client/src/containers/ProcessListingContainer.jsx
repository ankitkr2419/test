import React, { useEffect, useState } from "react";
import { useDispatch, useSelector } from "react-redux";
import ProcessListComponent from "components/ProcessListing";
import {
    changeProcessSequences,
    sortProcessListBySequence,
} from "components/ProcessListing/helper";
import { processListInitiated } from "action-creators/processActionCreators";
import { Loader } from "shared-components";
import { toast } from "react-toastify";

const ProcessListingContainer = (props) => {
    /*TODO: 1) get recipe details from reducer*/
    const [recipeDetails, setRecipeDetails] = useState({
        id: "8b5cd741-b6f7-443e-8e8b-5f1f1772d052",
        name: "test",
        created_at: "2021-04-29T11:52:11.171692Z",
        updated_at: "2021-04-29T11:52:11.171692Z",
    });

    const [processList, setProcessList] = useState([]);

    //if we have draggedProcessId, means user selected this process to change its sequence (move)
    const [draggedProcessId, setDraggedProcessId] = useState(null);

    const dispatch = useDispatch();
    const processListReducer = useSelector(
        (state) => state.processListReducer
    ).toJS();
    const { isLoading, error } = processListReducer;

    /**Get process list of a recipe */
    useEffect(() => {
        const recipeId = recipeDetails.id;
        if (recipeId) {
            dispatch(processListInitiated({ recipeId }));
        }
    }, [recipeDetails.id]);

    /** => store processList in local state (needed for change sequence and other local changes)
        => isOpen: represents that process menu should be open or not and its independently handled for each process
        => isOpen not coming from api, so adding default value */
    useEffect(() => {
        if (!isLoading && !error) {
            const list = processListReducer.processList?.map((obj) => ({
                ...obj,
                isOpen: false,
            }));
            setProcessList(list);
        }
    }, [isLoading, error]);

    //toggle isOpen field of process object to toggle process menu
    const toggleIsOpen = (processId) => {
        const newProcessList = processList.map((processObj) => {
            return processObj.id === processId
                ? { ...processObj, isOpen: !processObj.isOpen }
                : processObj;
        });
        setProcessList(newProcessList);
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
        setProcessList(sortedArr);
        toast.success("Process Moved");
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
            />
        </>
    );
};

export default React.memo(ProcessListingContainer);

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
    fetchProcessDataInitiated,
    fetchProcessDataReset,
    sequenceInitiated,
    sequenceReset,
    deleteProcessInitiated,
} from "action-creators/processActionCreators";
import { saveProcessReset } from "action-creators/processesActionCreators";
import { Loader } from "shared-components";
import { toast } from "react-toastify";
import { SELECT_PROCESS_PROPS, ROUTES } from "appConstants";
import { useHistory } from "react-router";

const ProcessListingContainer = (props) => {
    //if we have draggedProcessId, means user selected this process to change its sequence (move)
    const [draggedProcessId, setDraggedProcessId] = useState(null);

    //these fields tells us what to do next after sequence api success
    const [isFinish, setIsFinish] = useState(false);
    const [isAddProcess, setIsAddProcess] = useState(false);

    const dispatch = useDispatch();
    const history = useHistory();
    const processListReducer = useSelector(
        (state) => state.processListReducer
    ).toJS();
    const { isLoading, processList, error, sequenceError } = processListReducer;

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

    /**clear edit process reducer/ clear sequence data in reducer */
    useEffect(() => {
        dispatch(fetchProcessDataReset());
        dispatch(sequenceReset());
    }, []);

    /** after sequence api success  */
    useEffect(() => {
        if (sequenceError === false) {
            if (isFinish) {
                history.push(ROUTES.recipeListing);
            } else if (isAddProcess) {
                history.push(ROUTES.selectProcess);
            }
        }
    }, [sequenceError]);

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

    const handleEditProcess = (processObj) => {
        const processType = processObj.type;
        const token = activeDeckObj.token;
        //fetch process data, store in reducer
        dispatch(
            fetchProcessDataInitiated({
                processId: processObj.id,
                type: processType,
                token,
            })
        );
        //redirect to edit process depend on processType
        const routePathObj = SELECT_PROCESS_PROPS.find(
            (obj) => obj.processType === processType
        );
        const routePath = routePathObj?.route;
        history.push(routePath);
    };

    /** before finish/addProcess redirection, sequence api should be called and this redirection is handled in useeffect*/
    const onFinishConfirmation = () => {
        setIsFinish(true);
        setIsAddProcess(false);
        callSequenceApi();
    };

    const handleAddProcessClick = () => {
        setIsAddProcess(true);
        setIsFinish(false);
        callSequenceApi();
    };

    const callSequenceApi = () => {
        const token = activeDeckObj.token;
        //convert to required format and call api
        const newSequenceArray = processList.map((obj) => ({
            process_id: obj.id,
            sequence_num: obj.sequence_num,
        }));
        dispatch(
            sequenceInitiated({
                recipeId: recipeDetails?.id,
                processList: newSequenceArray,
                token,
            })
        );
    };

    const handleDeleteProcess = (processId) => {
        const token = activeDeckObj.token;
        dispatch(deleteProcessInitiated({ processId, token }));
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
                createDuplicateProcess={createDuplicateProcess} //copy
                handleEditProcess={handleEditProcess} //edit
                onFinishConfirmation={onFinishConfirmation} //finish
                handleAddProcessClick={handleAddProcessClick} //add process
                handleDeleteProcess={handleDeleteProcess} //delete process
            />
        </>
    );
    //redirect to edit process depend on processType
    const routePathObj = SELECT_PROCESS_PROPS.find(
      (obj) => obj.processType === processType
    );
    const routePath = routePathObj?.route;
    history.push(routePath);
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
        createDuplicateProcess={createDuplicateProcess} //copy
        handleEditProcess={handleEditProcess} //edit
      />
    </>
  );
};

export default React.memo(ProcessListingContainer);

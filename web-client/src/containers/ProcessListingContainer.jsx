import React, { useState } from "react";
import ProcessListComponent from "components/ProcessListing";
import {
    changeProcessSequences,
    sortProcessListBySequence,
} from "components/ProcessListing/helper";

const ProcessListingContainer = (props) => {
    /*TODO: 1) fetch dynamic process list from recipeId, also add isOpen field
     * 2) while storing process list in state, it should be sorted by sequence
     */
    const [recipeDetails, setRecipeDetails] = useState({
        name: "test",
        created_at: "2021-04-29T11:52:11.171692Z",
        updated_at: "2021-04-29T11:52:11.171692Z",
    });

    /**isOpen: represents that process menu should be open or not and its independently handled for each process
               also this field does not come from api, so adding isOepn:false by default*/
    const [processList, setProcessList] = useState([
        {
            id: "a5d058e3-7ce3-4a42-b2da-690e47139612",
            name: "AD-WW-c1-1-2",
            type: "AspireDispense",
            recipe_id: "6b7fcfa2-8337-4d79-829a-e9bd486a2de4",
            sequence_num: 1,
            created_at: "2021-03-11T17:34:16.507414Z",
            updated_at: "2021-03-11T17:34:16.507414Z",
            isOpen: false,
        },
        {
            id: "4fa5c4e3-699c-47bb-ac7a-b26d04efaeb5",
            name: "AD-WS-c1-1",
            type: "AspireDispense",
            recipe_id: "6b7fcfa2-8337-4d79-829a-e9bd486a2de4",
            sequence_num: 2,
            created_at: "2021-03-11T17:34:16.555738Z",
            updated_at: "2021-03-11T17:34:16.555738Z",
            isOpen: false,
        },
        {
            id: "fb88bada-ced7-4fa2-b845-4bb91e74341e",
            name: "AD-SW-c1-2",
            type: "AspireDispense",
            recipe_id: "6b7fcfa2-8337-4d79-829a-e9bd486a2de4",
            sequence_num: 3,
            created_at: "2021-03-11T17:34:16.580361Z",
            updated_at: "2021-03-11T17:34:16.580361Z",
            isOpen: false,
        },
    ]);

    //if we have processId, means user selected this process to change its sequence
    const [draggedProcessId, setDraggedProcessId] = useState(null);

    //toggle isOpen field of process object to toggle process menu
    const toggleIsOpen = (processId) => {
        const newProcessList = processList.map((processObj) => {
            return processObj.id === processId
                ? { ...processObj, isOpen: !processObj.isOpen }
                : processObj;
        });
        setProcessList(newProcessList);
    };

    /**setting this id will let us know that this process is dragged(move), to toggle process menu for drop operations */
    const handleDraggedProcessId = (processId) => {
        setDraggedProcessId(processId);
    };

    const handleChangeSequenceTo = (droppedProcessId) => {
        let arr = changeProcessSequences(
            processList,
            draggedProcessId,
            droppedProcessId
        );
        let sortedArr = sortProcessListBySequence(arr);
        setProcessList(sortedArr);

        //reset drag-drop (move)
        handleDraggedProcessId(null);
    };

    return (
        <ProcessListComponent
            recipeDetails={recipeDetails}
            processList={processList}
            toggleIsOpen={toggleIsOpen}
            draggedProcessId={draggedProcessId}
            setDraggedProcessId={handleDraggedProcessId}
            handleChangeSequenceTo={handleChangeSequenceTo}
        />
    );
};

export default React.memo(ProcessListingContainer);

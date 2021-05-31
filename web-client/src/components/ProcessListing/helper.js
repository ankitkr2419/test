/**returns array with exchanged sequence number of processes (with: indexOne, indexTwo) and sorted by sequence number */
export const changeProcessSequences = (array, sequenceProcessIdOne, sequenceProcessIdTwo) => {
    const processOfIndexOne = {
        ...array.find((pOne) => pOne.id === sequenceProcessIdOne),
    };
    const processOfIndexTwo = {
        ...array.find((pTwo) => pTwo.id === sequenceProcessIdTwo),
    };

    //change sequence logic: take process object of sequenceProcessIdOne and sequence number of sequenceProcessIdTwo (vice-versa)
    //for all objects, change isOpen: false to hide processMenu
    const sequenceExchangedArray = array.map((obj) => {
        return obj.id === sequenceProcessIdOne
            ? {
                  ...processOfIndexTwo,
                  sequence_num: processOfIndexOne.sequence_num,
                  isOpen: false,
              }
            : obj.id === sequenceProcessIdTwo
            ? {
                  ...processOfIndexOne,
                  sequence_num: processOfIndexTwo.sequence_num,
                  isOpen: false,
              }
            : {
                  ...obj,
                  isOpen: false,
              };
    });
    return sequenceExchangedArray;
};

export const sortProcessListBySequence = (array) => {
    const sortedArray = array.sort((a, b) => {
        return a.sequence_num - b.sequence_num;
    });
    return sortedArray;
};

//add or set isOpen: false to process-list objects
export const resetIsOpenInProcessList = (list) => {
    return list.map((obj) => ({ ...obj, isOpen: false }));
};

//will return processList array after deleting process with processId
//it will also re-arrange sequences as per index value
export const handleDeleteProcess = (processList, processId) => {
    const processListAfterDeleted = processList.filter(
        (obj) => obj.id !== processId
    );

    const rearrangedSequenceArr = processListAfterDeleted.map((obj, index) => {
        return {
            ...obj,
            sequence_num: index + 1,
        };
    });
    return rearrangedSequenceArr;
};

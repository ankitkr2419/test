/**returns array with exchanged sequence number of processes (with: indexOne, indexTwo) and sorted by sequence number */
export const changeProcessSequences = (array, idOne, idTwo) => {
    const processOfIndexOne = {
        ...array.find((pOne) => pOne.id === idOne),
    };
    const processOfIndexTwo = {
        ...array.find((pTwo) => pTwo.id === idTwo),
    };

    //change sequence logic: take process object of idOne and sequence number of idTwo (vice-versa)
    //for all objects, change isOpen: false to hide processMenu
    const sequenceExchangedArray = array.map((obj) => {
        return obj.id === idOne
            ? {
                  ...processOfIndexTwo,
                  sequence_num: processOfIndexOne.sequence_num,
                  isOpen: false,
              }
            : obj.id === idTwo
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
    return list.map((obj) => ({...obj, isOpen: false}))
}
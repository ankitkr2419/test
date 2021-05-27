import React from "react";
import { Text, ButtonIcon } from "shared-components";
import { ProcessCardBox } from "./ProcessCardBox";
import { InnerBox } from "./Styles";

const ProcessCard = (props) => {
    const {
        processId,
        processName,
        processIconName,
        isOpen,
        toggleIsOpen,
        draggedProcessId,
        setDraggedProcessId,
        handleChangeSequenceTo,
        handleProcessMove,
        createDuplicateProcess,
        handleEditProcess,
        handleDeleteProcess,
    } = props;

    const handleProcessMoveClick = () => {
        setDraggedProcessId(processId);
    };

    //when selected, previously selected process and this process will get sequence_number swapped.
    const handleDropHere = () => {
        handleChangeSequenceTo(processId);
    };

    const handleDuplicateProcessClick = () => {
        createDuplicateProcess(processId);
    };

    const getClassNameForMoveButton = () => {
        return `border-gray text-primary ml-2 ${
            draggedProcessId && draggedProcessId === processId
                ? "selected-button"
                : ""
        }`;
    };

    return (
        <div className="position-relative">
            <InnerBox>
                <ProcessCardBox
                    className={`d-flex justify-content-around flex-column bg-white ${
                        isOpen ? "selected-box " : ""
                    }`}
                >
                    <div className={`${isOpen ? "d-block" : "d-none"}`}>
                        {draggedProcessId && draggedProcessId !== processId ? (
                            <div className="d-flex justify-content-between align-items-center">
                                <Text
                                    className="drop-badge font-weight-bold text-white text-center"
                                    onClick={handleDropHere}
                                >
                                    Drop in this place
                                </Text>
                                <ButtonIcon
                                    position="absolute"
                                    placement="right"
                                    top={0}
                                    right={0}
                                    size={36}
                                    name="angle-down" //"cross"
                                    onClick={toggleIsOpen}
                                    className="border-0"
                                />
                            </div>
                        ) : (
                            <div className="d-flex justify-content-between align-items-center w-100 mb-2">
                                <div className="d-flex more-action">
                                    <ButtonIcon
                                        size={14}
                                        name="up"
                                        className="border-gray text-primary"
                                        onClick={() => handleProcessMove("up")}
                                    />
                                    <ButtonIcon
                                        size={14}
                                        name="down"
                                        className="border-gray text-primary ml-2"
                                        onClick={() =>
                                            handleProcessMove("down")
                                        }
                                    />
                                    <ButtonIcon
                                        size={24}
                                        name="move"
                                        className={getClassNameForMoveButton()}
                                        onClick={handleProcessMoveClick}
                                    />
                                    <ButtonIcon
                                        size={14}
                                        name="copy"
                                        className="border-gray text-primary ml-2"
                                        onClick={handleDuplicateProcessClick}
                                    />
                                </div>
                                <div className="d-flex more-action">
                                    <ButtonIcon
                                        size={14}
                                        name="edit-pencil"
                                        className="border-gray text-primary ml-2"
                                        onClick={handleEditProcess}
                                    />
                                    <ButtonIcon
                                        size={24}
                                        name="minus-1"
                                        className="border-gray text-primary ml-2"
                                        onClick={handleDeleteProcess}
                                    />
                                </div>
                            </div>
                        )}
                    </div>
                    <div
                        className="process-title d-flex align-items-center"
                        onClick={toggleIsOpen}
                    >
                        <ButtonIcon
                            size={14}
                            name={processIconName}
                            className="border-gray text-primary"
                        />
                        <Text Tag="label" className="mb-0">
                            {processName}
                        </Text>
                    </div>
                </ProcessCardBox>
            </InnerBox>
        </div>
    );
};

export default React.memo(ProcessCard);

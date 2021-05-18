import React from "react";
import { Card, CardBody, Row, Col } from "core-components";
import PaginationBox from "shared-components/PaginationBox";
import { Text } from "shared-components";
import ProcessCard from "./ProcessCard";
import { PROCESS_ICON_CONSTANTS } from "appConstants";

const ProcessListingCards = (props) => {
    let { processList, toggleIsOpen, draggedProcessId, setDraggedProcessId, handleChangeSequenceTo } =
        props;

    const getProcessIconName = (processType) => {
        let obj = PROCESS_ICON_CONSTANTS.find(
            (obj) => obj.processType === processType
        );

        let iconName = obj?.iconName
            ? obj.iconName
            : PROCESS_ICON_CONSTANTS.find(
                  (obj) => obj.processType === "default"
              ).iconName;
        return iconName;
    };

    return (
        <Card className="recipe-listing-cards">
            <CardBody className="p-5">
                <div className="d-flex justify-content-between align-items-center">
                    <Text Tag="span" className="recipe-name">
                        Total Processes: {processList?.length || 0}
                    </Text>

                    <div className="d-flex justify-content-end ml-auto">
                        <PaginationBox />
                    </div>
                </div>

                {/** Process List */}
                <Row className="pt-5">
                    {processList?.length > 0 ? (
                        processList.map((processObj) => {
                            return (
                                <Col md={4} key={processObj.id}>
                                    <ProcessCard
                                        processId={processObj.id}
                                        processName={processObj.name}
                                        processIconName={getProcessIconName(
                                            processObj.type
                                        )}
                                        isOpen={processObj.isOpen}
                                        toggleIsOpen={() =>
                                            toggleIsOpen(processObj.id)
                                        }
                                        draggedProcessId={draggedProcessId}
                                        setDraggedProcessId={
                                            setDraggedProcessId
                                        }
                                        handleChangeSequenceTo={handleChangeSequenceTo}
                                    />
                                </Col>
                            );
                        })
                    ) : (
                        <h4>No processes to show!</h4>
                    )}
                </Row>
            </CardBody>
        </Card>
    );
};

export default React.memo(ProcessListingCards);

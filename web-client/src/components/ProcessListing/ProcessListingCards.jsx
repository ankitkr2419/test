import React from "react";
import { Card, CardBody, Row, Col } from "core-components";
import PaginationBox from "shared-components/PaginationBox";
import { Text } from "shared-components";
import ProcessCard from "./ProcessCard";
import { PROCESS_ICON_CONSTANTS } from "appConstants";

const ProcessListingCards = (props) => {
    let {
        processList,
        toggleIsOpen,
        draggedProcessId,
        setDraggedProcessId,
        handleChangeSequenceTo,
        handleProcessMove,
    } = props;

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
                                        handleChangeSequenceTo={
                                            handleChangeSequenceTo
                                        }
                                        handleProcessMove={(direction) =>
                                            handleProcessMove(
                                                processObj.id,
                                                processObj.sequence_num,
                                                direction
                                            )
                                        }
                                    />
                                </Col>
                            );
                        })
                    ) : (
                        <h4>No processes to show!</h4>
                    )}
                </Row>

                {/**TODO: remove this after testing
                 *  - following is new UI alternative to process list,
                 *  and this need to try out with UI team */}
                {/*<div className="d-flex flex-column flex-wrap box pt-5">
                    {processList?.length > 0 ? (
                        processList.map((processObj) => {
                            return (
                                <div key={processObj.id}>
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
                                    
                                    />
                                </div>
                            );
                        })
                    ) : (
                        <h4>No processes to show!</h4>
                    )}
                </div>*/}

                {/**3rd solution */}
                {/*<div className="card-columns box">
                    {processList?.length > 0 ? (
                        processList.map((processObj) => {
                            return (
                                <div className="card" key={processObj.id}>
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
                                        
                                    />
                                </div>
                            );
                        })
                    ) : (
                        <h4>No processes to show!</h4>
                    )}
                    </div>*/}
            </CardBody>
        </Card>
    );
};

export default React.memo(ProcessListingCards);

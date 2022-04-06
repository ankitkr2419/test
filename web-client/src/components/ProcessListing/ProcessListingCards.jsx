import React, { useState, useEffect, useCallback } from "react";
import { Card, CardBody } from "core-components";
import PaginationBox from "shared-components/PaginationBox";
import { Text } from "shared-components";
import ProcessCard from "./ProcessCard";
import { SELECT_PROCESS_PROPS } from "appConstants";
import {
  paginator,
  initialPaginationStateProcessList,
} from "utils/paginationHelper";
import { useDispatch, useSelector } from "react-redux";
import {
  nextProcessPageBtn,
  prevProcessPageBtn,
} from "action-creators/PageActionCreators";

const ProcessListingCards = (props) => {
  let {
    deckName,
    processList,
    toggleIsOpen,
    draggedProcessId,
    setDraggedProcessId,
    handleChangeSequenceTo,
    handleProcessMove,
    createDuplicateProcess,
    handleEditProcess,
    handleDeleteProcess,
  } = props;

  const dispatch = useDispatch();
  const pageReducer = useSelector((state) => state.pageReducer);
  const page =
    deckName == "Deck A"
      ? pageReducer.processPageDeckA
      : pageReducer.processPageDeckB;
  const [paginatedData, setPaginatedData] = useState(
    initialPaginationStateProcessList
  );

  /**reset pagination when processList changed */
  useEffect(() => {
    findAndSetPagination();
  }, [processList]);

  /**reset pagination when page changed */
  useEffect(() => {
    findAndSetPagination();
  }, [page]);

  useEffect(() => {
    //if we dont have data on this page but having on previous page then go to previous page
    if (paginatedData?.list?.length === 0 && paginatedData?.total !== 0) {
      handlePrev();
    }
  }, [paginatedData]);

  const findAndSetPagination = () => {
    const data = paginator(processList, page, paginatedData.perPageItems);
    const newData = {
      ...paginatedData,
      page: data.page,
      prevPage: data.prePage,
      nextPage: data.nextPage,
      total: data.total,
      list: data.list,
      from: data.from,
      to: data.to,
    };
    setPaginatedData(newData);
  };

  const handleNext = useCallback(() => {
    if (paginatedData.nextPage) {
      dispatch(nextProcessPageBtn(deckName));
    }
  });

  const handlePrev = useCallback(() => {
    if (paginatedData.prevPage) {
      dispatch(prevProcessPageBtn(deckName));
    }
  });

  const getProcessIconName = useCallback((processType) => {
    let obj = SELECT_PROCESS_PROPS.find(
      (obj) => obj.processType === processType
    );

    let iconName = obj?.iconName
      ? obj.iconName
      : SELECT_PROCESS_PROPS.find((obj) => obj.processType === "default")
          .iconName;
    return iconName;
  });

  return (
    <Card className="recipe-listing-cards">
      <CardBody className="p-5">
        <div className="d-flex justify-content-between align-items-center">
          <Text Tag="span" className="recipe-name">
            Total Processes: {processList?.length || 0}
          </Text>

          <div className="d-flex justify-content-end ml-auto">
            <PaginationBox
              firstIndexOnPage={paginatedData.from}
              lastIndexOnPage={paginatedData.to}
              totalPages={paginatedData?.total || 0}
              handlePrev={handlePrev}
              handleNext={handleNext}
            />
          </div>
        </div>

        {/** Process List */}
        <div className="d-flex flex-column flex-wrap box py-4">
          {paginatedData?.list?.length > 0 ? (
            paginatedData?.list?.map((processObj, index) => {
              return (
                <div key={processObj.id}>
                  <ProcessCard
                    index={index}
                    page={page}
                    processId={processObj.id}
                    processName={processObj.name}
                    processIconName={getProcessIconName(processObj.type)}
                    isOpen={processObj.isOpen}
                    toggleIsOpen={() => toggleIsOpen(processObj.id)}
                    draggedProcessId={draggedProcessId}
                    setDraggedProcessId={setDraggedProcessId}
                    handleChangeSequenceTo={handleChangeSequenceTo}
                    handleProcessMove={(direction) =>
                      handleProcessMove(
                        processObj.id,
                        processObj.sequence_num,
                        direction
                      )
                    }
                    createDuplicateProcess={createDuplicateProcess}
                    handleEditProcess={() => handleEditProcess(processObj)}
                    handleDeleteProcess={() =>
                      handleDeleteProcess(processObj.id)
                    }
                  />
                </div>
              );
            })
          ) : (
            <h4>No processes to show!</h4>
          )}
        </div>
      </CardBody>
    </Card>
  );
};

export default React.memo(ProcessListingCards);

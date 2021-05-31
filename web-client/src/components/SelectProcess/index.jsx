import React from "react";

import { Row, Card, CardBody } from "core-components";
import { ButtonBar } from "shared-components";

import AppFooter from "components/AppFooter";
import Process from "./Process";
import { HeadingTitle, PageBody, ProcessOuterBox, TopContent } from "./Style";
import { SELECT_PROCESS_PROPS } from "appConstants";

const SelectProcessPageComponent = () => {
  return (
    <PageBody className="h-100">
      <ProcessOuterBox className="h-100">
        <div className="process-content select-process-bg">
          <TopContent className="d-flex justify-content-between align-items-center my-3 py-4">
            <div className="d-flex flex-column py-1">
              <HeadingTitle
                Tag="h5"
                className="text-primary font-weight-bold mb-0"
              >
                Select a process
              </HeadingTitle>
            </div>
          </TopContent>
          <Card className="process-content-box">
            <CardBody className="p-0">
              <Row className="row-small-gutter">
                {SELECT_PROCESS_PROPS.map((propObj) => {
                  return (
                    propObj.route && (
                      <Process
                        iconName={propObj.iconName}
                        processName={propObj.processName}
                        route={propObj.route}
                      />
                    )
                  );
                })}
              </Row>
            </CardBody>
          </Card>
          <ButtonBar />
        </div>
      </ProcessOuterBox>
      <AppFooter />
    </PageBody>
  );
};

SelectProcessPageComponent.propTypes = {};

export default SelectProcessPageComponent;

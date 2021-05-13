import React from "react";

import { Row, Card, CardBody, Col } from "core-components";
import { ButtonBar, ButtonIcon, Text } from "shared-components";

import AppFooter from "components/AppFooter";
import Process from "./Process";
import { HeadingTitle, PageBody, ProcessOuterBox, TopContent } from "./Style";

const SelectProcessPageComponent = () => {
  return (
    <>
      <PageBody>
        <ProcessOuterBox>
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
                  <Process iconName="piercing" processName="Piercing" />
                  <Process iconName="tip-pickup" processName="Tip Pickup" />
                  <Process
                    iconName="aspire-dispense"
                    processName="Aspire & Dispense"
                  />
                  <Process iconName="shaking" processName="Shaking" />
                  <Process iconName="heating" processName="Heating" />
                  <Process iconName="magnet" processName="Magnet" />
                  <Process iconName="tip-discard" processName="Tip Discard" />
                  <Process iconName="delay" processName="Delay" />
                  <Process iconName="tip-position" processName="Tip Position" />
                  {/* <Col md={4}>
									<div className="process-card bg-white d-flex align-items-center frame-icon">
										<ButtonIcon
												size={51}
												name='piercing'
												className="border-gray text-primary"
												//onClick={toggleExportDataModal}
										/>
										<Text Tag="span" className="ml-2 process-name">
											Piercing
										</Text>
									</div>
								</Col> */}
                </Row>
              </CardBody>
            </Card>
            <ButtonBar />
          </div>
        </ProcessOuterBox>
        <AppFooter />
      </PageBody>
    </>
  );
};

SelectProcessPageComponent.propTypes = {};

export default SelectProcessPageComponent;

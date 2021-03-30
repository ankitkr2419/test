import React from 'react';
import { ImageIcon, Center, Icon } from "shared-components";
import imgNoTemplate from "assets/images/video-thumbnail-poster.jpg";
import styled from "styled-components";
import { CardBody, Card } from "core-components";

const VideoPlayButton = styled.button`
  color: #7c7976;
  background-color: transparent;
  border: 0;
  outline: none;
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
`;

const VideoCard = () => {
  return (
    <Card className="card-video">
      <CardBody className="d-flex flex-column p-0">
        <Center className="video-thumbnail-wrapper">
          <ImageIcon
            src={imgNoTemplate}
            alt="No templates available"
            className="img-video-thumbnail"
          />
          <VideoPlayButton>
            <Icon name="play" size={124} />
          </VideoPlayButton>
        </Center>
      </CardBody>
    </Card>
  );
};

export default VideoCard;

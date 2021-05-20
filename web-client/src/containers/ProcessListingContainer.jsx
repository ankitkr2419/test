import React, { useEffect, useState } from "react";
import ProcessListComponent from "components/ProcessListing"

const ProcessListingContainer = (props) => {
    //TODO: fetch dynamic process list from recipeId
    const processList = [
        {
            id: "a5d058e3-7ce3-4a42-b2da-690e47139612",
            name: "AD-WW-c1-1-2",
            type: "AspireDispense",
            recipe_id: "6b7fcfa2-8337-4d79-829a-e9bd486a2de4",
            sequence_num: 1,
            created_at: "2021-03-11T17:34:16.507414Z",
            updated_at: "2021-03-11T17:34:16.507414Z",
        },
        {
            id: "4fa5c4e3-699c-47bb-ac7a-b26d04efaeb5",
            name: "AD-WS-c1-1",
            type: "AspireDispense",
            recipe_id: "6b7fcfa2-8337-4d79-829a-e9bd486a2de4",
            sequence_num: 2,
            created_at: "2021-03-11T17:34:16.555738Z",
            updated_at: "2021-03-11T17:34:16.555738Z",
        },
        {
            id: "fb88bada-ced7-4fa2-b845-4bb91e74341e",
            name: "AD-SW-c1-2",
            type: "AspireDispense",
            recipe_id: "6b7fcfa2-8337-4d79-829a-e9bd486a2de4",
            sequence_num: 3,
            created_at: "2021-03-11T17:34:16.580361Z",
            updated_at: "2021-03-11T17:34:16.580361Z",
        },
    ];

    return (
        <>
            <ProcessListComponent processList={processList} />
        </>
    );
};

export default React.memo(ProcessListingContainer);

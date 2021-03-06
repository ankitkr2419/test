import React from "react";
import { useHistory } from "react-router";
import moment from "moment";

import { Icon, ButtonIcon, Text } from "shared-components";
import { Button } from "core-components";
import { ROUTES } from "appConstants";

import { TopContent } from "./TopContent";
import { HeadingTitle } from "./HeadingTitle";

const TopContentComponent = (props) => {
    const {
        isProcessInProgress,
        onLogoutClicked,
        deckName,
        isAdmin,
        toggleAddNewRecipesModal,
        toggleTimeModal,
        toggleTrayDiscardModal,
        toggleLogoutModalVisibility,
        isProcessListingPage,
        recipeName,
        createdAt,
        updatedAt,
        processListBackButtonHandler,
    } = props;
    /**
     * prop: isProcessListingPage: false/undefined  means RecipeListing Page (default)
     * prop: isProcessListingPage: true  means ProcessListing Page
     */
    const history = useHistory();

    const classNameTopContent = () => {
        return `d-flex justify-content-between align-items-center mx-5 ${
            isProcessListingPage ? "process-listing-changes" : ""
        }`;
    };

    const handleManageUsersClick = () => {
        history.push(ROUTES.users);
    }

    return (
        <TopContent className={classNameTopContent()}>
            {/* Top content for RecipeListing page */}
            {isProcessInProgress || isProcessListingPage ? null : (
                <div className="d-flex align-items-center">
                    <div
                        style={{ cursor: "pointer" }}
                        onClick={onLogoutClicked}
                    >
                        <Icon
                            name="angle-left"
                            size={32}
                            className="text-white"
                        />
                    </div>
                    <HeadingTitle
                        Tag="h5"
                        className="text-white font-weight-bold ml-3 mb-0"
                    >
                        {`Select a Recipe for ${deckName}`}
                    </HeadingTitle>
                </div>
            )}

            {isProcessInProgress || isProcessListingPage ? null : (
                <div className="d-flex align-items-center ml-auto">
                    {isAdmin ? (
                        <>
                            <Button
                                color="secondary"
                                className="ml-2 border-primary btn-discard-tray bg-white"
                                onClick={toggleAddNewRecipesModal}
                            >
                                Add Recipe
                            </Button>
                            <Button
                                color="secondary"
                                className="ml-2 border-primary btn-discard-tray bg-white"
                                onClick={handleManageUsersClick}
                            >
                                Manage Users
                            </Button>
                        </>
                    ) : (
                        <>
                            {/* Might be needed later in v2.0 */}
                            {/* <ButtonIcon
                                name="download-1"
                                size={28}
                                className="bg-white border-primary"
                            /> */}
                            <Button
                                color="secondary"
                                className="ml-2 border-primary btn-clean-up bg-white"
                                onClick={toggleTimeModal}
                            >
                                {" "}
                                Clean Up
                            </Button>
                            <Button
                                color="secondary"
                                className="ml-2 border-primary btn-discard-tray bg-white"
                                onClick={toggleTrayDiscardModal}
                            >
                                Discard Tray
                            </Button>
                        </>
                    )}
                    <ButtonIcon
                        name="logout"
                        size={28}
                        className="ml-2 bg-white border-primary"
                        onClick={toggleLogoutModalVisibility}
                    />
                </div>
            )}

            {/* Top content for ProcessListing page */}
            {isProcessListingPage ? (
                <>
                    <div className="d-flex flex-column">
                        <div className="d-flex align-items-center">
                            <div
                                style={{ cursor: "pointer" }}
                                onClick={processListBackButtonHandler}
                            >
                                <Icon
                                    name="angle-left"
                                    size={32}
                                    className="text-white"
                                />
                            </div>
                            <div className="d-flex flex-column">
                                <Text className="text-white mb-0">Recipe</Text>
                                <HeadingTitle
                                    Tag="h5"
                                    className="text-white font-weight-bold mb-0"
                                >
                                    {recipeName}
                                </HeadingTitle>
                            </div>
                        </div>
                    </div>
                    <div className="d-flex justify-content-center align-items-center">
                        <div className="d-flex flex-column text-right">
                            {createdAt ? (
                                <Text className="text-white mb-0">
                                    Created on: {moment(createdAt).format("DD-MM-YYYY")}
                                </Text>
                            ) : null}
                            {updatedAt ? (
                                <Text className="text-white mb-0">
                                    Modified on: {moment(updatedAt).format("DD-MM-YYYY")}
                                </Text>
                            ) : null}
                        </div>
                    </div>
                </>
            ) : null}
        </TopContent>
    );
};

export default React.memo(TopContentComponent);

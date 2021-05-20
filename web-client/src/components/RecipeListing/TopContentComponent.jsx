import React from "react";
import { TopContent } from "./TopContent";
import { HeadingTitle } from "./HeadingTitle";
import { Icon, ButtonIcon } from "shared-components";
import { Button } from "core-components";

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
    } = props;

    return (
        <TopContent className="d-flex justify-content-between align-items-center mx-5">
            {isProcessInProgress ? null : (
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

            {isProcessInProgress ? null : (
                <div className="d-flex align-items-center ml-auto">
                    {isAdmin ? (
                        <Button
                            color="secondary"
                            className="ml-2 border-primary btn-discard-tray bg-white"
                            onClick={toggleAddNewRecipesModal}
                        >
                            Add Recipe
                        </Button>
                    ) : (
                        <>
                            <ButtonIcon
                                name="download-1"
                                size={28}
                                className="bg-white border-primary"
                            />
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
        </TopContent>
    );
};

export default React.memo(TopContentComponent);

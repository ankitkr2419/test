import React from "react";

import { Icon } from "shared-components";
import { NavItem, NavLink } from "reactstrap";
import classnames from "classnames";

import { LABWARE_ITEMS_NAME } from "appConstants";
import { updateAllTicks } from "./updateAllTicks";

const SideBarNavItems = (props) => {
  const { formik, activeTab, toggle } = props;

  return LABWARE_ITEMS_NAME.map((name, index) => {
    const currentState = formik.values;
    const key = Object.keys(currentState)[index];

    return (
      <NavItem key={key}>
        <NavLink
          className={classnames({ active: activeTab === `${index + 1}` })}
          onClick={() => {
            toggle(`${index + 1}`);
            updateAllTicks(formik);
          }}
        >
          {name}
          {currentState[key].isTicked ? (
            <Icon name="tick" size={12} className="ml-auto" />
          ) : null}
        </NavLink>
      </NavItem>
    );
  });
};

export default React.memo(SideBarNavItems);

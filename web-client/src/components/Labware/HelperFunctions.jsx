/**
 * old code for reference : will be deleted later!
 */

// export const getSideBarNavItems = (formik, activeTab, toggle) => {
//   const navItems = [];
//   LABWARE_ITEMS_NAME.forEach((name, index) => {
//     const currentState = formik.values;
//     const key = Object.keys(currentState)[index];
//     navItems.push(
//       <NavItem key={key}>
//         <NavLink
//           className={classnames({ active: activeTab === `${index + 1}` })}
//           onClick={() => {
//             toggle(`${index + 1}`);
//             updateAllTicks(formik);
//           }}
//         >
//           {name}
//           {currentState[key].isTicked ? (
//             <Icon name="tick" size={12} className="ml-auto" />
//           ) : null}
//         </NavLink>
//       </NavItem>
//     );
//   });
//   return navItems;
// };

// export const getTipsDropdown = (formik, options) => {
//   const tips = formik.values.tips;
//   const nDropdown = 3;
//   const tipsOptions = [];
//   for (let i = 0; i < nDropdown; i++) {
//     let tipPosition = tips.processDetails[`tipPosition${i + 1}`].id;
//     let index = options.map((item) => item.value).indexOf(tipPosition);
//     tipsOptions.push(
//       <FormGroup key={i} className="d-flex align-items-center mb-4">
//         <Label for={`tip-position-${i + 1}`} className="px-0 label-name">
//           Tip Position {i + 1}
//         </Label>
//         <div className="d-flex flex-column input-field position-relative">
//           <Select
//             placeholder="Select Option"
//             className=""
//             size="sm"
//             value={options[index]}
//             options={options}
//             onChange={(e) => {
//               formik.setFieldValue(
//                 `tips.processDetails.tipPosition${i + 1}.id`,
//                 e.value
//               );
//               formik.setFieldValue(
//                 `tips.processDetails.tipPosition${i + 1}.label`,
//                 e.label
//               );
//             }}
//           />
//           <FormError>Incorrect Tip Position {index + 1}</FormError>
//         </div>
//       </FormGroup>
//     );
//   }
//   return tipsOptions;
// };

// export const getTipsAtPosition = (position, formik, options) => {
//   const tips = formik.values.tips;
//   const tipPosition1Value = tips.processDetails.tipPosition1.id;
//   const tipPosition2Value = tips.processDetails.tipPosition2.id;
//   const tipPosition3Value = tips.processDetails.tipPosition3.id;

//   return (
//     <>
//       <div className="">
//         <div className="mb-3">
//           <FormGroup row>
//             <Label
//               for="select-tip-position"
//               md={12}
//               className="mb-3 font-weight-bold"
//             >
//               Select Tip Position
//             </Label>
//           </FormGroup>
//         </div>
//         <div className="">
//           {getTipsDropdown(formik, options)}
//           {/* <CommonField /> */}
//         </div>
//       </div>
//       <ProcessSetting>
//         <div className="tips-info">
//           <ul className="list-unstyled tip-position active">
//             {tipPosition1Value && (
//               <li className="highlighted tip-position-1"></li>
//             )}
//             {tipPosition2Value && (
//               <li className="highlighted tip-position-2 active"></li>
//             )}
//             {tipPosition3Value && (
//               <li className="highlighted tip-position-3 active"></li>
//             )}
//           </ul>
//           <ImageIcon src={labwareTips} alt="Tip Pickup Process" className="" />
//         </div>
//       </ProcessSetting>
//     </>
//   );
// };

// export const getTipPiercingCheckbox = (formik, nCheckboxes = 2) => {
//   const tipsPiercingCheckbox = [];
//   for (let index = 0; index < nCheckboxes; index++) {
//     let isChecked =
//       formik.values.tipPiercing.processDetails[`position${index + 1}`].id;
//     tipsPiercingCheckbox.push(
//       <CheckBox
//         id={`position${index + 1}`}
//         name={`position${index + 1}`}
//         label={`Position ${index + 1}`}
//         className={index > 0 ? "ml-4" : ""}
//         checked={isChecked ? true : false}
//         onChange={(e) => {
//           formik.setFieldValue(
//             `tipPiercing.processDetails.position${index + 1}.id`,
//             e.target.checked ? 3 : 0
//           );
//         }}
//       />
//     );
//   }
//   return tipsPiercingCheckbox;
// };

// export const getTipPiercingAtPosition = (position, formik) => {
//   const position1 = formik.values.tipPiercing.processDetails.position1.id;
//   const position2 = formik.values.tipPiercing.processDetails.position2.id;

//   return (
//     <>
//       <div className="mb-3">
//         <FormGroup row>
//           <Label
//             for="select-tip-piercing"
//             md={12}
//             className="mb-3 font-weight-bold"
//           >
//             Select Tip Piercing
//           </Label>
//         </FormGroup>
//       </div>

//       <div className="d-flex align-items-center">
//         {getTipPiercingCheckbox(formik)}
//         <ProcessSetting>
//           <div className="piercing-info">
//             <ul className="list-unstyled piercing-position active">
//               {position1 !== 0 && (
//                 <li className="highlighted piercing-position-1"></li>
//               )}
//               {position2 !== 0 && (
//                 <li className="highlighted piercing-position-2 active"></li>
//               )}
//             </ul>
//             <ImageIcon
//               src={labwarePiercing}
//               alt="Piercing Process"
//               className=""
//             />
//           </div>
//         </ProcessSetting>
//       </div>
//     </>
//   );
// };

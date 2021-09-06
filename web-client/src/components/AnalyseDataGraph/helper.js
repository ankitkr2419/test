/**
 *  Transform targets into (label, value) from (target_name, target_id,...)
 * to use in dropdown of analyse data filter
 */
export const generateTargetOptions = (targetList) => {
  return targetList?.map((target) => ({
    label: target.target_name,
    value: target.target_id,
    threshold: target.threshold,
  }));
};

export const lineConfigs = {
  fill: false,
  borderWidth: 2,
  pointRadius: 0,
  pointBorderColor: "rgba(148,147,147,1)",
  pointBackgroundColor: "#fff",
  pointBorderWidth: 0,
  pointHoverRadius: 0,
  pointHoverBackgroundColor: "rgba(148,147,147,1)",
  pointHoverBorderColor: "rgba(148,147,147,1)",
  pointHoverBorderWidth: 0,
  lineTension: 0.4,
  borderCapStyle: "butt",
};

export const lineConfigThreshold ={
  fill: false,
  borderWidth: 2,
  pointRadius: 0,
  pointBorderColor: "#a2ee95",
  borderDash: [10, 5],
  pointBackgroundColor: "#fff",
  pointBorderWidth: 0,
  pointHoverRadius: 0,
  pointHoverBackgroundColor: "#a2ee95",
  pointHoverBorderColor: "#a2ee95",
  pointHoverBorderWidth: 0,
}
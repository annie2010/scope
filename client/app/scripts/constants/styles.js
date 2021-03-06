
export const DETAILS_PANEL_WIDTH = 420;

export const DETAILS_PANEL_MARGINS = {
  top: 24,
  bottom: 48,
  right: 36
};

export const DETAILS_PANEL_OFFSET = 8;

export const CANVAS_METRIC_FONT_SIZE = 0.19;

export const CANVAS_MARGINS = {
  top: 160,
  left: 40,
  right: 40,
  bottom: 100,
};

//
// The base size the shapes were defined at matches nicely w/ a 14px font.
//
export const BASE_NODE_SIZE = 64;
export const MIN_NODE_SIZE = 24;
export const MAX_NODE_SIZE = 96;
export const BASE_NODE_LABEL_SIZE = 14;
export const MIN_NODE_LABEL_SIZE = 12;

// Node details table constants
export const NODE_DETAILS_TABLE_CW = {
  XS: '32px',
  S: '50px',
  M: '70px',
  L: '85px',
  XL: '120px',
  XXL: '140px',
  XXXL: '170px',
};

export const NODE_DETAILS_TABLE_COLUMN_WIDTHS = {
  count: NODE_DETAILS_TABLE_CW.XS,
  container: NODE_DETAILS_TABLE_CW.XS,
  docker_container_created: NODE_DETAILS_TABLE_CW.XXXL,
  docker_container_restart_count: NODE_DETAILS_TABLE_CW.M,
  docker_container_state_human: NODE_DETAILS_TABLE_CW.XXXL,
  docker_container_uptime: NODE_DETAILS_TABLE_CW.L,
  docker_cpu_total_usage: NODE_DETAILS_TABLE_CW.M,
  docker_memory_usage: NODE_DETAILS_TABLE_CW.M,
  open_files_count: NODE_DETAILS_TABLE_CW.M,
  pid: NODE_DETAILS_TABLE_CW.S,
  port: NODE_DETAILS_TABLE_CW.S,
  ppid: NODE_DETAILS_TABLE_CW.S,
  process_cpu_usage_percent: NODE_DETAILS_TABLE_CW.M,
  process_memory_usage_bytes: NODE_DETAILS_TABLE_CW.M,
  threads: NODE_DETAILS_TABLE_CW.M,

  // e.g. details panel > pods
  kubernetes_ip: NODE_DETAILS_TABLE_CW.XL,
  kubernetes_state: NODE_DETAILS_TABLE_CW.M,

  // weave connections
  weave_connection_connection: NODE_DETAILS_TABLE_CW.XXL,
  weave_connection_state: NODE_DETAILS_TABLE_CW.L,
  weave_connection_info: NODE_DETAILS_TABLE_CW.XL,
};

export const NODE_DETAILS_TABLE_XS_LABEL = {
  count: '#',
  // TODO: consider changing the name of this field on the BE
  container: '#',
};

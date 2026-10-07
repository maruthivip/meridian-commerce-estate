const core = require('@meridian/core');
module.exports = { load: (p) => ({ path: p, trace: core.traceId() }) };

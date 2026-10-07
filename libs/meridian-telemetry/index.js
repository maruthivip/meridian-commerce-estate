const cfg = require('@meridian/config');
module.exports = { emit: (e) => ({ e, cfg: cfg.load('telemetry') }) };

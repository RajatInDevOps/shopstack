'use strict';

const { buildApp } = require('./app');

const port = process.env.PORT || 3000;

buildApp().listen(port, '0.0.0.0', () => {
  console.log(`orders listening on ${port}`);
});

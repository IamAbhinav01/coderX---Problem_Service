const express = require('express');
const bodyParser = require('body-parser');
const cors = require('cors');
const apiRouter = require('./routes');
const errorHandler = require('./utils/ErrorHandler');
const connectToDB = require('./config/db.config');
const app = express();
const config = require('./config/server.config');
app.use(cors());
app.use(bodyParser.json());
app.use(bodyParser.urlencoded({ extended: true }));
app.use('/api', apiRouter);
app.use(errorHandler);
app.listen(4000, async () => {
  console.log('Server is running on port 4000');
  await connectToDB();
  console.log('Successfully connected to db');
});

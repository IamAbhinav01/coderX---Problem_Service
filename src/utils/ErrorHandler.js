const { StatusCodes } = require('http-status-codes');
const BaseError = require('../errors/Base.err');
function errorHandler(err, req, res, next) {
  
  
  
  
  if (err instanceof BaseError) {
    return res.status(err.statusCode).json({
      message: err.message,
      details: err.details,
      data: {},
    });
  }
  return res.status(StatusCodes.INTERNAL_SERVER_ERROR).json({
    success: false,
    message: err.message,
    error: err.details,
    data: {},
  });
}
module.exports = errorHandler;

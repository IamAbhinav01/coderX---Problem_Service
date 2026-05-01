const { StatusCodes } = require('http-status-codes');
const BaseError = require('./Base.err');

class NotFoundError extends BaseError {
  constructor(resourceName, details) {
    super('NotFoundError', StatusCodes.NOT_FOUND, `The requested resource ${resourceName} was not found`, details);
  }
}
module.exports = NotFoundError;

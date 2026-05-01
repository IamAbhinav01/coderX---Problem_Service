const { StatusCodes } = require('http-status-codes');
const BaseError = require('./Base.err');

class BadRequest extends BaseError {
  constructor(resourceName, details) {
    super('BadRequest', StatusCodes.BAD_REQUEST, `Invalid values for ${resourceName} provided`, details);
  }
}
module.exports = BadRequest;

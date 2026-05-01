const { StatusCodes } = require('http-status-codes');
const BaseError = require('./Base.err');

class NotImplemented extends BaseError {
  constructor(methodName) {
    super('NotImplemented', StatusCodes.NOT_IMPLEMENTED, `${methodName} is not implemented`, {});
  }
}
module.exports = NotImplemented;

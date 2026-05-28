const { Problem } = require('../models');
const NotFoundError = require('../errors/NotFound.err');
const BadRequest = require('../errors/BadRequest.err');

class ProblemRepository {
  async createProblem(problemData) {
    try {
      const problem = await Problem.create({
        ...problemData
      });
      return problem;
    } catch (err) {
      console.error('Error in creating problem:', err);
      throw new BadRequest('CreateProblem', err.message);
    }
  }

  async getProblem(problemId) {
    try {
      const problem = await Problem.findById(problemId);
      if (!problem) {
        throw new NotFoundError('Problem', { id: problemId });
      }
      return problem;
    } catch (err) {
      if (err instanceof NotFoundError) throw err;
      if (err.name === 'CastError') {
        throw new NotFoundError('Problem', { id: problemId });
      }
      console.error(`Error finding problem ${problemId}:`, err);
      throw err;
    }
  }

  async getProblems() {
    try {
      const problems = await Problem.find({});
      return problems;
    } catch (err) {
      console.error('Error finding problems:', err);
      throw new NotFoundError('Problems', err.message);
    }
  }

  async deleteProblem(problemId) {
    try {
      const result = await Problem.findByIdAndDelete(problemId);
      if (!result) {
        throw new NotFoundError('Problem', { id: problemId });
      }
      return result;
    } catch (err) {
      if (err instanceof NotFoundError) throw err;
      if (err.name === 'CastError') {
        throw new NotFoundError('Problem', { id: problemId });
      }
      console.error(`Error deleting problem ${problemId}:`, err);
      throw err;
    }
  }

  async updateProblem(problemId, problemData) {
    try {
      const problem = await Problem.findByIdAndUpdate(
        problemId,
        { $set: problemData },
        { new: true, runValidators: true }
      );
      if (!problem) {
        throw new NotFoundError('Problem', { id: problemId });
      }
      return problem;
    } catch (err) {
      if (err instanceof NotFoundError) throw err;
      if (err.name === 'CastError') {
        throw new NotFoundError('Problem', { id: problemId });
      }
      console.error(`Error updating problem ${problemId}:`, err);
      throw err;
    }
  }
}

module.exports = ProblemRepository;

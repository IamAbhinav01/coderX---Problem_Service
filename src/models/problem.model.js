const mongoose = require('mongoose');
const schema = mongoose.Schema;

const problemSchema = new schema({
  _id: {
    type: String,
    required: true,
    default: () => new mongoose.Types.ObjectId().toString(),
  },
  title: {
    type: String,
    required: [true, 'title is required'],
  },

  description: {
    type: String,
    required: [true, 'description is required'],
  },

  difficulty: {
    type: String,
    enum: ['easy', 'medium', 'hard'],
    required: [true, 'difficulty is required'],
    default: 'easy',
  },

  testCases: [
    {
      input: {
        type: String,
        required: true,
      },
      output: {
        type: String,
        required: true,
      },
    },
  ],

  codeSnippets: [
    {
      language: {
        type: String,
        enum: ['python', 'java', 'cpp', 'Python', 'Java', 'CPP'],
        required: true,
      },

      startSnippet: {
        type: String,
        required: true,
      },

      midSnippet: {
        type: String,
        default: '',
      },

      endSnippet: {
        type: String,
        required: true,
      },
    },
  ],

  editorial: {
    type: String,
  },

  topic: {
    type: String,
  },

  createdAt: {
    type: Date,
    default: Date.now,
  },
});

const Problem = mongoose.model('Problem', problemSchema);

module.exports = Problem;

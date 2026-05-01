const mongoose = require('mongoose');
const Problem = require('./src/models/problem.model');
require('dotenv').config();

async function checkDB() {
    try {
        await mongoose.connect(process.env.ATLAS_DB_URL);
        console.log('Connected to DB');
        const problems = await Problem.find({});
        console.log('Problems count:', problems.length);
        if (problems.length > 0) {
            console.log('Sample Problem ID:', problems[0]._id);
        } else {
            console.log('No problems found in DB');
        }
        process.exit(0);
    } catch (err) {
        console.error(err);
        process.exit(1);
    }
}

checkDB();

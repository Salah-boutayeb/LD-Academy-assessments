const express = require("express");
const cors = require("cors");
const app = express();
const db = require("./database/models");
const { errorHandler } = require("./middleware/middleware");
const dotenv = require("dotenv").config();
// express middleware
const bodyParser = require("body-parser");
app.use(bodyParser.urlencoded({ extended: false }));
app.use(bodyParser.json());
require("dotenv").config();
app.use(cors());

// use routes here

app.use("/api/users", require("./routes/userRoutes"));
app.use("/api/categories", require("./routes/categoryRoutes"));
app.use("/api/recipes", require("./routes/recipeRoutes"));

const port = 5000;
app.use(errorHandler);
db.sequelize.sync({ alter: true });
app.listen(port, () => {
  console.log(`server's runing on port ${port}`);
});

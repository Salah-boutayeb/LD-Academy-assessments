const asyncHandler = require("express-async-handler");
const bycrypt = require("bcryptjs");
const jwt = require("jsonwebtoken");
const { User } = require("../database/models");

const protect = asyncHandler(async (req, res, next) => {
  let token;
  if (
    req.headers.authorization &&
    req.headers.authorization.startsWith("Bearer")
  ) {
    try {
      token = req.headers.authorization.split(" ")[1];
      console.log(token);

      const decodedToken = jwt.verify(token, process.env.JWT_SECRET);

      req.user = await User.findByPk(decodedToken.id);

      next();
    } catch (error) {
      res.status(400);
      console.error("error authorization", error);
      throw new Error("no authorization");
    }
  }

  if (!token) {
    res.status(400);

    throw new Error("no authorization,no token");
  }
});

module.exports = { protect };

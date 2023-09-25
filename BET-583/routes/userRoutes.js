const express = require("express");
const router = express.Router();
const {
  registerUser,
  userLogin,
  getUserRecipes,
  addToFavourite,
  getUserFavoriteRecipes,
} = require("../controllers/userController");
const { protect } = require("../middleware/authMiddleware");
router.route("/signup").post(registerUser);
router.route("/login").post(userLogin);
router.route("/favorite_recipes/:id").post(protect, addToFavourite);
router.route("/favorite_recipes").get(protect, getUserFavoriteRecipes);
router.route("/recipes").get(protect, getUserRecipes);

module.exports = router;

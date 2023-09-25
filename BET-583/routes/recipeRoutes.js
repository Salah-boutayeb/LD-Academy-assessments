const express = require("express");
const router = express.Router();

const {
  getRecipes,
  getRecipeById,
  createRecipe,
  updateRecipe,
  deleteRecipe,
} = require("../controllers/recipeController");

router.route("/").get(getRecipes);
router.route("/:id").get(getRecipeById);
router.route("/").post(createRecipe);
router.route("/:id").put(updateRecipe);
router.route("/:id").delete(deleteRecipe);
module.exports = router;

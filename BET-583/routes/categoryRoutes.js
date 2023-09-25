const express = require("express");
const router = express.Router();
const {
  createCategory,
  updateCategory,
  deleteCategory,
  getCategories,
  getCategoryById,
  getCategoryRecipes,
} = require("../controllers/categoryController");

router.route("/").get(getCategories);
router.route("/:id").get(getCategoryById);
router.route("/:id/recipes").get(getCategoryRecipes);
router.route("/").post(createCategory);
router.route("/:id").put(updateCategory);
router.route("/:id").delete(deleteCategory);
module.exports = router;

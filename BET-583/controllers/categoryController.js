const asyncHandler = require("express-async-handler");
const { Category, Recipe } = require("../database/models");
const createCategory = asyncHandler(async (req, res) => {
  const { title, description, parentCategoryId } = req.body;

  await Category.create({
    title,
    description,
    parentCategoryId,
  })
    .then((result) => {
      return res.status(201).json({ result });
    })
    .catch((err) => {
      return res.status(500).json({ error: err });
    });
});

const updateCategory = asyncHandler(async (req, res) => {
  const { title, description, parentCategoryId } = req.body;
  const id = req.params.id;
  Category.findOne({ where: { id } })
    .then((category) => {
      if (!category) {
        throw new Error("No category found");
      }

      let values = {
        title,
        description,
        parentCategoryId,
      };

      category.update(values).then((updatedcategory) => {
        return res.status(201).json({ updatedcategory });
      });
    })
    .catch((error) => {
      return res.status(500).json({ error: err });
    });
});

const deleteCategory = asyncHandler(async (req, res) => {
  const id = req.params.id;
  Category.findOne({ where: { id } })
    .then((category) => {
      if (!category) {
        throw new Error("No category found");
      }

      category
        .destroy()
        .then(() =>
          res.status(202).json({ succes: true, message: "category deleted" })
        );
    })
    .catch((error) => {
      return res.status(500).json({ error: err });
    });
});

const getCategories = asyncHandler(async (req, res) => {
  await Category.findAll({
    include: {
      model: Category,
      as: "SubCategories",
      include: "SubCategories",
    },
  })
    .then((categories) => {
      return categories
        ? res.status(200).json({ categories, succes: true })
        : res.status(400).json({ succes: false, message: "not found" });
    })
    .catch((err) => {
      return res.status(400).json({ err });
    });
});
const getCategoryById = asyncHandler(async (req, res) => {
  const id = req.params.id;

  await Category.findByPk(id, {
    include: {
      model: Category,
      as: "SubCategories",
      include: "SubCategories",
    },
  })
    .then((category) => {
      return category
        ? res.status(200).json({ category, succes: true })
        : res.status(400).json({ succes: false, message: "not found" });
    })
    .catch((err) => {
      return res.status(400).json({ succes: false, message: "not found" });
    });
});

const getCategoryRecipes = asyncHandler(async (req, res) => {
  const id = req.params.id;

  await Category.findByPk(id, {
    include: {
      model: Recipe,
    },
  })
    .then((category) => {
      return category
        ? res.status(200).json({ category, succes: true })
        : res.status(400).json({ succes: false, message: "not found" });
    })
    .catch((err) => {
      return res.status(400).json({ succes: false, message: "not found" });
    });
});
module.exports = {
  createCategory,
  updateCategory,
  deleteCategory,
  getCategories,
  getCategoryById,
  getCategoryRecipes,
};

const asyncHandler = require("express-async-handler");
const { Recipe, Ingredient } = require("../database/models");
const createRecipe = asyncHandler(async (req, res) => {
  const { title, description, userId, categoryId, ingredients } = req.body;

  let recipe = await Recipe.create({
    title,
    description,
    userId,
    categoryId,
  });

  for (const ingredientData of ingredients) {
    await Ingredient.create({
      recipeId: recipe.id,
      name: ingredientData.name,
    });
  }

  recipe = await Recipe.findByPk(recipe.id, {
    include: Ingredient,
  });

  return res.status(201).json({ recipe });
});

const updateRecipe = asyncHandler(async (req, res) => {
  const { title, description, ingredients } = req.body;
  const id = req.params.id;
  Recipe.findOne({ where: { id } })
    .then((recipe) => {
      if (!recipe) {
        throw new Error("No category found");
      }

      let values = {
        title,
        description,
      };

      recipe.update(values).then(async (updatedRecipe) => {
        await Ingredient.destroy({ where: { recipeId: id } });

        for (const ingredientData of ingredients) {
          await Ingredient.create({
            name: ingredientData.name,
            quantity: ingredientData.quantity,
            recipeId: id,
          });
        }
      });
    })
    .catch((error) => {
      return res.status(500).json({ error: err });
    });
  const recipe = await Recipe.findByPk(id, { include: Ingredient });

  if (!recipe) {
    return res.status(404).json({ error: "Recipe not found" });
  }

  return res.status(200).json({ recipe });
});

const deleteRecipe = asyncHandler(async (req, res) => {
  const id = req.params.id;
  Recipe.findOne({ where: { id } })
    .then((recipe) => {
      if (!recipe) {
        return res.status(400).json({ succes: false, message: "not fount" });
      }

      recipe
        .destroy()
        .then(() =>
          res.status(202).json({ succes: true, message: "recipe deleted" })
        );
    })
    .catch((error) => {
      return res.status(500).json({ error });
    });
});

const getRecipes = asyncHandler(async (req, res) => {
  await Recipe.findAll({
    include: Ingredient,
  })
    .then((recipes) => {
      return recipes
        ? res.status(200).json({ recipes, succes: true })
        : res.status(400).json({ succes: false, message: "not found" });
    })
    .catch((err) => {
      return res.status(400).json({ err });
    });
});

const getRecipeById = asyncHandler(async (req, res) => {
  const id = req.params.id;

  await Recipe.findByPk(id, {
    include: Ingredient,
  })
    .then((recipe) => {
      return recipe
        ? res.status(200).json({ recipe, succes: true })
        : res.status(400).json({ succes: false, message: "not found" });
    })
    .catch((err) => {
      return res.status(400).json({ succes: false, message: "not found" });
    });
});
module.exports = {
  createRecipe,
  updateRecipe,
  deleteRecipe,
  getRecipes,
  getRecipeById,
};

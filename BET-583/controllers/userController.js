const asyncHandler = require("express-async-handler");
const bycrypt = require("bcryptjs");
const jwt = require("jsonwebtoken");
const { User, Recipe, Ingredient, Favorite } = require("../database/models");
const generateJwt = (id) => {
  return jwt.sign({ id }, process.env.JWT_SECRET, {
    expiresIn: "30d",
  });
};

const registerUser = asyncHandler(async (req, res) => {
  const { firstname, lastname, email, password } = req.body;
  if (!lastname || !firstname || !email || !password) {
    res.status(400);
    throw new Error("please fill all fields");
  }
  const userExists = await User.findOne({ where: { email } });
  if (userExists) {
    res.status(400);
    throw new Error("user already exists");
  }
  const salt = await bycrypt.genSalt(10);
  const hashPassword = await bycrypt.hash(password, salt);
  const user = await User.create({
    firstname,
    lastname,
    email,
    password: hashPassword,
  });
  /* create user */

  if (user) {
    res.status(201).json({
      id: user.id,
      name: user.name,
      email: user.email,
      score: user.score,
      token: generateJwt(user.id),
    });
  } else {
    res.status(400);
    throw new Error("invalid user data");
  }
});
/*  *********************************************************  */
const userLogin = asyncHandler(async (req, res) => {
  const { email, password } = req.body;
  if (!email || !password) {
    res.status(400);
    throw new Error("please fill all fields");
  }
  const userExists = await User.findOne({ where: { email } });
  if (userExists && (await bycrypt.compare(password, userExists.password))) {
    res.status(201).json({
      id: userExists.id,
      name: userExists.name,
      email: userExists.email,
      score: userExists.score,
      token: generateJwt(userExists.id),
    });
  } else {
    res.status(400);
    throw new Error("invalid credentials");
  }
});
const getUserRecipes = asyncHandler(async (req, res) => {
  const userId = req.user.id;

  await Recipe.findAll({
    where: { userId },
    include: {
      model: Ingredient,
    },
  })
    .then((recipes) => {
      return recipes
        ? res.status(200).json({ recipes, succes: true })
        : res.status(400).json({ succes: false, message: "not found" });
    })
    .catch((err) => {
      return res.status(400).json({ succes: false, message: "not found" });
    });
});

const addToFavourite = asyncHandler(async (req, res) => {
  try {
    const recipeId = req.params.id;
    const userId = req.user.id;
    const favoriteRecipe = await Favorite.findOne({
      where: { UserId: userId, RecipeId: recipeId },
    });
    if (favoriteRecipe) {
      return res.status(200).json({ message: "already in favorite" });
    }
    const recipe = await Recipe.findByPk(recipeId);

    if (!recipe) {
      return res.status(404).json({ error: "Recipe not found" });
    }

    recipe.nbr_likes += 1;

    await recipe.save();

    const user = await User.findByPk(userId);
    if (!user) {
      return res.status(404).json({ error: "User not found" });
    }
    await user.addFavoriteRecipe(recipeId);

    return res.status(201).json({ message: "Recipe added to favorites" });
  } catch (error) {
    console.error(error);
  }
});
const getUserFavoriteRecipes = async (req, res) => {
  const userId = req.user.id;

  const user = await User.findByPk(userId, {
    include: [{ model: Recipe, as: "FavoriteRecipes" }],
  });

  if (!user) {
    return res.status(404).json({ error: "User not found" });
  }

  const favoriteRecipes = user.FavoriteRecipes;

  return res.status(200).json({ favoriteRecipes });
};

module.exports = {
  registerUser,
  userLogin,
  getUserRecipes,
  addToFavourite,
  getUserFavoriteRecipes,
};

"use strict";
const { Model } = require("sequelize");

module.exports = (sequelize, DataTypes) => {
  class Recipe extends Model {
    /**
     * Helper method for defining associations.
     * This method is not a part of Sequelize lifecycle.
     * The `models/index` file will call this method automatically.
     */
    static associate(models) {
      // define association here
      Recipe.belongsTo(models.User, {
        foreignKey: "userId",
      });
      // Recipe.belongsToMany(models.User, { through: "favorite_recipe" });
      Recipe.hasMany(models.Ingredient, { foreignKey: "recipeId" });
    }
  }
  Recipe.init(
    {
      title: DataTypes.STRING,
      description: DataTypes.TEXT,
      nbr_likes: { type: DataTypes.INTEGER, defaultValue: 0 },
    },
    {
      sequelize,
      modelName: "Recipe",
    }
  );
  return Recipe;
};

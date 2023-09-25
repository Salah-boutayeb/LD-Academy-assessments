const { Model } = require("sequelize");

module.exports = (sequelize, DataTypes) => {
  class Favorite extends Model {}
  Favorite.init(
    {
      // You may include additional fields here if needed
      UserId: {
        type: DataTypes.INTEGER,
        allowNull: false,
      },
      RecipeId: {
        type: DataTypes.INTEGER,
        allowNull: false,
        unique: "favorite_unique_constraint", // Add a unique constraint
      },
    },
    {
      sequelize,
      modelName: "Favorite",
    }
  );

  return Favorite;
};

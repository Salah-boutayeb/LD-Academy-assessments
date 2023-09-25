"use strict";
const { Model } = require("sequelize");
module.exports = (sequelize, DataTypes) => {
  class Category extends Model {
    /**
     * Helper method for defining associations.
     * This method is not a part of Sequelize lifecycle.
     * The `models/index` file will call this method automatically.
     */
    static associate(models) {
      // define association here
      Category.hasMany(models.Recipe, { foreignKey: "categoryId" });

      // Self-relationship for sub-categories
      Category.hasMany(models.Category, {
        as: "SubCategories",
        foreignKey: "parentCategoryId",
      });
      Category.belongsTo(models.Category, {
        as: "ParentCategory",
        foreignKey: "parentCategoryId",
      });
    }
  }
  Category.init(
    {
      title: DataTypes.STRING,
      description: DataTypes.TEXT,

      parentCategoryId: DataTypes.INTEGER,
    },
    {
      sequelize,
      modelName: "Category",
    }
  );
  return Category;
};

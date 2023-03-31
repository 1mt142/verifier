go get -u golang.org/x/crypto/bcrypt 

go get -u github.com/golang-jwt/jwt/v5

In the example code I provided, there are three tables defined in the database: articles, tags, and categories. The relationships among these tables are defined as follows:

An article can have only one category, and a category can have many articles. This is a one-to-many relationship between categories and articles.
An article can have many tags, and a tag can be associated with many articles. This is a many-to-many relationship between articles and tags.
There is no direct relationship between categories and tags.
To implement these relationships using GORM, we use foreign key constraints and association tables. The Article model has a foreign key constraint to the Category model (CategoryID), which allows us to define the one-to-many relationship between categories and articles. The Article model also has a many-to-many relationship with the Tag model through an association table called article_tags, which allows us to associate tags with articles.
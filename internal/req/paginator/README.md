# Paginator

The paginator package adds pagination to a query. It consist of two functioning components that
will plug in to two layers of your application.

The `PaginatorControllerFacilitator` will plug in to your handler layer where it will convert the incoming 
data structure in to a paginator.

The `PaginatorRepositoryFacilitator` will plug in to your data retrieving layer where it will create 
the SQL `LIMIT, OFFSET` clause using the paginator and attach it to the passed in query.

However, `PaginatorRepositoryFacilitator` will generate queries **compatible** with
[kosatnkn/db](https://github.com/kosatnkn/db).
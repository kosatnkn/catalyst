# Filter

The filter package parses a populated struct as a slice of filters. It consist of two functioning components that
will plug in to two layers of your application.

The `FilterControllerFacilitator` will plug in to your handler layer where it will convert the incoming data structure
in to a filter slice.

The `FilterRepositoryFacilitator` will plug in to your data retrieving layer where it will create 
the SQL `WHERE` clause using the filter slice and attach it to the passed in query.

However, `FilterRepositoryFacilitator` will generate queries **compatible** with
[kosatnkn/db](https://github.com/kosatnkn/db).
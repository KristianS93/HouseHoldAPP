# HouseHoldApp

The household app is divided into three different microservices.

1. **GroceryList** - a service that maintains a list of grocery items, items can be added, delete, individually edited or the list can be cleared.
2. **UserSystem** - a service maintaining connection to a database responsible for all users in the system and their associated information.
3. **MealPlanner** - a service that suggests meals based on a user defined lists of meals - ingredients for the meals can also be added to the users grocery list easily.

Additionally, there is a central server responsible for the front end and connection to other services.

Currently, GroceryList and MealPlanner run in Docker containers.
 





# BookStoreCRUDOperations
The book store application is a simple web-based application that allows users to create, read, update, and delete books. The application is built using the Gorilla Mux router, which makes it easy to handle incoming HTTP requests and direct them to the appropriate handler function.

The application has four main endpoints:

+GET /books - This endpoint retrieves a list of all the books in the store.

+GET /books/{id} - This endpoint retrieves a specific book by its ID.

+POST /books - This endpoint adds a new book to the store.

+PUT /books/{id} - This endpoint updates an existing book by its ID.

DELETE /books/{id} - This endpoint deletes a book by its ID.

Each endpoint is handled by a specific function in the application, which uses the appropriate HTTP method (GET, POST, PUT, or DELETE) to perform the corresponding CRUD operation on the book store data.

To store the book data, the application uses a simple in-memory slice of book structs. Each book struct contains an ID, a title, an author, and a year of publication. When a new book is added to the store, it is assigned a unique ID and appended to the slice. When a book is updated or deleted, its ID is used to locate it in the slice and modify or remove it as necessary.

Overall, the book store application is a simple but functional example of how to use Gorilla Mux and CRUD operations in Go to build a web-based application that interacts with data.




### Exercise 1

We build RESTful API's to deliver financial and credit services to our clients.

Due to our business's nature, we maintain a PCI Level 1 Service Provider environment for the transmission, storage, and retrieval of sensitive data. Our engineers go through yearly security training, and our infrastructure and code go through regular security audits.

The first step in building a new consumer-facing web application is the successful storage of authentication credentials. We would like you to create a web application that allows users to sign up for an imaginary website and then log in securely.

##### The objectives are:

- Create a web application in a private repository called `<lastname>-backend-1`.
- Build the necessary REST endpoints that allow a user to sign up with an email address and password, login with an existing email address and password, and log out.
- Design a database schema for your preferred database engine to store and query the credentials. Make sure to include the SQL table creation scripts in your repository.
- Create a basic index page at the root `/` of the project that allows a user to login.
- Also, create a `/members` area that informs the user they are logged in and allows them to log out. Redirect the user to this page upon successful login.
- Inform the user when their login credentials are not valid.
- Create a `Dockerfile` and `docker-compose.yml` that allows us to build and start your project via `docker-compose up -d`. 

##### When complete:

Provide a link to your private repository to your recruiter or company representative. Make sure your `README` includes how to start your application and anything we should know. Additionally, please include notes on any design choices you made throughout the exercise (dependencies, libraries, assumptions, etc.).

Lastly, in your `README`, please note how many hours you spent on the exercise.
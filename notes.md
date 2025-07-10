Day 2: Simple Go Registration Form App

Task description: 
The web application should be a basic user registration. The registration process consists of 4 separate steps. Only one step is shown at a time to the customer. The user is able to leave the registration on every step/view until he finished the whole registration successfully. Accordingly, he should be redirected to the last opened step when he's joining the process again.

View 1: Insert Personal information
- Firstname, lastname, email

View 2: Insert address information
- Address including street, house number, zip code, city

View 3: Insert payment information
- Adress including street, house number, zip code, city

View 4: Success page
- Success message, when data are successful saved
- Show the "paymentDataId" returned by the payment api you will also create.


Guidelines: PostgreSQL, Hexagonal Architecture, Domain Driven Design, Model View Controller, Redis

Firstly implement only the backend in Go

- Sessions, Tokens, Interfaces, Cookies
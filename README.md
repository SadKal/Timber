# Timber

Timber is a real-time messaging app, with the possibility of creating various chats/groups with other users.

**Timber** is the app I am developing for my final project.

To get it running, first you need to create the backend's .env with the info provided in the .env.example. You will need a MySQL database set up.

Same with fronend. Fill .env with the data it needs.

Then, on the frontend, you need to use:

`npm install`

Then, open a terminal tab for backend and another one for frontend.

In the backend one, run:

`go run .`

and in the frontend one:

`npm run dev`

Open the direction the console says and if everything is configured correctly,
you'll be able to use the app.

Another option to get it running is having docker compose installed.
Then only filling the frontend's .env is required. Just write there localhost:"the port mapped in the dockerfile" and use:
`docker-compose up`

Please do have in mind that this is just for local use so JWT secret key and all env variables are not that relevant.

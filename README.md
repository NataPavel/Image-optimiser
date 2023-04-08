# Image-optimizer
> ### Image optimizer is an application that resizes to 100%, 75%, 50%, 25% of the original image
> ### Several facts that you should know about app:
> - #### To run the application, write in the terminal: `go run cmd/main.go`
> - #### The app runs on port :8080
> - #### Images stores in MySql database. The app uses [goose package](https://github.com/pressly/goose) for migrations
> - #### To migrate tables type `goose -dir migrations mysql "<path to database>" up`
> - #### Migrations are stored in `migrations` folder
> - #### Passwords to RabbitMq and the database stored in the `.env' file included in .gitignore
> - #### The applications uses [resize package](https://github.com/nfnt/resize) to resize images
>-------------------------------------------------------------------

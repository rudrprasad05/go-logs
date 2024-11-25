# What is it?
Just something i made up in one day...so def got some bugs. Is a working product tho. Look at main.go for an exmaple

# How to use it
router := http.NewServeMux()
loggedHandler := logs.LoggingMiddleware(logger, router)

Just wrap your router in the logger. Acts as middleware. Intecepts HTTP request and logs them. 

# Feel free to edit this

while true
do
    # Your action or command goes here
    git config --global user.email "you@example.com"
    git config --global user.name "Your Name"
    git add .
    git commit -m "fixing css"
    git push
    
    # Sleep for 20 minutes (1200 seconds)
    sleep 60
done
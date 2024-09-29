MACHINE_NAME=daily-go-bike

fly machine create . --name $MACHINE_NAME --schedule daily

MACHINE_ID=$(fly machine list | grep daily-go-bike | awk '{ print $1 }')